package closer

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// Closer вызывает функции в порядке LIFO.
type Closer struct {
	mu        sync.Mutex
	callbacks []callback
}

// New создает Closer.
func New() *Closer {
	return &Closer{}
}

// Add добавляет коллбек для закрытия.
func (c *Closer) Add(fn func() error, opts ...CallbackOption) {
	c.mu.Lock()
	c.callbacks = append(c.callbacks, newCallback(fn, opts...))
	c.mu.Unlock()
}

// Close вызывает все коллбеки.
func (c *Closer) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	var errs []error
	for i := len(c.callbacks) - 1; i >= 0; i-- {
		cb := c.callbacks[i]

		ch := make(chan error, 1)
		go func() {
			defer close(ch)
			ch <- cb.fn()
		}()

		select {
		case <-time.After(cb.config.timeout):
			errs = append(errs, fmt.Errorf("[%s]: timed out [%s]", cb.config.name, cb.config.timeout))
		case err := <-ch:
			if err != nil {
				errs = append(errs, fmt.Errorf("[%s]: %w", cb.config.name, err))
			}
		}
	}

	return errors.Join(errs...)
}
