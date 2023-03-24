package always

// Must is used by wrapping a call to a constructor style function, it will panic with the error if it is non-nil.
func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
