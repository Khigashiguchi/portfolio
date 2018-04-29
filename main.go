package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// Routes
	e.Static("/", "public")

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
