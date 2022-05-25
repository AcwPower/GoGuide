package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ss := make([][]uint8, dy)
	for i := range ss {
		s := make([]uint8, dx)
		for j := range s {
			s[j] = uint8(dx * dy)
		}
		ss[i] = s
	}
	return ss
}

func main() {
	pic.Show(Pic)
}
