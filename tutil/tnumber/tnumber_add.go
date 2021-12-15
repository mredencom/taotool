package tnumber

type Add interface {
	AddF32ToF32(p1 float32, p2 float32) float64
	AddF32ToF64(p1 float32, p2 float64) float64
}
