package prospects

import (
	"encoding/json"

	"demoService/helpers"

	uuid "github.com/google/uuid"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct {
}

// GetSingleProspects returns single prospects
func (s *Server) GetSingleProspects(ctx context.Context, req *GetSingleProspectsRequest) (*GetSingleProspectsResponse, error) {
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
		if req.Filter.Status != nil {
			if req.Filter.Status.Type != "" {
				filter["status"] = req.Filter.Status.Value
			}
		}
		if req.Filter.UserId != nil {
			if req.Filter.UserId.Type != "" {
				filter["userId"] = req.Filter.UserId.Value
			}
		}
		if req.Filter.Phone != nil {
			if req.Filter.Phone.Type != "" {
				filter["phone"] = req.Filter.Phone.Value
			}
		}
	}
	data, err := helpers.GetRecord("prospects", filter)
	if err != nil {
		return &GetSingleProspectsResponse{Success: false, Error: &Error{Error: "record_not_found"}}, nil
	}
	var record Prospects

	record.Id = data["id"].(string)

	record.Name = data["name"].(string)

	record.Email = data["email"].(string)

	record.Status = data["status"].(string)

	record.UserId = data["userId"].(string)

	record.Phone = data["phone"].(string)

	return &GetSingleProspectsResponse{Success: true, Data: &record}, nil
}

// GetMultipleProspects fuctions returns list of all prospects by a specific filter
func (s *Server) GetMultipleProspects(ctx context.Context, req *GetMultipleProspectsRequest) (*GetMultipleProspectsResponse, error) {

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
		if req.Filter.Status != nil {
			if req.Filter.Status.Type != "" {
				filter["status"] = req.Filter.Status.Value
			}
		}
		if req.Filter.UserId != nil {
			if req.Filter.UserId.Type != "" {
				filter["userId"] = req.Filter.UserId.Value
			}
		}
		if req.Filter.Phone != nil {
			if req.Filter.Phone.Type != "" {
				filter["phone"] = req.Filter.Phone.Value
			}
		}
	}

	var records []*Prospects
	data, err := helpers.GetRecords("prospects", filter)
	if err != nil {
		return &GetMultipleProspectsResponse{Success: false, Error: &Error{Error: "records_not_found"}}, nil
	}
	for i := 0; i < len(data); i++ {
		var record Prospects

		record.Id = data[i]["id"].(string)

		record.Name = data[i]["name"].(string)

		record.Email = data[i]["email"].(string)

		record.Status = data[i]["status"].(string)

		record.UserId = data[i]["userId"].(string)

		record.Phone = data[i]["phone"].(string)

		records = append(records, &record)
	}
	return &GetMultipleProspectsResponse{Success: true, Data: records}, nil
}

// CreateSingleProspects stores new prospects in database and returns id
func (s *Server) CreateSingleProspects(ctx context.Context, req *CreateSingleProspectsRequest) (*CreateSingleProspectsResponse, error) {
	req.Data.Id = uuid.New().String()

	err := helpers.InsertRecord("prospects", req.Data)
	if err != nil {
		return &CreateSingleProspectsResponse{Success: false, Error: &Error{Error: "unable_to_insert_record"}}, nil
	}
	return &CreateSingleProspectsResponse{Success: true, Id: req.Data.Id}, nil
}

// CreateMultipleProspects stores multiple prospects in database and returns ids
func (s *Server) CreateMultipleProspects(ctx context.Context, req *CreateMultipleProspectsRequest) (*CreateMultipleProspectsResponse, error) {
	var records []interface{}
	var insertedIDs []string
	for _, record := range req.Data {
		record.Id = uuid.New().String()
		insertedIDs = append(insertedIDs, record.Id)
		records = append(records, record)
	}
	err := helpers.InsertRecords("prospects", records)
	if err != nil {
		return &CreateMultipleProspectsResponse{Success: false, Error: &Error{Error: "unable_to_insert_record"}}, nil
	}
	return &CreateMultipleProspectsResponse{Success: true, Ids: insertedIDs}, nil
}

