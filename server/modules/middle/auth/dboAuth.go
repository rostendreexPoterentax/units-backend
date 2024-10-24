package auth

import (
	"sever/modules/low/database"
	"time"

	"github.com/rotisserie/eris"
)

type DboUser struct {
	Id             int       `gorm:"primaryKey;column:id"`
	Name           string    `gorm:"column:name"`
	Email          string    `gorm:"column:email"`
	Password       string    `gorm:"column:hashed_password"`
	CreationMoment time.Time `gorm:"column:creation_moment"`
}

func AddUser(user *DboUser) (int, error) {
	res := database.Db.Table("dbo.user").Create(user)
	if res.Error != nil {
		return -1, res.Error
	}
	return user.Id, nil
}

func GetUser(email string) (*DboUser, error) {
	var user DboUser
	res := database.Db.Table("dbo.user").Find(&user, email)
	if res.Error != nil {
		return nil, eris.Wrapf(res.Error, "fauled to load user from database by email: %v", email)
	}
	return &user, nil
}
