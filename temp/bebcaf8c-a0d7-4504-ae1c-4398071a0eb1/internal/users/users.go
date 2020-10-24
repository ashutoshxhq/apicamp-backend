package users

import (
	"encoding/json"

	"demoService/helpers"

	uuid "github.com/google/uuid"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct {
}

// GetSingleUsers returns single users
func (s *Server) GetSingleUsers(ctx context.Context, req *GetSingleUsersRequest) (*GetSingleUsersResponse, error) {
	filter := make(map[string]interface{})
	if req.Filter != nil {

		if req.Filter.Id != nil {
			if req.Filter.Id.Type != "" {
				filter["id"] = req.Filter.Id.Value
			}
		}
		if req.Filter.Name != nil {
			if req.Filter.Name.Type != "" {
				filter["name"] = req.Filter.Name.Value
			}
		}
		if req.Filter.Email != nil {
			if req.Filter.Email.Type != "" {
				filter["email"] = req.Filter.Email.Value
			}
		}
	}
	data, err := helpers.GetRecord("users", filter)
	if err != nil {
		return &GetSingleUsersResponse{Success: false, Error: &Error{Error: "record_not_found"}}, nil
	}
	var record Users

	record.Id = data["id"].(string)

	record.Name = data["name"].(string)

	record.Email = data["email"].(string)

	return &GetSingleUsersResponse{Success: true, Data: &record}, nil
}

// GetMultipleUsers fuctions returns list of all users by a specific filter
func (s *Server) GetMultipleUsers(ctx context.Context, req *GetMultipleUsersRequest) (*GetMultipleUsersResponse, error) {

	filter := make(map[string]interface{})
	if req.Filter != nil {
		if len(req.Ids) > 0 {
			ids := make(map[string]interface{})
			ids["$in"] = req.Ids
			filter["id"] = ids
		}
		if req.Filter.Id != nil {
			if req.Filter.Id.Type != "" {
				filter["id"] = req.Filter.Id.Value
			}
		}
		if req.Filter.Name != nil {
			if req.Filter.Name.Type != "" {
				filter["name"] = req.Filter.Name.Value
			}
		}
		if req.Filter.Email != nil {
			if req.Filter.Email.Type != "" {
				filter["email"] = req.Filter.Email.Value
			}
		}
	}

	var records []*Users
	data, err := helpers.GetRecords("users", filter)
	if err != nil {
		return &GetMultipleUsersResponse{Success: false, Error: &Error{Error: "records_not_found"}}, nil
	}
	for i := 0; i < len(data); i++ {
		var record Users

		record.Id = data[i]["id"].(string)

		record.Name = data[i]["name"].(string)

		record.Email = data[i]["email"].(string)

		records = append(records, &record)
	}
	return &GetMultipleUsersResponse{Success: true, Data: records}, nil
}

// CreateSingleUsers stores new users in database and returns id
func (s *Server) CreateSingleUsers(ctx context.Context, req *CreateSingleUsersRequest) (*CreateSingleUsersResponse, error) {
	req.Data.Id = uuid.New().String()

	err := helpers.InsertRecord("users", req.Data)
	if err != nil {
		return &CreateSingleUsersResponse{Success: false, Error: &Error{Error: "unable_to_insert_record"}}, nil
	}
	return &CreateSingleUsersResponse{Success: true, Id: req.Data.Id}, nil
}

// CreateMultipleUsers stores multiple users in database and returns ids
func (s *Server) CreateMultipleUsers(ctx context.Context, req *CreateMultipleUsersRequest) (*CreateMultipleUsersResponse, error) {
	var records []interface{}
	var insertedIDs []string
	for _, record := range req.Data {
		record.Id = uuid.New().String()
		insertedIDs = append(insertedIDs, record.Id)
		records = append(records, record)
	}
	err := helpers.InsertRecords("users", records)
	if err != nil {
		return &CreateMultipleUsersResponse{Success: false, Error: &Error{Error: "unable_to_insert_record"}}, nil
	}
	return &CreateMultipleUsersResponse{Success: true, Ids: insertedIDs}, nil
}

// UpdateSingleUsers updates a users and returns success state
func (s *Server) UpdateSingleUsers(ctx context.Context, req *UpdateSingleUsersRequest) (*UpdateSingleUsersResponse, error) {
	filter := make(map[string]interface{})
	if req.Filter != nil {

		if req.Filter.Id != nil {
			if req.Filter.Id.Type != "" {
				filter["id"] = req.Filter.Id.Value
			}
		}
		if req.Filter.Name != nil {
			if req.Filter.Name.Type != "" {
				filter["name"] = req.Filter.Name.Value
			}
		}
		if req.Filter.Email != nil {
			if req.Filter.Email.Type != "" {
				filter["email"] = req.Filter.Email.Value
			}
		}
	}

	update := make(map[string]interface{})
	jsonData, _ := json.Marshal(req.Data)
	json.Unmarshal(jsonData, &update)

	err := helpers.UpdateRecord("users", filter, update)
	if err != nil {
		return &UpdateSingleUsersResponse{Success: false, Error: &Error{Error: "unable_to_update_record"}}, nil
	}
	return &UpdateSingleUsersResponse{Success: true}, nil
}

// UpdateMultipleUsers updates multiple users and returns success state
func (s *Server) UpdateMultipleUsers(ctx context.Context, req *UpdateMultipleUsersRequest) (*UpdateMultipleUsersResponse, error) {
	filter := make(map[string]interface{})
	if req.Filter != nil {
		if len(req.Ids) > 0 {
			ids := make(map[string]interface{})
			ids["$in"] = req.Ids
			filter["id"] = ids
		}
		if req.Filter.Id != nil {
			if req.Filter.Id.Type != "" {
				filter["id"] = req.Filter.Id.Value
			}
		}
		if req.Filter.Name != nil {
			if req.Filter.Name.Type != "" {
				filter["name"] = req.Filter.Name.Value
			}
		}
		if req.Filter.Email != nil {
			if req.Filter.Email.Type != "" {
				filter["email"] = req.Filter.Email.Value
			}
		}
	}

	update := make(map[string]interface{})
	jsonData, _ := json.Marshal(req.Data)
	json.Unmarshal(jsonData, &update)

	err := helpers.UpdateRecords("users", filter, update)
	if err != nil {
		return &UpdateMultipleUsersResponse{Success: false, Error: &Error{Error: "unable_to_update_records"}}, nil
	}
	return &UpdateMultipleUsersResponse{Success: true}, nil
}

// DeleteSingleUsers deletes a users by id
func (s *Server) DeleteSingleUsers(ctx context.Context, req *DeleteSingleUsersRequest) (*DeleteSingleUsersResponse, error) {
	filter := make(map[string]interface{})
	if req.Filter != nil {

		if req.Filter.Id != nil {
			if req.Filter.Id.Type != "" {
				filter["id"] = req.Filter.Id.Value
			}
		}
		if req.Filter.Name != nil {
			if req.Filter.Name.Type != "" {
				filter["name"] = req.Filter.Name.Value
			}
		}
		if req.Filter.Email != nil {
			if req.Filter.Email.Type != "" {
				filter["email"] = req.Filter.Email.Value
			}
		}
	}

	err := helpers.DeleteRecord("users", filter)
	if err != nil {
		return &DeleteSingleUsersResponse{Success: false, Error: &Error{Error: "unable_to_delete_record"}}, nil
	}
	return &DeleteSingleUsersResponse{Success: true}, nil
}

// DeleteMultipleUsers deletes multiple users by ids or filter
func (s *Server) DeleteMultipleUsers(ctx context.Context, req *DeleteMultipleUsersRequest) (*DeleteMultipleUsersResponse, error) {
	filter := make(map[string]interface{})
	if req.Filter != nil {
		if len(req.Ids) > 0 {
			ids := make(map[string]interface{})
			ids["$in"] = req.Ids
			filter["id"] = ids
		}
		if req.Filter.Id != nil {
			if req.Filter.Id.Type != "" {
				filter["id"] = req.Filter.Id.Value
			}
		}
		if req.Filter.Name != nil {
			if req.Filter.Name.Type != "" {
				filter["name"] = req.Filter.Name.Value
			}
		}
		if req.Filter.Email != nil {
			if req.Filter.Email.Type != "" {
				filter["email"] = req.Filter.Email.Value
			}
		}
	}

	err := helpers.DeleteRecords("users", filter)
	if err != nil {
		return &DeleteMultipleUsersResponse{Success: false, Error: &Error{Error: "unable_to_delete_records"}}, nil
	}
	return &DeleteMultipleUsersResponse{Success: true}, nil
}
