package proc

import (
	"runtime"
	"strconv"
	"strings"
	"sync/atomic"
)

// ProcessID represents a process identifier
type ProcessID = uint32

// Addr represents a memory address
type Addr = uintptr

// AddrFromString converts a string into an Addr
// Notice that it will not return any error so if anything went wrong, Addr will be 0.
func AddrFromString(addr string) Addr {
	a, _ := strconv.ParseUint(addr, 16, 0)
	return Addr(a)
}

// Process definition
type Process struct {
	platform_process
	ID       ProcessID
	ParentID ProcessID
	Name     string
	opened   atomic.Bool
}

func (proc *Process) open() error {
	if proc.opened.Load() {
		return nil
	}

	if err := proc.init(); err != nil {
		return err
	}

	proc.opened.Store(true)

	// Make sure that proc gets stopped correctly whenever it's garbage collected
	runtime.SetFinalizer(proc, func(obj any) {
		obj.(*Process).close()
	})

	return nil
}

// Module retrieves a particular dynamic module based on the name
func (proc *Process) Module(name string) (*Module, error) {
	modules, err := proc.AllModules()
	if err != nil {
		return nil, err
	}

	for _, module := range modules {
		if strings.EqualFold(module.Name, name) {
			return module, nil
		}
	}

	return nil, ErrModuleNotFound
}

// ProcessByName retrieves a list of processes that matches the specified name
func ProcessByName(name string) ([]*Process, error) {
	processes, err := AllProcesses()
	if err != nil {
		return nil, err
	}

	matches := []*Process{}

	for _, process := range processes {
		if process.Name == name {
			matches = append(matches, process)
		}
	}

	return matches, nil
}
