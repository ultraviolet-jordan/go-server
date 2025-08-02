package math

const (
	BitsRange   = 33
	BitsPairs   = 0x55555555
	BitsQuads   = 0x33333333
	BitsNibbles = 0x0f0f0f0f
	BitsSum     = 0x01010101
)

// ----

type Bits struct {
	masks [BitsRange]int32
}

func NewBits() *Bits {
	bits := &Bits{}
	incr := int32(2)
	for i := 1; i < BitsRange; i++ {
		bits.masks[i] = incr - 1
		incr += incr
	}
	return bits
}

func BitCount(num int32) int32 {
	one := num - ((num >> 1) & BitsPairs)
	two := (one & BitsQuads) + ((one >> 2) & BitsQuads)
	return (((two + (two >> 4)) & BitsNibbles) * BitsSum) >> 24
}

// ----

func (b *Bits) SetBitRange(num int32, start int32, end int32) int32 {
	return num | (b.masks[end-start+1] << start)
}

func (b *Bits) SetBitRangeToInt(num int32, value int32, start int32, end int32) int32 {
	cleared := b.ClearBitRange(num, start, end)
	max := b.masks[end-start+1]
	if value > max {
		value = max
	}
	return cleared | (value << start)
}

func (b *Bits) ClearBitRange(num int32, start int32, end int32) int32 {
	return num & ^(b.masks[end-start+1] << start)
}
