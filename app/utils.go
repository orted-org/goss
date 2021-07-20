package main

type GeneralResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func GenerateGeneralResponse(status int, message string) *GeneralResponse {
	return &GeneralResponse{
		Status:  status,
		Message: message,
	}
}
