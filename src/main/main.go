package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	tasks := []task{getTaobaofedArticles, getJerryQuArticles, getWuBoyArticles}

	var wg sync.WaitGroup
	wg.Add(len(tasks))

	var allArticles articleList

	var mutex sync.Mutex
	for _, t := range tasks {
		go func(t task) {
			defer wg.Done()
			articles := t()
			mutex.Lock()
			allArticles = append(allArticles, articles...)
			mutex.Unlock()
		}(t)
	}

	wg.Wait()

	sort.Sort(allArticles)

	for _, article := range allArticles {
		fmt.Println(article)
	}
}
