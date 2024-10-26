package main

import (
	"log"
	"os/exec"
	"testing"
)

func TestProcessFile(t *testing.T) {
	basePath := "testdata/versions"
	env := "dev"
	_, err := ProcessFile(basePath, env)
	if err != nil {
		t.Fatalf("ProcessFile returned an error: %v", err)
	}
}

func teardown() {
	cmd := exec.Command("git", "restore", "testdata")
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to restore testdata: %v\n", err)
	}
}
