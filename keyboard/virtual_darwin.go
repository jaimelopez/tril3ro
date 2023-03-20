package keyboard

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa

#import <Foundation/Foundation.h>
#import <ApplicationServices/ApplicationServices.h>
*/
import "C"

var modififiers = map[Key]C.CGEventFlags{
	KeyCommand:      C.kCGEventFlagMaskCommand,
	KeyRightCommand: C.kCGEventFlagMaskCommand,
	KeyShift:        C.kCGEventFlagMaskShift,
	KeyRightShift:   C.kCGEventFlagMaskShift,
	KeyOption:       C.kCGEventFlagMaskAlternate,
	KeyRightOption:  C.kCGEventFlagMaskAlternate,
	KeyControl:      C.kCGEventFlagMaskControl,
	KeyRightControl: C.kCGEventFlagMaskControl,
	KeyCapsLock:     C.kCGEventFlagMaskAlphaShift,
	KeyFunction:     C.kCGEventFlagMaskSecondaryFn,
}

type platform_virtual struct{}

func (v *virtual) init() (*virtual, error) {
	return v, nil
}

func (virtual) sequence(value string) error {
	source := C.CGEventSourceCreate(C.kCGEventSourceStateHIDSystemState)
	press := C.CGEventCreateKeyboardEvent(source, 0, true)
	release := C.CGEventCreateKeyboardEvent(source, 0, false)

	uni := []C.UniChar{}

	for i := 0; i < len(value); i++ {
		uni = append(uni, C.UniChar(value[i]))
	}

	C.CGEventKeyboardSetUnicodeString(press, C.UniCharCount(len(value)), &uni[0])
	C.CGEventKeyboardSetUnicodeString(release, C.UniCharCount(len(value)), &uni[0])

	C.CGEventPost(C.kCGAnnotatedSessionEventTap, press)
	C.CGEventPost(C.kCGAnnotatedSessionEventTap, release)

	C.CFRelease(C.CFTypeRef(press))
	C.CFRelease(C.CFTypeRef(release))

	return nil
}

func (v virtual) press(keys ...Key) error {
	v.eventKeys(EventKeyPress, keys...)

	return nil
}

func (v virtual) release(keys ...Key) error {
	v.eventKeys(EventKeyRelease, keys...)

	return nil
}

func (virtual) eventKeys(t EventType, keys ...Key) {
	source := C.CGEventSourceCreate(C.kCGEventSourceStateHIDSystemState)
	nks := []Key{}
	mods := []C.CGEventFlags{}

	press := true
	if t == EventKeyRelease {
		press = false
	}

	for _, key := range keys {
		if val, ok := modififiers[key]; ok {
			mods = append(mods, val)
			continue
		}

		nks = append(nks, key)
	}

	for _, key := range nks {
		event := C.CGEventCreateKeyboardEvent(source, C.CGKeyCode(key), C.bool(press))

		for _, flag := range mods {
			C.CGEventSetFlags(event, flag)
		}

		C.CGEventPost(C.kCGAnnotatedSessionEventTap, event)
		C.CFRelease(C.CFTypeRef(event))
	}
}
