package handlers

import (
	"database/sql"
	"github.com/labstack/echo"
	"net/http"
	"github.com/Khigashiguchi/profile-site/models"
)

func GetBlogs(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetBlogs(db))
	}
}