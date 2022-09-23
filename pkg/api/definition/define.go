package definition

import (
	"github.com/heketi/heketi/executors"
)

type ApiInfo struct {
	Name           string
	Description    string
	RequestDefine  *ApiRequest
	ResponseDefine *ApiResponse
	Executor       executors.Executor
}

type ApiRequest struct {
	Path   string
	Method MethodType
	Fields []*ApiRequestField
}

type ApiResponse struct {
}

type ApiRequestField struct {
	Name string

	Validators []Validator
}
