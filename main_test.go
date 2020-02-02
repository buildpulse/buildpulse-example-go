package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("Expected 1 + 2 to equal 3")
	}
}
