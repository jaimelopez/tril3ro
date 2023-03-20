package keyboard

import "golang.org/x/exp/slices"

type virtual struct {
	platform_virtual
}

// NewVirtual instiates a virtual keyboard to interact with
func NewVirtual() (*virtual, error) {
	return (&virtual{}).init()
}

// Type a literal in the virtual keyboard
// TODO: Maybe implement a delay to emulate a real typing for every keystroke
func (v virtual) Type(value string) error {
	return v.sequence(value)
}

// Press and release a sequence of keys in the virtual keyboard
// Duplicated keys will be executed only once
func (v virtual) PressAndRelease(keys ...Key) error {
	if err := v.Press(keys...); err != nil {
		return err
	}

	if err := v.Release(keys...); err != nil {
		return err
	}

	return nil
}

// Press a sequence of keys in the virtual keyboard
// Duplicated keys will be pressed only once
func (v virtual) Press(keys ...Key) error {
	return v.press(slices.Compact(keys)...)
}

// Release a sequence of keys in the virtual keyboard
// Duplicated keys will be release only once
func (v virtual) Release(keys ...Key) error {
	return v.release(slices.Compact(keys)...)
}
