package users

import (
	"demoService/helpers"
	uuid "github.com/google/uuid"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct {
}

// GetSingleUsers returns single users
func (s *Server) GetSingleUsers(ctx context.Context, req *GetSingleUsersRequest) (*GetSingleUsersResponse, error) {
	var where string
	for i, filter := range req.Filters {
		if i != 0 {
			where = where + " AND"
		} else {
			where = " WHERE "
		}
		if filter.Type == "eq" { 
			where  = where+" "+filter.Field+ " = "+ filter.Value
		} else if filter.Type == "gt" {
			where  = where+" "+filter.Field+ " > "+ filter.Value
		} else if filter.Type == "lt" {
			where  = where+" "+filter.Field+ " < "+ filter.Value
		} else if filter.Type == "gte" {
			where  = where+" "+filter.Field+ " >= "+ filter.Value
		} else if filter.Type == "lte" {
			where  = where+" "+filter.Field+ " <= "+ filter.Value
		} else if filter.Type == "ne" {
			where  = where+" "+filter.Field+ " != "+ filter.Value
		} else if filter.Type == "like" {
			where  = where+" "+filter.Field+ " LIKE "+ filter.Value
		} else if filter.Type == "in" {
			where  = where+" "+filter.Field+ " IN "+ filter.Value
		}
	}
	conn, err := helpers.DatabasePool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	rows, err := conn.Query(context.Background(), "SELECT * FROM users"+where+" LIMIT 1")
	if err != nil {
		return nil, err
	}
	var record Users
	for rows.Next() {
		err := rows.Scan( &record.Id,   &record.Name,   &record.Email )
		if err != nil {
			return nil, err
		}
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return &GetSingleUsersResponse{Success: true, Data: &record}, nil
}

// GetMultipleUsers fuctions returns list of all users by a specific filter
func (s *Server) GetMultipleUsers(ctx context.Context, req *GetMultipleUsersRequest) (*GetMultipleUsersResponse, error) {

	var where string
	for i, filter := range req.Filters {
		if i != 0 {
			where = where + " AND"
		}
		if filter.Type == "eq" { 
			where  = where+" "+filter.Field+ " = "+ filter.Value
		} else if filter.Type == "gt" {
			where  = where+" "+filter.Field+ " > "+ filter.Value
		} else if filter.Type == "lt" {
			where  = where+" "+filter.Field+ " < "+ filter.Value
		} else if filter.Type == "gte" {
			where  = where+" "+filter.Field+ " >= "+ filter.Value
		} else if filter.Type == "lte" {
			where  = where+" "+filter.Field+ " <= "+ filter.Value
		} else if filter.Type == "ne" {
			where  = where+" "+filter.Field+ " != "+ filter.Value
		} else if filter.Type == "like" {
			where  = where+" "+filter.Field+ " LIKE "+ filter.Value
		} else if filter.Type == "in" {
			where  = where+" "+filter.Field+ " IN "+ filter.Value
		}
	}

	
	return &GetMultipleUsersResponse{Success: true}, nil
}

// CreateSingleUsers stores new users in database and returns id
func (s *Server) CreateSingleUsers(ctx context.Context, req *CreateSingleUsersRequest) (*CreateSingleUsersResponse, error) {
	// req.Data.Id = uuid.New().String()
	
	
	return &CreateSingleUsersResponse{Success: true}, nil
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
	
	return &CreateMultipleUsersResponse{Success: true}, nil
}

// UpdateSingleUsers updates a users and returns success state
func (s *Server) UpdateSingleUsers(ctx context.Context, req *UpdateSingleUsersRequest) (*UpdateSingleUsersResponse, error) {
	var where string
	for i, filter := range req.Filters {
		if i != 0 {
			where = where + " AND"
		}
		if filter.Type == "eq" { 
			where  = where+" "+filter.Field+ " = "+ filter.Value
		} else if filter.Type == "gt" {
			where  = where+" "+filter.Field+ " > "+ filter.Value
		} else if filter.Type == "lt" {
			where  = where+" "+filter.Field+ " < "+ filter.Value
		} else if filter.Type == "gte" {
			where  = where+" "+filter.Field+ " >= "+ filter.Value
		} else if filter.Type == "lte" {
			where  = where+" "+filter.Field+ " <= "+ filter.Value
		} else if filter.Type == "ne" {
			where  = where+" "+filter.Field+ " != "+ filter.Value
		} else if filter.Type == "like" {
			where  = where+" "+filter.Field+ " LIKE "+ filter.Value
		} else if filter.Type == "in" {
			where  = where+" "+filter.Field+ " IN "+ filter.Value
		}
	}

	
	return &UpdateSingleUsersResponse{Success: true}, nil
}

// UpdateMultipleUsers updates multiple users and returns success state
func (s *Server) UpdateMultipleUsers(ctx context.Context, req *UpdateMultipleUsersRequest) (*UpdateMultipleUsersResponse, error) {
	var where string
	for i, filter := range req.Filters {
		if i != 0 {
			where = where + " AND"
		}
		if filter.Type == "eq" { 
			where  = where+" "+filter.Field+ " = "+ filter.Value
		} else if filter.Type == "gt" {
			where  = where+" "+filter.Field+ " > "+ filter.Value
		} else if filter.Type == "lt" {
			where  = where+" "+filter.Field+ " < "+ filter.Value
		} else if filter.Type == "gte" {
			where  = where+" "+filter.Field+ " >= "+ filter.Value
		} else if filter.Type == "lte" {
			where  = where+" "+filter.Field+ " <= "+ filter.Value
		} else if filter.Type == "ne" {
			where  = where+" "+filter.Field+ " != "+ filter.Value
		} else if filter.Type == "like" {
			where  = where+" "+filter.Field+ " LIKE "+ filter.Value
		} else if filter.Type == "in" {
			where  = where+" "+filter.Field+ " IN "+ filter.Value
		}
	}

	
	return &UpdateMultipleUsersResponse{Success: true}, nil
}

// DeleteSingleUsers deletes a users by id
func (s *Server) DeleteSingleUsers(ctx context.Context, req *DeleteSingleUsersRequest) (*DeleteSingleUsersResponse, error) {
	var where string
	for i, filter := range req.Filters {
		if i != 0 {
			where = where + " AND"
		}
		if filter.Type == "eq" { 
			where  = where+" "+filter.Field+ " = "+ filter.Value
		} else if filter.Type == "gt" {
			where  = where+" "+filter.Field+ " > "+ filter.Value
		} else if filter.Type == "lt" {
			where  = where+" "+filter.Field+ " < "+ filter.Value
		} else if filter.Type == "gte" {
			where  = where+" "+filter.Field+ " >= "+ filter.Value
		} else if filter.Type == "lte" {
			where  = where+" "+filter.Field+ " <= "+ filter.Value
		} else if filter.Type == "ne" {
			where  = where+" "+filter.Field+ " != "+ filter.Value
		} else if filter.Type == "like" {
			where  = where+" "+filter.Field+ " LIKE "+ filter.Value
		} else if filter.Type == "in" {
			where  = where+" "+filter.Field+ " IN "+ filter.Value
		}
	}
	
	return &DeleteSingleUsersResponse{Success: true}, nil
}

// DeleteMultipleUsers deletes multiple users by ids or filter
func (s *Server) DeleteMultipleUsers(ctx context.Context, req *DeleteMultipleUsersRequest) (*DeleteMultipleUsersResponse, error) {
	var where string
	for i, filter := range req.Filters {
		if i != 0 {
			where = where + " AND"
		}
		if filter.Type == "eq" { 
			where  = where+" "+filter.Field+ " = "+ filter.Value
		} else if filter.Type == "gt" {
			where  = where+" "+filter.Field+ " > "+ filter.Value
		} else if filter.Type == "lt" {
			where  = where+" "+filter.Field+ " < "+ filter.Value
		} else if filter.Type == "gte" {
			where  = where+" "+filter.Field+ " >= "+ filter.Value
		} else if filter.Type == "lte" {
			where  = where+" "+filter.Field+ " <= "+ filter.Value
		} else if filter.Type == "ne" {
			where  = where+" "+filter.Field+ " != "+ filter.Value
		} else if filter.Type == "like" {
			where  = where+" "+filter.Field+ " LIKE "+ filter.Value
		} else if filter.Type == "in" {
			where  = where+" "+filter.Field+ " IN "+ filter.Value
		}
	}
	
	return &DeleteMultipleUsersResponse{Success: true}, nil
}

