package main

import (
	"fmt"
)

func FirstToSecond(f, s chan int) {
	for val := range f {
		s <- val * 2
	}
}

//Разработать конвейер чисел.
//Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
//после чего данные из второго канала должны выводиться в stdout.
func main() {
	var n int
	fmt.Scanln(&n)
	InputSlice := make([]int, n)
	for i := range InputSlice {
		fmt.Scan(&InputSlice[i])
	}
	firstCh := make(chan int)
	secondCh := make(chan int)
	defer close(firstCh)
	defer close(secondCh)
	go FirstToSecond(firstCh, secondCh)
	for i := range InputSlice {
		firstCh <- InputSlice[i]
		fmt.Println(<-secondCh)
	}
}
