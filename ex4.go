package main

import (
	"fmt"
	"os"
	"os/signal"
)

func CheckSignal() bool {
	s := make(chan os.Signal)
	signal.Notify(s, os.Interrupt)
	if <-s == os.Interrupt {
		return true
	}
	return false
}

func Workers(w int, c chan string) {
	for res := range c {
		fmt.Println(w, res)
	}
}

//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.
//Программа должна завершаться по нажатию Ctrl+C.
//Выбрать и обосновать способ завершения работы всех воркеров.
func main() {
	var NumWorkers int
	WorkChan := make(chan string)
	fmt.Scanln(&NumWorkers)
	for NumWorkers != 0 {
		go Workers(NumWorkers, WorkChan)
		NumWorkers--
	}
	f := false
	go func() {
		f = CheckSignal()
	}()
	for {
		if f {
			fmt.Println("finish")
			close(WorkChan)
			break
		} else {
			WorkChan <- "!!!work!!!"
		}
	}
}
