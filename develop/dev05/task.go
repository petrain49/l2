package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	pattern := os.Args[len(os.Args)-2]
	fileName := os.Args[len(os.Args)-1]

	after := flag.Int("A", -1, "печатать +N строк после совпадения")
	before := flag.Int("B", -1, "печатать +N строк до совпадения")
	context := flag.Int("C", -1, "(A+B) печатать ±N строк вокруг совпадения")
	count := flag.Bool("c", false, "количество строк")
	ignore := flag.Bool("i", false, "игнорировать регистр")
	invert := flag.Bool("v", false, "вместо совпадения, исключать")
	fixed := flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	lineNum := flag.Bool("n", false, "печатать номер строки")
	flag.Parse()

	file, err := getString(fileName)
	if err != nil {
		log.Fatal(err)
	}

	res := grep(file, pattern, *after, *before, *context, *count, *ignore, *invert, *fixed, *lineNum)
	log.Println(res)
}

func grep(file string, pattern string, after int, before int, context int, count bool, ignore bool, invert bool, fixed bool, lineNum bool) string {
	actualFileStrings := strings.Split(file, "\n")

	var fileStrings []string
	if ignore {
		fileStrings = strings.Split(strings.ToLower(file), "\n")
	} else {
		fileStrings = strings.Split(file, "\n")
	}

	stringNums := make(map[int]struct{})

	for n, str := range fileStrings {
		if fixed && pattern == str {
			stringNums[n] = struct{}{}
		} else if !fixed {
			if ok, _ := regexp.MatchString(pattern, str); ok {
				stringNums[n] = struct{}{}
			}
		}
	}

	var res strings.Builder

	if count {
		res.WriteString(fmt.Sprintf("%d", len(stringNums)))
		return res.String()
	}

	if invert {
		for x := 0; x < len(actualFileStrings); x++ {
			if _, ok := stringNums[x]; !ok {
				res.WriteString(formatString(actualFileStrings, x, lineNum))
			}
		}
		return res.String()
	}

	if after != -1 {
		for n := range stringNums {
			res.WriteString(formatString(actualFileStrings, n, lineNum))

			for x := 0; x < after; x++ {
				if n+x < len(actualFileStrings) {
					res.WriteString(formatString(actualFileStrings, n+x, lineNum))
				}
			}
			res.WriteString("\n")
		}
		return res.String()
	}

	if before != -1 {
		for n := range stringNums {
			for x := 0; x < n; x++ {
				if n-x <= before {
					res.WriteString(formatString(actualFileStrings, n-x, lineNum))
				}
			}
			res.WriteString(formatString(actualFileStrings, n, lineNum))
			res.WriteString("\n")
		}
		return res.String()
	}

	if context != -1 {
		for n := range stringNums {
			for x := 0; x < n; x++ {
				if n-x <= before {
					res.WriteString(formatString(actualFileStrings, n-x, lineNum))
				}
			}

			res.WriteString(formatString(actualFileStrings, n, lineNum))

			for x := 0; x < after; x++ {
				if n+x < len(actualFileStrings) {
					res.WriteString(formatString(actualFileStrings, n+x, lineNum))
				}
			}
			res.WriteString("\n")
		}
		return res.String()
	}

	for n := range stringNums {
		res.WriteString(formatString(actualFileStrings, n, lineNum))
	}
	return res.String()
}

func formatString(str []string, n int, lineNum bool) string {
	if !lineNum {
		return str[n] + "\n"
	}
	return fmt.Sprintf("%d: %s\n", n, str[n])
}

func getString(fileName string) (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}

	file, err := ioutil.ReadFile(path + "/" + fileName)
	if err != nil {
		return "", err
	}

	if len(file) == 0 {
		return "", errors.New("empty string")
	}

	return string(file), nil
}
