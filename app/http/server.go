package http

import (
	config2 "goss/app/config"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

var (
	SessionTtl = config2.DEF_SESSION_TTL
)
var Logger *log.Logger = log.Default()

func StartServer() {
	e := echo.New()
	e.GET("/dev", func(c echo.Context) error {
		return c.JSON(http.StatusOK, GenerateGeneralResponse(http.StatusOK, "Dev Endpoint Running!"))
	})
	RoutesInit(e)
	getSessionTTL, err := strconv.Atoi(os.Getenv("SESSION_ENV"))
	if err != nil {
		Logger.Print("Session TTL not detected in environment, set to default for 7 days for fallback if session ttl not provided in request")
	} else {
		SessionTtl = getSessionTTL
	}
	e.Logger.Fatal(e.Start(":4000"))
}
