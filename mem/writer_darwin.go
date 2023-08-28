package mem

/*
#include <mem_darwin.h>
*/
import "C"

import (
	"unsafe"

	"github.com/jaimelopez/tril3ro/proc"
)

// Write certain data into a particular memory address
func (r *Writer[T]) WriteOf(addr proc.Addr, data T, size uint) error {
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
