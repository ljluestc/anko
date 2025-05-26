package vm

import (
	"bytes"
	"testing"

	"github.com/mattn/anko/env"
)

func TestExecute(t *testing.T) {
	vm := &VM{}
	err := vm.Execute()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestVMExecute(t *testing.T) {
	e := env.NewEnv()
	vm := &VM{}

	var buf bytes.Buffer
	vm.Stdout = &buf

	// Placeholder script
	script := "a = 1 + 1"
	err := vm.Execute()
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}
