package keyboard

import (
	"bytes"
	"context"
	"encoding/binary"
	"os"
	"syscall"
	"unsafe"

	"github.com/jaimelopez/tril3ro/internal/device"
	"github.com/jaimelopez/tril3ro/internal/execution"
)

const inputEventSize = int(unsafe.Sizeof(inputEvent{}))

type eventReceiver func() (Key, EventType)

type platform_listener struct {
	events   []string
	cancelFn context.CancelFunc
	state    *state
}

func (l *listener) init() (*listener, error) {
	if !execution.IsRoot() {
		return nil, ErrInsufficientPrivileges
	}

	events, err := device.FindKeyboardEvents()
	if err != nil {
		return nil, err
	}

	l.events = events
	l.state = NewState()

	return l, nil
}

func (l *listener) start() error {
	ctx, cancelFn := context.WithCancel(context.Background())
	l.cancelFn = cancelFn
	rec := make(chan eventReceiver)

	l.wg.Add(1)

	go func(ctx context.Context, rec chan eventReceiver) {
		defer l.wg.Done()

		for {
			select {
			case fn := <-rec:
				l.handle(fn())
			case <-ctx.Done():
				return
			}
		}
	}(ctx, rec)

	for _, event := range l.events {
		l.wg.Add(1)

		go func(event string) {
			defer l.wg.Done()

			l.read(ctx, event, rec)
		}(event)
	}

	return nil
}

func (l *listener) stop() {
	l.cancelFn()
	l.wg.Wait()
}

func (l *listener) listen(ch receiver) error {
	return nil
}

func (l *listener) unlisten(ch receiver) {}

func (l *listener) read(ctx context.Context, event string, rec chan<- eventReceiver) {
	fd, err := os.OpenFile(event, os.O_RDONLY|syscall.O_NONBLOCK, os.ModeDevice)
	if err != nil {
		return
	}

	defer fd.Close()

	listening := true

	go func(listening *bool) {
		for *listening {
			event := inputEvent{}
			buffer := make([]byte, inputEventSize)

			// Will also return err when fd is closed
			_, err := fd.Read(buffer)
			if err != nil {
				return
			}

			// Not interested in other events more than key change
			err = binary.Read(bytes.NewBuffer(buffer), binary.LittleEndian, &event)
			if err != nil || event.EV != evKeyChange {
				continue
			}

			rec <- func() (Key, EventType) {
				return event.Key, event.Type
			}
		}
	}(&listening)

	<-ctx.Done()

	listening = false
}

func (l *listener) handle(key Key, t EventType) {
	keys := []Key{key}

	switch t {
	case EventKeyPress:
		l.state.Press(key)
		keys = l.state.Pressed()
	case EventKeyRelease:
		l.state.Release(key)
	default:
		return
	}

	l.lock.RLock()
	defer l.lock.RUnlock()

	for _, re := range l.receivers {
		re <- Event{Type: t, Keys: keys}
	}
}
