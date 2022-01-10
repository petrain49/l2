package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	words := strings.Split(text, " ")
	log.Println(searchAnagrams(words))
}

func searchAnagrams(words []string) map[string][]string {
	words = sort.StringSlice(words)

	res := make(map[string][]string)

	for x, w := range words {
		words[x] = strings.ToLower(w)
	}
	
	for i := 0; i < len(words); i++ {
		if words[i] == "" {
			continue
		}

		res[words[i]] = make([]string, 0, len(words))

		for j := i + 1; j < len(words); j++ {
			if words[j] == "" {
				continue
			}
			

			if isAnagram(words[i], words[j]) {
				res[words[i]] = append(res[words[i]], words[j])
				words[j] = ""
			}
		}
	}

	for k, v := range res {
		if len(v) == 0 {
			delete(res, k)
		}
	}
	return res
}

func isAnagram(a string, b string) bool {
	if len(a) == 0 || len(b) == 0 || len(a) != len(b) {
		return false
	}

	aRunes := []rune(a)
	bRune := []rune(b)

	aSet := make(map[rune]byte)
	bSet := make(map[rune]byte)

	var currentARune rune
	var currentBRune rune

	for x := 0; x < len(aRunes); x++ {
		currentARune = aRunes[x]
		currentBRune = bRune[x]

		_, ok := aSet[currentARune]
		if !ok {
			aSet[currentARune] = 1
		} else {
			aSet[currentARune]++
		}

		_, ok = bSet[currentBRune]
		if !ok {
			bSet[currentBRune] = 1
		} else {
			bSet[currentBRune]++
		}
	}

	for ar, an := range aSet {
		bn, ok := bSet[ar]
		if !ok || an != bn {
			return false
		}
	}
	return true
}