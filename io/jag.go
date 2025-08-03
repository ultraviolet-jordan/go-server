package io

import (
	"fmt"
	"strings"
)

// ----

type JagFile struct {
	FileCount   int32
	FileHashes  []int32
	FileUnpacks []int32
	FilePacks   []int32
	FileOffsets []int32
	Data        []int8
	Unpacked    bool
}

func Hash(name string) int32 {
	hash := int32(0)
	upper := strings.ToUpper(name)
	for i := range len(upper) {
		hash = hash*61 + int32(upper[i]) - 32
	}
	return hash
}

func NewJagFile(buf *Packet) (*JagFile, error) {
	// buf := FromBytes(bytes)
	unpacked := buf.G3()
	packed := buf.G3()

	decompressed := false

	if packed != unpacked {
		decompressedData, err := Bz2Decompress(buf.Data, unpacked, true, 6)
		if err != nil {
			return nil, err
		}
		buf = FromBytes(decompressedData)
		decompressed = true
	}

	fileCount := buf.G2()
	fileHashes := make([]int32, fileCount)
	fileUnpacks := make([]int32, fileCount)
	filePacks := make([]int32, fileCount)
	fileOffsets := make([]int32, fileCount)

	pos := buf.Pos + fileCount*10
	for index := range fileCount {
		fileHashes[index] = buf.G4()
		fileUnpacks[index] = buf.G3()
		filePacks[index] = buf.G3()
		fileOffsets[index] = pos
		pos += filePacks[index]
	}

	return &JagFile{
		FileCount:   fileCount,
		FileHashes:  fileHashes,
		FileUnpacks: fileUnpacks,
		FilePacks:   filePacks,
		FileOffsets: fileOffsets,
		Data:        buf.Data,
		Unpacked:    decompressed,
	}, nil
}

// ----

func (j *JagFile) Read(name string) (*Packet, error) {
	hash := Hash(name)
	for i, fileHash := range j.FileHashes {
		if fileHash == hash {
			return j.Get(int32(i))
		}
	}
	return nil, fmt.Errorf("file not found: %s", name)
}

func (j *JagFile) Get(index int32) (*Packet, error) {
	if index < 0 || index >= j.FileCount {
		return nil, fmt.Errorf("index out of bounds: %d", index)
	}

	if int(j.FileOffsets[index]) >= len(j.Data) {
		return nil, fmt.Errorf("offset out of data bounds: %d", j.FileOffsets[index])
	}

	start := j.FileOffsets[index]
	end := start + j.FilePacks[index]
	if int(end) > len(j.Data) {
		return nil, fmt.Errorf("file end out of bounds: %d", end)
	}

	fileData := j.Data[start:end]

	if j.Unpacked {
		return FromBytes(fileData), nil
	}

	decompressed, err := Bz2Decompress(fileData, j.FileUnpacks[index], true, 0)
	if err != nil {
		return nil, err
	}

	return FromBytes(decompressed), nil
}
