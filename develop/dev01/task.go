package main

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

import (
	"fmt"
	"log"
	"os"

	"github.com/beevik/ntp"
)

const ntpServer = "ru.pool.ntp.org"

func main() {
	printTime()
}

func printTime() error {
	time, err := ntp.Time(ntpServer)
	errCheck(err)

	fmt.Println(time)

	return err
}

func errCheck(err error) {
	if err == nil {
		return
	}

	log.SetOutput(os.Stderr)
	log.Fatal(err)
}