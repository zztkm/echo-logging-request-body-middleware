package main

import (
	"bytes"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func logBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		buf, _ := io.ReadAll(c.Request().Body)
		c.Logger().Info(string(buf))
		c.Request().Body = io.NopCloser(bytes.NewBuffer(buf))
		return next(c)
	}
}

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Use(logBody)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"message": "ショートカットに捧げよ"})
	})
	e.Logger.Fatal(e.Start(":1323"))
}
