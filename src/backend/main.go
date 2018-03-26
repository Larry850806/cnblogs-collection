package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	// Get page amount
	homePageURL := "https://taobaofed.org/categories/Node-js/"
	doc, err := goquery.NewDocument(homePageURL)
	if err != nil {
		panic(err)
	}
	pageAmount, err := strconv.Atoi(doc.Find("a.page-number").Last().Text())
	if err != nil {
		panic(err)
	}

	// Generate urls
	urls := make([]string, pageAmount)
	for i := 1; i <= pageAmount; i++ {
		var url string
		if i == 1 {
			url = homePageURL
		} else {
			url = fmt.Sprintf("%spage/%d/", homePageURL, i)
		}
		urls[i-1] = url

	}

	// Get titles
	var wg sync.WaitGroup
	wg.Add(pageAmount)

	// each page can contains 10 articles
	titles := make([]string, 0, pageAmount*10)
	for _, url := range urls {
		go func(url string) {
			doc, err := goquery.NewDocument(url)
			if err != nil {
				panic(err)
			}
			doc.Find("h2.article-title").Each(func(_ int, s *goquery.Selection) {
				title := strings.TrimSpace(s.Text())
				titles = append(titles, title)
			})
			wg.Done()
		}(url)
	}

	wg.Wait()

	// Print
	for _, title := range titles {
		fmt.Println(title)
	}
}
