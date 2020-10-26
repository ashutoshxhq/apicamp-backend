package templates

//Service is template for service file generation
var Service string = `package {{.Name}}

import (
	"{{.ServiceName}}/helpers"
	uuid "github.com/google/uuid"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct {
}

// GetSingle{{.Name | Title}} returns single {{.Name}}
func (s *Server) GetSingle{{.Name | Title}}(ctx context.Context, req *GetSingle{{.Name | Title}}Request) (*GetSingle{{.Name | Title}}Response, error) {
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
	rows, err := conn.Query(context.Background(), "SELECT * FROM {{.Name}}"+where+" LIMIT 1")
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
	return &GetSingle{{.Name | Title}}Response{Success: true, Data: &record}, nil
}

// GetMultiple{{.Name | Title}} fuctions returns list of all {{.Name}} by a specific filter
func (s *Server) GetMultiple{{.Name | Title}}(ctx context.Context, req *GetMultiple{{.Name | Title}}Request) (*GetMultiple{{.Name | Title}}Response, error) {

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

	conn, err := helpers.DatabasePool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	rows, err := conn.Query(context.Background(), "SELECT * FROM {{.Name}}"+where+" LIMIT 100")
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
	return &GetMultiple{{.Name | Title}}Response{Success: true, Data: records}, nil
}

// CreateSingle{{.Name | Title}} stores new {{.Name}} in database and returns id
func (s *Server) CreateSingle{{.Name | Title}}(ctx context.Context, req *CreateSingle{{.Name | Title}}Request) (*CreateSingle{{.Name | Title}}Response, error) {
	// req.Data.Id = uuid.New().String()
	
	
	return &CreateSingle{{.Name | Title}}Response{Success: true}, nil
}

// CreateMultiple{{.Name | Title}} stores multiple {{.Name}} in database and returns ids
func (s *Server) CreateMultiple{{.Name | Title}}(ctx context.Context, req *CreateMultiple{{.Name | Title}}Request) (*CreateMultiple{{.Name | Title}}Response, error) {
	var records []interface{}
	var insertedIDs []string
	for _, record := range req.Data {
		record.Id = uuid.New().String()
		insertedIDs = append(insertedIDs, record.Id)
		records = append(records, record)
	}
	
	return &CreateMultiple{{.Name | Title}}Response{Success: true}, nil
}

// UpdateSingle{{.Name | Title}} updates a {{.Name}} and returns success state
func (s *Server) UpdateSingle{{.Name | Title}}(ctx context.Context, req *UpdateSingle{{.Name | Title}}Request) (*UpdateSingle{{.Name | Title}}Response, error) {
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

	
	return &UpdateSingle{{.Name | Title}}Response{Success: true}, nil
}

// UpdateMultiple{{.Name | Title}} updates multiple {{.Name}} and returns success state
func (s *Server) UpdateMultiple{{.Name | Title}}(ctx context.Context, req *UpdateMultiple{{.Name | Title}}Request) (*UpdateMultiple{{.Name | Title}}Response, error) {
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

	
	return &UpdateMultiple{{.Name | Title}}Response{Success: true}, nil
}

// DeleteSingle{{.Name | Title}} deletes a {{.Name}} by id
func (s *Server) DeleteSingle{{.Name | Title}}(ctx context.Context, req *DeleteSingle{{.Name | Title}}Request) (*DeleteSingle{{.Name | Title}}Response, error) {
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
	
	return &DeleteSingle{{.Name | Title}}Response{Success: true}, nil
}

// DeleteMultiple{{.Name | Title}} deletes multiple {{.Name}} by ids or filter
func (s *Server) DeleteMultiple{{.Name | Title}}(ctx context.Context, req *DeleteMultiple{{.Name | Title}}Request) (*DeleteMultiple{{.Name | Title}}Response, error) {
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
	
	return &DeleteMultiple{{.Name | Title}}Response{Success: true}, nil
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
