package io

import (
	"bytes"
	"fmt"
	"github.com/dsnet/compress/bzip2"
	"io"
	"unsafe"
)

var Bz2Header = []int8{'B', 'Z', 'h', '1'}

func Bz2Decompress(src []int8, decompressLength int32, prependHeader bool, offset int32) ([]int8, error) {
	var checked []int8
	if prependHeader {
		checked = append(Bz2Header, src[offset:]...)
	} else {
		checked = src
	}

	// Unsafe reinterpret: []int8 → []byte
	//goland:noinspection GoRedundantConversion
	data := unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(checked))), len(checked))

	reader, err := bzip2.NewReader(bytes.NewReader(data), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create bzip2 reader: %w", err)
	}

	decompressedBytes := make([]byte, decompressLength)

	if _, err := io.ReadFull(reader, decompressedBytes); err != nil {
		return nil, fmt.Errorf("bzip2 decompression failed during read: %w", err)
	}

	if err := reader.Close(); err != nil {
		return nil, fmt.Errorf("bzip2 decompression failed during close: %w", err)
	}

	// Unsafe reinterpret: []byte → []int8
	//goland:noinspection GoRedundantConversion
	return unsafe.Slice((*int8)(unsafe.Pointer(unsafe.SliceData(decompressedBytes))), len(decompressedBytes)), nil
}

func Bz2Compress(src []int8) ([]int8, error) {
	// Unsafe reinterpret: []int8 → []byte
	//goland:noinspection GoRedundantConversion
	data := unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(src))), len(src))

	var buf bytes.Buffer
	writer, err := bzip2.NewWriter(&buf, &bzip2.WriterConfig{Level: bzip2.BestCompression})
	if err != nil {
		return nil, fmt.Errorf("failed to create bzip2 writer: %w", err)
	}

	if _, err := writer.Write(data); err != nil {
		return nil, fmt.Errorf("bzip2 compression failed during write: %w", err)
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("bzip2 compression failed during close: %w", err)
	}

	compressedBytes := buf.Bytes()

	// Unsafe reinterpret: []byte → []int8
	//goland:noinspection GoRedundantConversion
	return unsafe.Slice((*int8)(unsafe.Pointer(unsafe.SliceData(compressedBytes))), len(compressedBytes)), nil
}
