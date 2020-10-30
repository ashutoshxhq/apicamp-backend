package templates

//Service is template for service file generation
var Service string = `package {{.Name}}

import (
	"{{.ServiceName}}/helpers"
	"strings"
	"strconv"
	"errors"
	uuid "github.com/google/uuid"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct {
}

// GetSingle{{.Name | Title}} returns single {{.Name}}
func (s *Server) GetSingle{{.Name | Title}}(ctx context.Context, req *GetSingle{{.Name | Title}}Request) (*GetSingle{{.Name | Title}}Response, error) {
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
	rows, err := conn.Query(context.Background(), "SELECT * FROM {{.Name}}"+where+" LIMIT 1",whereValues...)
	if err != nil {
		return nil, err
	}
	var record {{.Name | Title}}
	for rows.Next() {
		err := rows.Scan({{range $index,$field := .Fields}} &record.{{$field.Name  | Title}}{{if (eq $index ($.NoOfFields | subOne))}}{{ else }}, {{end}} {{end}})
		if err != nil {
			return nil, err
		}
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return &GetSingle{{.Name | Title}}Response{Data: &record}, nil
}

// GetMultiple{{.Name | Title}} fuctions returns list of all {{.Name}} by a specific filter
func (s *Server) GetMultiple{{.Name | Title}}(ctx context.Context, req *GetMultiple{{.Name | Title}}Request) (*GetMultiple{{.Name | Title}}Response, error) {
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
	rows, err := conn.Query(context.Background(), "SELECT * FROM {{.Name}}"+where+" LIMIT 100",whereValues...)
	if err != nil {
		return nil, err
	}
	var records []*{{.Name | Title}}
	for rows.Next() {
		var record {{.Name | Title}}

		err := rows.Scan({{range $index,$field := .Fields}} &record.{{$field.Name  | Title}}{{if (eq $index ($.NoOfFields | subOne))}}{{ else }}, {{end}} {{end}})
		if err != nil {
			return nil, err
		}
		records = append(records, &record)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return &GetMultiple{{.Name | Title}}Response{Data: records}, nil
}

// CreateSingle{{.Name | Title}} stores new {{.Name}} in database and returns id
func (s *Server) CreateSingle{{.Name | Title}}(ctx context.Context, req *CreateSingle{{.Name | Title}}Request) (*CreateSingle{{.Name | Title}}Response, error) {
	{{range $index,$field := .Fields}}{{if (eq $field.Default "RANDOM_UUID")}}req.Data.{{$field.Name | Title}} = uuid.New().String(){{end}}{{end}}
	
	conn, err := helpers.DatabasePool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	_, err = conn.Exec(context.Background(), "INSERT INTO {{.Name}} ({{range $index,$field := .Fields}}{{$field.Name}}{{if (eq $index ($.NoOfFields | subOne))}}{{ else }},{{end}} {{end}}) VALUES({{range $index,$field := .Fields}}${{$index | addOne}}{{if (eq $index ($.NoOfFields | subOne))}}{{ else }},{{end}} {{end}})", {{range $index,$field := .Fields}}req.Data.{{$field.Name | Title}}{{if (eq $index ($.NoOfFields | subOne))}}{{ else }},{{end}} {{end}} )
	if err != nil {
		return nil, err
	}
	
	return &CreateSingle{{.Name | Title}}Response{}, nil
}

// CreateMultiple{{.Name | Title}} stores multiple {{.Name}} in database and returns ids
func (s *Server) CreateMultiple{{.Name | Title}}(ctx context.Context, req *CreateMultiple{{.Name | Title}}Request) (*CreateMultiple{{.Name | Title}}Response, error) {
	var values string
	for i, record := range req.Data {
		if i == 0{
		values = values +"("
		} else{
			values = values +", ("
		}
		{{range $index,$field := .Fields}}{{if (eq $field.Default "RANDOM_UUID")}}
		req.Data[i].{{$field.Name | Title}} = uuid.New().String(){{end}}
		values = values+ record.{{$field.Name | Title}}{{if (eq $index ($.NoOfFields | subOne))}}{{ else }}+","{{end}}{{end}}
	
		values = values +")"
	}	
	conn, err := helpers.DatabasePool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	_, err = conn.Exec(context.Background(), "INSERT INTO {{.Name}} ({{range $index,$field := .Fields}}{{$field.Name}}{{if (eq $index ($.NoOfFields | subOne))}}{{ else }},{{end}} {{end}}) VALUES "+values)
	if err != nil {
		return nil, err
	}
	
	return &CreateMultiple{{.Name | Title}}Response{}, nil
}

// Update{{.Name | Title}} updates a {{.Name}} and returns success state
func (s *Server) Update{{.Name | Title}}(ctx context.Context, req *Update{{.Name | Title}}Request) (*Update{{.Name | Title}}Response, error) {
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
	var sql string = "UPDATE {{.Name}} " + setString + where
	conn, err := helpers.DatabasePool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	_, err = conn.Exec(context.Background(), sql, append(setValues, whereValues...)...)
	if err != nil {
		return nil, err
	}
	return &Update{{.Name | Title}}Response{}, nil
}

// Delete{{.Name | Title}} deletes multiple {{.Name}} by ids or filter
func (s *Server) Delete{{.Name | Title}}(ctx context.Context, req *Delete{{.Name | Title}}Request) (*Delete{{.Name | Title}}Response, error) {
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
	var sql string = "DELETE FROM {{.Name}} " + where
	_, err = conn.Exec(context.Background(), sql, whereValues...)
	if err != nil {
		return nil, err
	}
	return &Delete{{.Name | Title}}Response{}, nil
}

`

