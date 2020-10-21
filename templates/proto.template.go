package templates

//ProtoHeader ...
var ProtoHeader string = `syntax = "proto3";
import "google/api/annotations.proto";
package {{.Name}};
option go_package = ".;{{.Name}}";

message Error {
	string error = 1;
	string message = 2;
}

`

//ProtoMessages ...
var ProtoMessages string = `

message FilterField {
	string type = 1;
	string value = 2;
}

message Filter {
	FilterField id = 1;{{range $index,$field := .Fields}}
	FilterField {{$field.Name}} = {{$index | addTwo }};{{end}}
}

message {{.Name | Title}} {
	string id = 1;{{range $index,$field := .Fields}}
	{{$field.Type}} {{$field.Name}} = {{$index | addTwo }};{{end}}
}

message Get{{.Name | Title}}Request{
	Filter filter = 1;
	{{if .IncludesPassword}}bool includePassword = 2;{{end}}
}
  
message Get{{.Name | Title}}Response {
	{{.Name | Title}} data = 1;
	bool success = 2;
	Error error = 3;
}

message GetMultiple{{.Name | Title}}Request {
	repeated string ids = 1;
	Filter filter = 2;
	{{if .IncludesPassword}}bool includePassword = 3;{{end}}
}
  
message GetMultiple{{.Name | Title}}Response {
	repeated {{.Name | Title}} data = 1;
	bool success = 2;
	Error error = 3;
}

message Create{{.Name | Title}}Request {
	{{.Name | Title}} data = 1;
}

message Create{{.Name | Title}}Response {
	string id = 1;
	bool success = 2;
	Error error = 3;
}

message CreateMultiple{{.Name | Title}}Request {
	repeated {{.Name | Title}} data = 1;
}

message CreateMultiple{{.Name | Title}}Response {
	repeated string ids = 1;
	bool success = 2;
	Error error = 3;
}

message Update{{.Name | Title}}Request {
	Filter filter = 1;
	{{.Name | Title}} data = 2;
}

message Update{{.Name | Title}}Response {
	bool success = 1;
	Error error = 2;
}

message UpdateMultiple{{.Name | Title}}Request {
	repeated string ids = 1;
	Filter filter = 2;
	{{.Name | Title}} data = 3;
}

message UpdateMultiple{{.Name | Title}}Response {
	bool success = 1;
	Error error = 2;
}

message Delete{{.Name | Title}}Request{
	string id = 1;
	Filter filter = 2;
}

message Delete{{.Name | Title}}Response{
	bool success = 1;
	Error error = 2;
}

message DeleteMultiple{{.Name | Title}}Request{
	repeated string ids = 1;
	Filter filter = 2;
}

message DeleteMultiple{{.Name | Title}}Response{
	bool success = 1;
	Error error = 2;
}

`

//ProtoServiceHeader ...
var ProtoServiceHeader string = `
service {{.Name}}Service {
`

//ProtoServiceMethods ...
var ProtoServiceMethods string = `
rpc Get{{.Name | Title}}(Get{{.Name | Title}}Request) returns (Get{{.Name | Title}}Response) {
	option (google.api.http) = {
		post: "/v1/{{.Name}}/get"
		body: "*"
	};
}
rpc GetMultiple{{.Name | Title}}(GetMultiple{{.Name | Title}}Request) returns (GetMultiple{{.Name | Title}}Response) {
	option (google.api.http) = {
		post: "/v1/{{.Name}}/getMultiple"
		body: "*"
	};
}
rpc Create{{.Name | Title}}(Create{{.Name | Title}}Request) returns (Create{{.Name | Title}}Response) {
	option (google.api.http) = {
		post: "/v1/{{.Name}}/create"
		body: "*"
	};
}
rpc CreateMultiple{{.Name | Title}}(CreateMultiple{{.Name | Title}}Request) returns (CreateMultiple{{.Name | Title}}Response) {
	option (google.api.http) = {
		post: "/v1/{{.Name}}/createMultiple"
		body: "*"
	};
}
rpc Update{{.Name | Title}}(Update{{.Name | Title}}Request) returns (Update{{.Name | Title}}Response) {
	option (google.api.http) = {
		post: "/v1/{{.Name}}/update"
		body: "*"
	};
}
rpc UpdateMultiple{{.Name | Title}}(UpdateMultiple{{.Name | Title}}Request) returns (UpdateMultiple{{.Name | Title}}Response) {
	option (google.api.http) = {
		post: "/v1/{{.Name}}/updateMultiple"
		body: "*"
	};
}
rpc Delete{{.Name | Title}}(Delete{{.Name | Title}}Request) returns (Delete{{.Name | Title}}Response) {
	option (google.api.http) = {
		post: "/v1/{{.Name}}/delete"
		body: "*"
	};
}
rpc DeleteMultiple{{.Name | Title}}(DeleteMultiple{{.Name | Title}}Request) returns (DeleteMultiple{{.Name | Title}}Response) {
	option (google.api.http) = {
		post: "/v1/{{.Name}}/deleteMultiple"
		body: "*"
	};
}
`

//ProtoServiceFooter ...
var ProtoServiceFooter string = `
}
`
