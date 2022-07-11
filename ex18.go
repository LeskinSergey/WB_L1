package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.RWMutex
	c uint64
}

func NewCoounter() *Counter {
	return &Counter{}
}
func (count *Counter) increment(wg *sync.WaitGroup) {
	count.Lock()
	defer count.Unlock()
	count.c++
	wg.Done()
}

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика.
func main() {
	var n, i uint64
	var wg sync.WaitGroup
	count := NewCoounter()
	fmt.Scanln(&n)
	for i < n {
		wg.Add(1)
		go count.increment(&wg)
		i++
	}
	wg.Wait()
	fmt.Println(count.c)
}
