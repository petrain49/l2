package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
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
			err := os.Chdir(cmd[1])
			if err != nil {
				log.Println(err)
				continue
			}

		case "pwd":
			path, err := os.Getwd()
			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Println(path)

		case "echo":
			if len(cmd) > 1 {
				fmt.Println(cmd[1:])
			}

		case "kill":
			pid, err := strconv.Atoi(cmd[1])
			if err != nil {
				log.Println(err)
				continue
			}

			process, err := os.FindProcess(pid)
			if err != nil {
				log.Println(err)
				continue
			}

			err = process.Kill()
			if err != nil {
				log.Println(err)
				continue
			}
		case "ps":
			matches, err := filepath.Glob("/proc/*/exe")
			if err != nil {
				log.Println(err)
				continue
			}

    		for _, file := range matches {
				target, err := os.Readlink(file)
				if err != nil {
					log.Println(err)
					continue
				}

				if len(target) > 0 {
					fmt.Printf("%s\n", filepath.Base(target))
				}
    		}
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
