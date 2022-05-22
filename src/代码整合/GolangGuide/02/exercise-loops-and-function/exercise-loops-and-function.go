package main

import (
	"fmt"
)

func Sqrt(x float64)  {
	var z float64 = 1.0
	for i:=0; i<=int(x)/2; i++ {
		z+=0.1
		if z*z >=x {
			fmt.Println(z)
			break
		}
	}
}

func main() {
	Sqrt(1252)
}
