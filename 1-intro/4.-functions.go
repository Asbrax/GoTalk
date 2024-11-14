package main

import "fmt"

//func add(x, y int) int {
func suma(x, y int) string {

	fmt.Println(y + x)
	return "Hola Huichapan"
}

func main() {
	fmt.Println(suma(10, 15))
}
