package main

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

// Object --------------------------------------------
type Object struct {
	firstState State
	secondState State

	currentState State
}

func NewObject() *Object {
	object := Object{}
	object.firstState = &FirstStateObject{&object}
	object.secondState = &SecondStateObject{&object}
	object.SetState(object.firstState)
	return &object
}

func (o *Object) FirstState() string {
	return o.currentState.FirstState()
}

func (o *Object) SecondState() string {
	return o.currentState.SecondState()
}

func (o *Object) SetState(s State) {
	o.currentState = s
}

// States --------------------------------------------
type State interface {
	FirstState() string
	SecondState() string
}

// First ---------------------------------------------
type FirstStateObject struct {
	obj *Object
}

func (o *FirstStateObject) FirstState() string {
	o.obj.SetState(o.obj.firstState)
	return "First"
}

func (o *FirstStateObject) SecondState() string {
	o.obj.SetState(o.obj.secondState)
	return "Second"
}

// Second --------------------------------------------
type SecondStateObject struct{
	obj *Object
}

func (o *SecondStateObject) FirstState() string {
	o.obj.SetState(o.obj.firstState)
	return "First"
}

func (o *SecondStateObject) SecondState() string {
	o.obj.SetState(o.obj.secondState)
	return "Second"
}

// Main ----------------------------------------------
/*
Состояние — это поведенческий паттерн, позволяющий динамически изменять поведение объекта при смене его состояния.

	+Избавляет от множества больших условных операторов машины состояний.
	+Концентрирует в одном месте код, связанный с определённым состоянием.
	+Упрощает код контекста.

	-Может неоправданно усложнить код, если состояний мало и они редко меняются.
*/
func main() {
	object := NewObject()

	fmt.Println(object.FirstState(), object.currentState)
	fmt.Println(object.SecondState(), object.currentState)

}