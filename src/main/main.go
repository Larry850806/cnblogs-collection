package main

import (
	"fmt"
	"strings"
	"sync"

	"database/sql"

	_ "github.com/lib/pq"
)

func getAllArticles() []article {
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

	fmt.Println(len(allArticles))

	return allArticles
}

func escape(s string) string {
	return strings.Replace(s, `'`, `''`, -1)
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

	allArticles := getAllArticles()

	values := make([]string, 0, len(allArticles))
	for _, article := range allArticles {
		date := fmt.Sprintf("%d-%d-%d", article.date.Year(), int(article.date.Month()), article.date.Day())
		str := fmt.Sprintf(`('%s', '%s', '%s', '%s')`, escape(article.title), date, escape(article.url), article.author)
		values = append(values, str)
	}

	sqlStatement := fmt.Sprintf(`
		INSERT INTO article (title, date, url, author)
		VALUES %s
	`, strings.Join(values, ",\n"))

	fmt.Println(sqlStatement)

	stmt, err := db.Prepare(sqlStatement)
	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
