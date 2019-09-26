package main

import "fmt"

func main() {
	var m = map[string]int{
		"hello":   1,
		"goodbye": 2,
		"clap":    3,
	}

	fmt.Println("default", m)
	for v, a := range m {
		fmt.Println(v, a)
	}

	if _, ok := m["dk"]; ok {
		fmt.Println("valid")

	} else {
		fmt.Println("No")
	}
}
