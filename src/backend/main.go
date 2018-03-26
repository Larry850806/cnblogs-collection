package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/araddon/dateparse"
)

func getTaobaofedArticles() []article {
	// STEP 1: Get page amount
	homePageURL := "https://taobaofed.org/categories/Node-js/"
	doc, err := goquery.NewDocument(homePageURL)
	if err != nil {
		panic(err)
	}
	pageAmount, err := strconv.Atoi(doc.Find("a.page-number").Last().Text())
	if err != nil {
		panic(err)
	}

	// STEP 2: Generate urls
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

	// STEP 3: Get articles
	var wg sync.WaitGroup
	wg.Add(pageAmount)

	// each page at most can contains 10 articles
	articles := make([]article, 0, pageAmount*10)
	for _, url := range urls {
		go func(url string) {
			doc, err := goquery.NewDocument(url)
			if err != nil {
				panic(err)
			}
			doc.Find("div.article-summary-inner").Each(func(_ int, s *goquery.Selection) {
				title := strings.TrimSpace(s.Find("h2.article-title").Text())

				publishDate, err := dateparse.ParseAny(s.Find("time").Text())
				if err != nil {
					panic(err)
				}

				articlePrefix := "https://taobaofed.org"
				articleRelativeURL, _ := s.Find("a").Attr("href")
				articleFullURL := fmt.Sprintf("%s%s", articlePrefix, articleRelativeURL)

				a := article{title: title, date: publishDate, author: "掏寶前端團隊", url: articleFullURL}
				// a := types.Article{Date: publishDate}

				articles = append(articles, a)
			})
			wg.Done()
		}(url)
	}

	wg.Wait()
	return articles
}

func main() {
	tasks := []task{getTaobaofedArticles}

	var wg sync.WaitGroup
	wg.Add(len(tasks))

	allArticles := make([]article, 0)
	for _, t := range tasks {
		go func(t task) {
			articles := t()
			allArticles = append(allArticles, articles...)
			wg.Done()
		}(t)
	}

	wg.Wait()

	// Print
	for _, article := range allArticles {
		fmt.Println(article)
	}
}
