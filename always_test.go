package always_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/wafer-bw/always"
)

func TestMust(t *testing.T) {
	t.Run("returns result & does not panic when there is no error", func(t *testing.T) {
		expectResult := "foo"
		input := func() (string, error) {
			return expectResult, nil
		}

		defer func() {
			if r := recover(); r != nil {
				t.Errorf(fmt.Sprintf("func always.Must(%T) should not panic\n\tPanic value: %v", input, r))
			}
		}()

		result := always.Must(input())
		if result != expectResult {
			t.Errorf("func always.Must(%T) = %v, want %v", input, result, expectResult)
		}
	})

	t.Run("panic when there is an error", func(t *testing.T) {
		expectResult, expectPanic := "", "oh no"
		input := func() (string, error) {
			return expectResult, errors.New(expectPanic)
		}

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("func always.Must(%T) should panic", input)
			} else if r.(error).Error() != expectPanic {
				t.Errorf("func always.Must(%T) should panic with %v, got %v", input, expectPanic, r)
			}
		}()

		result := always.Must(input())
		if result != expectResult {
			t.Errorf("func always.Must(%T) = %v, want %v", input, result, expectResult)
		}
	})
}

func TestMustDo(t *testing.T) {
	t.Run("does not panic when there is no error", func(t *testing.T) {
		input := func() error {
			return nil
		}

		defer func() {
			if r := recover(); r != nil {
				t.Errorf(fmt.Sprintf("func always.Must(%T) should not panic\n\tPanic value: %v", input, r))
			}
		}()

		always.MustDo(input())
	})

	t.Run("panic when there is an error", func(t *testing.T) {
		expectPanic := "oh no"
		input := func() error {
			return errors.New(expectPanic)
		}

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("func always.Must(%T) should panic", input)
			} else if r.(error).Error() != expectPanic {
				t.Errorf("func always.Must(%T) should panic with %v, got %v", input, expectPanic, r)
			}
		}()

		always.MustDo(input())
	})
}
