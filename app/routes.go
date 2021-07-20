package main

import (
	"github.com/labstack/echo/v4"
)

func RoutesInit(e *echo.Echo) {
	e.POST("/create-session", CreateSession)
	e.GET("/get-session", GetSession)
	e.DELETE("/delete-session", DeleteSession)
}
