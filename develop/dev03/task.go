package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца TODO
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов TODO

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	fileName := os.Args[len(os.Args)-1]

	column := flag.Int("k", -1, "column number")
	byNumber := flag.Bool("n", false, "sort by number")
	reverse := flag.Bool("r", false, "reverse sort")
	unique := flag.Bool("u", false, "no duplicates")
	trimSpace := flag.Bool("b", false, "trim spaces")
	check := flag.Bool("c", false, "check sort")
	flag.Parse()
	
	fileString, err := getString(fileName)
	if err != nil {
		log.Fatal(err)
	}

	res, err := sortStrings(fileString, *column, *byNumber, *reverse, *unique, *trimSpace, *check)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}

func sortStrings(file string, column int, byNumber bool, reverse bool, unique bool, trimSpace, check bool) (string, error) {
	if trimSpace {
		log.Println("-b")
		file = strings.TrimSpace(file)
	}

	fileStrings := strings.Split(file, "\n")

	if check && sort.StringsAreSorted(fileStrings) {
			fmt.Println("-c file sorted")
		}

	if column != -1 {
		log.Println("-k", column)

		sort.Slice(fileStrings, func(i, j int) bool {
			h := strings.Split(fileStrings[i], " ")
			d := strings.Split(fileStrings[j], " ")
			if len(h) <= column || len(d) <= column {
				return h[0] < d[0]
			}
			return h[column] < d[column]
		})
	}

	if unique {
		log.Println("-u")

		set := map[string]struct{}{}

		for _, str := range fileStrings {
			set[str] = struct{}{}
		}

		fileStrings = make([]string, 0, len(set))
		for x := range set {
			fileStrings = append(fileStrings, x)
		}
	}

	if byNumber {
		log.Println("-n")
		nums := make([]int, 0, len(fileStrings))

		for x := range fileStrings {
			n, err := strconv.Atoi(fileStrings[x])
			if err != nil {
				log.Fatal(err)
			}

			nums = append(nums, n)
		}

		if reverse {
			log.Println("-r")
			sort.Sort(sort.Reverse(sort.IntSlice(nums)))
		} else {
			sort.Ints(nums)
		}

		fileStrings = make([]string, 0, len(nums))
		for _, num := range nums {
			fileStrings = append(fileStrings, strconv.Itoa(num))
		}
	}

	if reverse {
		log.Println("-r")
		sort.Sort(sort.Reverse(sort.StringSlice(fileStrings)))
	}

	return sliceToString(fileStrings), nil
}

func sliceToString(slice []string) string {
	var res strings.Builder

	for _, s := range slice {
		res.WriteString(s)
		res.WriteString("\n")
	}

	return res.String()
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

	return string(file), nil
}