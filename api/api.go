package api // 声明 包 api

import (
	"net/http"

	"github.com/labstack/echo"
)

func Http()  {
	println("it's module api func Http")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello, api http!!")
	})
	e.Logger.Fatal(e.Start(":8001"))
}