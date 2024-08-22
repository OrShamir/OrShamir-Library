package book

type BookDTO struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	Topic      string `json:"topic"`
	Year       int    `json:"year"`
	Popularity int    `json:"popularity"`
}

type BookSearchDTO struct {
	Query string `json:"query"`
}
