package main

import (
	"bufio"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

const fileNameHTML = "downloadedHTML"
const fileNameLinks = "downloadedLinks"

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

	fileHTML, err := os.OpenFile(path+"/"+fileNameHTML, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer fileHTML.Close()

	fileLinks, err := os.OpenFile(path+"/"+fileNameLinks, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer fileLinks.Close()

	_, err = io.Copy(fileHTML, response.Body)
	if err != nil {
		return err
	}

	str, _ := getString(fileNameHTML)
	findLinks(str, fileLinks)

	return err
}

func findLinks(str string, file *os.File) {
	href := "href"
	src := "src"

	lines := strings.Split(str, ">")

	attrHref := regexp.MustCompile(`href="[^"]+"`)
	attrSrc := regexp.MustCompile(`src="[^"]+"`)

	for _, l := range lines {
		if strings.Contains(l, href) {
			resHref := attrHref.FindAllString(str, -1)
			for x := 0; x < len(resHref); x++ {
				file.WriteString(resHref[x][len(`href="`)-1:len(resHref[x])-1] + "\n")
			}
		}

		if strings.Contains(l, src) {
			resSrc := attrSrc.FindAllString(str, -1)
			for x := 0; x < len(resSrc); x++ {
				file.WriteString(resSrc[x][len(`href="`)-1:len(resSrc[x])-1] + "\n")
			}
		}
	}
}

func getString(fileName string) (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}

	file, err := ioutil.ReadFile(path + "/" + fileName)
	if err != nil {
		return "", err
	}

	if len(file) == 0 {
		return "", errors.New("empty string")
	}

	return string(file), nil
}
