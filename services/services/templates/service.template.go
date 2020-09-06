package templates

//Service is template for service file generation
var Service string = `package {{.Name}}

import (
	"encoding/json"
	
	"{{.Package}}/services/{{.Name}}/src/utils"
	uuid "github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct {
}
{{range $i,$model := .Models}}
// Get{{$model.Name | Title}} returns single {{$model.Name}}
func (s *Server) Get{{$model.Name | Title}}(ctx context.Context, req *Get{{$model.Name | Title}}Request) (*Get{{$model.Name | Title}}Response, error) {
	filter := make(map[string]interface{})
	if req.Filter != nil {
		if req.Filter.Id != nil {
			if req.Filter.Id.Type != "" {
				filter["id"] = req.Filter.Id.Value
			}
		} {{range $index,$field := $model.Fields}} {{if $model.IncludesPassword }} {{if (eq $model.PasswordField $field.Name)}}  {{ else }} 
		if req.Filter.{{$field.Name | Title}} != nil {
			if req.Filter.{{$field.Name | Title}}.Type != "" { 
				filter["{{$field.Name}}"] = req.Filter.{{$field.Name | Title}}.Value 
			}
		} {{end}} {{ else }} {{if (eq $model.PasswordField $field.Name)}} {{ else }} if req.Filter.{{$field.Name | Title}} != nil {
			if req.Filter.{{$field.Name | Title}}.Type != "" { 
				filter["{{$field.Name}}"] = req.Filter.{{$field.Name | Title}}.Value 
			}
		} {{end}} {{end}} {{end}} 
	}
	data, err := utils.GetRecord("{{$model.Name}}", filter)
	if err != nil {
		return &Get{{$model.Name | Title}}Response{Success: false, Error: &Error{Error: "record_not_found"}}, nil
	}
	var record {{$model.Name | Title}}
	record.Id = data["id"].(string){{range $index,$field := $model.Fields}}
	{{ if or (eq $field.Type "bool") (eq $field.Type "int32") (eq $field.Type "int64") (eq $field.Type "uint32") (eq $field.Type "uint64") (eq $field.Type "string") }} {{if $model.IncludesPassword }} {{if (eq $model.PasswordField $field.Name)}} if req.IncludePassword {
		record.{{$field.Name | Title}} = data["{{$field.Name}}"].({{$field.Type}})
		
	} {{ else }} record.{{$field.Name | Title}} = data["{{$field.Name}}"].({{$field.Type}}) {{end}} {{ else }} {{if (eq $model.PasswordField $field.Name)}} {{ else }}  record.{{$field.Name | Title}} = data["{{$field.Name}}"].({{$field.Type}}) {{end}} {{end}} {{end}} {{ if (eq $field.Type "double") }}record.{{$field.Name | Title}} = data["{{$field.Name}}"].(float64){{end}} {{ if (eq $field.Type "float") }}record.{{$field.Name | Title}} = data["{{$field.Name}}"].(float32){{end}} {{ if (eq $field.Type "float") }}record.{{$field.Name | Title}} = data["{{$field.Name}}"].(float32){{end}} {{end}}

	return &Get{{$model.Name | Title}}Response{Success: true, Data: &record}, nil
}

// GetMultiple{{$model.Name | Title}} fuctions returns list of all {{$model.Name}} by a specific filter
func (s *Server) GetMultiple{{$model.Name | Title}}(ctx context.Context, req *GetMultiple{{$model.Name | Title}}Request) (*GetMultiple{{$model.Name | Title}}Response, error) {

	filter := make(map[string]interface{})
	if req.Filter != nil {
		if len(req.Ids) > 0 {
			ids := make(map[string]interface{})
			ids["$in"] = req.Ids
			filter["id"] = ids
		} else if req.Filter.Id != nil {
			if req.Filter.Id.Type != "" {
				filter["id"] = req.Filter.Id.Value
			}
		} {{range $index,$field := $model.Fields}} {{if $model.IncludesPassword }} {{if (eq $model.PasswordField $field.Name)}}  {{ else }} 
		if req.Filter.{{$field.Name | Title}} != nil {
			if req.Filter.{{$field.Name | Title}}.Type != "" { 
				filter["{{$field.Name}}"] = req.Filter.{{$field.Name | Title}}.Value 
			}
		} {{end}} {{ else }} {{if (eq $model.PasswordField $field.Name)}} {{ else }} if req.Filter.{{$field.Name | Title}} != nil {
			if req.Filter.{{$field.Name | Title}}.Type != "" { 
				filter["{{$field.Name}}"] = req.Filter.{{$field.Name | Title}}.Value 
			}
		} {{end}} {{end}} {{end}} 
	}

	var records []*{{$model.Name | Title}}
	data, err := utils.GetRecords("{{$model.Name}}", filter)
	if err != nil {
		return &GetMultiple{{$model.Name | Title}}Response{Success: false, Error: &Error{Error: "records_not_found"}}, nil
	}
	for i := 0; i < len(data); i++ {
		var record {{$model.Name | Title}}
		record.Id = data[i]["id"].(string){{range $index,$field := $model.Fields}}
		{{ if or (eq $field.Type "bool") (eq $field.Type "int32") (eq $field.Type "int64") (eq $field.Type "uint32") (eq $field.Type "uint64") (eq $field.Type "string") }} {{if $model.IncludesPassword }} {{if (eq $model.PasswordField $field.Name)}} if req.IncludePassword {
			record.{{$field.Name | Title}} = data[i]["{{$field.Name}}"].({{$field.Type}})
		} {{ else }} record.{{$field.Name | Title}} = data[i]["{{$field.Name}}"].({{$field.Type}}) {{end}} {{ else }} {{if (eq $model.PasswordField $field.Name)}} {{ else }}  record.{{$field.Name | Title}} = data[i]["{{$field.Name}}"].({{$field.Type}}) {{end}} {{end}} {{end}} {{ if (eq $field.Type "double") }}record.{{$field.Name | Title}} = data[i]["{{$field.Name}}"].(float64){{end}} {{ if (eq $field.Type "float") }}record.{{$field.Name | Title}} = data[i]["{{$field.Name}}"].(float32){{end}} {{ if (eq $field.Type "float") }}record.{{$field.Name | Title}} = data[i]["{{$field.Name}}"].(float32){{end}} {{end}}
		records = append(records, &record)
	}
	return &GetMultiple{{$model.Name | Title}}Response{Success: true, Data: records}, nil
}

// Create{{$model.Name | Title}} stores new {{$model.Name}} in database and returns id
func (s *Server) Create{{$model.Name | Title}}(ctx context.Context, req *Create{{$model.Name | Title}}Request) (*Create{{$model.Name | Title}}Response, error) {
	req.Data.Id = uuid.New().String()
	bytes, err := bcrypt.GenerateFromPassword([]byte(req.Data.{{$model.PasswordField | Title}}), 14)
	if err != nil {
		return &Create{{$model.Name | Title}}Response{Success: false, Error: &Error{Error: "unable_to_parse_password"}}, nil
	}
	req.Data.{{$model.PasswordField | Title}} = string(bytes)

	err = utils.InsertRecord("{{$model.Name}}", req.Data)
	if err != nil {
		return &Create{{$model.Name | Title}}Response{Success: false, Error: &Error{Error: "unable_to_insert_record"}}, nil
	}
	return &Create{{$model.Name | Title}}Response{Success: true, Id: req.Data.Id}, nil
}

// CreateMultiple{{$model.Name | Title}} stores multiple {{$model.Name}} in database and returns ids
func (s *Server) CreateMultiple{{$model.Name | Title}}(ctx context.Context, req *CreateMultiple{{$model.Name | Title}}Request) (*CreateMultiple{{$model.Name | Title}}Response, error) {
	var records []interface{}
	var insertedIDs []string
	for _, record := range req.Data {
		record.Id = uuid.New().String()
		insertedIDs = append(insertedIDs, record.Id)
		bytes, _ := bcrypt.GenerateFromPassword([]byte(record.{{$model.PasswordField | Title}}), 14)
		record.{{$model.PasswordField | Title}} = string(bytes)
		records = append(records, record)
	}
	err := utils.InsertRecords("{{$model.Name}}", records)
	if err != nil {
		return &CreateMultiple{{$model.Name | Title}}Response{Success: false, Error: &Error{Error: "unable_to_insert_record"}}, nil
	}
	return &CreateMultiple{{$model.Name | Title}}Response{Success: true, Ids: insertedIDs}, nil
}

// Update{{$model.Name | Title}} updates a {{$model.Name}} and returns success state
func (s *Server) Update{{$model.Name | Title}}(ctx context.Context, req *Update{{$model.Name | Title}}Request) (*Update{{$model.Name | Title}}Response, error) {
	filter := make(map[string]interface{})
	if req.Filter != nil {
		if req.Filter.Id != nil {
			if req.Filter.Id.Type != "" {
				filter["id"] = req.Filter.Id.Value
			}
		} {{range $index,$field := $model.Fields}} {{if $model.IncludesPassword }} {{if (eq $model.PasswordField $field.Name)}}  {{ else }} 
		if req.Filter.{{$field.Name | Title}} != nil {
			if req.Filter.{{$field.Name | Title}}.Type != "" { 
				filter["{{$field.Name}}"] = req.Filter.{{$field.Name | Title}}.Value 
			}
		} {{end}} {{ else }} {{if (eq $model.PasswordField $field.Name)}} {{ else }} if req.Filter.{{$field.Name | Title}} != nil {
			if req.Filter.{{$field.Name | Title}}.Type != "" { 
				filter["{{$field.Name}}"] = req.Filter.{{$field.Name | Title}}.Value 
			}
		} {{end}} {{end}} {{end}} 
	}

	update := make(map[string]interface{})
	jsonData, _ := json.Marshal(req.Data)
	json.Unmarshal(jsonData, &update)
	
	err := utils.UpdateRecord("{{$model.Name}}", filter, update)
	if err != nil {
		return &Update{{$model.Name | Title}}Response{Success: false, Error: &Error{Error: "unable_to_update_record"}}, nil
	}
	return &Update{{$model.Name | Title}}Response{Success: true}, nil
}

// UpdateMultiple{{$model.Name | Title}} updates multiple {{$model.Name}} and returns success state
func (s *Server) UpdateMultiple{{$model.Name | Title}}(ctx context.Context, req *UpdateMultiple{{$model.Name | Title}}Request) (*UpdateMultiple{{$model.Name | Title}}Response, error) {
	filter := make(map[string]interface{})
	if req.Filter != nil {
		if len(req.Ids) > 0 {
			ids := make(map[string]interface{})
			ids["$in"] = req.Ids
			filter["id"] = ids
		} else if req.Filter.Id != nil {
			if req.Filter.Id.Type != "" {
				filter["id"] = req.Filter.Id.Value
			}
		} {{range $index,$field := $model.Fields}} {{if $model.IncludesPassword }} {{if (eq $model.PasswordField $field.Name)}}  {{ else }} 
		if req.Filter.{{$field.Name | Title}} != nil {
			if req.Filter.{{$field.Name | Title}}.Type != "" { 
				filter["{{$field.Name}}"] = req.Filter.{{$field.Name | Title}}.Value 
			}
		} {{end}} {{ else }} {{if (eq $model.PasswordField $field.Name)}} {{ else }} if req.Filter.{{$field.Name | Title}} != nil {
			if req.Filter.{{$field.Name | Title}}.Type != "" { 
				filter["{{$field.Name}}"] = req.Filter.{{$field.Name | Title}}.Value 
			}
		} {{end}} {{end}} {{end}} 
	}

	update := make(map[string]interface{})
	jsonData, _ := json.Marshal(req.Data)
	json.Unmarshal(jsonData, &update)

	err := utils.UpdateRecords("{{$model.Name}}", filter, update)
	if err != nil {
		return &UpdateMultiple{{$model.Name | Title}}Response{Success: false, Error: &Error{Error: "unable_to_update_records"}}, nil
	}
	return &UpdateMultiple{{$model.Name | Title}}Response{Success: true}, nil
}

// Delete{{$model.Name | Title}} deletes a {{$model.Name}} by id
func (s *Server) Delete{{$model.Name | Title}}(ctx context.Context, req *Delete{{$model.Name | Title}}Request) (*Delete{{$model.Name | Title}}Response, error) {
	filter := make(map[string]interface{})
	if req.Filter != nil {
		if req.Filter.Id != nil {
			if req.Filter.Id.Type != "" {
				filter["id"] = req.Filter.Id.Value
			}
		} {{range $index,$field := $model.Fields}} {{if $model.IncludesPassword }} {{if (eq $model.PasswordField $field.Name)}}  {{ else }} 
		if req.Filter.{{$field.Name | Title}} != nil {
			if req.Filter.{{$field.Name | Title}}.Type != "" { 
				filter["{{$field.Name}}"] = req.Filter.{{$field.Name | Title}}.Value 
			}
		} {{end}} {{ else }} {{if (eq $model.PasswordField $field.Name)}} {{ else }} if req.Filter.{{$field.Name | Title}} != nil {
			if req.Filter.{{$field.Name | Title}}.Type != "" { 
				filter["{{$field.Name}}"] = req.Filter.{{$field.Name | Title}}.Value 
			}
		} {{end}} {{end}} {{end}} 
	}

	err := utils.DeleteRecord("{{$model.Name}}", filter)
	if err != nil {
		return &Delete{{$model.Name | Title}}Response{Success: false, Error: &Error{Error: "unable_to_delete_record"}}, nil
	}
	return &Delete{{$model.Name | Title}}Response{Success: true}, nil
}

// DeleteMultiple{{$model.Name | Title}} deletes multiple {{$model.Name}} by ids or filter
func (s *Server) DeleteMultiple{{$model.Name | Title}}(ctx context.Context, req *DeleteMultiple{{$model.Name | Title}}Request) (*DeleteMultiple{{$model.Name | Title}}Response, error) {
	filter := make(map[string]interface{})
	if req.Filter != nil {
		if len(req.Ids) > 0 {
			ids := make(map[string]interface{})
			ids["$in"] = req.Ids
			filter["id"] = ids
		} else if req.Filter.Id != nil {
			if req.Filter.Id.Type != "" {
				filter["id"] = req.Filter.Id.Value
			}
		} {{range $index,$field := $model.Fields}} {{if $model.IncludesPassword }} {{if (eq $model.PasswordField $field.Name)}}  {{ else }} 
		if req.Filter.{{$field.Name | Title}} != nil {
			if req.Filter.{{$field.Name | Title}}.Type != "" { 
				filter["{{$field.Name}}"] = req.Filter.{{$field.Name | Title}}.Value 
			}
		} {{end}} {{ else }} {{if (eq $model.PasswordField $field.Name)}} {{ else }} if req.Filter.{{$field.Name | Title}} != nil {
			if req.Filter.{{$field.Name | Title}}.Type != "" { 
				filter["{{$field.Name}}"] = req.Filter.{{$field.Name | Title}}.Value 
			}
		} {{end}} {{end}} {{end}} 
	}

	err := utils.DeleteRecords("{{$model.Name}}", filter)
	if err != nil {
		return &DeleteMultiple{{$model.Name | Title}}Response{Success: false, Error: &Error{Error: "unable_to_delete_records"}}, nil
	}
	return &DeleteMultiple{{$model.Name | Title}}Response{Success: true}, nil
}
{{end}}
`

