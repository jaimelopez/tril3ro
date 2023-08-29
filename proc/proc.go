package proc

import (
	"strings"

	"github.com/jaimelopez/tril3ro/common"
)

// Process definition
type Process struct {
	ID       common.ProcessID
	ParentID common.ProcessID
	Name     string
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
