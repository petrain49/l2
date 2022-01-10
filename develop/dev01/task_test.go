package main

import "testing"

func TestTime(t *testing.T) {
	err := printTime()
	if err != nil {
		t.Log(err)
	}
}