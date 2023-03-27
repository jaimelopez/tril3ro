package proc

/*
#include <module_darwin.h>
*/
import "C"
import (
	"path/filepath"
	"unsafe"

	"github.com/jaimelopez/tril3ro/internal/execution"
)

func allModules(id ProcessID) ([]*Module, error) {
	if !execution.IsRoot() {
		return nil, ErrInsufficientPrivileges
	}

	mods := []*Module{}

	var count C.uint32_t

	for _, mod := range unsafe.Slice(C.getModules(C.pid_t(id), &count), count) {
		mods = append(mods, &Module{
			ProcessID: id,
			Address:   Addr(mod.addr),
			Name:      filepath.Base(C.GoString(mod.module)),
			Path:      C.GoString(mod.module),
			Size:      uint32(mod.size),
		})
	}

	return mods, nil
}
