package notes

type Note struct {
	title   string
	content string
}

func New(title, content string) Note {
	return Note{
		title:   title,
		content: content,
	}
}

func (n Note) Title() string {
	return n.title
}

func (n Note) Content() string {
	return n.content
}
