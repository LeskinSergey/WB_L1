package main

import "fmt"

//Реализовать пересечение двух неупорядоченных множеств.
func main() {
	a1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a2 := []int{11, 23, 13, 24, 5, 6, 7, 8, 9}
	a3 := append(a1, a2...)
	m := make(map[int]int)
	intersection := make([]int, 0)
	for _, value := range a3 {
		_, ok := m[value]
		if !ok {
			m[value] = 1
		} else {
			m[value]++
		}
	}
	for key, val := range m {
		if val > 1 {
			intersection = append(intersection, key)
		}
	}
	fmt.Println(intersection)
}
