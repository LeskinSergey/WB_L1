package main

import "fmt"

//Дана переменная int64. Разработать программу, которая устанавливает i-й бит в 1 или 0.
func main() {
	var num, i, ZeroOne int64
	fmt.Scanln(&num, &i, &ZeroOne)
	fmt.Println("before", num)
	if i == 1 {
		num |= 1 << i
	} else {
		num &^= 1 << i
	}
	fmt.Println("after", num)
}
