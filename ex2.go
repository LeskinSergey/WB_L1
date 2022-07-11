package main

import (
	"fmt"
	"sync"
)

func square(a int, wg *sync.WaitGroup) {
	fmt.Println(a * a)
	wg.Done()
}

//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
//и выведет их квадраты в stdout.
func main() {
	var nums = []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	for i := range nums {
		wg.Add(1)
		go square(nums[i], &wg)
	}
	wg.Wait()
}
