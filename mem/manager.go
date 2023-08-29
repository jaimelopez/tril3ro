package mem

import "github.com/jaimelopez/tril3ro/common"

type manager struct {
	*handler
}

func newManager(opts ...Option) (*manager, error) {
	m := &manager{}

	for _, option := range opts {
		err := option(m)
		if err != nil {
			return nil, err
		}
	}

	if m.handler == nil {
		return nil, ErrHandlerNotSpecified
	}

	return m, nil
}

// Option represents a function that help to configure the manager
type Option func(*manager) error

// WithDefaultHandler option includes a default handler for the specified process id
func WithDefaultHandler(processID common.ProcessID) Option {
	return func(m *manager) error {
		h, err := NewHandler(processID)
		m.handler = h

		return err
	}
}

// WithtHandler option allows to reuse previously defined handlers
func WithtHandler(h *handler) Option {
	return func(m *manager) error {
		m.handler = h

		return nil
	}
}
