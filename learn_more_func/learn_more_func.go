package main

import "fmt"

func makeAdd(b int) func(int) int {
	return func(c int) int {
		fmt.Println("inside1", b, "inside c", c)
		return c + b
	}
}

func makeDouble(f func(int) int) func(int) int {
	return func(a int) int {
		b := f(a)
		fmt.Println("double ", a, "double b", b)
		return b * 2
	}
}

func main() {
	//addOne := makeAdd(1)
	//fmt.Println(makeAdd(2))
	addTwo := makeAdd(2)
	doubletwo := makeDouble(addTwo)

	fmt.Println(doubletwo(2))
}

/*
func stringadd1(a string) string {
	return a + "abc"
}

func stringadd2(a string) string {
	return a + "bcd"
}

func main() {
	arr := [2]string{"hello", "world"}
	printstring := func(arr [2]string, f func(string) string) {

		fmt.Println(f(arr[0]))
	}
	printstring(arr, stringadd1)
}
*/
