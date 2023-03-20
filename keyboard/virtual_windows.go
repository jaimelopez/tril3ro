package keyboard

// #cgo pkg-config: libevdev
// #include <libevdev/libevdev.h>
// #include <libevdev/libevdev-uinput.h>
import "C"

var (
	procKeyBdEvent = user32.NewProc("keybd_event")
	procVkKeyScan  = user32.NewProc("VkKeyScanA")
)

type platform_virtual struct{}

func (v *virtual) init() (*virtual, error) {
	return v, nil
}

func (v virtual) sequence(value string) error {
	for i := 0; i < len(value); i++ {
		char, _, err := procVkKeyScan.Call(uintptr(value[i]))
		if char == 0 {
			return err
		}

		keys := []Key{Key(char)}

		if ((char >> 8) & 0xF) == scanShiftModifier {
			keys = []Key{Key(KeyShift), Key(char)}
		}

		if err := v.press(keys...); err != nil {
			return err
		}

		if err := v.release(keys...); err != nil {
			return err
		}
	}

	return nil
}

func (v virtual) press(keys ...Key) error {
	for _, key := range keys {
		ret, _, err := procKeyBdEvent.Call(uintptr(key), 0, 0x0000, 0)
		if ret != 0 {
			return err
		}
	}

	return nil
}

func (v virtual) release(keys ...Key) error {
	for _, key := range keys {
		ret, _, err := procKeyBdEvent.Call(uintptr(key), 0, 0x0002, 0)
		if ret != 0 {
			return err
		}
	}

	return nil
}
