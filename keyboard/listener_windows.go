package keyboard

import (
	"sync"
	"syscall"
	"unsafe"
)

var (
	user32                  = syscall.NewLazyDLL("user32.dll")
	procSetWindowsHookEx    = user32.NewProc("SetWindowsHookExW")
	procCallNextHookEx      = user32.NewProc("CallNextHookEx")
	procUnhookWindowsHookEx = user32.NewProc("UnhookWindowsHookEx")
	procGetMessage          = user32.NewProc("GetMessageW")
)

type handler func(int, uintptr, uintptr) uintptr

type platform_listener struct {
	state state
	hooks sync.Map
}

func (l *listener) init() (*listener, error) {
	l.state = NewState()

	return l, nil
}

func (l *listener) start() error {
	return nil
}

func (l *listener) stop() {}

func (l *listener) listen(ch receiver) error {
	cherr := make(chan error, 1)

	go func() {
		var hook uintptr
		var err error

		hook, _, _ = procSetWindowsHookEx.Call(hookKeyboardLL, syscall.NewCallback(l.handle(ch, &hook)), 0, 0)
		if hook == 0 {
			err = ErrUnableToCreateListener
		}

		cherr <- err

		l.hooks.Store(ch, hook)

		for {
			if _, exists := l.hooks.Load(ch); !exists {
				break
			}

			_, _, _ = procGetMessage.Call(hook, 0, 0, 0)
		}
	}()

	return <-cherr
}

func (l *listener) unlisten(ch receiver) {
	if hook, exists := l.hooks.LoadAndDelete(ch); exists {
		procUnhookWindowsHookEx.Call(hook.(uintptr))
	}
}

func (l *listener) handle(ch receiver, hook *uintptr) handler {
	return func(code int, wparam uintptr, lparam uintptr) uintptr {
		ret, _, _ := procCallNextHookEx.Call(*hook, uintptr(code), wparam, lparam)

		if code != 0 {
			return ret
		}

		t := EventType(wparam)
		key := Key(*(*Key)(unsafe.Pointer(lparam))()
		keys := []Key{key}

		switch t {
		case EventKeyPress:
			l.state.Press(key)
			keys = l.state.Pressed()
		case EventKeyRelease:
			l.state.Release(key)
		default:
			return ret
		}

		ch <- Event{Type: t, Keys: keys}

		return ret
	}
}
