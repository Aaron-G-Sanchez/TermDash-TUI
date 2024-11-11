package main

import "testing"

// Nonsensical test for CI set up.
func TestSetup(t *testing.T) {
	got := 1
	if got != 1 {
		t.Errorf("Expected %v to equal 1: want 1", got)
	}
}
