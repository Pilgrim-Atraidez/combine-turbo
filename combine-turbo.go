package main

import (
	"bufio"
	"fmt"
	"os"
)

type HttpResponse struct {
	url      string
	response *http.Response
	err      error
}

func getUrlList(fileName string) []string {
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

	return urlList
}

func main() {
	fileName := "urls.txt"
	urlList := getUrlList(fileName)

	ch := make(chan *http.Response, len(urlList))
	for _, url := range urlList {
		go func(url string) {
			fmt.Printf("Fetching %s\n", url)
			response, err := http.Get(url)
			ch <- &HttpResponse{url, response, err}
		}(url)
	}

	responses := make([]*HttpResponse, len(urlList))
	responses = <-ch

	for _, err := range responses {
		fmt.Println(err)
	}
}
