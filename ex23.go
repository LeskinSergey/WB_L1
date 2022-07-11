package main

import "fmt"

func removeElementByIndex(a []int, i int) []int {
	return append(a[:i], a[i+1:]...)
}

//Удалить i-ый элемент из слайса.
func main() {
	var n, i int
	fmt.Scanln(&n)
	nums := make([]int, n)
	for i := range nums {
		fmt.Scan(&nums[i])
	}
	fmt.Println("какой элемент удалить?")
	fmt.Scanln(&i)
	fmt.Println(removeElementByIndex(nums, i-1))
}
