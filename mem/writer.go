package mem

import (
	"unsafe"

	"github.com/jaimelopez/tril3ro/common"
)

// writer struct to manage memory writing operations
type writer[T any] struct {
	*manager
}

// Write certain data into a particular memory address
func (w *writer[T]) Write(addr common.Addr, data T) error {
	var et T

	return w.WriteOf(addr, data, uint(unsafe.Sizeof(et)))
}

// NewWriter instantiates a new memory writer for specified data struct
func NewWriter[T any](opts ...Option) (*writer[T], error) {
	m, err := newManager(opts...)
	if err != nil {
		return nil, err
	}

	return &writer[T]{m}, nil
}

// NewWriterForProc instantiates a new memory writer for specified data struct using the default handler
func NewWriterForProc[T any](processID uint32) (*writer[T], error) {
	return NewWriter[T](WithDefaultHandler(processID))
}