//UserService is template for service file generation
var UserService string = `package {{.Name}}

// Do your Implementations here...
`

//Database is template for database helper function generation
var Database string = `package utils

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// Database is global database variable
	Database *mongo.Database
	// URI is mongo connection string
	URI string = "mongodb+srv://egnite:Aqbfjotld9@cluster0-wtkg5.mongodb.net/egnite?retryWrites=true&w=majority"
)

func init() {

	clientOptions := options.Client().ApplyURI(URI)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	Database = client.Database("egnite")
}

// GetRecord will fetch single record
func GetRecord(db string, filter map[string]interface{}) (map[string]interface{}, error) {
	var record map[string]interface{}
	collection := Database.Collection(db)
	documentReturned := collection.FindOne(context.TODO(), filter)
	err := documentReturned.Decode(&record)
	if err != nil {
		return record, err
	}
	return record, nil
}

// GetRecords will fetch multiple records
func GetRecords(db string, filter map[string]interface{}) ([]map[string]interface{}, error) {
	var records []map[string]interface{}
	collection := Database.Collection(db)
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return records, err
	}
	for cur.Next(context.TODO()) {
		var record map[string]interface{}
		_ = cur.Decode(&record)
		records = append(records, record)
	}
	return records, nil
}

// InsertRecord will insert single record
func InsertRecord(db string, record interface{}) error {
	collection := Database.Collection(db)
	_, err := collection.InsertOne(context.TODO(), record)
	if err != nil {
		return err
	}
	return nil
}

// InsertRecords will insert multiple records
func InsertRecords(db string, records []interface{}) error {
	collection := Database.Collection(db)
	_, err := collection.InsertMany(context.TODO(), records)
	if err != nil {
		return err
	}
	return nil
}

// UpdateRecord will update single record
func UpdateRecord(db string, filter map[string]interface{}, update map[string]interface{}) error {
	collection := Database.Collection(db)
	_, err := collection.UpdateOne(context.TODO(), filter, bson.D{{Key: "$set", Value: update}})
	if err != nil {
		return err
	}
	return nil
}

// UpdateRecords will update multiple records
func UpdateRecords(db string, filter map[string]interface{}, update map[string]interface{}) error {
	collection := Database.Collection(db)
	_, err := collection.UpdateMany(context.TODO(), filter, bson.D{{Key: "$set", Value: update}})
	if err != nil {
		return err
	}
	return nil
}

// DeleteRecord will delete single record
func DeleteRecord(db string, filter interface{}) error {
	collection := Database.Collection(db)
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

// DeleteRecords will delete multiple records
func DeleteRecords(db string, filter map[string]interface{}) error {
	collection := Database.Collection(db)
	_, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
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

	{{.Name}} "{{.Package}}/services/{{.Name}}/src"
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

	{{.Name}}Server := {{.Name}}.Server{}

	{{.Name}}.Register{{.Name | Title}}ServiceServer(grpcServer, &{{.Name}}Server)
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

	{{.Name}}Client := {{.Name}}.New{{.Name | Title}}ServiceClient(conn)
	err = {{.Name}}.Register{{.Name | Title}}ServiceHandlerClient(ctx, rmux, {{.Name}}Client)
	if err != nil {
		log.Fatal(err)
	}
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

	wg.Add(1)
	go startGRPC(&wg)

	wg.Add(1)
	go startHTTP(&wg)

	wg.Wait()
}
`
