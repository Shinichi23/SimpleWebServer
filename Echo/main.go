package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()

	e.GET("/", func(ctx echo.Context) {
		return ctx.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
