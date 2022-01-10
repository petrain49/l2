package main

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

type Check interface {
	check(*Object)
	setNext(Check)
}

type Object struct {
	firstCheck  bool
	secondCheck bool
	thirdCheck  bool
}

// First check ---------------------------------------
type FirstCheck struct {
	next Check
}

func (c *FirstCheck) check(o *Object) {
	o.firstCheck = true
	fmt.Println("first")

	c.next.check(o)
}

func (c *FirstCheck) setNext(ch Check) {
	c.next = ch
}

// Second check --------------------------------------
type SecondCheck struct {
	next Check
}

func (c *SecondCheck) check(o *Object) {
	o.secondCheck = true
	fmt.Println("second")

	c.next.check(o)
}

func (c *SecondCheck) setNext(ch Check) {
	c.next = ch
}

// Third Check ---------------------------------------
type ThirdCheck struct {
	next Check
}

func (c *ThirdCheck) check(o *Object) {
	o.thirdCheck = true
	fmt.Println("third")
}

func (c *ThirdCheck) setNext(ch Check) {
	c.next = ch
}

// Main ----------------------------------------------
// Цепочка обязанностей — это поведенческий паттерн, позволяющий передавать запрос по цепочке потенциальных обработчиков,
// пока один из них не обработает запрос.
func main() {
	third := &ThirdCheck{}

	second := &SecondCheck{}
	second.setNext(third)

	first := &FirstCheck{}
	first.setNext(second)

	obj := &Object{}
	first.check(obj)

	fmt.Println(obj)
}
