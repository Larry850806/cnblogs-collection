package main

import (
	"article"
	"fmt"
	"sort"
	"sync"
	"task"
)

func main() {
	tasks := []task.Task{getTaobaofedArticles, getJerryQuArticles, getWuBoyArticles}

	var wg sync.WaitGroup
	wg.Add(len(tasks))

	var allArticles article.List

	var mutex sync.Mutex
	for _, t := range tasks {
		go func(t task.Task) {
			defer wg.Done()
			articles := t.Run()
			mutex.Lock()
			allArticles = allArticles.Append(articles)
			mutex.Unlock()
		}(t)
	}

	wg.Wait()

	sort.Sort(allArticles)

	for _, article := range allArticles {
		fmt.Println(article)
	}
}
