package mem

import (
	"unsafe"

	"github.com/jaimelopez/tril3ro/proc"
	"golang.org/x/sys/windows"
)

// Write certain data into a particular memory address
func (w *writer[T]) WriteOf(addr proc.Addr, data T, size uint) error {
	_ = w.open()

	dtw := (*[]byte)(unsafe.Pointer(&data))

	if err := windows.WriteProcessMemory(
		w.handle,
		uintptr(addr),
		(*byte)(unsafe.Pointer(dtw)),
		uintptr(size),
		nil,
	); err != nil {
		return err
	}

	return nil
}
