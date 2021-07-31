package http

import (
	"encoding/json"
	"net/http"

	"github.com/himanshu-sah/goss/app/service"
	"github.com/himanshu-sah/goss/app/utils"

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
	// setting the default session in case neither ENV is set nor request contains the TTL
	var ttl = SessionTtl
	if body.TTL != 0 {
		// if session is provided in the request
		ttl = utils.GetTimeFor(body.TTL)
	}

	// converting the session data to json
	session, _ := json.Marshal(body.Session)

	// storing the session into redis store
	sessionId, err := service.CreateSession(string(session), ttl)
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
	// getting the data from redis store
	res, err := service.GetSession(sessionId)
	if err != nil {
		return c.JSON(http.StatusNotFound, GenerateGeneralResponse(http.StatusNotFound, err.Error()))
	}
	// converting the redis value string to json format
	var session interface{}
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
func TruncateStore(c echo.Context) error {
	err := service.TruncateStore()
	if err != nil {
		c.JSON(http.StatusInternalServerError, GenerateGeneralResponse(http.StatusInternalServerError, err.Error()))
	}
	return c.JSON(http.StatusOK, GenerateGeneralResponse(http.StatusOK, "truncated session store"))
}
