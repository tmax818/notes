package main

import "fmt"

var z string = "hello scoped world"

func main() {
	var x string = "hello world"
	y := "hello world"
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
