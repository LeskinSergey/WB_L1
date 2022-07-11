package main

import (
	"fmt"
	"unicode/utf8"
)

func ReversStr(s string) string {
	TmpS := []rune(s)
	res := make([]rune, 0)
	for i := utf8.RuneCountInString(s) - 1; i >= 0; i-- {
		res = append(res, TmpS[i])
	}
	return string(res)
}

//Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
//Символы могут быть unicode.
func main() {
	var s string
	fmt.Scanln(&s)
	res := ReversStr(s)
	fmt.Println(res)
}
