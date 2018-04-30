package handlers

import (
	"database/sql"
	"github.com/labstack/echo"
	"net/http"
	"github.com/Khigashiguchi/profile-site/models"
)

type H map[string]interface{}

func GetBlogs(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetBlogs(db))
	}
}

func PostBlog(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        var blog models.Blog
        c.Bind(&blog)

        id, err := models.PostBlog(db, blog.Title, blog.URL, blog.Published)
        if err != nil {
            return err
        }
        return c.JSON(http.StatusCreated, H{
            "creted": id,
        })
    }
}
