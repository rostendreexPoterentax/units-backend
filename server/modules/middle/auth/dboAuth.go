package auth

import (
	"sever/modules/middle/low/database"
	"time"
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
