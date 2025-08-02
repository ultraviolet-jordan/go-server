package io

import "testing"

func TestIsaacNoSeed(t *testing.T) {
	isaac := NewIsaac([4]int32{})
	for index := 0; index < 1_000_000; index++ {
		isaac.Next()
	}
	result := isaac.Next()
	if result != 1536048213 {
		t.Errorf("TestIsaac0 expected 1536048213, got %d", result)
	}
}

func TestIsaacSeed(t *testing.T) {
	isaac := NewIsaac([4]int32{1, 2, 3, 4})
	for index := 0; index < 1_000_000; index++ {
		isaac.Next()
	}
	result := isaac.Next()
	if result != -107094133 {
		t.Errorf("TestIsaac0 expected -107094133, got %d", result)
	}
}
