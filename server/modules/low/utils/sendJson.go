package utils

import (
	"encoding/json"
	"net/http"
)

type ApiRsponse struct {
	Success  bool        `json:"success"`
	SendData interface{} `json:"sended_data"`
	Status   int         `json:"status"`
	Message  string      `json:"message"`
}

func SendResponse(
	w http.ResponseWriter,
	statusCode int,
	data interface{},
	mesage string) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := ApiRsponse{
		Success:  statusCode >= 200 && statusCode <= 300,
		SendData: data,
		Status:   statusCode,
		Message:  mesage,
	}
	er := json.NewEncoder(w).Encode(response)
	if er != nil {
		return er
	}
	return nil
}
