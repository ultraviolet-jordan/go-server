package io

import "testing"

func TestDecompress(t *testing.T) {
	original := "Hello world!"

	// Convert original string to []int8 inline
	input := make([]int8, len(original))
	for i := range original {
		input[i] = int8(original[i])
	}

	compressed, err := Bz2Compress(input)
	if err != nil {
		t.Fatalf("compression failed: %v", err)
	}

	decompressed, err := Bz2Decompress(compressed, int32(len(original)), false, 0)
	if err != nil {
		t.Fatalf("decompression failed: %v", err)
	}

	if len(decompressed) != len(input) {
		t.Fatalf("decompressed length mismatch: got %d, want %d", len(decompressed), len(input))
	}
	for i := range decompressed {
		if decompressed[i] != input[i] {
			t.Errorf("byte mismatch at index %d: got %v, want %v", i, decompressed[i], input[i])
		}
	}
}

func TestCompress(t *testing.T) {
	original := "Hello world!"

	// Inline conversion from string to []int8
	input := make([]int8, len(original))
	for i := range original {
		input[i] = int8(original[i])
	}

	compressed, err := Bz2Compress(input)
	if err != nil {
		t.Fatalf("compression failed: %v", err)
	}

	decompressed, err := Bz2Decompress(compressed, 12, false, 0)
	if err != nil {
		t.Fatalf("decompression failed: %v", err)
	}

	if len(decompressed) != len(input) {
		t.Fatalf("decompressed length mismatch: got %d, want %d", len(decompressed), len(input))
	}
	for i := range decompressed {
		if decompressed[i] != input[i] {
			t.Errorf("byte mismatch at index %d: got %v, want %v", i, decompressed[i], input[i])
		}
	}
}

func BenchmarkBz2(b *testing.B) {
	original := "Hello world!"

	input := make([]int8, len(original))
	for i := range original {
		input[i] = int8(original[i])
	}

	b.ResetTimer()

	for b.Loop() {
		compressed, _ := Bz2Compress(input)
		_, _ = Bz2Decompress(compressed, int32(len(original)), false, 0)
	}
}
