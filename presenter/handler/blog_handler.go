package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Khigashiguchi/profile-site/domain/model"
	"github.com/Khigashiguchi/profile-site/domain/repository"
)

type BlogHandler struct {
	DB *sql.DB
}

func (s *BlogHandler) GetBlogsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	blogs, err := repository.GetBlogs(s.DB)
	if err != nil {
		fmt.Printf("Error occurs while fetch blogs %s", err)
	}

	if err := json.NewEncoder(w).Encode(blogs); err != nil {
		fmt.Printf("Error occurs while encoding response %s", err)
	}
}

func (s *BlogHandler) PostBlogHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Get Body json data
	var blog model.Blog
	if err := json.NewDecoder(r.Body).Decode(blog); err != nil {
		fmt.Printf("Error occurs while encoding response %s", err)
	}

	id, err := repository.PostBlog(s.DB, blog.Title, blog.URL, blog.Published)
	if err != nil {
		fmt.Printf("Error occurs while save posted blog record %s", err)
	}

	res := struct {
		result    int
		createdID int64
	}{
		result:    http.StatusCreated,
		createdID: id,
	}
	if err := json.NewEncoder(w).Encode(res); err != nil {
		fmt.Printf("Error occurs while encoding response %s", err)
	}
}
