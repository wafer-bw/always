package always_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wafer-bw/always"
)

func TestMust(t *testing.T) {
	t.Run("returns result & does not panic when there is no error", func(t *testing.T) {
		require.NotPanics(t, func() {
			res := always.Must(func() (string, error) { return "foo", nil }())
			require.Equal(t, "foo", res)
		})
	})

	t.Run("panic when there is an error", func(t *testing.T) {
		require.Panics(t, func() {
			always.Must(func() (string, error) { return "", errors.New("oh no") }())
		})
	})
}
