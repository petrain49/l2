package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	url := scanner.Text()

	err := wget(url)
	if err != nil {
		log.Println(err)
	}
}

func wget(url string) error {
	response, err := http.Get(url)
	if err != nil {
		log.Println(response.Status)
		return err
	}
	defer response.Body.Close()

	path, err := os.Getwd()
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path + "/" + "downloaded", os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return err
}
