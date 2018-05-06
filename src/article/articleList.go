package article

// List is a list of articles
type List []Article

func (al List) Len() int {
	return len(al)
}

func (al List) Less(i int, j int) bool {
	article1, article2 := al[i], al[j]
	// move latest article to first element
	return article1.isAfter(&article2)
}

func (al List) Swap(i int, j int) {
	al[i], al[j] = al[j], al[i]
}

// Append returns the merged List
func (al List) Append(al2 List) List {
	return append(al, al2...)
}
