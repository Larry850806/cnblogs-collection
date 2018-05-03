package main

import (
	"fmt"
	"time"
)

type article struct {
	title  string
	date   time.Time
	url    string
	author string
}

func (a article) String() string {
	date := fmt.Sprintf("%d-%02d-%02d", a.date.Year(), int(a.date.Month()), a.date.Day())
	return fmt.Sprintf("%s by %s - %s", date, a.author, a.title)
}

type articleList []article

func (al articleList) Len() int {
	return len(al)
}

func (al articleList) Less(i int, j int) bool {
	article1, article2 := al[i], al[j]
	date1, date2 := article1.date, article2.date
	// move latest article to first element
	return date1.After(date2)
}

func (al articleList) Swap(i int, j int) {
	al[i], al[j] = al[j], al[i]
}

type task func() articleList
