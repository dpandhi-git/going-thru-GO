package main

import "fmt"

func main() {
	var a byte = 8
	var b int8 = 20
	var c bool = false
	fmt.Println(a + byte(b))
	fmt.Println(c)
}
