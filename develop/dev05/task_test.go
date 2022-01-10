package main

import "testing"

func TestGrep(t *testing.T) {
	want := "1:  I have walked out in rain - and back in rain.\n"
	file, _ := getString("test")

	got := grep(file, "rain", -1, -1, -1, false, false, false, false, true)
	if got != want {
		t.Fatal(got)
	}
}