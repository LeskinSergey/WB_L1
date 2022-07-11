package main

import "fmt"

//Поменять местами два числа без создания временной переменной.
func main() {
	a := 10
	b := 5
	a, b = b, a
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}
