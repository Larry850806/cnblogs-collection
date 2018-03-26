package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getHTML(url string) string {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	htmlByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	html := fmt.Sprintf("%s", htmlByte)
	return html
}

func main() {
	url := "http://tour.golang.org/welcome/1"
	html := getHTML(url)
	fmt.Printf(html)
}
