package mem

type platform_process struct {
	handle windows.Handle
}

func (h *handler)) init() error {
	handle, err := windows.OpenProcess(windows.PROCESS_VM_OPERATION|windows.PROCESS_VM_READ|windows.PROCESS_VM_WRITE, false, uint32(h.ProcessID))
	if err != nil {
		return err
	}

	h.handle = handle

	return nil
}

func (h *handler)) close() {
	_ = windows.CloseHandle(h.handle)
}
