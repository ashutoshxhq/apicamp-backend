package users

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

// GetSingleUsers returns single users
func (s *Server) GetSingleUsers(ctx context.Context, req *GetSingleUsersRequest) (*GetSingleUsersResponse, error) {
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
	rows, err := conn.Query(context.Background(), "SELECT * FROM users"+where+" LIMIT 1",whereValues...)
	if err != nil {
		return nil, err
	}
	var record Users
	for rows.Next() {
		err := rows.Scan( &record.Name,   &record.Email,   &record.Id )
		if err != nil {
			return nil, err
		}
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return &GetSingleUsersResponse{Data: &record}, nil
}

// GetMultipleUsers fuctions returns list of all users by a specific filter
func (s *Server) GetMultipleUsers(ctx context.Context, req *GetMultipleUsersRequest) (*GetMultipleUsersResponse, error) {
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
	rows, err := conn.Query(context.Background(), "SELECT * FROM users"+where+" LIMIT 100",whereValues...)
	if err != nil {
		return nil, err
	}
	var records []*Users
	for rows.Next() {
		var record Users

		err := rows.Scan( &record.Name,   &record.Email,   &record.Id )
		if err != nil {
			return nil, err
		}
		records = append(records, &record)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return &GetMultipleUsersResponse{Data: records}, nil
}

// CreateSingleUsers stores new users in database and returns id
func (s *Server) CreateSingleUsers(ctx context.Context, req *CreateSingleUsersRequest) (*CreateSingleUsersResponse, error) {
	
	
	conn, err := helpers.DatabasePool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	_, err = conn.Exec(context.Background(), "INSERT INTO users (name, email, id ) VALUES($1, $2, $3 )", req.Data.Name, req.Data.Email, req.Data.Id  )
	if err != nil {
		return nil, err
	}
	
	return &CreateSingleUsersResponse{}, nil
}

// CreateMultipleUsers stores multiple users in database and returns ids
func (s *Server) CreateMultipleUsers(ctx context.Context, req *CreateMultipleUsersRequest) (*CreateMultipleUsersResponse, error) {
	var values string
	for i, record := range req.Data {
		if i == 0{
		values = values +"("
		} else{
			values = values +", ("
		}
		
		values = values+ record.Name+","
		values = values+ record.Email+","
		values = values+ record.Id
	
		values = values +")"
	}	
	conn, err := helpers.DatabasePool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	_, err = conn.Exec(context.Background(), "INSERT INTO users (name, email, id ) VALUES "+values)
	if err != nil {
		return nil, err
	}
	
	return &CreateMultipleUsersResponse{}, nil
}

// UpdateUsers updates a users and returns success state
func (s *Server) UpdateUsers(ctx context.Context, req *UpdateUsersRequest) (*UpdateUsersResponse, error) {
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
	var sql string = "UPDATE users " + setString + where
	conn, err := helpers.DatabasePool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	_, err = conn.Exec(context.Background(), sql, append(setValues, whereValues...)...)
	if err != nil {
		return nil, err
	}
	return &UpdateUsersResponse{}, nil
}

// DeleteUsers deletes multiple users by ids or filter
func (s *Server) DeleteUsers(ctx context.Context, req *DeleteUsersRequest) (*DeleteUsersResponse, error) {
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
	var sql string = "DELETE FROM users " + where
	_, err = conn.Exec(context.Background(), sql, whereValues...)
	if err != nil {
		return nil, err
	}
	return &DeleteUsersResponse{}, nil
}

