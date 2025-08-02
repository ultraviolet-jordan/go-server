package io

import "math"

const (
	IsaacRatio  = int32(-1640531527) // 0x9E3779B9
	IsaacRange  = 256
	IsaacOffset = IsaacRange - 1
)

// ----

type Isaac struct {
	rsl   [IsaacRange]int32
	mem   [IsaacRange]int32
	count int32
	a     int32
	b     int32
	c     int32
}

func NewIsaac(seed [4]int32) *Isaac {
	isaac := &Isaac{}
	for i := range seed {
		isaac.rsl[i] = seed[i]
	}
	isaac.init()
	return isaac
}

// ----

func (i *Isaac) Next() int32 {
	count := i.count
	i.count--
	if count == 0 {
		i.isaac()
		i.count = IsaacOffset
	}
	return i.rsl[i.count]
}

// ----

func (i *Isaac) init() {
	a := IsaacRatio
	b := IsaacRatio
	c := IsaacRatio
	d := IsaacRatio
	e := IsaacRatio
	f := IsaacRatio
	g := IsaacRatio
	h := IsaacRatio

	for range 4 {
		a ^= b << 11
		d += a
		b += c
		b ^= int32(uint32(c) >> 2)
		e += b
		c += d
		c ^= d << 8
		f += c
		d += e
		d ^= int32(uint32(e) >> 16)
		g += d
		e += f
		e ^= f << 10
		h += e
		f += g
		f ^= int32(uint32(g) >> 4)
		a += f
		g += h
		g ^= h << 8
		b += g
		h += a
		h ^= int32(uint32(a) >> 9)
		c += h
		a += b
	}

	for index := 0; index < IsaacRange; index += 8 {
		a += i.rsl[index]
		b += i.rsl[index+1]
		c += i.rsl[index+2]
		d += i.rsl[index+3]
		e += i.rsl[index+4]
		f += i.rsl[index+5]
		g += i.rsl[index+6]
		h += i.rsl[index+7]

		a ^= b << 11
		d += a
		b += c
		b ^= int32(uint32(c) >> 2)
		e += b
		c += d
		c ^= d << 8
		f += c
		d += e
		d ^= int32(uint32(e) >> 16)
		g += d
		e += f
		e ^= f << 10
		h += e
		f += g
		f ^= int32(uint32(g) >> 4)
		a += f
		g += h
		g ^= h << 8
		b += g
		h += a
		h ^= int32(uint32(a) >> 9)
		c += h
		a += b

		i.mem[index] = a
		i.mem[index+1] = b
		i.mem[index+2] = c
		i.mem[index+3] = d
		i.mem[index+4] = e
		i.mem[index+5] = f
		i.mem[index+6] = g
		i.mem[index+7] = h
	}

	for index := 0; index < IsaacRange; index += 8 {
		a += i.mem[index]
		b += i.mem[index+1]
		c += i.mem[index+2]
		d += i.mem[index+3]
		e += i.mem[index+4]
		f += i.mem[index+5]
		g += i.mem[index+6]
		h += i.mem[index+7]

		a ^= b << 11
		d += a
		b += c
		b ^= int32(uint32(c) >> 2)
		e += b
		c += d
		c ^= d << 8
		f += c
		d += e
		d ^= int32(uint32(e) >> 16)
		g += d
		e += f
		e ^= f << 10
		h += e
		f += g
		f ^= int32(uint32(g) >> 4)
		a += f
		g += h
		g ^= h << 8
		b += g
		h += a
		h ^= int32(uint32(a) >> 9)
		c += h
		a += b

		i.mem[index] = a
		i.mem[index+1] = b
		i.mem[index+2] = c
		i.mem[index+3] = d
		i.mem[index+4] = e
		i.mem[index+5] = f
		i.mem[index+6] = g
		i.mem[index+7] = h
	}

	i.isaac()
	i.count = IsaacRange
}

func (i *Isaac) isaac() {
	i.c++
	i.b += i.c
	for index := range IsaacRange {
		x := i.mem[index]
		switch index & 0x3 {
		case 0:
			i.a ^= i.a << 13
		case 1:
			i.a ^= int32(uint32(i.a) >> 6)
		case 2:
			i.a ^= i.a << 2
		case 3:
			i.a ^= int32(uint32(i.a) >> 16)
		}
		i.a += i.mem[(index+128)&math.MaxUint8]
		y := i.mem[(x>>2)&math.MaxUint8] + i.a + i.b
		i.mem[index] = y
		i.b = i.mem[(y>>10)&math.MaxUint8] + x
		i.rsl[index] = i.b
	}
}