// UpdateSingleProspects updates a prospects and returns success state
func (s *Server) UpdateSingleProspects(ctx context.Context, req *UpdateSingleProspectsRequest) (*UpdateSingleProspectsResponse, error) {
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
		if req.Filter.Status != nil {
			if req.Filter.Status.Type != "" {
				filter["status"] = req.Filter.Status.Value
			}
		}
		if req.Filter.UserId != nil {
			if req.Filter.UserId.Type != "" {
				filter["userId"] = req.Filter.UserId.Value
			}
		}
		if req.Filter.Phone != nil {
			if req.Filter.Phone.Type != "" {
				filter["phone"] = req.Filter.Phone.Value
			}
		}
	}

	update := make(map[string]interface{})
	jsonData, _ := json.Marshal(req.Data)
	json.Unmarshal(jsonData, &update)

	err := helpers.UpdateRecord("prospects", filter, update)
	if err != nil {
		return &UpdateSingleProspectsResponse{Success: false, Error: &Error{Error: "unable_to_update_record"}}, nil
	}
	return &UpdateSingleProspectsResponse{Success: true}, nil
}

// UpdateMultipleProspects updates multiple prospects and returns success state
func (s *Server) UpdateMultipleProspects(ctx context.Context, req *UpdateMultipleProspectsRequest) (*UpdateMultipleProspectsResponse, error) {
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
		if req.Filter.Status != nil {
			if req.Filter.Status.Type != "" {
				filter["status"] = req.Filter.Status.Value
			}
		}
		if req.Filter.UserId != nil {
			if req.Filter.UserId.Type != "" {
				filter["userId"] = req.Filter.UserId.Value
			}
		}
		if req.Filter.Phone != nil {
			if req.Filter.Phone.Type != "" {
				filter["phone"] = req.Filter.Phone.Value
			}
		}
	}

	update := make(map[string]interface{})
	jsonData, _ := json.Marshal(req.Data)
	json.Unmarshal(jsonData, &update)

	err := helpers.UpdateRecords("prospects", filter, update)
	if err != nil {
		return &UpdateMultipleProspectsResponse{Success: false, Error: &Error{Error: "unable_to_update_records"}}, nil
	}
	return &UpdateMultipleProspectsResponse{Success: true}, nil
}

// DeleteSingleProspects deletes a prospects by id
func (s *Server) DeleteSingleProspects(ctx context.Context, req *DeleteSingleProspectsRequest) (*DeleteSingleProspectsResponse, error) {
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
		if req.Filter.Status != nil {
			if req.Filter.Status.Type != "" {
				filter["status"] = req.Filter.Status.Value
			}
		}
		if req.Filter.UserId != nil {
			if req.Filter.UserId.Type != "" {
				filter["userId"] = req.Filter.UserId.Value
			}
		}
		if req.Filter.Phone != nil {
			if req.Filter.Phone.Type != "" {
				filter["phone"] = req.Filter.Phone.Value
			}
		}
	}

	err := helpers.DeleteRecord("prospects", filter)
	if err != nil {
		return &DeleteSingleProspectsResponse{Success: false, Error: &Error{Error: "unable_to_delete_record"}}, nil
	}
	return &DeleteSingleProspectsResponse{Success: true}, nil
}

// DeleteMultipleProspects deletes multiple prospects by ids or filter
func (s *Server) DeleteMultipleProspects(ctx context.Context, req *DeleteMultipleProspectsRequest) (*DeleteMultipleProspectsResponse, error) {
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
		if req.Filter.Status != nil {
			if req.Filter.Status.Type != "" {
				filter["status"] = req.Filter.Status.Value
			}
		}
		if req.Filter.UserId != nil {
			if req.Filter.UserId.Type != "" {
				filter["userId"] = req.Filter.UserId.Value
			}
		}
		if req.Filter.Phone != nil {
			if req.Filter.Phone.Type != "" {
				filter["phone"] = req.Filter.Phone.Value
			}
		}
	}

	err := helpers.DeleteRecords("prospects", filter)
	if err != nil {
		return &DeleteMultipleProspectsResponse{Success: false, Error: &Error{Error: "unable_to_delete_records"}}, nil
	}
	return &DeleteMultipleProspectsResponse{Success: true}, nil
}
