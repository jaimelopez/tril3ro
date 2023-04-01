package keyboard

/*
 #cgo CFLAGS: -x objective-c
 #cgo LDFLAGS: -framework Cocoa

 #import <Foundation/Foundation.h>
 #import <ApplicationServices/ApplicationServices.h>

extern CGEventRef handleEvent(CGEventTapProxy proxy, CGEventType type, CGEventRef event, void *refcon);

static inline CFMachPortRef CGCreateEventTap() {
	CGEventMask eventMask = {
		CGEventMaskBit(kCGEventKeyDown) |
		CGEventMaskBit(kCGEventKeyUp) |
		CGEventMaskBit(kCGEventFlagsChanged)
	};

	CFMachPortRef eventTap = CGEventTapCreate(
		kCGSessionEventTap, kCGHeadInsertEventTap, 0, eventMask, handleEvent, NULL
	);

	CGEventTapEnable(eventTap, true);

	return eventTap;
}
*/
import "C"

import (
	"sync"
	"unsafe"

	"golang.org/x/exp/slices"
)

var (
	relock    sync.RWMutex
	receivers = []receiver{}
	current   = NewState()
	tap       C.CFMachPortRef
	ref       C.CFRunLoopSourceRef
)

type platform_listener struct {
	quit chan struct{}
}

func (l *listener) init() (*listener, error) {
	l.quit = make(chan struct{}, 1)

	return l, nil
}

func (l *listener) start() error {
	tap = C.CGCreateEventTap()

	if unsafe.Pointer(&tap) == nil {
		return ErrUnableToCreateListener
	}

	ref = C.CFMachPortCreateRunLoopSource(C.kCFAllocatorDefault, tap, 0)

	if unsafe.Pointer(&ref) == nil {
		return ErrUnableToCreateListener
	}

	go func() {
		C.CFRunLoopAddSource(C.CFRunLoopGetCurrent(), ref, C.kCFRunLoopCommonModes)
		C.CFRunLoopRun()

		<-l.quit
	}()

	return nil
}

func (l *listener) stop() {
	C.CGEventTapEnable(tap, false)
	C.CFRunLoopRemoveSource(C.CFRunLoopGetCurrent(), ref, C.kCFRunLoopCommonModes)
	C.CFRunLoopStop(C.CFRunLoopGetCurrent())

	l.quit <- struct{}{}
}

func (l *listener) listen(ch receiver) error {
	relock.Lock()
	defer relock.Unlock()

	receivers = append(receivers, ch)

	return nil
}

func (l *listener) unlisten(ch receiver) {
	relock.Lock()
	defer relock.Unlock()

	for idx, rch := range receivers {
		if ch == rch {
			receivers = slices.Delete(receivers, idx, idx+1)
		}
	}
}

//export handleEvent
func handleEvent(proxy C.CGEventTapProxy, et C.CGEventType, event C.CGEventRef, refcon *C.void) C.CGEventRef {
	t := EventType(et)
	key := Key(C.CGEventGetIntegerValueField(event, C.kCGKeyboardEventKeycode))
	keys := []Key{key}

	switch t {
	case EventKeyPress:
		current.Press(key)
	case EventKeyRelease:
		current.Release(key)
	case eventKeyModifier:
		current.Toggle(key)

		if current.IsPressed(key) {
			t = EventKeyPress
		} else {
			t = EventKeyRelease
		}
	default:
		return event
	}

	if t == EventKeyPress {
		keys = current.Pressed()
	}

	relock.RLock()
	defer relock.RUnlock()

	for _, re := range receivers {
		re <- Event{Type: t, Keys: keys}
	}

	return event
}
