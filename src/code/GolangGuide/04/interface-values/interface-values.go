package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type myfloat float64

func (f myfloat) M() {
	fmt.Println(f)
}

func describe(i I)  {
	fmt.Printf("%v  %T\n",i,i)
}
func main() {
	var i I
	i=&T{"hello world"}
	i.M()
	describe(i)

	i=myfloat(math.Pi)
	i.M()
	describe(i)
}
