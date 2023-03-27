package keyboard

import (
	"encoding/binary"
	"os"
	"runtime"
	"syscall"

	"github.com/jaimelopez/tril3ro/internal/device"
	"github.com/jaimelopez/tril3ro/internal/execution"
)

type platform_virtual struct {
	fds []*os.File
}

func (v *virtual) init() (*virtual, error) {
	if !execution.IsRoot() {
		return nil, ErrInsufficientPrivileges
	}

	events, err := device.FindKeyboardEvents()
	if err != nil {
		return nil, err
	}

	for _, event := range events {
		fd, err := os.OpenFile(event, os.O_WRONLY|syscall.O_NONBLOCK, os.ModeDevice)
		if err != nil {
			v.stop()
			return nil, err
		}

		v.fds = append(v.fds, fd)
	}

	// Make sure we close fd when virtual gets garbage collected
	runtime.SetFinalizer(v, func(obj any) {
		v := obj.(*virtual)
		v.stop()
	})

	return v, nil
}

func (v virtual) sequence(value string) error {
	for i := 0; i < len(value); i++ {
		keys, found := charKeymap[value[i]]
		if !found {
			continue
		}

		v.press(keys...)
		v.release(keys...)
	}

	return nil
}

func (v virtual) press(keys ...Key) error {
	for _, key := range keys {
		err := v.send(inputEvent{
			EV:   evKeyChange,
			Key:  key,
			Type: EventKeyPress,
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func (v virtual) release(keys ...Key) error {
	for _, key := range keys {
		err := v.send(inputEvent{
			EV:   evKeyChange,
			Key:  key,
			Type: EventKeyRelease,
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func (v virtual) send(ie inputEvent) error {
	for _, fd := range v.fds {
		err := binary.Write(fd, binary.LittleEndian, ie)
		if err != nil {
			return err
		}

		err = binary.Write(fd, binary.LittleEndian, inputEvent{EV: evSync})
		if err != nil {
			return err
		}
	}

	return nil
}

func (v *virtual) stop() {
	for _, fd := range v.fds {
		fd.Close()
	}

	v.fds = []*os.File{}
}
