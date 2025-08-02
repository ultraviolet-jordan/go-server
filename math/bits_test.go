package math

import "testing"

func TestBitCount(t *testing.T) {
	result := BitCount(15) // 15 in binary is 1111, so result is 4
	if result != 4 {
		t.Errorf("Expected 4, got %d", result)
	}
}

func TestSetBitRange(t *testing.T) {
	bits := NewBits()
	result := bits.SetBitRange(0, 1, 3) // Sets bits 1 to 3 in 0, resulting in 14 (1110 in binary)
	if result != 14 {
		t.Errorf("Expected 14, got %d", result)
	}
}

func TestSetBitRangeToInt(t *testing.T) {
	bits := NewBits()
	result := bits.SetBitRangeToInt(0, 3, 1, 3) // Sets bits 1 to 3 to match 3, resulting in 6
	if result != 6 {
		t.Errorf("Expected 6, got %d", result)
	}
}

func TestClearBitRange(t *testing.T) {
	bits := NewBits()
	result := bits.ClearBitRange(15, 1, 3) // Clears bits 1 to 3 in 15, resulting in 1
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}
