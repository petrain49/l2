package main

import (
	"flag"
	"fmt"
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

	k := flag.Int("k", -1, "column number")
	n := flag.Bool("n", false, "sort by number")
	r := flag.Bool("r", false, "reverse sort")
	u := flag.Bool("u", false, "no duplicates")
	m := flag.Bool("M", false, "sort by month")
	b := flag.Bool("b", false, "trim spaces")
	c := flag.Bool("c", false, "check sort")
	flag.Parse()

	keys := keys{
		column:    *k,
		byNumber:  *n,
		reverse:   *r,
		unique:    *u,
		byMonth:   *m,
		trimSpace: *b,
		check:     *c,
	}
	log.Println(keys)

	fileString, err := getString(fileName)
	if err != nil {
		log.Fatal(err)
	}

	res, err := sortStrings(fileString, keys)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}

func sortStrings(file string, keys keys) (string, error) {
	if keys.trimSpace {
		log.Println("-b")
		file = strings.TrimSpace(file)
	}

	fileStrings := strings.Split(file, "\n")

	if keys.check && sort.StringsAreSorted(fileStrings) {
		fmt.Println("-c file sorted")
	}

	if keys.column != -1 {
		log.Println("-k", keys.column)

		sort.Slice(fileStrings, func(i, j int) bool {
			h := strings.Split(fileStrings[i], " ")
			d := strings.Split(fileStrings[j], " ")
			if len(h) <= keys.column || len(d) <= keys.column {
				return h[0] < d[0]
			}
			return h[keys.column] < d[keys.column]
		})
	}

	if keys.unique {
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

	if keys.byNumber {
		log.Println("-n")
		nums := make([]int, 0, len(fileStrings))

		for x := range fileStrings {
			n, err := strconv.Atoi(fileStrings[x])
			if err != nil {
				log.Fatal(err)
			}

			nums = append(nums, n)
		}

		if keys.reverse {
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

	if keys.reverse {
		log.Println("-r")
		sort.Sort(sort.Reverse(sort.StringSlice(fileStrings)))
	}

	if keys.column == -1 && !keys.byNumber && !keys.reverse && !keys.unique && !keys.trimSpace && !keys.check {
		log.Println("none")
		sort.Sort(sort.StringSlice(fileStrings))
	}

	return sliceToString(fileStrings), nil
}
