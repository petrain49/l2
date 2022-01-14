package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===


Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*


Так же требуется поддерживать функционал fork/exec-команд

Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var cmd []string
	done := false

	for scanner.Scan() && !done {
		cmd = strings.Split(scanner.Text(), " ")

		switch cmd[0] {
		case "cd":
			if cmd[1] == ".." {
				path := pwd()
				index := strings.LastIndex(path, "/")

				err := os.Chdir(path[:index])
				if err != nil {
					log.Println(err)
				}
				continue
			}
			err := os.Chdir(cmd[1])
			if err != nil {
				log.Println(err)
				continue
			}

		case "pwd":
			path := pwd()

			fmt.Printf("pwd %s\n", path)

		case "echo":
			if len(cmd) > 1 {
				fmt.Println(cmd[1:])
			}

		case "kill":
			pid := kill(cmd[1])

			fmt.Printf("killed proc %d\n", pid)

		case "ps":
			processes := ps()
			fmt.Println(processes)

		case "quit":
			done = true

		default:
			command := exec.Command(cmd[0], cmd[1:]...)
			command.Stdout = os.Stdout
			command.Stderr = os.Stderr
			command.Run()

		}
	}
}
