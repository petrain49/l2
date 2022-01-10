package main

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

type iObject interface {
	setFirstField(int)
	setSecondField(int)
}

// Object --------------------------------------------
type Object struct {
	firstField int
	secondField int
}

func (o *Object) setFirstField(n int) {
	o.firstField = n
}

func (o *Object) setSecondField(n int) {
	o.secondField = n
}

// First object --------------------------------------
type FirstObject struct {
	Object
}

func NewFirstObject() iObject {
	return &FirstObject{}
}

// Second object -------------------------------------
type SecondObject struct {
	Object
}

func NewSecondObject() iObject {
	return &SecondObject{}
}

// Main ----------------------------------------------
// Фабричный метод — это порождающий паттерн проектирования, который решает проблему создания различных продуктов,
//без указания конкретных классов продуктов.
func main() {
	first := getObject("first")
	first.setFirstField(1)
	first.setSecondField(2)

	second := getObject("second")
	second.setFirstField(3)
	second.setSecondField(4)

	fmt.Println(first, second)
}

func getObject(s string) iObject {
	switch s {
	case "first":
		return NewFirstObject()
	case "second":
		return NewSecondObject()
	default:
		return nil
	}
}