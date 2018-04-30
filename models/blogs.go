package models

import "database/sql"

type Blog struct {
	ID int `json:"id"`
	Title string `json:"title"`
	URL string `json:"url"`
	Published string `json:"published"`
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
		err := rows.Scan(&blog.ID, &blog.Title, &blog.URL, &blog.Published)
		if err != nil {
			panic(err)
		}
		result.Blogs = append(result.Blogs, blog)
	}
	return result
}

func PostBlog(db *sql.DB, title string, url string, published string) (int64, error) {
    sql := "INSERT INTO blogs(title, url, published) VALUES(?, ?, ?)"

    stmt, err := db.Prepare(sql)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()

    result, err := stmt.Exec(title, url, published)
    if err != nil {
        panic(err)
    }

    return result.LastInsertId()
}
