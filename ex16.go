package main

import (
	"fmt"
	"sort"
)

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка
func main() {
	var n int
	fmt.Scanln(&n)
	nums := make([]int, n)
	for i := range nums {
		fmt.Scan(&nums[i])
	}
	sort.Ints(nums)
	fmt.Println(nums)
}
