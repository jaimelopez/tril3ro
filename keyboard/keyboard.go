package keyboard

import (
	"golang.org/x/exp/slices"
)

// Key representation in a keyboard
type Key uint16

// EventType represents the type of a keyboard event
type EventType uint32

// Event represents an interaction with the keyboard
type Event struct {
	// Type specifies which kind of event it stores
	Type EventType
	// Keys represents a sequence of keyboard keys
	Keys []Key
}

// Matches checks if the event matches with a certan group of keys
func (event Event) Matches(keys ...Key) bool {
	return slices.Equal(keys, event.Keys)
}

// Contans evaluates if the event contains at least the specified sequence
func (event Event) Contains(keys ...Key) bool {
	for _, val := range keys {
		if !slices.Contains(event.Keys, val) {
			return false
		}
	}

	return true
}

// NewKeyPressEvent generates a key press event
func NewKeyPressEvent(keys ...Key) Event {
	return Event{Type: EventKeyPress, Keys: keys}
}

// NewKeReleaseEvent generates a key release event
// Be aware that a release event can only contain one key as there´s no way to release several keys at the same time
func NewKeReleaseEvent(key Key) Event {
	return Event{Type: EventKeyRelease, Keys: []Key{key}}
}
