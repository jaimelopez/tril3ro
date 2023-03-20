package file

import "errors"

var (
	// ErrNotReferencedObject its raised when a non pointer is received trying to parse into a struct
	ErrNotReferencedObject = errors.New("object not a pointer or reference")

	// ErrNotReferencedObject its raised when a non pointer is received trying to parse into a struct
	ErrInvalidObject = errors.New("object specified is not valid")

	// ErrNonStructObject its raised when the underlaying object is not an struct
	ErrNonStructObject = errors.New("underñaying object is not a struct")
)
