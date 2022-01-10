package main

import "testing"

func TestIsAnagram(t *testing.T) {
	a := "пятка"
	b := "тяпка"

	c := "слиток"
	d := "листок"

	e := "шарик"
	f := "широк"

	g := "пень"
	h := "тень"

	if !isAnagram(a, b) {
		t.Fail()
	}

	if !isAnagram(c, d) {
		t.Fail()
	}

	if isAnagram(e, f) {
		t.Fail()
	}

	if isAnagram(g, h) {
		t.Fail()
	}
} 

func TestSearchAnagram(t *testing.T) {
	word := []string{"тяпка", "листок", "пятка", "слиток", "пятак", "столик", "лето"}

	res := searchAnagrams(word)
	t.Log(res)
}