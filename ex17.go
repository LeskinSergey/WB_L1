package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, item int) int {
	left := 0 // определяем границы поиска
	right := len(arr) - 1

	for left <= right { // пока границы не сократились до одного
		mid := left + right // проверяем середину
		guess := arr[mid]
		if guess == item { // если нашли возвращаем индекс
			return mid
		}
		if guess > item { // если больше искомого числа
			right = mid - 1
		} else { // если меньше
			left = mid + 1
		}
	}
	return -1 // не нашли число
}

//Реализовать бинарный поиск встроенными методами языка.
func main() {
	var n int
	fmt.Scanln(&n)
	nums := make([]int, n)
	for i := range nums {
		fmt.Scan(&nums[i])
	}
	// сортируем его (обязательное условие для бинарного поиска)
	sort.Ints(nums)
	//вводим число, которое будем искать
	var find int
	fmt.Scanln(&find)
	index := binarySearch(nums, find)
	if index == -1 {
		fmt.Println("Number not found")
	} else {
		fmt.Println(nums, "number", nums[index], "at index", index)
	}
}
