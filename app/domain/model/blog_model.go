package model

type Blog struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	URL       string `json:"url"`
	Published string `json:"published"`
}

type BlogCollection struct {
	Blogs []Blog `json:"items"`
}
