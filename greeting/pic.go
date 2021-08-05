package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	row := make([]uint8, dx)
	rec := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		row := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			row[x] = uint8(x ^ y)
		}
		rec[y] = row
	}

	return rec
}

func main() {
	pic.Show(Pic)
}
