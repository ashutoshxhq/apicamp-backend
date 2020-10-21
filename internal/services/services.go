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

	jsonData := map[string]string{
		"query": `
            { 
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
        `,
	}
	jsonValue, _ := json.Marshal(jsonData)
	request, err := http.NewRequest("POST", "https://apicamp-graphql.herokuapp.com/v1/graphql", bytes.NewBuffer(jsonValue))
	request.Header.Add("x-hasura-admin-secret", `r5862n`)
	client := &http.Client{Timeout: time.Second * 100}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
	return &GenerateServiceCodeResponse{Success: "generated service"}, nil
}
