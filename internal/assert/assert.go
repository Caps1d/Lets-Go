package assert

import (
	"testing"
)

func Equal[T comparable](t *testing.T, actual, expected T) {
	// this indicates to the Go test runner that
	// our Equal() func is a test helper.
	// This means that when t.Errorf() is called from here
	// the Go test runner will report the filename
	// and line number of the code which called our Equal()
	// func in the output.
	t.Helper()

	if actual != expected {
		t.Errorf("got %v; want %v", actual, expected)
	}
}
