package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
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

func getResponses(urlList []string) []*HttpResponse {
	ch := make(chan *HttpResponse)
	responses := []*HttpResponse{}
	for _, url := range urlList {
		go func(url string) {
			fmt.Printf("Fetching %s\n", url)
			response, err := http.Get(url)
			ch <- &HttpResponse{url, response, err}
		}(url)
	}

	for {
		select {
		case r := <-ch:
			fmt.Printf("%s was fetched\n", r.url)
			responses = append(responses, r)
			if len(responses) == len(urlList) {
				return responses
			}
		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}
	}
}

func main() {
	fileName := "urls.txt"
	urlList := getUrlList(fileName)
	responses := getResponses(urlList)

	for _, err := range responses {
		fmt.Println(err)
	}
}
