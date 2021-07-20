package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

func CreateSession(c echo.Context) error {
	var body SessionData
	// creating a uuid for session
	uuid := uuid.NewV4()
	body.SessionID = uuid.String()
	//getting the body of the request
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, GenerateGeneralResponse(http.StatusBadRequest, "Invalid Data For Session"))
	}
	// if session is not pressent in the body
	if body.Session == nil {
		return c.JSON(http.StatusBadRequest, GenerateGeneralResponse(http.StatusBadRequest, "Please Provide Session Data"))
	}
	// converting the json body to struct
	value, err := json.Marshal(body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, GenerateGeneralResponse(http.StatusBadRequest, "Invalid Data For Session"))
	}

	// storing the session into redis store
	err = RedisClient.Set(context.Background(), uuid.String(), string(value), time.Duration(SESSION_TTL)).Err()
	if err != nil {
		return c.JSON(http.StatusOK, GenerateGeneralResponse(http.StatusInternalServerError, "Something Went Wrong"))
	}
	return c.JSON(http.StatusCreated, body)
}
func GetSession(c echo.Context) error {
	// getting the session id from the request
	sessionId := c.QueryParam("sessionId")

	// checking if session id is present
	if len(sessionId) == 0 {
		return c.JSON(http.StatusBadRequest, GenerateGeneralResponse(http.StatusBadRequest, "Session ID missing from query parameter"))
	}
	// gettin the data from redis store
	res, err := RedisClient.Get(context.Background(), sessionId).Result()
	if err != nil {
		return c.JSON(http.StatusNotFound, GenerateGeneralResponse(http.StatusNotFound, "Session Not Found"))
	}
	// converting the redis value string to json format
	var session SessionData
	err = json.Unmarshal([]byte(res), &session)
	if err != nil {
		fmt.Print("Json error")
	}
	return c.JSON(http.StatusOK, session)
}

func DeleteSession(c echo.Context) error {
	// getting the session id from the request
	sessionId := c.QueryParam("sessionId")

	// checking if session id is present
	if len(sessionId) == 0 {
		return c.JSON(http.StatusBadRequest, GenerateGeneralResponse(http.StatusBadRequest, "Session ID missing from query parameter"))
	}
	// removing session from the redis store
	err := RedisClient.Del(context.Background(), sessionId).Err()
	if err != nil {
		return c.JSON(http.StatusOK, GenerateGeneralResponse(http.StatusInternalServerError, "Session Could Not Be Deleted"))
	}
	return c.JSON(http.StatusOK, GenerateGeneralResponse(http.StatusOK, "Session Deleted Successfully"))
}
