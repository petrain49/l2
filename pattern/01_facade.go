package main

import (
	"fmt"
	"strings"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

type Cat struct {
	meow *Meow
	purr *Purr
	hiss *Hiss
}

func NewCat() *Cat {
	return &Cat {
		meow: &Meow{},
		purr: &Purr{},
		hiss: &Hiss{},
	}
}

func (c *Cat) todo() string {
	res := []string {
		c.meow.meow(),
		c.purr.purr(),
		c.hiss.hiss(),
	}
	return strings.Join(res, "\n")
}

type Meow struct{}

func (m *Meow) meow() string {
	fmt.Println("meow")
	return "meow"
}

type Purr struct{}

func (p *Purr) purr() string {
	fmt.Println("purr")
	return "purr"
}

type Hiss struct{}

func (h *Hiss) hiss() string {
	fmt.Println("hiss")
	return "hiss"
}

// Main ----------------------------------------------
// Фасад — это структурный паттерн, который предоставляет простой (но урезанный) интерфейс к сложной системе объектов, библиотеке или фреймворку.
func main() {
	cat := NewCat()
	cat.todo()
}