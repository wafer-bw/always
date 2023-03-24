package always

import "errors"

var (
	ErrNotOK = errors.New("not ok")
)

// Must is used to wrap a call to a constructor style function, will panic with the error if it is non-nil
func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
