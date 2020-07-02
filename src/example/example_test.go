// Test filename must be ended with _test.go

package main

import "testing"

// Tesy function name must begin with Test*
func TestAdd(t *testing.T) {
	res := Add(5, 5)
	if res != 10 {
		t.Errorf("Sum was incorrect, got: %f, want: %f.", res, 10.0)
	}
}
