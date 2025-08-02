package math

import (
	"math"
	"testing"
)

func TestTrigRadians90(t *testing.T) {
	result := Radians(90)
	if result != 0.03451456772742693 {
		t.Errorf("TestTrigRadians90 expected 0.03451456772742693, got %f", result)
	}
}

func TestTrigRadians360(t *testing.T) {
	result := Radians(360)
	if result != 0.1380582709097077 {
		t.Errorf("TestTrigRadians360 expected 0.1380582709097077, got %f", result)
	}
}

func TestTrigAtan26144(t *testing.T) {
	result := Atan2(1, -1)
	if result != 6144 {
		t.Errorf("TestTrigAtan26144 expected 6144, got %d", result)
	}
}

func TestTrigAtan212288(t *testing.T) {
	result := Atan2(-1, 0)
	if result != 12288 {
		t.Errorf("TestTrigAtan212288 expected 12288, got %d", result)
	}
}

func TestTrigSin(t *testing.T) {
	trig := NewTrig()
	result := trig.Sin(int32(math.Round(math.Pi)))
	if result != 18 {
		t.Errorf("TestTrigSin expected 18, got %d", result)
	}
}

func TestTrigCos(t *testing.T) {
	trig := NewTrig()
	result := trig.Cos(int32(math.Round(math.Pi)))
	if result != 16383 {
		t.Errorf("TestTrigCos expected 16383, got %d", result)
	}
}
