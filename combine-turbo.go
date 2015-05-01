package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := "urls.txt"
	var urlList []string

	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		urlList = append(urlList, scanner.Text())
	}

	for _, url := range urlList {
		fmt.Println(url)
	}
}
