package mem

import (
	"runtime"
	"sync/atomic"

	"github.com/jaimelopez/tril3ro/common"
)

type handler struct {
	platform_handler
	processID common.ProcessID
	opened    atomic.Bool
}

// NewHandler instantiates and open a new memory handler pointing to a particular process id
// It will error in case the handler can not attach to the process no matter the cause
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
