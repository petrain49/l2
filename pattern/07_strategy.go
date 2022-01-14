package main

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

type iAlgo interface {
	execute(*Object)
}

// Algorithms ----------------------------------------
type FirstAlgo struct{}

func (a *FirstAlgo) execute(o *Object) {
	fmt.Println("First")
}

type SecondAlgo struct{}

func (a *SecondAlgo) execute(o *Object) {
	fmt.Println("Second")
}

type ThirdAlgo struct{}

func (a *ThirdAlgo) execute(o *Object) {
	fmt.Println("Third")
}

// Object --------------------------------------------
type Object struct{
	algorithm iAlgo
}

func (o *Object) setAlgo(alg iAlgo) {
	o.algorithm = alg
}

func (o *Object) execute() {
	o.algorithm.execute(o)
}

// Main ----------------------------------------------
/*
Стратегия — это поведенческий паттерн, выносит набор алгоритмов в собственные классы и делает их взаимозаменимыми.
	+Горячая замена алгоритмов на лету.
	+Изолирует код и данные алгоритмов от остальных классов.
	+Уход от наследования к делегированию.
	+Реализует принцип открытости/закрытости.

	-Усложняет программу за счёт дополнительных классов.
	-Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.
*/
func main() {
	obj := &Object{}

	first := &FirstAlgo{}
	second := &SecondAlgo{}
	third := &ThirdAlgo{}

	obj.setAlgo(first)
	obj.execute()

	obj.setAlgo(second)
	obj.execute()

	obj.setAlgo(third)
	obj.execute()
}
