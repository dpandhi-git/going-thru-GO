package main

import "fmt"

type Foo struct {
	a int
	b string
}

func main() {
	f := Foo{10, "hi"}

	var f3 *Foo = &f
	f3.a = 20

	fmt.Println(f)
	fmt.Println(*f3)

}
