package always

// Must is a helper that wraps a call to a function returning (T, error)
// and panics if the error is non-nil.
func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

// MustDo is a helper that wraps a call to a function returning (error) and
// panics if the error is non-nil.
func MustDo(err error) {
	if err != nil {
		panic(err)
	}
}
