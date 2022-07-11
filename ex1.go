package main

import "fmt"

type Human struct {
	a int
	b int
}
type Action struct {
	Human
}

func (h Human) Summa() {
	fmt.Println(h.a + h.b)
}

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
func main() {
	var res Action
	res.a = 3
	res.b = 4
	res.Summa()
}
