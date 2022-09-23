package definition

type Validator interface {
	Validate(val interface{}) bool
}