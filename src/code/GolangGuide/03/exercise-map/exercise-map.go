package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	fmt.Println(strings.Fields(s))
	for i := 0; i < len(strings.Fields(s)); i++ {
		m[strings.Fields(s)[i]] += 1
	}
	return m
}

func main() {

	wc.Test(WordCount)
}
