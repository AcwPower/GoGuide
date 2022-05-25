package main

import "fmt"

func main() {
	name := [4]string{
		"rossbao",
		"jetliao",
		"jackeyliu",
		"xiaozeng",
	}
	a :=name[0:2]
	b :=name[1:3]
	fmt.Println(a,b)
	b[0]="atenzou"
	fmt.Println(a,b)
	fmt.Println(name)
}