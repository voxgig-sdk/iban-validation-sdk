package core

type IbanValidationError struct {
	IsIbanValidationError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewIbanValidationError(code string, msg string, ctx *Context) *IbanValidationError {
	return &IbanValidationError{
		IsIbanValidationError: true,
		Sdk:              "IbanValidation",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *IbanValidationError) Error() string {
	return e.Msg
}
