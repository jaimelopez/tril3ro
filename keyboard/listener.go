package keyboard

import (
	"runtime"
	"sync"
)

// Handler represents a function to be executed under certain even
type Handler func()

// EventHandler represents a handler for events
type EventHandler func(event Event)

type receiver chan Event

type listener struct {
	platform_listener
	lock      sync.RWMutex
	wg        sync.WaitGroup
	receivers []receiver
}

// NewListener instiates a new keyboard listener
func NewListener() (*listener, error) {
	l, err := (&listener{}).init()

	// Make sure that listener gets stopped correctly whenever it's garbage collected
	runtime.SetFinalizer(l, func(obj any) {
		l := obj.(*listener)
		l.Stop()
	})

	return l, err
}

// Bind certain event
// Be aware that a release event can only contain one key as there´s no way to release several keys at the same time
func (l *listener) Bind(h Handler, be Event) error {
	if be.Type == EventKeyPress && len(be.Keys) < 1 {
		return ErrInvalidPressEvent
	}

	if be.Type == EventKeyRelease && len(be.Keys) != 1 {
		return ErrInvalidReleaseEvent
	}

	return l.Hook(func(event Event) {
		if event.Type == be.Type && event.Matches(be.Keys...) {
			h()
		}
	})
}

// Hook all the events and forward them to the specified event handler
func (l *listener) Hook(h EventHandler) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	if len(l.receivers) == 0 {
		if err := l.start(); err != nil {
			return err
		}
	}

	ch := make(receiver)

	err := l.listen(ch)
	if err != nil {
		return err
	}

	l.receivers = append(l.receivers, ch)

	l.wg.Add(1)

	go func() {
		defer l.wg.Done()

		for event := range ch {
			h(event)
		}
	}()

	return nil
}

// Stop listener so no more event will be received
func (l *listener) Stop() {
	l.lock.Lock()
	defer l.lock.Unlock()

	for _, channel := range l.receivers {
		l.unlisten(channel)

		close(channel)
	}

	l.receivers = []receiver{}

	l.stop()

	l.wg.Wait()
}
