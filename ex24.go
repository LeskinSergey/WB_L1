package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

func NewPoint(x, y int) *Point {
	return &Point{
		x: x,
		y: y,
	}
}
func GetLength(d1, d2 *Point) float64 {
	DeltaX := d2.x - d1.x
	DeltaY := d2.y - d1.y
	return math.Sqrt(float64(DeltaX*DeltaX + DeltaY*DeltaY))
}

//Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
func main() {
	var x1, y1, x2, y2 int
	fmt.Scanln(&x1, &y1)
	fmt.Scanln(&x2, &y2)
	dot1 := NewPoint(x1, y1)
	dot2 := NewPoint(x2, y2)
	l := GetLength(dot1, dot2)
	fmt.Println("len :", l)
}
