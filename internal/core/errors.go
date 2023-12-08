package core

import "fmt"

type InvalidDataError struct {
	Status  int
	Message string
	Example interface{}
}

type NotFoundError struct {
	Status  int
	Message string
}

type CantCreateErr struct {
	Status  int
	Message string
}

func NewNotFundErr(status int, message string) NotFoundError {
	return NotFoundError{
		Status:  status,
		Message: message,
	}
}

func NewCantCreateErr(status int, message string) CantCreateErr {
	return CantCreateErr{
		Status:  status,
		Message: message,
	}
}

func NewInvalidDataErr(status int, message string, exmp interface{}) InvalidDataError {
	return InvalidDataError{
		Status:  status,
		Message: message,
		Example: exmp,
	}
}

func (nfe NotFoundError) Error() string {
	return fmt.Sprintf("%v not found", nfe.Message)
}

func (cce CantCreateErr) Error() string {
	return fmt.Sprintf("%v can not create", cce.Message)
}

func (invD InvalidDataError) Error() string {
	return fmt.Sprintf("invalide data: %s", invD.Message)
}
