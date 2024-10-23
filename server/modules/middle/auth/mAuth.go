package auth

import (
	"time"

	"github.com/rotisserie/eris"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(dtoPassword string) (string, error) {
	hashedPassword, er := bcrypt.GenerateFromPassword([]byte(dtoPassword), bcrypt.DefaultCost)
	if er != nil {
		return "", er
	}
	return string(hashedPassword), nil
}

func addNewUser(newUser *DtoRegister) (int, error) {
	hashedPassword, er := hashPassword(newUser.Password)
	if er != nil {
		return -1, eris.Wrap(er, "failed to create hashed password")
	}
	user := &DboUser{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: hashedPassword,
		CreationMoment: time.Now().UTC(),
	}
	newUId, er := AddUser(user) 
	if er != nil {
		return -1, eris.Wrap(er, "failed to add new user to database")
	}
	return newUId, nil
}
