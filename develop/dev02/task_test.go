package main

import (
	"strings"
	"testing"
)

func TestUnpack(t *testing.T) {
	testSet := []string {`a4bc2d5e`, `a4`, `abcd`, `45`, ``, `qwe\4\5`, `qwe\45`, `qwe\\5`, `\4\52`}


	for _, s := range testSet {
		t.Log(unpack(s))
	}


	//t.Log(unpack(testSet[0]))
}

func TestWriteNRunes(t *testing.T) {
	var res strings.Builder
	n := []rune{'1', '0'}
	writeNRunes(&res, 'z', &n)
	t.Log(res.String(), n)
}

func TestDigitCheck(t *testing.T) {
	nonDigit := rune('g')
	digit := rune('8')

	ok := isDigit(nonDigit)
	if ok {
		t.Fail()
	}

	ok = isDigit(digit)
	if !ok {
		t.Fail()
	}
}

func TestValidate(t *testing.T) {
	str1 := ""
	str2 := "1ff4"
	str3 := `\1f`

	_, err := validate(str1)
	if err == nil {
		t.Fail()
	}
	_, err = validate(str2)
	if err == nil {
		t.Fail()
	}
	_, err = validate(str3)
	if err != nil {
		t.Fail()
	}


}
