package main

import (
	"errors"
	"fmt"
	"strings"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

const (
	symbolState byte = iota
	digitState
	escapeState
)

func main() {
	var s string
	fmt.Scanf("%s", &s)

	res, err := unpack(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	
}

func unpack(str string) (string, error) {
	runes, err := validate(str)

	var res strings.Builder

	var lastSymbol rune
	var currentSymbol rune
	var currentNumber []rune

	currentState := symbolState
	previousState := symbolState

	for x := 0; x < len(runes); x++ {
		currentSymbol = runes[x]

		if previousState != escapeState {
			currentState = switchState(currentSymbol)
		} else {
			currentState = symbolState
		}

		switch currentState {
		case symbolState:
			writeNRunes(&res, lastSymbol, &currentNumber) 

			res.WriteRune(currentSymbol)
			lastSymbol = currentSymbol

			previousState = currentState

		case digitState:
			currentNumber = append(currentNumber, currentSymbol)
			if x == len(runes) - 1 {
				writeNRunes(&res, lastSymbol, &currentNumber)
			}

			previousState = currentState

		case escapeState:
			previousState = currentState
		}
	}

	return res.String(), err
}

func validate(str string) ([]rune, error) {
	stringErr := errors.New("incorrect string")

	if len(str) == 0 {
		return []rune{}, nil
	}

	runes := []rune(str)

	if isDigit(runes[0]) {
		return []rune{}, stringErr
	}

	return runes, nil
}

func isDigit(r rune) bool {
	return (r >= '0' && r <= '9')
}

func toDigit(runes []rune) int {
	res := 0
	for r, m := len(runes)-1, 1; r >= 0; r, m = r-1, m*10 {
		res += int(runes[r]-'0') * m
	}
	return res
}

func switchState(r rune) byte {
	switch {
	case isDigit(r):
		return digitState
	case r == '\\':
		return escapeState
	default:
		return symbolState
	}
}

// if previous runes were digits - write last symbol to string.builder <n - 1> times
func writeNRunes(builder *strings.Builder, r rune, n *[]rune) {
	if len(*n) == 0 {
		return
	}

	nInt := toDigit(*n)
	for x := 0; x < nInt-1; x++ {
		builder.WriteRune(r)
	}

	*n = []rune{}
}
