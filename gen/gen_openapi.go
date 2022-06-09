package gen

//go:generate oapi-codegen -generate "chi-server" -o openapi/server.go resources/openapi/v1.yaml
//go:generate oapi-codegen -generate "types" -o openapi/type.go resources/openapi/v1.yaml
//go:generate oapi-codegen -generate "spec" -o openapi/spec.go resources/openapi/v1.yaml
