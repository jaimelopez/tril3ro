package proc

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

func allModules(id ProcessID) ([]*Module, error) {
	handle, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPMODULE|windows.TH32CS_SNAPMODULE32, id)
	if err != nil {
		return nil, err
	}

	var entry windows.ModuleEntry32
	entry.Size = uint32(unsafe.Sizeof(entry))

	mods := []*Module{}

	for err := windows.Module32First(handle, &entry); err == nil; err = windows.Module32Next(handle, &entry) {
		mods = append(mods, &Module{
			ProcessID: id,
			Address:   Addr(entry.ModBaseAddr),
			Size:      entry.ModBaseSize,
			Name:      syscall.UTF16ToString(entry.Module[:]),
			Path:      syscall.UTF16ToString(entry.ExePath[:]),
		})
	}

	return mods, nil
}
