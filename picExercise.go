package main

import "golang.org/x/tour/pic"

func createRow(x, y int) []uint8 {
	row := make([]uint8, y)
	for i := range row {
		row[i] = uint8(i*(x ^ y)/2)
	}
	return row
}

func Pic(dx, dy int) [][]uint8 {
	v := make([][]uint8, dy)
	for i := range v {
		v[i] = createRow(i, dy)
	}
	return v
}

func main() {
	pic.Show(Pic)
}
