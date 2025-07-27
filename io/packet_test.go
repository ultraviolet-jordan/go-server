package io

import (
	"reflect"
	"testing"
)

func TestPacket1(t *testing.T) {
	packet := &Packet{
		Data: make([]int8, 1),
	}
	packet.P1(255)
	packet.Pos = 0
	result := packet.G1()
	if result != 255 {
		t.Errorf("TestPacket1 expected 255, got %d", result)
	}
}

func TestPacket1S(t *testing.T) {
	packet := &Packet{
		Data: make([]int8, 1),
	}
	packet.P1(255)
	packet.Pos = 0
	result := packet.G1S()
	if result != -1 {
		t.Errorf("TestPacket1S expected -1, got %d", result)
	}
}

func TestPacket2(t *testing.T) {
	packet := &Packet{
		Data: make([]int8, 2),
	}
	packet.P2(65535)
	packet.Pos = 0
	result := packet.G2()
	if result != 65535 {
		t.Errorf("TestPacket2 expected 65535, got %d", result)
	}
}

func TestPacket2S(t *testing.T) {
	packet := &Packet{
		Data: make([]int8, 2),
	}
	packet.P2(65535)
	packet.Pos = 0
	result := packet.G2S()
	if result != -1 {
		t.Errorf("TestPacket2S expected -1, got %d", result)
	}
}

func TestPacketI2(t *testing.T) {
	packet := &Packet{
		Data: make([]int8, 2),
	}
	packet.IP2(65535)
	packet.Pos = 0
	result := packet.IG2()
	if result != 65535 {
		t.Errorf("TestPacketI2 expected 65535, got %d", result)
	}
}

func TestPacket3(t *testing.T) {
	packet := &Packet{
		Data: make([]int8, 3),
	}
	packet.P3(16777215)
	packet.Pos = 0
	result := packet.G3()
	if result != 16777215 {
		t.Errorf("TestPacket3 expected 16777215, got %d", result)
	}
}

func TestPacketI3(t *testing.T) {
	packet := &Packet{
		Data: make([]int8, 3),
	}
	packet.IP3(16777215)
	packet.Pos = 0
	result := packet.IG3()
	if result != 16777215 {
		t.Errorf("TestPacketI3 expected 16777215, got %d", result)
	}
}

func TestPacket4(t *testing.T) {
	packet := &Packet{
		Data: make([]int8, 4),
	}
	packet.P4(2147483647)
	packet.Pos = 0
	result := packet.G4()
	if result != 2147483647 {
		t.Errorf("TestPacket4 expected 2147483647, got %d", result)
	}
}

func TestPacketI4(t *testing.T) {
	packet := &Packet{
		Data: make([]int8, 4),
	}
	packet.IP4(2147483647)
	packet.Pos = 0
	result := packet.IG4()
	if result != 2147483647 {
		t.Errorf("TestPacketI4 expected 2147483647, got %d", result)
	}
}

func TestPacket8(t *testing.T) {
	packet := &Packet{
		Data: make([]int8, 8),
	}
	packet.P8(9223372036854775807)
	packet.Pos = 0
	result := packet.G8()
	if result != 9223372036854775807 {
		t.Errorf("TestPacket8 expected 9223372036854775807, got %d", result)
	}
}

func BenchmarkPacket8(b *testing.B) {
	packet := &Packet{
		Data: make([]int8, 8),
	}

	b.ResetTimer()

	for b.Loop() {
		packet.P8(9223372036854775807)
		packet.Pos = 0
		packet.G8()
		packet.Pos = 0
	}
}

func TestPacketSTR(t *testing.T) {
	str := "Hello World!"
	packet := &Packet{
		Data: make([]int8, len(str)+1),
	}
	packet.PSTR(str, 10)
	packet.Pos = 0
	result := packet.GSTR(10)
	if result != str {
		t.Errorf("TestPacketSTR expected Hello World!, got %s", result)
	}
}

func BenchmarkPacketSTR(b *testing.B) {
	str := "Hello World!"
	packet := &Packet{
		Data: make([]int8, len(str)+1),
	}

	b.ResetTimer()

	for b.Loop() {
		packet.PSTR(str, 10)
		packet.Pos = 0
		packet.GSTR(10)
		packet.Pos = 0
	}
}

func TestPacketDATA(t *testing.T) {
	packet := &Packet{
		Data: make([]int8, 3),
	}

	packet.PDATA([]int8{1, 2, 3, 4, 5}, 1, 3)
	packet.Pos = 0
	dst := make([]int8, 3)
	packet.GDATA(dst, 0, int32(len(dst)))
	if !reflect.DeepEqual(packet.Data, dst) {
		t.Errorf("TestPacketDATA expected [2, 3, 4], got %v", dst)
	}
}
