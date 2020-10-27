package templates

//Proto ...
var Proto string = `syntax = "proto3";
import "google/api/annotations.proto";
import "google/protobuf/struct.proto";

package {{.Name}};
option go_package = ".;{{.Name}}";

message {{.Name | Title}} {
	{{range $index,$field := .Fields}}
	{{$field.Type}} {{$field.Name}} = {{$index | addOne }};{{end}}
}

message GetSingle{{.Name | Title}}Request{
	map<string, google.protobuf.Value> filters = 1;
}
  
message GetSingle{{.Name | Title}}Response {
	{{.Name | Title}} data = 1;
}

message GetMultiple{{.Name | Title}}Request {
	map<string, google.protobuf.Value> filters = 1;
}
  
message GetMultiple{{.Name | Title}}Response {
	repeated {{.Name | Title}} data = 1;
}

message CreateSingle{{.Name | Title}}Request {
	{{.Name | Title}} data = 1;
}

message CreateSingle{{.Name | Title}}Response {
	string message = 2;
}

message CreateMultiple{{.Name | Title}}Request {
	repeated {{.Name | Title}} data = 1;
}

message CreateMultiple{{.Name | Title}}Response {
	string message = 2;
}

message Update{{.Name | Title}}Request {
	map<string, google.protobuf.Value> filters = 1;
	map<string, google.protobuf.Value> data = 2;
}

message Update{{.Name | Title}}Response {
	string message = 2;
}

message Delete{{.Name | Title}}Request{
	map<string, google.protobuf.Value> filters = 1;
}

message Delete{{.Name | Title}}Response{
	string message = 2;
}

service {{.Name}}Service {
	rpc GetSingle{{.Name | Title}}(GetSingle{{.Name | Title}}Request) returns (GetSingle{{.Name | Title}}Response) {
		option (google.api.http) = {
			post: "/v1/{{.Name}}/getSingle"
			body: "*"
		};
	}
	rpc GetMultiple{{.Name | Title}}(GetMultiple{{.Name | Title}}Request) returns (GetMultiple{{.Name | Title}}Response) {
		option (google.api.http) = {
			post: "/v1/{{.Name}}"
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
	rpc Update{{.Name | Title}}(Update{{.Name | Title}}Request) returns (Update{{.Name | Title}}Response) {
		option (google.api.http) = {
			post: "/v1/{{.Name}}/update"
			body: "*"
		};
	}
	rpc Delete{{.Name | Title}}(Delete{{.Name | Title}}Request) returns (Delete{{.Name | Title}}Response) {
		option (google.api.http) = {
			post: "/v1/{{.Name}}/delete"
			body: "*"
		};
	}
	
}
`
