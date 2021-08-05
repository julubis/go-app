package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Version(context echo.Context) error {
	return context.JSON(http.StatusOK, map[string]interface{}{"version": 1})
}

func main() {
	e := echo.New()
	e.GET("/", Version)
	address := fmt.Sprintf("%s:%s", "localhost", "3000")
	fmt.Println(address)
	e.Start(address)

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	// e.Logger.Fatal(e.Start(""))
}
