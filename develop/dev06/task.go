package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	fields := flag.Int("f", 1, "выбрать поля (колонки)")
	delimeter := flag.String("d", "\t", "использовать другой разделитель")
	separated := flag.Bool("s", false, "только строки с разделителем")
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		stdinHandle(*fields, *delimeter, *separated)
		return
	}

	name := os.Args[len(os.Args)-1]
	file, err := getStrings(name) 
	if err != nil {
		log.Fatal(err)
	}

	for _, str := range file {
		fmt.Println(cut(str, *fields, *delimeter, *separated))
	}
}

func stdinHandle(fields int, delimeter string, separated bool) string {
	scanner := bufio.NewScanner(os.Stdin)
	str := ""

	for {
		scanner.Scan()
		str = scanner.Text()
		if str == "" {
			return ""
		}

		fmt.Println(cut(str, fields, delimeter, separated))
	}
}

func getStrings(fileName string) ([]string, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file, err := ioutil.ReadFile(path + "/" + fileName)
	if err != nil {
		return nil, err
	}

	res := strings.Split(string(file), "\n")
	return res, nil
}

func cut(str string, fields int, delimeter string, separated bool) string {
	if separated && !strings.Contains(str, delimeter) {
		return ""
	}

	words := strings.Split(str, delimeter)

	if fields <= len(words) {
		return words[fields-1] + "\n"
	} 
	return ""
}