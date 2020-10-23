package templates

//Proto ...
var Proto string = `syntax = "proto3";
import "google/api/annotations.proto";
package {{.Name}};
option go_package = ".;{{.Name}}";

message Error {
	string error = 1;
	string message = 2;
}

message FilterField {
	string type = 1;
	string value = 2;
}

message Filter {
	{{range $index,$field := .Fields}}
	FilterField {{$field.Name}} = {{$index | addOne }};{{end}}
}

message {{.Name | Title}} {
	{{range $index,$field := .Fields}}
	{{$field.Type}} {{$field.Name}} = {{$index | addOne }};{{end}}
}

message GetSingle{{.Name | Title}}Request{
	Filter filter = 1;
}
  
message GetSingle{{.Name | Title}}Response {
	{{.Name | Title}} data = 1;
	bool success = 2;
	Error error = 3;
}

message GetMultiple{{.Name | Title}}Request {
	repeated string ids = 1;
	Filter filter = 2;
}
  
message GetMultiple{{.Name | Title}}Response {
	repeated {{.Name | Title}} data = 1;
	bool success = 2;
	Error error = 3;
}

message CreateSingle{{.Name | Title}}Request {
	{{.Name | Title}} data = 1;
}

message CreateSingle{{.Name | Title}}Response {
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

message UpdateSingle{{.Name | Title}}Request {
	Filter filter = 1;
	{{.Name | Title}} data = 2;
}

message UpdateSingle{{.Name | Title}}Response {
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

message DeleteSingle{{.Name | Title}}Request{
	string id = 1;
	Filter filter = 2;
}

message DeleteSingle{{.Name | Title}}Response{
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

service {{.Name}}Service {
	rpc GetSingle{{.Name | Title}}(GetSingle{{.Name | Title}}Request) returns (GetSingle{{.Name | Title}}Response) {
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
	rpc CreateSingle{{.Name | Title}}(CreateSingle{{.Name | Title}}Request) returns (CreateSingle{{.Name | Title}}Response) {
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
	rpc UpdateSingle{{.Name | Title}}(UpdateSingle{{.Name | Title}}Request) returns (UpdateSingle{{.Name | Title}}Response) {
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
	rpc DeleteSingle{{.Name | Title}}(DeleteSingle{{.Name | Title}}Request) returns (DeleteSingle{{.Name | Title}}Response) {
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
}
`
