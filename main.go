package main

import (
	"fmt"
	"log"
	"os"
)

const BasePath = "/app/config"

func main() {
	_, err := ProcessFile(BasePath, "dev")
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func ProcessFileHardToTest(environment string) error {
	filePath := fmt.Sprintf("%s/%s.cue", BasePath, environment)
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
	filePath := fmt.Sprintf("%s%s.cue", basePath, environment)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	// Processing data.
	s := string(data)
	fmt.Println(s)

	return s, nil
}
