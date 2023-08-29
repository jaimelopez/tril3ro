package mem

import (
	"unsafe"

	"github.com/jaimelopez/tril3ro/proc"
)

// Writer struct to manage memory writing operations
type Writer[T any] struct {
	*manager
}

// Write certain data into a particular memory address
func (w *Writer[T]) Write(addr proc.Addr, data T) error {
	var et T

	return w.WriteOf(addr, data, uint(unsafe.Sizeof(et)))
}

// NewWriter instantiates a new memory writer for specified data struct
func NewWriter[T any](opts ...Option) (*Writer[T], error) {
	m, err := newManager(opts...)
	if err != nil {
		return nil, err
	}

	return &Writer[T]{m}, nil
}
