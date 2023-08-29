package mem

/*
#include <mem_darwin.h>
*/
import "C"

import (
	"unsafe"

	"github.com/jaimelopez/tril3ro/proc"
)

// Read certain memory address
func (r *reader[T]) ReadOf(addr proc.Addr, into *T, size uint) error {
	sz := C.uint(size)

	C.read_process_memory_bytes(C.uint(r.task), C.mach_vm_address_t(addr), unsafe.Pointer(into), &sz)

	return nil
}
