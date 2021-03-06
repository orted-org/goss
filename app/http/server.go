package http

import (
	"log"
	"net/http"
	"os"
	"strconv"

	config "github.com/himanshu-sah/goss/app/config"
	"github.com/himanshu-sah/goss/app/utils"

	"github.com/labstack/echo/v4"
)

var (
	SessionTtl = config.DEF_SESSION_TTL
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
		SessionTtl = utils.GetTimeFor(getSessionTTL)
	}
	e.Logger.Fatal(e.Start(":1211"))
}
