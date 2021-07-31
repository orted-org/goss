package http

import (
	"github.com/labstack/echo/v4"
)

func RoutesInit(e *echo.Echo) {
	e.POST("/create", CreateSession)
	e.GET("/get", GetSession)
	e.DELETE("/delete", DeleteSession)
	e.DELETE("/truncate", TruncateStore)
}
