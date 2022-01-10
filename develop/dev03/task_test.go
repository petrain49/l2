package main

import "testing"


func TestGetString(t *testing.T) {
	_, err := getString("./strings")
	if err != nil {
		t.Fail()
	}
}

func TestSliceToString(t *testing.T) {
	slice := []string{"hello", "world!"}

	str := sliceToString(slice)

	if str != "hello\nworld!\n" {
		t.Log(str)
		t.Fail()
	}
}