package auth

import (
	"encoding/json"
	"net/http"
	"sever/modules/low/utils"
)

func ApiRegisterUser(w http.ResponseWriter, r *http.Request) {
	var req *DtoRegister

	er := json.NewDecoder(r.Body).Decode(req)
	if er != nil {
		http.Error(w, "invalid input data", http.StatusBadRequest)
		return
	}

	newUserId, er := addNewUser(req)
	if er != nil {
		badData := map[string]int{"id": -1}
		utils.SendResponse(w, 400, badData, "Произошла ошибка. Попробуйте позже.")
		return
	}

	data := map[string]int{"id": newUserId}

	er = utils.SendResponse(w, 200, data, "Аккаунт успешно зарегестрирован")
	if er != nil {
		badData := map[string]int{"id": -1}
		utils.SendResponse(w, 400, badData, "Произошла ошибка. Попробуйте позже.")
		return
	}
}
