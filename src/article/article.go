package article

import (
	"fmt"
	"time"
)

// Article is a struct for store a single article
type Article struct {
	Title  string
	Date   time.Time
	URL    string
	Author string
}

func (a Article) String() string {
	date := fmt.Sprintf("%d-%02d-%02d", a.Date.Year(), int(a.Date.Month()), a.Date.Day())
	return fmt.Sprintf("%s by %s：%s", date, a.Author, a.Title)
}

func (a *Article) isAfter(a2 *Article) bool {
	d1, d2 := a.Date, a2.Date
	return d1.After(d2)
}
