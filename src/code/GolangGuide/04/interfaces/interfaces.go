package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type myfloat float64

func (f myfloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
func main() {
	var a Abser
	v := Vertex{3, 4}
	f := myfloat(-math.Sqrt(2))

	a = f
	a = &v
	fmt.Println(v, a.Abs())
}
