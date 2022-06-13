package main

import "fmt"

type IPAddr [4]byte

// TODO: 给 IPAddr 添加一个 "String() string" 方法
func String(ip IPAddr) {
	for i:=0;i<3;i++{
		fmt.Print(ip[i])
		fmt.Print(".")
	}
		fmt.Print(ip[3])
}
func main(){
	
}