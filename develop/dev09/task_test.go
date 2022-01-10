package main

import "testing"

func TestWget(t *testing.T) {
	url := "https://en.wikipedia.org/wiki/%22Hello,_World!%22_program"

	wget(url)
}