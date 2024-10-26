package main

import (
	"fmt"
	"os"
	"testing"
)

// Now, you can pass a mock path or a temporary test directory in your unit tests.
func TestProcessFile(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp(".", "test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up after the test

	// Define environment and file path
	env := "dev"
	filePath := fmt.Sprintf("%s/%s.yaml", tempDir, env)

	// Write test data to the temporary file
	content := []byte("hello, world")
	if err := os.WriteFile(filePath, content, 0o660); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	// Call the ProcessFile function with the temporary directory as basePath
	err = ProcessFile(tempDir, env)
	if err != nil {
		t.Fatalf("ProcessFile returned an error: %v", err)
	}
}
