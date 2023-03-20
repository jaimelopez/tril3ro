package keyboard

import "errors"

var (
	// ErrUnableToCreateListener represents an error when listener hook cannot be created
	ErrUnableToCreateListener = errors.New("unable to create hook")

	// ErrInvalidPressEvent represents an error when an invalid press event type is defined
	ErrInvalidPressEvent = errors.New("a press event type has to contain at least one key")

	// ErrInvalidReleaseEvent represents an error when an invalid release event type is defined
	ErrInvalidReleaseEvent = errors.New("a release event type can contain just one key")

	// ErrInsufficientPrivileges represents an error when an operation is requested with not enough privileges
	ErrInsufficientPrivileges = errors.New("insufficient privileges to perform operation")
)
