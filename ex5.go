package main

import (
	"fmt"
	"math/rand"
	"time"
)

func SendData(ch chan int) {
	for {
		ch <- rand.Int()
	}
}
func ReadData(ch chan int) {
	for data := range ch {
		fmt.Println(data)
	}
}

func CountTime(t int, end chan bool) {
	time.AfterFunc(time.Duration(t)*time.Second, func() {
		fmt.Println("time is up!")
		end <- true
	})
}

//Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
//По истечению N секунд программа должна завершаться.
func main() {
	var timer int
	ch := make(chan int)
	end := make(chan bool)
	fmt.Scanln(&timer)
	go SendData(ch)
	go ReadData(ch)
	go CountTime(timer, end)
	c := <-end
	if c {
		close(ch)
	}
}
