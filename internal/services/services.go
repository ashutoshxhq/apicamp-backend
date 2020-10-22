package services

import (
	"bytes"
	context "context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Server represents the gRPC server
type Server struct {
}

// GenerateServiceCode fuctions generates service code
func (s *Server) GenerateServiceCode(ctx context.Context, req *GenerateServiceCodeRequest) (*GenerateServiceCodeResponse, error) {

	query := map[string]string{
		"query": `
            { 
				services(where: {id: {_eq: "bebcaf8c-a0d7-4504-ae1c-4398071a0eb1"}}){
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

	// fmt.Println(len(jsonData["data"].(map[string]interface{})["services"].([]interface{})[0].(map[string]interface{})["models"].([]interface{})))
	service.Id = jsonData["data"].(map[string]interface{})["services"].([]interface{})[0].(map[string]interface{})["id"].(string)
	service.Name = jsonData["data"].(map[string]interface{})["services"].([]interface{})[0].(map[string]interface{})["name"].(string)

	for i, model := range jsonData["data"].(map[string]interface{})["services"].([]interface{})[0].(map[string]interface{})["models"].([]interface{}) {
		service.Models = append(service.Models, &Model{
			Id:   model.(map[string]interface{})["id"].(string),
			Name: model.(map[string]interface{})["name"].(string),
		})

		for _, field := range model.(map[string]interface{})["fields"].([]interface{}) {
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
	fmt.Print(len(service.Models[0].Fields))
	return &GenerateServiceCodeResponse{Success: "generated service"}, nil
}
