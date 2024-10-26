package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

const BasePath = "/app/config"

func main() {
	err := ProcessFile(BasePath, "dev")
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func ProcessFileHardToTest(environment string) error {
	filePath := fmt.Sprintf("%s/%s.yaml", BasePath, environment)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	fmt.Print(string(data))

	return nil
}

// Refactor the code.
//
// To make it more testable, you can refactor the code as follows:
//
// Remove the constant BasePath from the function's internals.
// Accept the base path as a parameter in ProcessFile instead.

func ProcessFile(basePath, environment string) error {
	filePath := fmt.Sprintf("%s/%s.yaml", basePath, environment)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Processing data.
	s := string(data)
	fmt.Println(s)

	return nil
}

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
