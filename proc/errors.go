package proc

import "errors"

var (
	// ErrOperationNotAllowed its raised whenever a operation is denied
	ErrOperationNotAllowed = errors.New("operation not allowed")

	// ErrProcessNotFound represents a process not found error
	ErrProcessNotFound = errors.New("process not found")

	// ErrMainTaskNotFound represents a task not found error
	ErrMainTaskNotFound = errors.New("main task not found")

	// ErrModuleNotFound represents a module not found error
	ErrModuleNotFound = errors.New("module not found")

	// ErrUnexpectedConversion represents generic unexpected error
	ErrUnexpectedResult = errors.New("unexpected result error")

	// ErrInsufficientPrivileges represents an error when an operation is requested with not enough privileges
	ErrInsufficientPrivileges = errors.New("insufficient privileges to perform operation")
)
