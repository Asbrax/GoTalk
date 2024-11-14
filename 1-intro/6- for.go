package main

import "fmt"

func main() {

	// The most basic type, with a single condition.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("--------")
	for j := 3; j > 0; j-- {
		fmt.Println(j)
	}

	fmt.Println("--------")
	for {
		fmt.Println("loop")
		break
	}

	fmt.Println("--------")
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
