package mem

import (
	"runtime"
	"sync/atomic"
)

type handler struct {
	platform_process
	processID uint32
	opened    atomic.Bool
}

func NewHandler(processID uint32) (*handler, error) {
	h := &handler{processID: processID}

	err := h.open()
	if err != nil {
		return nil, err
	}

	return h, nil
}

func (h *handler) open() error {
	if h.opened.Load() {
		return nil
	}

	if err := h.init(); err != nil {
		return err
	}

	h.opened.Store(true)

	// Make sure that proc gets stopped correctly whenever it's garbage collected
	runtime.SetFinalizer(h, func(obj any) {
		obj.(*handler).close()
	})

	return nil
}
