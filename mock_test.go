package main

import (
	"os"
	"testing"
)

func TestGetAddress(t *testing.T) {
	tempFile, err := os.CreateTemp("./", "test-input")
	if err != nil {
		t.Fatalf("Error creating temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	input := "123 Main st, Austin Texas\n"
	if _, err := tempFile.Write([]byte(input)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	if _, err := tempFile.Seek(0, 0); err != nil {
		t.Fatalf("Failed to seek in temp file: %v", err)
	}

	originalStdin := os.Stdin
	os.Stdin = tempFile
	defer func() {
		os.Stdin = originalStdin
	}()

	got := getAddress()
	want := "123 Main st, Austin Texas"
	if got != want {
		t.Errorf("Expected %v, but got %v", want, got)
	}
}
