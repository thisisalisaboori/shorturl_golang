package main

import (
	"fmt"
	controllers "shorturl/Controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
)

func main() {
	fmt.Println("hello")

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	// Routes
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/api/link/:link", controllers.GetLink)
	e.POST("/api/link", controllers.SetLink)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
