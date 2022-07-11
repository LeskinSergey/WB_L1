package main

import (
	"fmt"
	"strconv"
)

type Work interface {
	Addition()
}

func (n ABInt) Addition() {
	fmt.Println(n.a + n.b)
}

type ABInt struct {
	a, b int64
}
type ABString struct {
	a, b string
}

type AdapterString struct {
	s *ABString
}

func (s *ABString) ConvertToInt() (int64, int64) {
	n1, _ := strconv.ParseInt(s.a, 10, 64)
	n2, _ := strconv.ParseInt(s.b, 10, 64)
	return n1, n2
}

func (adapt *AdapterString) Addition() {
	a, b := adapt.s.ConvertToInt()
	fmt.Println(a + b)
}

func NewString(a, b string) *ABString {
	return &ABString{
		a: a,
		b: b,
	}
}

func NewInt(a, b int64) *ABInt {
	return &ABInt{
		a: a,
		b: b,
	}
}
func NewAdapter(a, b string) *AdapterString {
	return &AdapterString{
		s: NewString(a, b),
	}
}

//Реализовать паттерн «адаптер» на любом примере.
func main() {
	I := NewInt(10, 11)
	StrAdapt := NewAdapter("10", "11")
	fmt.Println("ABSInt")
	I.Addition()
	fmt.Println("Adapter")
	StrAdapt.Addition()
}
