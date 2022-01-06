package validator

type Validator interface {
	Validate(data interface{}) error
}
