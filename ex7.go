package main

import (
	"fmt"
	"sync"
)

type DefenseMap struct {
	sync.RWMutex
	m map[interface{}]interface{}
}

func NewM() *DefenseMap {
	return &DefenseMap{
		m: make(map[interface{}]interface{}),
	}
}
func (obj *DefenseMap) Add(key, value interface{}) {
	obj.Lock()
	defer obj.Unlock()
	obj.m[key] = value
}
func (obj *DefenseMap) Delete(key interface{}) {
	obj.Lock()
	defer obj.Unlock()
	delete(obj.m, key)
}
func (obj *DefenseMap) Get(key interface{}) interface{} {
	obj.Lock()
	defer obj.Unlock()
	return obj.m[key]
}

//Реализовать конкурентную запись данных в map.
func main() {
	NewMap := NewM()
	NewMap.Add("hello", "world!")
	res := NewMap.Get("hello")
	fmt.Println(res)
	NewMap.Delete("hello")
}
