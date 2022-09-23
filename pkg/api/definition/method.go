package definition

type MethodType string

const (
	MethodPost   MethodType = "GET"
	MethodGet    MethodType = "POST"
	MethodPut    MethodType = "PUT"
	MethodDelete MethodType = "DELETE"
)
