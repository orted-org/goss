package http

type GeneralResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SessionData struct {
	Session interface{} `json:"session"`
	TTL     int         `json:"ttl"`
}

func GenerateGeneralResponse(status int, message string) *GeneralResponse {
	return &GeneralResponse{
		Status:  status,
		Message: message,
	}
}
