package auth

import (
	"encoding/json"
	"net/http"
	"sever/modules/low/utils"
)

func ApiAuth(w http.ResponseWriter, r *http.Request) {
	var req DtoAuthUserRequest
	er := json.NewDecoder(r.Body).Decode(&req)
	if er != nil {
		utils.SendResponse(w, 502, nil, "Произошла ошибка при попытке распарсить json. Попробуйте позже")
	}
	response, er := loginUser(&req)
	if er != nil {
		utils.SendResponse(w, 502, nil, "Произошла ошибка Попробуйте позже")
	}
	utils.SendResponse(w, 200, response, "success")
}
