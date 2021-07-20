package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

var (
	SESSION_TTL = int(time.Hour * 24 * 7)
)

var RedisClient *redis.Client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

type SessionData struct {
	SessionID string      `json:"sessionId"`
	Session   interface{} `json:"session"`
}

func main() {
	e := echo.New()
	e.GET("/dev", func(c echo.Context) error {
		return c.JSON(http.StatusOK, GenerateGeneralResponse(http.StatusOK, "Dev Endpoint Running!"))
	})
	RoutesInit(e)
	getSessionTTL, err := strconv.Atoi(os.Getenv("SESSION_ENV"))
	if err != nil {
		e.Logger.Error("TTL not set, configured with default TTL of 7 days")
		getSessionTTL = int(time.Hour * 24 * 7)
	} else {
		SESSION_TTL = getSessionTTL
	}
	e.Logger.Fatal(e.Start(":4000"))
}
