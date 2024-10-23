package auth

type DtoRegister struct {
	Name     string `json:"Name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
