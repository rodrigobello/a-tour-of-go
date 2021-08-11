package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	dySlice := make([][]uint8, dy)
	dxSlice := make([]uint8, dx)

	for y := range dySlice {
		for x := range dxSlice {
			dxSlice[x] = uint8((x + y) / 2)
		}
		dySlice[y] = dxSlice
	}
	return dySlice
}

func runSlicesExercise() {
	pic.Show(Pic)
}
