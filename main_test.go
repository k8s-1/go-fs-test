package main

import (
	"testing"
)

func TestProcessFile(t *testing.T) {
	basePath := "testdata/versions/"
	env := "dev"

	err := ProcessFile(basePath, env)
	if err != nil {
		t.Fatalf("ProcessFile returned an error: %v", err)
	}
}
