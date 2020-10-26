package prospects

import (
	"demoService/helpers"
	uuid "github.com/google/uuid"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct {
}

// GetSingleProspects returns single prospects
func (s *Server) GetSingleProspects(ctx context.Context, req *GetSingleProspectsRequest) (*GetSingleProspectsResponse, error) {
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
	rows, err := conn.Query(context.Background(), "SELECT * FROM prospects"+where+" LIMIT 1")
	if err != nil {
		return nil, err
	}
	var record Prospects
	for rows.Next() {
		err := rows.Scan( &record.Id,   &record.Name,   &record.Email,   &record.Status,   &record.UserId,   &record.Phone )
		if err != nil {
			return nil, err
		}
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return &GetSingleProspectsResponse{Success: true, Data: &record}, nil
}

// GetMultipleProspects fuctions returns list of all prospects by a specific filter
func (s *Server) GetMultipleProspects(ctx context.Context, req *GetMultipleProspectsRequest) (*GetMultipleProspectsResponse, error) {

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

	
	return &GetMultipleProspectsResponse{Success: true}, nil
}

// CreateSingleProspects stores new prospects in database and returns id
func (s *Server) CreateSingleProspects(ctx context.Context, req *CreateSingleProspectsRequest) (*CreateSingleProspectsResponse, error) {
	// req.Data.Id = uuid.New().String()
	
	
	return &CreateSingleProspectsResponse{Success: true}, nil
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
	
	return &CreateMultipleProspectsResponse{Success: true}, nil
}

// UpdateSingleProspects updates a prospects and returns success state
func (s *Server) UpdateSingleProspects(ctx context.Context, req *UpdateSingleProspectsRequest) (*UpdateSingleProspectsResponse, error) {
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

	
	return &UpdateSingleProspectsResponse{Success: true}, nil
}

// UpdateMultipleProspects updates multiple prospects and returns success state
func (s *Server) UpdateMultipleProspects(ctx context.Context, req *UpdateMultipleProspectsRequest) (*UpdateMultipleProspectsResponse, error) {
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

	
	return &UpdateMultipleProspectsResponse{Success: true}, nil
}

// DeleteSingleProspects deletes a prospects by id
func (s *Server) DeleteSingleProspects(ctx context.Context, req *DeleteSingleProspectsRequest) (*DeleteSingleProspectsResponse, error) {
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
	
	return &DeleteSingleProspectsResponse{Success: true}, nil
}

// DeleteMultipleProspects deletes multiple prospects by ids or filter
func (s *Server) DeleteMultipleProspects(ctx context.Context, req *DeleteMultipleProspectsRequest) (*DeleteMultipleProspectsResponse, error) {
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
	
	return &DeleteMultipleProspectsResponse{Success: true}, nil
}

