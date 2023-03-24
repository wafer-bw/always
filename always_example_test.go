package always_test

import (
	"errors"
	"fmt"

	"github.com/wafer-bw/always"
)

func ExampleMust_panics() {
	always.Must(func() (string, error) { return "", errors.New("oh no") }())
}

func ExampleMust_notPanics() {
	res := always.Must(func() (string, error) { return "hello", nil }())
	fmt.Println(res)
	// Output:
	// hello
}
