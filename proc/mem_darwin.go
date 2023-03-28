package proc

/*
#include <mem_darwin.h>
*/
import "C"

import (
	"unsafe"
)

// Read certain memory address
func (r *Reader[T]) Read(addr Addr) (*T, error) {
	var data T

	size := C.uint(unsafe.Sizeof(data))
	dataAddr := uintptr(C.readProcessMemory(C.int(r.ID), C.mach_vm_address_t(addr), &size))

	bs := (*[]byte)(unsafe.Pointer(&data))
	vs := (*[]byte)(unsafe.Pointer(dataAddr))
	*bs = *vs

	return &data, nil
}

// Write certain data into a particular memory address
func (r *Writer[T]) Write(addr Addr, data T) error {
	if !C.writeProcessMemory(
		C.int(r.ID),
		C.mach_vm_address_t(addr),
		C.vm_offset_t(uintptr(unsafe.Pointer(&data))),
		C.uint(unsafe.Sizeof(data)),
	) {
		return ErrUnexpectedResult
	}

	return nil
}
