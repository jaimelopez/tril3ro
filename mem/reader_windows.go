package mem

import (
	"unsafe"

	"github.com/jaimelopez/tril3ro/common"
	"golang.org/x/sys/windows"
)

// Read certain memory address
func (r *reader[T]) ReadOf(addr common.Addr, into *T, size uint) error {
	_ = r.open()

	buffer := (*[]byte)(unsafe.Pointer(into))

	if err := windows.ReadProcessMemory(
		r.handle,
		uintptr(addr),
		(*byte)(unsafe.Pointer(buffer)),
		uintptr(size),
		nil,
	); err != nil {
		return err
	}

	return nil
}
