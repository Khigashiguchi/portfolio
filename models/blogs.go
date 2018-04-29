package models

import "database/sql"

type Blog struct {
	ID int `json:"id"`
	Title string `json:"name"`
	URL string `json:"url"`
}

type BlogCollection struct {
	Blogs []Blog `json:"items"`
}

func GetBlogs(db *sql.DB) BlogCollection {
	sql := "SELECT * FROM blogs"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	result := BlogCollection{}
	for rows.Next() {
		blog := Blog{}
		err := rows.Scan(&blog.ID, &blog.Title, &blog.URL)
		if err != nil {
			panic(err)
		}
		result.Blogs = append(result.Blogs, blog)
	}
	return result
}