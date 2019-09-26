package main

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Println(os.Args[0])
	word := os.Args[1]
	greet := "farewell"
	switch l := len(word); {
	case word == "hello":
		fmt.Println("Good morning")
		fallthrough
	case word == "hru":
		fmt.Println("how are you")
	case word == greet:
	case l == 2:
		fmt.Println("out of my league")
	case word == "nice":
		fmt.Println("Nice to meet you!!")
	default:
		fmt.Println("good bye word length", l)
	}
}
