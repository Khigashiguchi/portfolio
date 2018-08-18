package repository

import (
	"database/sql"

	"github.com/Khigashiguchi/portfolio/app/domain/model"
)

func GetBlogs(db *sql.DB) (model.BlogCollection, error) {
	result := model.BlogCollection{}

	sql := "SELECT id, title, url, published FROM blogs"
	rows, err := db.Query(sql)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		blog := model.Blog{}
		err := rows.Scan(&blog.ID, &blog.Title, &blog.URL, &blog.Published)
		if err != nil {
			return result, err
		}
		result.Blogs = append(result.Blogs, blog)
	}
	return result, nil
}

func PostBlog(db *sql.DB, title string, url string, published string) (int64, error) {
	sql := "INSERT INTO blogs(title, url, published) VALUES(?, ?, ?)"

	stmt, err := db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(title, url, published)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}
