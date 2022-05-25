package main

import (
	"fmt"
)


type vertex struct{
	a,b int
}

var m =map[string]vertex{
	"abao":{100,200},
	"bbao":{300,400},
}

func main(){
	fmt.Println(m["abao"])
}