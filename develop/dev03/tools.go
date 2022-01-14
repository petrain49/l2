package main

import (
	"io/ioutil"
	"os"
	"strings"
)

type keys struct {
	column    int
	byNumber  bool
	reverse   bool
	unique    bool
	byMonth   bool
	trimSpace bool
	check     bool
	suffix    bool
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

	return string(file), nil
}

func sliceToString(slice []string) string {
	var res strings.Builder

	for _, s := range slice {
		res.WriteString(s)
		res.WriteString("\n")
	}

	return res.String()
}
