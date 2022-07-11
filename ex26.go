package main

import (
	"fmt"
	"strings"
)

func CheckUnique(s string) bool {
	unique := make(map[string]bool)
	m := strings.Split(s, "")
	for _, val := range m {
		_, ok := unique[val]
		if !ok {
			unique[val] = true
		} else {
			return false
		}
	}
	return true
}

//Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
//Функция проверки должна быть регистронезависимой.
//
//Например:
//abcd — true
//abCdefAaf — false
//	aabcd — false
func main() {
	var s string
	fmt.Scanln(&s)
	r := CheckUnique(s)
	fmt.Println(s, r)
}
