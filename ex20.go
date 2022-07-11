package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RevSentence(s string) string {
	s = strings.Trim(s, "\n")
	tmp := strings.Split(s, " ")
	var res string
	for _, i := range tmp {
		res = i + " " + res
	}
	return res
}

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».
func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	res := RevSentence(s)
	fmt.Println(res)
}
