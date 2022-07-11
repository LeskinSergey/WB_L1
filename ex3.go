package main

import (
	"fmt"
	"sync"
)

func SumSquare(a int, wg *sync.WaitGroup, s *int) {
	*s += a * a
	wg.Done()
}

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….)
//с использованием конкурентных вычислений.
func main() {
	var nums = []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	var sum int
	for i := range nums {
		wg.Add(1)
		go SumSquare(nums[i], &wg, &sum)
	}
	wg.Wait()
	fmt.Println(sum)
}
