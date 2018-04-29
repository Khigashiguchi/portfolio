package main

import (
	"github.com/labstack/echo"
	"golang.org/x/crypto/acme/autocert"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	//e.Static("/", "public")
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `
<h1>Welcome to Portfolio!</h1>
<h3>TLS certificates automatically installed.
`)
	})
	//e.Start(":8080")
	e.Logger.Fatal(e.StartAutoTLS(":4433"))
}
