package keyboard

import "sync"

type state struct {
	mods sync.Map
}

// NewState instantiates a new keyboard state
func NewState() *state {
	return &state{}
}

// Press marks a particular key as pressed
func (s *state) Press(key Key) {
	s.assign(key, true)
}

// Release marks a particular key as released
func (s *state) Release(key Key) {
	s.assign(key, false)
}

// Toggle the state of a key
func (s *state) Toggle(key Key) {
	val, _ := s.mods.LoadOrStore(key, false)

	s.mods.Store(key, !val.(bool))
}

// IsPressed checks whether a particular key it´s already marked as pressed or not
func (s *state) IsPressed(key Key) bool {
	_, exists := s.mods.Load(key)

	return exists
}

// Pressed returns a snapshot of the keys that were marked as pressed
func (s *state) Pressed() []Key {
	pressed := []Key{}

	s.mods.Range(func(key any, value any) bool {
		if value.(bool) {
			pressed = append(pressed, key.(Key))
		}

		return true
	})

	return pressed
}

func (s *state) assign(key Key, press bool) {
	s.mods.Store(key, press)
}
