package io

import "testing"

func TestGnomeballButtonsHash(t *testing.T) {
	result := Hash("gnomeball_buttons.dat")
	if result != 22834782 {
		t.Errorf("TestGnomeballButtons expected 22834782, got %d", result)
	}
}

func TestHeadiconsHash(t *testing.T) {
	result := Hash("headicons.dat")
	if result != -288954319 {
		t.Errorf("TestHeadicons expected -288954319, got %d", result)
	}
}

func TestHitmarksHash(t *testing.T) {
	result := Hash("hitmarks.dat")
	if result != -1502153170 {
		t.Errorf("TestHitmarks expected -1502153170, got %d", result)
	}
}
