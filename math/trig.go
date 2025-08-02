package math

import "math"

const (
	Tau                    = 2 * math.Pi
	DegreesRange           = 16384
	DegreesMask            = DegreesRange - 1
	RadiansToDegreesFactor = float64(DegreesRange) / Tau
	TrigRange              = math.Pi / (float64(DegreesRange) / 2)
)

// ----

type Trig struct {
	sin [DegreesRange]int32
	cos [DegreesRange]int32
}

func NewTrig() *Trig {
	trig := &Trig{}
	for i := range DegreesRange {
		s, c := math.Sincos(float64(i) * TrigRange)
		trig.sin[i] = int32(s * DegreesRange)
		trig.cos[i] = int32(c * DegreesRange)
	}
	return trig
}

func Radians(x int32) float64 {
	return (float64(x&DegreesMask) / DegreesRange) * Tau
}

func Atan2(y int32, x int32) int32 {
	return int32(math.Round(math.Atan2(float64(y), float64(x))*RadiansToDegreesFactor)) & DegreesMask
}

// ----

func (t *Trig) Sin(value int32) int32 {
	return t.sin[value&DegreesMask]
}

func (t *Trig) Cos(value int32) int32 {
	return t.cos[value&DegreesMask]
}
