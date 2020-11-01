package prospects

import (
	"demoService/helpers"
	"strings"
	"strconv"
	"errors"
	uuid "github.com/google/uuid"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct {
}

// GetSingleProspects returns single prospects
func (s *Server) GetSingleProspects(ctx context.Context, req *GetSingleProspectsRequest) (*GetSingleProspectsResponse, error) {
	var where string
	var whereValues []interface{}
	var i int = 0
	for key, value := range req.Filters {
		if i != 0 {
			where = where + " AND"
		} else {
			where = " WHERE "
		}
		if len(strings.Split(key, "_")) > 1 && key != "_" {
			if strings.TrimSpace(strings.Split(key, "_")[0]) != "" {
				if strings.Split(key, "_")[1] == "eq" {
					where = where + " " + strings.Split(key, "_")[0] + " = " + strconv.Itoa(i+1)
				} else if strings.Split(key, "_")[1] == "gt" {
					where = where + " " + strings.Split(key, "_")[0] + " > " + strconv.Itoa(i+1)
				} else if strings.Split(key, "_")[1] == "lt" {
					where = where + " " + strings.Split(key, "_")[0] + " < " + strconv.Itoa(i+1)
				} else if strings.Split(key, "_")[1] == "gte" {
					where = where + " " + strings.Split(key, "_")[0] + " >= " + strconv.Itoa(i+1)
				} else if strings.Split(key, "_")[1] == "lte" {
					where = where + " " + strings.Split(key, "_")[0] + " <= " + strconv.Itoa(i+1)
				} else if strings.Split(key, "_")[1] == "ne" {
					where = where + " " + strings.Split(key, "_")[0] + " != " + strconv.Itoa(i+1)
				} else if strings.Split(key, "_")[1] == "like" {
					where = where + " " + strings.Split(key, "_")[0] + " LIKE " + strconv.Itoa(i+1)
				} else if strings.Split(key, "_")[1] == "in" {
					where = where + " " + strings.Split(key, "_")[0] + " IN " + strconv.Itoa(i+1)
				} else {
					return nil, errors.New("Error: Unable to parse the filter type")
				}
			} else {
				return nil, errors.New("Error: Unable to parse filter field name")
			}
		} else if key != "_" {
			where = where + " " + strings.Split(key, "_")[0] + " = " + strconv.Itoa(i+1)
		} else {
			return nil, errors.New("Error: Unable to parse filter field name")
		}
		whereValues = append(whereValues, value.AsInterface())
		i++
	}
	conn, err := helpers.DatabasePool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	rows, err := conn.Query(context.Background(), "SELECT * FROM prospects"+where+" LIMIT 1",whereValues...)
	if err != nil {
		return nil, err
	}
	var record Prospects
	for rows.Next() {
		err := rows.Scan( &record.Name,   &record.Email,   &record.Status,   &record.UserId,   &record.Phone,   &record.Id )
		if err != nil {
			return nil, err
		}
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return &GetSingleProspectsResponse{Data: &record}, nil
}

// GetMultipleProspects fuctions returns list of all prospects by a specific filter
func (s *Server) GetMultipleProspects(ctx context.Context, req *GetMultipleProspectsRequest) (*GetMultipleProspectsResponse, error) {
	var where string
	var whereValues []interface{}
	var i int = 0
	for key, value := range req.Filters {
		if i != 0 {
			where = where + " AND"
		} else {
			where = " WHERE "
		}
		if len(strings.Split(key, "_")) > 1 && key != "_" {
			if strings.TrimSpace(strings.Split(key, "_")[0]) != "" {
				if strings.Split(key, "_")[1] == "eq" {
					where = where + " " + strings.Split(key, "_")[0] + " = " + strconv.Itoa(i+1)
				} else if strings.Split(key, "_")[1] == "gt" {
					where = where + " " + strings.Split(key, "_")[0] + " > " + strconv.Itoa(i+1)
				} else if strings.Split(key, "_")[1] == "lt" {
					where = where + " " + strings.Split(key, "_")[0] + " < " + strconv.Itoa(i+1)
				} else if strings.Split(key, "_")[1] == "gte" {
					where = where + " " + strings.Split(key, "_")[0] + " >= " + strconv.Itoa(i+1)
				} else if strings.Split(key, "_")[1] == "lte" {
					where = where + " " + strings.Split(key, "_")[0] + " <= " + strconv.Itoa(i+1)
				} else if strings.Split(key, "_")[1] == "ne" {
					where = where + " " + strings.Split(key, "_")[0] + " != " + strconv.Itoa(i+1)
				} else if strings.Split(key, "_")[1] == "like" {
					where = where + " " + strings.Split(key, "_")[0] + " LIKE " + strconv.Itoa(i+1)
				} else if strings.Split(key, "_")[1] == "in" {
					where = where + " " + strings.Split(key, "_")[0] + " IN " + strconv.Itoa(i+1)
				} else {
					return nil, errors.New("Error: Unable to parse the filter type")
				}
			} else {
				return nil, errors.New("Error: Unable to parse filter field name")
			}
		} else if key != "_" {
			where = where + " " + strings.Split(key, "_")[0] + " = " + strconv.Itoa(i+1)
		} else {
			return nil, errors.New("Error: Unable to parse filter field name")
		}
		whereValues = append(whereValues, value.AsInterface())
		i++
	}
	conn, err := helpers.DatabasePool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	rows, err := conn.Query(context.Background(), "SELECT * FROM prospects"+where+" LIMIT 100",whereValues...)
	if err != nil {
		return nil, err
	}
	var records []*Prospects
	for rows.Next() {
		var record Prospects

		err := rows.Scan( &record.Name,   &record.Email,   &record.Status,   &record.UserId,   &record.Phone,   &record.Id )
		if err != nil {
			return nil, err
		}
		records = append(records, &record)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return &GetMultipleProspectsResponse{Data: records}, nil
}

// CreateSingleProspects stores new prospects in database and returns id
func (s *Server) CreateSingleProspects(ctx context.Context, req *CreateSingleProspectsRequest) (*CreateSingleProspectsResponse, error) {
	req.Data.Id = uuid.New().String()
	
	conn, err := helpers.DatabasePool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	_, err = conn.Exec(context.Background(), "INSERT INTO prospects (name, email, status, userId, phone, id ) VALUES($1, $2, $3, $4, $5, $6 )", req.Data.Name, req.Data.Email, req.Data.Status, req.Data.UserId, req.Data.Phone, req.Data.Id  )
	if err != nil {
		return nil, err
	}
	
	return &CreateSingleProspectsResponse{}, nil
}

// CreateMultipleProspects stores multiple prospects in database and returns ids
func (s *Server) CreateMultipleProspects(ctx context.Context, req *CreateMultipleProspectsRequest) (*CreateMultipleProspectsResponse, error) {
	var values string
	for i, record := range req.Data {
		if i == 0{
		values = values +"("
		} else{
			values = values +", ("
		}
		
		values = values+ record.Name+","
		values = values+ record.Email+","
		values = values+ record.Status+","
		values = values+ record.UserId+","
		values = values+ record.Phone+","
		req.Data[i].Id = uuid.New().String()
		values = values+ record.Id
	
		values = values +")"
	}	
	conn, err := helpers.DatabasePool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	_, err = conn.Exec(context.Background(), "INSERT INTO prospects (name, email, status, userId, phone, id ) VALUES "+values)
	if err != nil {
		return nil, err
	}
	
	return &CreateMultipleProspectsResponse{}, nil
}

// UpdateProspects updates a prospects and returns success state
func (s *Server) UpdateProspects(ctx context.Context, req *UpdateProspectsRequest) (*UpdateProspectsResponse, error) {
	var where string
	var whereValues []interface{}
	var setString string
	var setValues []interface{}
	var i int = 0
	var j int = 0
	for key, value := range req.Data {
		if i != 0 {
			setString = setString + ", "
		} else {
			setString = " SET "
		}

		if strings.TrimSpace(key) != "" {
			setString = setString + key + " = $" + strconv.Itoa(i+1)
			setValues = append(setValues, value.AsInterface())
		} else {
			return nil, errors.New("Error: Unable to parse data fieldname")
		}
		i++
	}

	for key, value := range req.Filters {
		if j != 0 {
			where = where + " AND"
		} else {
			where = " WHERE "
		}
		if len(strings.Split(key, "_")) > 1 && key != "_" {
			if strings.TrimSpace(strings.Split(key, "_")[0]) != "" {
				if strings.Split(key, "_")[1] == "eq" {
					where = where + " " + strings.Split(key, "_")[0] + " = $" + strconv.Itoa(i+1)
				} else if strings.Split(key, "_")[1] == "gt" {
					where = where + " " + strings.Split(key, "_")[0] + " > $" + strconv.Itoa(i+1)
				} else if strings.Split(key, "_")[1] == "lt" {
					where = where + " " + strings.Split(key, "_")[0] + " < $" + strconv.Itoa(i+1)
				} else if strings.Split(key, "_")[1] == "gte" {
					where = where + " " + strings.Split(key, "_")[0] + " >= $" + strconv.Itoa(i+1)
				} else if strings.Split(key, "_")[1] == "lte" {
					where = where + " " + strings.Split(key, "_")[0] + " <= $" + strconv.Itoa(i+1)
				} else if strings.Split(key, "_")[1] == "ne" {
					where = where + " " + strings.Split(key, "_")[0] + " != $" + strconv.Itoa(i+1)
				} else if strings.Split(key, "_")[1] == "like" {
					where = where + " " + strings.Split(key, "_")[0] + " LIKE $" + strconv.Itoa(i+1)
				} else if strings.Split(key, "_")[1] == "in" {
					where = where + " " + strings.Split(key, "_")[0] + " IN $" + strconv.Itoa(i+1)
				} else {
					return nil, errors.New("Error: Unable to parse the filter type")
				}
			} else {
				return nil, errors.New("Error: Unable to parse filter field name")
			}
		} else if key != "_" {
			where = where + " " + strings.Split(key, "_")[0] + " = " + value.GetStringValue()
		} else {
			return nil, errors.New("Error: Unable to parse filter field name")
		}
		whereValues = append(whereValues, value.AsInterface())
		i++
		j++
	}
	var sql string = "UPDATE prospects " + setString + where
	conn, err := helpers.DatabasePool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	_, err = conn.Exec(context.Background(), sql, append(setValues, whereValues...)...)
	if err != nil {
		return nil, err
	}
	return &UpdateProspectsResponse{}, nil
}

// DeleteProspects deletes multiple prospects by ids or filter
func (s *Server) DeleteProspects(ctx context.Context, req *DeleteProspectsRequest) (*DeleteProspectsResponse, error) {
	var where string
	var whereValues []interface{}
	var i int = 0
	for key, value := range req.Filters {
		if i != 0 {
			where = where + " AND"
		} else {
			where = " WHERE "
		}
		if len(strings.Split(key, "_")) > 1 && key != "_" {
			if strings.TrimSpace(strings.Split(key, "_")[0]) != "" {
				if strings.Split(key, "_")[1] == "eq" {
					where = where + " " + strings.Split(key, "_")[0] + " = " + strconv.Itoa(i+1)
				} else if strings.Split(key, "_")[1] == "gt" {
					where = where + " " + strings.Split(key, "_")[0] + " > " + strconv.Itoa(i+1)
				} else if strings.Split(key, "_")[1] == "lt" {
					where = where + " " + strings.Split(key, "_")[0] + " < " + strconv.Itoa(i+1)
				} else if strings.Split(key, "_")[1] == "gte" {
					where = where + " " + strings.Split(key, "_")[0] + " >= " + strconv.Itoa(i+1)
				} else if strings.Split(key, "_")[1] == "lte" {
					where = where + " " + strings.Split(key, "_")[0] + " <= " + strconv.Itoa(i+1)
				} else if strings.Split(key, "_")[1] == "ne" {
					where = where + " " + strings.Split(key, "_")[0] + " != " + strconv.Itoa(i+1)
				} else if strings.Split(key, "_")[1] == "like" {
					where = where + " " + strings.Split(key, "_")[0] + " LIKE " + strconv.Itoa(i+1)
				} else if strings.Split(key, "_")[1] == "in" {
					where = where + " " + strings.Split(key, "_")[0] + " IN " + strconv.Itoa(i+1)
				} else {
					return nil, errors.New("Error: Unable to parse the filter type")
				}
			} else {
				return nil, errors.New("Error: Unable to parse filter field name")
			}
		} else if key != "_" {
			where = where + " " + strings.Split(key, "_")[0] + " = " + strconv.Itoa(i+1)
		} else {
			return nil, errors.New("Error: Unable to parse filter field name")
		}
		whereValues = append(whereValues, value.AsInterface())
		i++
	}
	conn, err := helpers.DatabasePool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	var sql string = "DELETE FROM prospects " + where
	_, err = conn.Exec(context.Background(), sql, whereValues...)
	if err != nil {
		return nil, err
	}
	return &DeleteProspectsResponse{}, nil
}

