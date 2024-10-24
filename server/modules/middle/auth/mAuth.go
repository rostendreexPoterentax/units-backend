package auth

import (
	"time"

	"github.com/rotisserie/eris"
	"golang.org/x/crypto/bcrypt"
)

// addNewUser Регистрация пользователя. Добавляет его в бд.
func addNewUser(newUser *DtoRegister) (int, error) {
	hashedPassword, er := hashPassword(newUser.Password)
	if er != nil {
		return -1, eris.Wrap(er, "failed to create hashed password")
	}
	user := &DboUser{
		Name:           newUser.Name,
		Email:          newUser.Email,
		Password:       hashedPassword,
		CreationMoment: time.Now().UTC(),
	}
	newUId, er := AddUser(user)
	if er != nil {
		return -1, eris.Wrap(er, "failed to add new user to database")
	}
	return newUId, nil
}
//	loginUser Авторизация юзера
func loginUser(loginData *DtoAuthUserRequest) (*DtoAuthUserResponse, error) {
	user, er := GetUser(loginData.Email)
	if er != nil {
		return nil, er
	}
	er = checkPassword(loginData.Password)
	if er != nil {
		return nil, er
	}

	token, er := GenerateToken(user.Id, user.Email)
	if er != nil {
		return nil, eris.Wrapf(er, "failed to generate token for user with email: %v", user.Email)
	}
	return &DtoAuthUserResponse{
		UserId:    user.Id,
		UserToken: token,
	}, nil
}

// hashPassword Получает пароль из запроса при регистрации и возвращает его hash-версию.
func hashPassword(dtoPassword string) (string, error) {
	hashedPassword, er := bcrypt.GenerateFromPassword([]byte(dtoPassword), bcrypt.DefaultCost)
	if er != nil {
		return "", er
	}
	return string(hashedPassword), nil
}

func checkPassword(password string) error {
	hashedPassword, er := hashPassword(password)
	if er != nil {
		return er
	}
	er = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if er != nil {
		return er
	}
	return nil
}
