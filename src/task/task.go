package task

import "article"

// Task is a function that get article.List from website
type Task func() article.List

// Run is to run this task an return a article.List
func (t Task) Run() article.List {
	return t()
}
