package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x
	for {
		tmp := z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
		if math.Abs(z-tmp) < 1e-12 {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(12))
}