//UserService is template for service file generation
var UserService string = `package {{.Name}}

// Do your Implementations here...
`

//Postgres ...
var Postgres string = `package helpers

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	// DatabasePool is global database pool
	DatabasePool *pgxpool.Pool
)

//InitializeDatabase ...
func InitializeDatabase() {
	dbpool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	DatabasePool = dbpool
}
`

//DockerFile is template for DockerFile generation
var DockerFile string = `
FROM golang:alpine as builder
RUN mkdir /build 
WORKDIR /build
COPY go.mod ./
COPY . .
RUN go mod tidy
RUN go build -o main .
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
WORKDIR /app
EXPOSE 8000
EXPOSE 9000
CMD ["./main"]
`

//ServerFile ...
var ServerFile string = `package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"
	"{{.Name}}/helpers"
	{{range $index,$model := .Models}}"{{.ServiceName}}/internal/{{$model.Name}}"
	{{end}} 
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
)

func startGRPC(wg *sync.WaitGroup) {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	grpcServer := grpc.NewServer()
	{{range $index,$model := .Models}}
	{{$model.Name}}Server := {{$model.Name}}.Server{}
	{{$model.Name}}.Register{{$model.Name | Title}}ServiceServer(grpcServer, &{{$model.Name}}Server)
	{{end}} 
	log.Println("gRPC server ready...")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	wg.Done()
}

func startHTTP(wg *sync.WaitGroup) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Connect to the GRPC server
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	// Register grpc-gateway
	rmux := runtime.NewServeMux()
	{{range $index,$model := .Models}}
	{{$model.Name}}Client := {{$model.Name}}.New{{$model.Name | Title}}ServiceClient(conn)
	err = {{$model.Name}}.Register{{$model.Name | Title}}ServiceHandlerClient(ctx, rmux, {{$model.Name}}Client)
	if err != nil {
		log.Fatal(err)
	}
	{{end}}
	
	handler := cors.Default().Handler(rmux)
	log.Println("rest server ready...")

	err = http.ListenAndServe(":8000", handler)
	if err != nil {
		log.Fatal(err)
	}
	wg.Done()

}

func main() {
	var wg sync.WaitGroup
	helpers.InitializeDatabase()
	wg.Add(1)
	go startGRPC(&wg)

	wg.Add(1)
	go startHTTP(&wg)

	wg.Wait()
}
`
