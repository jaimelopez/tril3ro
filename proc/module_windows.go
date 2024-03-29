package proc

import (
	"syscall"
	"unsafe"

	"github.com/jaimelopez/tril3ro/common"
	"golang.org/x/sys/windows"
)

// AllModules retrieves all dynamic modules for the process
func (proc *Process) AllModules() ([]*Module, error) {
	handle, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPMODULE|windows.TH32CS_SNAPMODULE32, proc.ID)
	if err != nil {
		return nil, err
	}

	var entry windows.ModuleEntry32
	entry.Size = uint32(unsafe.Sizeof(entry))

	mods := []*Module{}

	for err := windows.Module32First(handle, &entry); err == nil; err = windows.Module32Next(handle, &entry) {
		mods = append(mods, &Module{
			ProcessID: proc.ID,
			Address:   common.Addr(entry.ModBaseAddr),
			Size:      entry.ModBaseSize,
			Name:      syscall.UTF16ToString(entry.Module[:]),
			Path:      syscall.UTF16ToString(entry.ExePath[:]),
		})
	}

	return mods, nil
}
