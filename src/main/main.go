package main

import (
	"fmt"
	"sync"
	"time"

	"database/sql"

	_ "github.com/lib/pq"
)

func printAllArticles() {
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

func main() {
	// connect to postgresql
	connStr := "postgres://postgres:dev@localhost/cnblogs?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println(db)

	sqlStatement := `
		INSERT INTO article (title, date, url, author)
		VALUES ($1, $2, $3, $4)
	`

	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec("標題", time.Now(), "https://url", "盧承億")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
