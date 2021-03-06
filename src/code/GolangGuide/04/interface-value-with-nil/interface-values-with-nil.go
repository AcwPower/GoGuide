package main

import "fmt"

type I interface {
	M()
}
type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}
func describe(i I) {
	fmt.Printf("%v %T\n", i, i)
}
func main() {
	var i I
	var t *T
	i = t
	i.M()
	describe(i)

	t = &T{"hello nil"}
	i = t
	i.M()
	describe(i)
}
