package io

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"math"
	"unsafe"
)

var strDecoder = charmap.Windows1252.NewDecoder()
var strEncoder = charmap.Windows1252.NewEncoder()

// ----

type Packet struct {
	Data   []int8
	Pos    int32
	BitPos int32
}

func NewPacket(len int32) *Packet {
	return &Packet{
		Data: make([]int8, len),
	}
}

func FromBytes(bytes []int8) *Packet {
	return &Packet{
		Data: bytes,
	}
}

// ----

func (p *Packet) G1() int32 {
	p.Pos += 1
	return int32(p.Data[p.Pos-1]) & math.MaxUint8
}

func (p *Packet) G1S() int32 {
	p.Pos += 1
	return int32(p.Data[p.Pos-1])
}

func (p *Packet) G2() int32 {
	return (p.G1() << 8) | p.G1()
}

func (p *Packet) G2S() int32 {
	return (p.G1S() << 8) | p.G1S()
}

func (p *Packet) IG2() int32 {
	return p.G1() | (p.G1() << 8)
}

func (p *Packet) G3() int32 {
	return (p.G1() << 16) | p.G2()
}

func (p *Packet) IG3() int32 {
	return p.IG2() | (p.G1() << 16)
}

func (p *Packet) G4() int32 {
	return (p.G1() << 24) | p.G3()
}

func (p *Packet) IG4() int32 {
	return p.IG3() | (p.G1() << 24)
}

func (p *Packet) G8() int64 {
	return (int64(p.G4()) & math.MaxUint32 << 32) | (int64(p.G4()) & math.MaxUint32)
}

func (p *Packet) GSTR(terminator byte) string {
	start := p.Data[p.Pos:]
	dst := *(*[]byte)(unsafe.Pointer(&start))
	index := bytes.IndexByte(dst, terminator)
	data := dst[:index]
	str, _ := strDecoder.String(unsafe.String(&data[0], len(data)))
	p.Pos += int32(index + 1)
	return str
}

func (p *Packet) GDATA(dst []int8, offset int32, length int32) {
	copy(dst[offset:offset+length], p.Data[p.Pos:p.Pos+length])
	p.Pos += length
}

// ----

func (p *Packet) P1(val int32) {
	p.Pos += 1
	p.Data[p.Pos-1] = int8(val)
}

func (p *Packet) P2(val int32) {
	p.P1(val >> 8)
	p.P1(val)
}

func (p *Packet) IP2(val int32) {
	p.P1(val)
	p.P1(val >> 8)
}

func (p *Packet) P3(val int32) {
	p.P1(val >> 16)
	p.P2(val)
}

func (p *Packet) IP3(val int32) {
	p.IP2(val)
	p.P1(val >> 16)
}

func (p *Packet) P4(val int32) {
	p.P1(val >> 24)
	p.P3(val)
}

func (p *Packet) IP4(val int32) {
	p.IP3(val)
	p.P1(val >> 24)
}

func (p *Packet) P8(val int64) {
	p.P4(int32((val >> 32) & math.MaxUint32))
	p.P4(int32(val & math.MaxUint32))
}

func (p *Packet) PSTR(str string, terminator byte) {
	start := p.Data[p.Pos:]
	n, _, _ := strEncoder.Transform(
		*(*[]byte)(unsafe.Pointer(&start)),
		unsafe.Slice(unsafe.StringData(str), len(str)),
		true,
	)
	p.Pos += int32(n)
	p.P1(int32(terminator))
}

func (p *Packet) PDATA(src []int8, offset int32, length int32) {
	copy(p.Data[p.Pos:], src[offset:offset+length])
	p.Pos += length
}

// ----

func (p *Packet) PSIZE4(len int32) {
	p.Data[p.Pos-len-4] = int8(len >> 24)
	p.Data[p.Pos-len-3] = int8(len >> 16)
	p.Data[p.Pos-len-2] = int8(len >> 8)
	p.Data[p.Pos-len-1] = int8(len)
}

func (p *Packet) PSIZE2(len int32) {
	p.Data[p.Pos-len-2] = int8(len >> 8)
	p.Data[p.Pos-len-1] = int8(len)
}

func (p *Packet) PSIZE1(len int32) {
	p.Data[p.Pos-len-1] = int8(len)
}

// ----

func (p *Packet) GSMARTS() int32 {
	if uint8(p.Data[p.Pos]) < 128 {
		return p.G1() - 64
	}
	return p.G2() - 0xc000
}

func (p *Packet) PSMARTS(val int32) {
	if val < 64 && val >= -64 {
		p.P1(val + 64)
	} else if val < 16384 && val >= -16384 {
		p.P2(val + 0xc000)
	} else {
		panic(fmt.Sprintf("Error PSMARTS out of range: %v", val))
	}
}

func (p *Packet) GSMART() int32 {
	if uint8(p.Data[p.Pos]) < 128 {
		return p.G1()
	}
	return p.G2() - 0x8000
}

func (p *Packet) PSMART(val int32) {
	if val >= 0 && val < 128 {
		p.P1(val)
	} else if val >= 0 && val < 32768 {
		p.P2(val + 0x8000)
	} else {
		panic(fmt.Sprintf("Error PSMART out of range: %v", val))
	}
}

// ----

func (p *Packet) BITS() {
	p.BitPos = p.Pos << 3
}

func (p *Packet) BYTES() {
	p.BitPos = (p.Pos + 7) >> 3
}
