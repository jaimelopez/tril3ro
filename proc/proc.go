package proc

import (
	"strconv"
	"strings"
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
	ID       ProcessID
	ParentID ProcessID
	Name     string
}

func (proc *Process) Open() error {
	return nil
}

func (proc *Process) Close() error {
	return nil
}

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

func (proc *Process) AllModules() ([]*Module, error) {
	return allModules(proc.ID)
}

// ByID retrieves a process that matches the specified ID
func ByID(id ProcessID) (*Process, error) {
	return process(id)
}

// ByName retrieves a list of processes that matches the specified name
func ByName(name string) ([]*Process, error) {
	processes, err := allProcesses()
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

// All the running processes
func All() ([]*Process, error) {
	return allProcesses()
}

// AllIDs retrieves all the running processes IDs
func AllIDs() ([]ProcessID, error) {
	return allProcessesIDs()
}
