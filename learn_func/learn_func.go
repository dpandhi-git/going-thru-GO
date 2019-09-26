package main

import "fmt"

func main() {
	//fmt.Println(maniof2(2, 3))

	arr := [2]int{2, 3}
	fmt.Println(arr)
}

func maniof2(a int, b int) (int, int) {
	return a + b, a - b
}
