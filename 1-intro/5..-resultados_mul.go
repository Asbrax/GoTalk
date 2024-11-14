package main

import "fmt"

func swap(x, y string) (string, string, int) {
	var num = 22
	return y, x, num
}

func main() {
	a, b, c := swap("hola", "mundo")
	fmt.Println(a, b, c)
}
