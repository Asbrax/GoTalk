package main

import "fmt"

func swap(x, y string) (string, string, int) {
	return y, x ,5
}

func main() {
	a, b , c:= swap("hola", "mundo")
	fmt.Println(a, b, c)
}
