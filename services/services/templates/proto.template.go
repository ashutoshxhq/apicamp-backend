package templates

//Proto is template for proto file generation
var Proto string = `syntax = "proto3";
import "google/api/annotations.proto";
package {{.Name}};
option go_package = ".;{{.Name}}";

message Error {
	string error = 1;
	string message = 2;
}
{{range $i,$model := .Models}}

message FilterField {
	string type = 1;
	string value = 2;
}

message Filter {
	FilterField id = 1;{{range $index,$field := $model.Fields}}
	FilterField {{$field.Name}} = {{$index | addTwo }};{{end}}
}

message {{$model.Name | Title}} {
	string id = 1;{{range $index,$field := $model.Fields}}
	{{$field.Type}} {{$field.Name}} = {{$index | addTwo }};{{end}}
}

message Get{{$model.Name | Title}}Request{
	Filter filter = 1;
	{{if $model.IncludesPassword}}bool includePassword = 2;{{end}}
}
  
message Get{{$model.Name | Title}}Response {
	{{$model.Name | Title}} data = 1;
	bool success = 2;
	Error error = 3;
}

message GetMultiple{{$model.Name | Title}}Request {
	repeated string ids = 1;
	Filter filter = 2;
	{{if $model.IncludesPassword}}bool includePassword = 3;{{end}}
}
  
message GetMultiple{{$model.Name | Title}}Response {
	repeated {{$model.Name | Title}} data = 1;
	bool success = 2;
	Error error = 3;
}

message Create{{$model.Name | Title}}Request {
	{{$model.Name | Title}} data = 1;
}

message Create{{$model.Name | Title}}Response {
	string id = 1;
	bool success = 2;
	Error error = 3;
}

message CreateMultiple{{$model.Name | Title}}Request {
	repeated {{$model.Name | Title}} data = 1;
}

message CreateMultiple{{$model.Name | Title}}Response {
	repeated string ids = 1;
	bool success = 2;
	Error error = 3;
}

message Update{{$model.Name | Title}}Request {
	Filter filter = 1;
	{{$model.Name | Title}} data = 2;
}

message Update{{$model.Name | Title}}Response {
	bool success = 1;
	Error error = 2;
}

message UpdateMultiple{{$model.Name | Title}}Request {
	repeated string ids = 1;
	Filter filter = 2;
	{{$model.Name | Title}} data = 3;
}

message UpdateMultiple{{$model.Name | Title}}Response {
	bool success = 1;
	Error error = 2;
}

message Delete{{$model.Name | Title}}Request{
	string id = 1;
	Filter filter = 2;
}

message Delete{{$model.Name | Title}}Response{
	bool success = 1;
	Error error = 2;
}

message DeleteMultiple{{$model.Name | Title}}Request{
	repeated string ids = 1;
	Filter filter = 2;
}

message DeleteMultiple{{$model.Name | Title}}Response{
	bool success = 1;
	Error error = 2;
}
{{end}}

service {{.Name}}Service {
{{range $i,$model := .Models}}
{{if $model.Crudgen }}
	rpc Get{{$model.Name | Title}}(Get{{$model.Name | Title}}Request) returns (Get{{$model.Name | Title}}Response) {
		option (google.api.http) = {
			post: "/v1/{{$model.Name}}/get"
			body: "*"
		};
	}
	rpc GetMultiple{{$model.Name | Title}}(GetMultiple{{$model.Name | Title}}Request) returns (GetMultiple{{$model.Name | Title}}Response) {
		option (google.api.http) = {
			post: "/v1/{{$model.Name}}/getMultiple"
			body: "*"
		};
	}
	rpc Create{{$model.Name | Title}}(Create{{$model.Name | Title}}Request) returns (Create{{$model.Name | Title}}Response) {
		option (google.api.http) = {
			post: "/v1/{{$model.Name}}/create"
			body: "*"
		};
	}
	rpc CreateMultiple{{$model.Name | Title}}(CreateMultiple{{$model.Name | Title}}Request) returns (CreateMultiple{{$model.Name | Title}}Response) {
		option (google.api.http) = {
			post: "/v1/{{$model.Name}}/createMultiple"
			body: "*"
		};
	}
	rpc Update{{$model.Name | Title}}(Update{{$model.Name | Title}}Request) returns (Update{{$model.Name | Title}}Response) {
		option (google.api.http) = {
			post: "/v1/{{$model.Name}}/update"
			body: "*"
		};
	}
	rpc UpdateMultiple{{$model.Name | Title}}(UpdateMultiple{{$model.Name | Title}}Request) returns (UpdateMultiple{{$model.Name | Title}}Response) {
		option (google.api.http) = {
			post: "/v1/{{$model.Name}}/updateMultiple"
			body: "*"
		};
	}
	rpc Delete{{$model.Name | Title}}(Delete{{$model.Name | Title}}Request) returns (Delete{{$model.Name | Title}}Response) {
		option (google.api.http) = {
			post: "/v1/{{$model.Name}}/delete"
			body: "*"
		};
	}
	rpc DeleteMultiple{{$model.Name | Title}}(DeleteMultiple{{$model.Name | Title}}Request) returns (DeleteMultiple{{$model.Name | Title}}Response) {
		option (google.api.http) = {
			post: "/v1/{{$model.Name}}/deleteMultiple"
			body: "*"
		};
	}
	{{end}}
{{end}}
}
`
