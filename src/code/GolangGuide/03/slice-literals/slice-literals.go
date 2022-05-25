package main

import "fmt"

func main() {
	a := []int{1, 3, 5, 7, 11, 13}
	b := []string{"rossbao", "jetliao", "xiaozeng"}
	c := []bool{true, false, true, false}
	s := []struct {
		a int
		s string
	}{
		{1, "a"},
		{2, "b"},
		{3, "c"},
	}

	 slipt:= []int{}
	if slipt == nil {
		fmt.Println("nil!")
	}
	fmt.Println(a, b, c, s, slipt)
}
