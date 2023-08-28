package mem

/*
 #cgo CFLAGS: -x objective-c
 #cgo LDFLAGS: -framework Cocoa

 #include <libproc.h>
 #include <mach-o/dyld_images.h>
 #include <mach/mach_traps.h>
 #include <mach/mach_init.h>
*/
import "C"

type platform_handler struct {
	task uint32
}

func (h *handler) init() error {
	var task C.task_t

	C.task_for_pid(C.mach_task_self_, C.int(h.processID), &task)

	h.task = uint32(task)

	return nil
}

func (h *handler) close() {
	h.task = 0
}
