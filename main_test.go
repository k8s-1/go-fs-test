package main

import (
	"testing"
)

func TestProcessFile(t *testing.T) {
	// Define the test base path and environment
	basePath := "testdata/config"
	env := "dev"

	// Call ProcessFile using the testdata directory
	err := ProcessFile(basePath, env)
	if err != nil {
		t.Fatalf("ProcessFile returned an error: %v", err)
	}
}
