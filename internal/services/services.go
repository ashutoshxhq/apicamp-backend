package services

import (
	"bytes"
	context "context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/apicamp/backend/helpers"
	"github.com/apicamp/backend/templates"
)

// Server represents the gRPC server
type Server struct {
}

// GenerateServiceCode fuctions generates service code
func (s *Server) GenerateServiceCode(ctx context.Context, req *GenerateServiceCodeRequest) (*GenerateServiceCodeResponse, error) {

	query := map[string]string{
		"query": `
            { 
				services(where: {id: {_eq: "` + req.ServiceId + `"}}){
					id
					name
					models {
						name
						service_id
						updated_at
						id
						relationships {
						  created_at
						  deleted
						  id
						  model_field_id
						  model_id
						  name
						  relationship_model_field_id
						  relationship_model_id
						  type
						  updated_at
						  relationshipModel {
							id
							name
						  }
						  relationshipModelField {
							id
							name
							type
						  }
						  modelField {
							id
							name
							type
						  }
						}
						fields {
						  id
						  type
						  name
						  key
						  created_at
						  default
						  deleted
						  model_id
						  null_value
						  updated_at
						}
					}
				}   
            }
        `,
	}
	jsonValue, _ := json.Marshal(query)
	request, err := http.NewRequest("POST", "https://apicamp-graphql.herokuapp.com/v1/graphql", bytes.NewBuffer(jsonValue))
	request.Header.Add("x-hasura-admin-secret", `r5862n`)
	client := &http.Client{Timeout: time.Second * 100}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)
	var jsonData map[string]interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println("error:", err)
	}
	var service Service

	service.Id = jsonData["data"].(map[string]interface{})["services"].([]interface{})[0].(map[string]interface{})["id"].(string)
	service.Name = jsonData["data"].(map[string]interface{})["services"].([]interface{})[0].(map[string]interface{})["name"].(string)

	for i, model := range jsonData["data"].(map[string]interface{})["services"].([]interface{})[0].(map[string]interface{})["models"].([]interface{}) {
		service.Models = append(service.Models, &Model{
			Id:          model.(map[string]interface{})["id"].(string),
			Name:        model.(map[string]interface{})["name"].(string),
			NoOfFields:  int32(len(model.(map[string]interface{})["fields"].([]interface{}))),
			ServiceName: jsonData["data"].(map[string]interface{})["services"].([]interface{})[0].(map[string]interface{})["name"].(string),
		})

		for _, field := range model.(map[string]interface{})["fields"].([]interface{}) {
			if field.(map[string]interface{})["type"].(string) == "uuid" || field.(map[string]interface{})["type"].(string) == "hashed_string" {
				field.(map[string]interface{})["type"] = "string"
			}
			service.Models[i].Fields = append(service.Models[i].Fields, &Field{
				Id:        field.(map[string]interface{})["id"].(string),
				Name:      field.(map[string]interface{})["name"].(string),
				Type:      field.(map[string]interface{})["type"].(string),
				Default:   field.(map[string]interface{})["default"].(string),
				Key:       field.(map[string]interface{})["key"].(string),
				NullValue: field.(map[string]interface{})["null_value"].(string),
			})
		}

		for _, relationship := range model.(map[string]interface{})["relationships"].([]interface{}) {

			service.Models[i].Relationships = append(service.Models[i].Relationships, &Relationship{
				Id:                       relationship.(map[string]interface{})["id"].(string),
				Name:                     relationship.(map[string]interface{})["name"].(string),
				Type:                     relationship.(map[string]interface{})["type"].(string),
				RelationshipModelId:      relationship.(map[string]interface{})["relationship_model_id"].(string),
				RelationshipModelFieldId: relationship.(map[string]interface{})["relationship_model_field_id"].(string),
				CurrentModelId:           relationship.(map[string]interface{})["model_id"].(string),
				CurrentModelFieldId:      relationship.(map[string]interface{})["model_field_id"].(string),
			})
		}
	}
	funcMap := template.FuncMap{
		"Title": strings.Title,
		"addOne": func(n int) int {
			return n + 1
		},
		"subOne": func(n int32) int32 {
			return n - 1
		},
	}
	_, err = os.Stat("./temp/" + service.Id + "/proto")
	if os.IsNotExist(err) {
		errDir := os.MkdirAll("./temp/"+service.Id+"/proto", 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}

	_, err = os.Stat("./temp/" + service.Id + "/helpers")
	if os.IsNotExist(err) {
		errDir := os.MkdirAll("./temp/"+service.Id+"/helpers", 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}

	helpers.CopyDirectory("./google", "./temp/"+service.Id+"/google/")

	for _, model := range service.Models {
		_, err = os.Stat("./temp/" + service.Id + "/internal" + model.Name)
		if os.IsNotExist(err) {
			errDir := os.MkdirAll("./temp/"+service.Id+"/internal/"+model.Name, 0755)
			if errDir != nil {
				log.Fatal(err)
			}
		}
		protoTemplate, err := template.New("proto").Funcs(funcMap).Parse(templates.Proto)
		if err != nil {
			panic(err)
		}
		protoFile, err := os.Create("./temp/" + service.Id + "/proto/" + model.Name + ".proto")
		if err != nil {
			panic(err)
		}
		err = protoTemplate.Execute(protoFile, model)
		if err != nil {
			panic(err)
		}
		helpers.ExecuteCommand(`protoc -I=. -I=./temp/` + service.Id + `/proto --go_out=plugins=grpc:./temp/` + service.Id + `/internal/` + model.Name + ` --go_opt=paths=source_relative ` + model.Name + `.proto`)
		helpers.ExecuteCommand(`protoc -I=. -I=./temp/` + service.Id + `/proto --grpc-gateway_out=logtostderr=true,paths=source_relative:./temp/` + service.Id + `/internal/` + model.Name + ` ` + model.Name + `.proto`)
		helpers.ExecuteCommand(`protoc -I=. -I=./temp/` + service.Id + `/proto --swagger_out=logtostderr=true:./temp/` + service.Id + `/internal/` + model.Name + ` ` + model.Name + `.proto`)
		serviceTemplate, err := template.New("service").Funcs(funcMap).Parse(templates.Service)
		if err != nil {
			panic(err)
		}
		serviceFile, err := os.Create("./temp/" + service.Id + "/internal/" + model.Name + "/" + model.Name + ".go")
		if err != nil {
			panic(err)
		}
		err = serviceTemplate.Execute(serviceFile, model)
		if err != nil {
			panic(err)
		}
	}

	serverTemplate, err := template.New("service").Funcs(funcMap).Parse(templates.ServerFile)
	if err != nil {
		panic(err)
	}
	serverFile, err := os.Create("./temp/" + service.Id + "/main.go")
	if err != nil {
		panic(err)
	}
	err = serverTemplate.Execute(serverFile, &service)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("./temp/"+service.Id+"/DockerFile", []byte(templates.DockerFile), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	err = ioutil.WriteFile("./temp/"+service.Id+"/helpers/postgres.go", []byte(templates.Postgres), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	helpers.ExecuteCommandInDirectory("go mod init "+service.Name, "./temp/"+service.Id)
	return &GenerateServiceCodeResponse{Success: "generated service"}, nil
}
