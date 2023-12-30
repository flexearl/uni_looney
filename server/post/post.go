package post

type Post struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        string `json:"body"`
	PostID      int
}

func (post *Post) New(title, description, body string) Post {
	id := 1
	return Post{title, description, body, id}
}
