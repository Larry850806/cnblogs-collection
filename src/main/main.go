package main

import (
	"fmt"
	"sync"
)

func main() {
	tasks := []task{getTaobaofedArticles, getJerryQuArticles, getWuBoyArticles}

	var wg sync.WaitGroup
	wg.Add(len(tasks))

	allArticles := make([]article, 0)
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

	// Print
	for _, article := range allArticles {
		fmt.Println(article)
	}
	fmt.Println(len(allArticles))
}