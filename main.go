package main

import (
	"fmt"
	"log"
	"os"
  "path/filepath"
)

const BasePath = "config"

func main() {
	_, err := ProcessFile(BasePath, "dev")
	if err != nil {
		log.Fatalf("%v", err)
	}
}

// Using a global constant inside a function makes
// testing difficult. So don't.
func ProcessFileHardToTest(environment string) error {
	filePath := filepath.Join(BasePath, environment)
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

func ProcessFile(basePath, environment string) (string, error) {
  f := environment + ".cue"
	filePath := filepath.Join(BasePath, f)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	// Processing data.
	s := string(data)
	fmt.Println(s)

	return s, nil
}
