package main

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

type Rectangle struct {
	x int
	y int
}

func (r *Rectangle) GetX() int {
	return r.x
}

func (r *Rectangle) GetY() int {
	return r.y
}

func (r *Rectangle) Accept(v visitor) {
	v.visitRectangle(r)
}

// Visitor -------------------------------------------
type visitor interface {
	visitRectangle(*Rectangle)
}

type AreaCalc struct {
	area int
}

func (a *AreaCalc) visitRectangle(r *Rectangle) {
	a.area = r.x * r.y
}

// Main ----------------------------------------------
// Посетитель — это поведенческий паттерн, который позволяет добавить новую операцию для целой иерархии классов, не изменяя код этих классов.
func main() {
	rect := &Rectangle{4, 5}

	ac := &AreaCalc{}
	rect.Accept(ac)

	fmt.Println(ac.area)
}
