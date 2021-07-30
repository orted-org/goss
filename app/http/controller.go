package http

import (
	"encoding/json"
	"goss/app/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateSession(c echo.Context) error {
	var body SessionData

	//getting the body of the request and converting to struct
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, GenerateGeneralResponse(http.StatusBadRequest, "Invalid Data For Session"))
	}
	// if session is not pressent in the body
	if body.Session == nil {
		return c.JSON(http.StatusBadRequest, GenerateGeneralResponse(http.StatusBadRequest, "Please Provide Session Data"))
	}
	if body.TTL == 0 {
		body.TTL = SessionTtl
	}
	// converting the struct to json
	session, _ := json.Marshal(body)

	// storing the session into redis store
	sessionId, err := service.CreateSession(string(session), body.TTL)
	if err != nil {
		return c.JSON(http.StatusOK, GenerateGeneralResponse(http.StatusInternalServerError, err.Error()))
	}
	return c.JSON(http.StatusCreated, GenerateGeneralResponse(http.StatusCreated, sessionId))
}
func GetSession(c echo.Context) error {
	// getting the session id from the request
	sessionId := c.QueryParam("sessionId")

	// checking if session id is present
	if len(sessionId) == 0 {
		return c.JSON(http.StatusBadRequest, GenerateGeneralResponse(http.StatusBadRequest, "Session ID missing from query parameter"))
	}
	// gettin the data from redis store
	res, err := service.GetSession(sessionId)
	if err != nil {
		return c.JSON(http.StatusNotFound, GenerateGeneralResponse(http.StatusNotFound, err.Error()))
	}
	// converting the redis value string to json format
	var session SessionData
	json.Unmarshal([]byte(res), &session)

	// retuning json
	return c.JSON(http.StatusOK, session)
}

func DeleteSession(c echo.Context) error {
	// getting the session id from the request
	sessionId := c.QueryParam("sessionId")

	// checking if session id is present
	if len(sessionId) == 0 {
		return c.JSON(http.StatusBadRequest, GenerateGeneralResponse(http.StatusBadRequest, "Session ID missing from query parameter"))
	}

	err := service.DeleteSession(sessionId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, GenerateGeneralResponse(http.StatusInternalServerError, err.Error()))
	}
	return c.JSON(http.StatusOK, GenerateGeneralResponse(http.StatusOK, "Session Deleted Successfully"))
}
