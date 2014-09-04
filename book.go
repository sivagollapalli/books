package books

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b *Book) CategoryByLength() string {
	if b.Pages > 100 {
		return "NOVEL"
	} else {
		return "SHORT STORY"
	}
}
