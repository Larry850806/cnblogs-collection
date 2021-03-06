package main

import (
	"article"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/araddon/dateparse"
)

func getTaobaofedArticles() article.List {
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
	articles := make(article.List, 0, pageAmount*10)
	var mutex sync.Mutex
	for _, url := range urls {
		go func(url string) {
			defer wg.Done()
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

				a := article.Article{Title: title, Date: publishDate, Author: "掏寶前端團隊", URL: articleFullURL}

				mutex.Lock()
				articles = append(articles, a)
				mutex.Unlock()
			})
		}(url)
	}

	wg.Wait()
	return articles
}

func getJerryQuArticles() article.List {
	url := "https://imququ.com/post/series.html"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}

	var articles article.List
	doc.Find(".entry-content > ul > li").Each(func(_ int, s *goquery.Selection) {
		link := s.Find("a:not(:last-child)")

		urlRelativePath, _ := link.Attr("href")
		urlFullPath := fmt.Sprintf("%s%s", "https://imququ.com", urlRelativePath)

		title := strings.TrimSpace(link.Text())

		// fmt.Println(dateparse.ParseAny("dec 23, 2011"))
		publishDate, err := dateparse.ParseAny(strings.Trim(s.Find("span.date").Text(), "()"))
		if err != nil {
			panic(err)
		}

		a := article.Article{Title: title, Date: publishDate, Author: "Jerry Qu", URL: urlFullPath}
		articles = append(articles, a)
	})

	return articles
}

func getWuBoyArticles() article.List {
	// STEP 1: Get page amount
	homePageURL := "https://blog.wu-boy.com/category/電腦技術/"
	doc, err := goquery.NewDocument(homePageURL)
	if err != nil {
		panic(err)
	}

	doc.Find("span.meta-nav.screen-reader-text").Remove()
	pageAmount, err := strconv.Atoi(doc.Find("nav > div > a:nth-child(4)").Text())
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

	// each page at most can contains 8 articles
	articles := make(article.List, 0, pageAmount*8)
	var mutex sync.Mutex
	for _, url := range urls {
		go func(url string) {
			defer wg.Done()
			doc, err := goquery.NewDocument(url)
			if err != nil {
				panic(err)
			}
			doc.Find("article").Each(func(_ int, s *goquery.Selection) {
				titleElement := s.Find("h2.entry-title")

				title := titleElement.Text()
				articleURL, _ := titleElement.Find("a").Attr("href")

				publishDate, err := dateparse.ParseAny(s.Find("span.posted-on time").First().Text())
				if err != nil {
					panic(err)
				}

				a := article.Article{Title: title, Date: publishDate, Author: "AppleBOY", URL: articleURL}

				mutex.Lock()
				articles = append(articles, a)
				mutex.Unlock()
			})
		}(url)
	}

	wg.Wait()
	return articles
}
