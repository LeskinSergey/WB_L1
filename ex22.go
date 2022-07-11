package main

import (
	"fmt"
	"math/big"
)

//Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
func main() {
	var n1, n2 string
	fmt.Println("введите первое число :")
	fmt.Scanln(&n1)
	fmt.Println("введите второе число :")
	fmt.Scanln(&n2)
	result := new(big.Int)
	a, ok1 := new(big.Int).SetString(n1, 10)
	b, ok2 := new(big.Int).SetString(n2, 10)
	if !ok1 || !ok2 {
		fmt.Println("Error args")
	}
	fmt.Println("Сложение: ", result.Add(a, b))
	fmt.Println("Вычитание: ", result.Sub(a, b))
	fmt.Println("Деление: ", result.Div(a, b))
	fmt.Println("Умножение: ", result.Mul(a, b))
}
