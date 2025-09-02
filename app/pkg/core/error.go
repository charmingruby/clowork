package core

import "fmt"

type NotFoundError struct {
	message string
}

func NewNotFoundError(resource string) NotFoundError {
	return NotFoundError{
		message: fmt.Sprintf("%s not found", resource),
	}
}

func (e NotFoundError) Error() string {
	return e.message
}

type ConflictError struct {
	message string
}

func NewConflictError(resource string) ConflictError {
	return ConflictError{
		message: fmt.Sprintf("%s already exists", resource),
	}
}

func (e ConflictError) Error() string {
	return e.message
}

type DatabaseError struct {
	originalErr error
	message     string
}

func NewDatabaseError(err error) DatabaseError {
	return DatabaseError{
		message: "database error",
	}
}

func (e DatabaseError) Error() string {
	return e.message
}

func (e DatabaseError) Unwrap() error {
	return e.originalErr
}

type InvalidCredentialsError struct {
	message string
}

func NewInvalidCredentialsError() InvalidCredentialsError {
	return InvalidCredentialsError{
		message: "invalid credentials",
	}
}

func (e InvalidCredentialsError) Error() string {
	return e.message
}
