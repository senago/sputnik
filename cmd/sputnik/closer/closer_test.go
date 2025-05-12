package closer

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCloserSequence(t *testing.T) {
	t.Parallel()

	const size = 10
	counter := size
	closer := New()

	for i := 1; i <= size; i++ {
		idx := i
		closer.Add(
			func() error {
				assert.Equal(t, counter, idx)
				counter--
				return nil
			},
			WithCallbackName(fmt.Sprint(idx)),
		)
	}
	require.NoError(t, closer.Close())

	require.Equal(t, 0, counter)
}

func TestCloserTimeout(t *testing.T) {
	t.Parallel()

	closer := New()

	closer.Add(
		func() error {
			time.Sleep(50 * time.Millisecond)
			return nil
		},
		WithCallbackTimeout(time.Millisecond),
	)

	require.Error(t, closer.Close())
}

func TestCloserError(t *testing.T) {
	t.Parallel()

	closer := New()

	closer.Add(
		func() error {
			return assert.AnError
		},
		WithCallbackTimeout(time.Second),
	)

	require.ErrorIs(t, closer.Close(), assert.AnError)
}
