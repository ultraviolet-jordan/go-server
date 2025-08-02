package io

import "math"

// CRC32B Reversed CRC-32 polynomial for Cyclic Redundancy Check (CRC).
// This is sometimes referred to as CRC32B.
const (
	CRCRange = 256
	CRC32B   = int32(-306674912) // 0xEDB88320
)

// ----

type CRC struct {
	table [CRCRange]int32
}

func NewCRC() *CRC {
	crc := &CRC{}
	for b := range CRCRange {
		remainder := int32(b)
		for range 8 {
			if remainder&0x1 == 1 {
				remainder = int32(uint32(remainder)>>1) ^ CRC32B
			} else {
				remainder = int32(uint32(remainder) >> 1)
			}
		}
		crc.table[b] = remainder
	}
	return crc
}

// ----

func (c *CRC) GetCRC(src []int8, offset int, length int) int32 {
	crc := int32(-1)
	for i := offset; i < length; i++ {
		crc = (int32(uint32(crc) >> 8)) ^ c.table[(crc^int32(src[i]))&math.MaxUint8]
	}
	return ^crc
}
