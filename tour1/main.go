package main

import (
	"fmt"
	"math"
)

const precision = 0.0001

func Sqrt(x float64) (float64, int) {
	mathZ := math.Sqrt(x)
	z := 1.0
	count := 0
	for i := 0; i < 10; i++ {
		count++
		z -= (z * z - x) / (2 * z)
		if math.Abs(z - mathZ) < precision {
			break
		}
	}
	return z, count
}

// Exercise: Loops and Functions
func main() {
	for i := 1; i <= 10; i++ {
		x := float64(i)
		z, count := Sqrt(x)
		m := math.Sqrt(x)
		fmt.Printf("x=%v, count=%v, Sqrt(%v)=%v, math.Sqrt(%v)=%v\n", x, count, x, z, x, m)
	}
}