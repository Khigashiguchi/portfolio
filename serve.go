package main

import (
    "github.com/labstack/echo"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "github.com/Khigashiguchi/profile-site/handlers"
)

func main() {
    db := initDB("storage.db")
    migrate(db)

    e := echo.New()

    // Routes
    e.Static("/", "public")
    e.GET("/blogs", handlers.GetBlogs(db))
    e.POST("/blogs", handlers.PostBlog(db))

    // Start server
    e.Logger.Fatal(e.Start(":8080"))
}

func initDB(filepath string) *sql.DB {
    db, err := sql.Open("sqlite3", filepath)

    if err != nil {
        panic(err)
    }

    if db == nil {
        panic("database is nil")
    }

    return db
}

func migrate(db *sql.DB) {
    sql := `
	CREATE TABLE IF NOT EXISTS blogs(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title VARCHAR NOT NULL,
		url VARCHAR NOT NULL,
		published DATETIME
	);
	`

    _, err := db.Exec(sql)

    if err != nil {
        panic(err)
    }
}
