package mem

import (
	"unsafe"

	"github.com/jaimelopez/tril3ro/proc"
)

// Writer struct to manage memory writing operations
type Writer[T any] struct {
	*handler
}

// Write certain data into a particular memory address
func (r *Writer[T]) Write(addr proc.Addr, data T) error {
	var et T

	return r.WriteOf(addr, data, uint(unsafe.Sizeof(et)))
}

// NewWriter instantiates a new memory writer for specified data struct
func NewWriter[T any](processID uint32) (*Writer[T], error) {
	h, err := NewHandler(processID)
	if err != nil {
		return nil, err
	}

	return &Writer[T]{h}, nil
}
