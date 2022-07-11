package main

import "fmt"

func CheckType(tmp interface{}) {
	switch tmp.(type) {
	case int:
		fmt.Println("type is int")
	case string:
		fmt.Println("type is string")
	case bool:
		fmt.Println("type is bool")
	case float64:
		fmt.Println("type is float")
	case chan int:
		fmt.Println("type is chan int")
	default:
		fmt.Println("type is none of the above")
	}
}

//Разработать программу, которая в рантайме способна определить тип переменной:
//int, string, bool, channel из переменной типа interface{}.
func main() {
	ch := make(chan int)
	ch1 := make(chan string)
	CheckType(123)
	CheckType("123")
	CheckType(true)
	CheckType(12.3)
	CheckType(ch)
	CheckType(ch1)
}
