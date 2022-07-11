package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func UseContext(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		fmt.Println("1st: work goroutines is ended by context")
	}
}

func UseChanel(ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case <-ch:
		fmt.Println("2nd: work goroutines is ended by chanel")

	}
}

func UseTimer(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		fmt.Println("3rd: work goroutines is ended by ContextTimeOut")
	}
}

//Реализовать все возможные способы остановки выполнения горутины.
func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	//1
	ctxCancel, cancel := context.WithCancel(context.Background())
	go UseContext(ctxCancel, &wg)
	cancel()
	//2
	ch := make(chan bool)
	defer close(ch)
	go UseChanel(ch, &wg)
	ch <- true
	//3
	ctxTime, timeout := context.WithTimeout(context.Background(), 3*time.Second)
	go UseTimer(ctxTime, &wg)
	defer timeout()
	wg.Wait()
}
