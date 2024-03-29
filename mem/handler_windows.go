package mem

import (
	"golang.org/x/sys/windows"
)

type platform_handler struct {
	handle windows.Handle
}

func (h *handler) init() error {
	handle, err := windows.OpenProcess(windows.PROCESS_VM_OPERATION|windows.PROCESS_VM_READ|windows.PROCESS_VM_WRITE, false, uint32(h.processID))
	if err != nil {
		return err
	}

	h.handle = handle

	return nil
}

func (h *handler) close() {
	_ = windows.CloseHandle(h.handle)
}
