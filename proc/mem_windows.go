package proc

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

// Read certain memory address
func (r *Reader[T]) Read(addr Addr, into *T) error {
	// TODO: Maybe move it to an open/close proc
	handle, err := windows.OpenProcess(windows.PROCESS_VM_OPERATION|windows.PROCESS_VM_READ, false, uint32(r.ID))
	if err != nil {
		return err
	}

	defer windows.CloseHandle(handle)

	var data T
	buffer := (*[]byte)(unsafe.Pointer(into))

	if err := windows.ReadProcessMemory(
		handle,
		uintptr(addr),
		(*byte)(unsafe.Pointer(buffer)),
		uintptr(unsafe.Sizeof(data)),
		nil,
	); err != nil {
		return err
	}

	return nil
}

// Write certain data into a particular memory address
func (r *Writer[T]) Write(addr Addr, data T) error {
	// TODO: Maybe move it to an open/close proc
	handle, err := windows.OpenProcess(windows.PROCESS_VM_OPERATION|windows.PROCESS_VM_WRITE, false, uint32(r.ID))
	if err != nil {
		return err
	}

	defer windows.CloseHandle(handle)

	dtw := (*[]byte)(unsafe.Pointer(&data))

	if err := windows.WriteProcessMemory(
		handle,
		uintptr(addr),
		(*byte)(unsafe.Pointer(dtw)),
		uintptr(unsafe.Sizeof(data)),
		nil,
	); err != nil {
		return err
	}

	return nil
}
