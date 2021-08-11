package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %.2f", e)
}

func SqrtHandlingNegatives(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(-2)
	}
	z := x
	for {
		tmp := z
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-tmp) < 1e-12 {
			break
		}
	}
	return z, nil
}

func main() {
	fmt.Println(SqrtHandlingNegatives(2))
	fmt.Println(SqrtHandlingNegatives(-2))
}
