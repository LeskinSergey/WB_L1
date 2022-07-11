package main

import (
	"fmt"
	"time"
)

//Реализовать собственную функцию sleep.
func main() {
	var t int
	fmt.Scanln(&t)
	MySleep(t)
	fmt.Println("прошло", t, "секунд, заканчиваем программу")
}

func MySleep(t int) {
	<-time.After(time.Duration(t) * time.Second)
}
