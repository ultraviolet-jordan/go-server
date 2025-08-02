package io

import "testing"

func TestCRC(t *testing.T) {
	crc := NewCRC()
	result := crc.GetCRC([]int8{1, 2, 3, 4}, 0, 4)
	if result != -1237517363 {
		t.Errorf("TestCRC expected -1237517363, got %d", result)
	}
}
