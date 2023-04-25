package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	controllers "shorturl/Controllers"
)

func main() {
	fmt.Println("hello")

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/api/link/:link", controllers.GetLink)
	e.POST("/api/link", controllers.SetLink)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
