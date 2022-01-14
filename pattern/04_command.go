package main

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

// Sender --------------------------------------------
type Button struct {
	command Command
}

func (b *Button) Press() {
	b.command.Execute()
}

// First command -------------------------------------
type Command interface {
	Execute()
}

type OnCommand struct {
	device Device
}

func (c *OnCommand) Execute() {
	c.device.On()
}

// Second command ------------------------------------
type OffCommand struct {
	device Device
}

func (c *OffCommand) Execute() {
	c.device.Off()
}

// Object --------------------------------------------
type Device interface {
	On()
	Off()
}

type Lamp struct {
	Light bool
}

func (l *Lamp) On() {
	l.Light = true
	fmt.Println("light on")
}

func (l *Lamp) Off() {
	l.Light = false
	fmt.Println("light off")
}

// Main ----------------------------------------------
/*
Команда — это поведенческий паттерн, позволяющий заворачивать запросы или простые операции в отдельные объекты.
	+Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
	+Позволяет реализовать простую отмену и повтор операций.
	+Позволяет реализовать отложенный запуск операций.
	+Позволяет собирать сложные команды из простых.
	+Реализует принцип открытости/закрытости.

	-Усложняет код программы из-за введения множества дополнительных классов.
*/
func main() {
	lamp := &Lamp{}

	onCommand := &OnCommand{
		device: lamp,
	}

	offCommand := &OffCommand{
		device: lamp,
	}

	onButton := &Button{
		command: onCommand,
	}
	onButton.Press()

	offButton := &Button{
		command: offCommand,
	}
	offButton.Press()
}
