package proc

/*
#include <mem_darwin.h>
*/
import "C"

import (
	"unsafe"
)

// Read certain memory address
func (r *Reader[T]) ReadOf(addr Addr, into *T, size uint) error {
	sz := C.uint(size)

	_ = r.open()

	C.read_process_memory_bytes(C.uint(r.task), C.mach_vm_address_t(addr), unsafe.Pointer(into), &sz)

	return nil
}

// Write certain data into a particular memory address
func (r *Writer[T]) WriteOf(addr Addr, data T, size uint) error {
	_ = r.open()

	if !C.write_process_memory(
		C.uint(r.task),
		C.mach_vm_address_t(addr),
		C.vm_offset_t(uintptr(unsafe.Pointer(&data))),
		C.uint(size),
	) {
		return ErrUnexpectedResult
	}

	return nil
}
