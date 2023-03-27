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

	list := C.getModules(C.pid_t(id), &count)
	defer C.free(unsafe.Pointer(list))

	for _, mod := range unsafe.Slice(list, count) {
		mods = append(mods, &Module{
			ProcessID: id,
			Address:   Addr(mod.addr),
			Size:      uint32(mod.size),
			Name:      filepath.Base(C.GoString(mod.module)),
			Path:      C.GoString(mod.module),
		})
	}

	return mods, nil
}
