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
