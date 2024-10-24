package auth

type DtoRegister struct {
	Name     string `json:"Name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type DtoAuthUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DtoAuthUserResponse struct {
	UserId    int    `json:"user_id"`
	UserToken string `json:user_token"`
}
