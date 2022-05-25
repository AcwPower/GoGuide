package main

import "fmt"

func printslice(s []int) {
	fmt.Printf("len:%d,cap:%d,%v\n", len(s), cap(s), s)
}
func main() {
	var s []int
	printslice(s)

	s=append(s,1)
	printslice(s)

	s=append(s, 2,3,4,7)
	printslice(s)
}
