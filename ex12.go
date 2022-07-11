package main

import "fmt"

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
func main() {
	s := []string{"cat", "dog", "tree", "cat", "fish", "cat"}
	m := make(map[string]int)
	res := make([]string, 0)
	for _, val := range s {
		m[val] += 1
		if m[val] == 1 {
			res = append(res, val)
		}
	}
	fmt.Println("s =", s)
	fmt.Println("result =", res)
}
