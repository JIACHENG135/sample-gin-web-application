package dao

import (
	db "sample-gin-web-app/database"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserName string
	Age      int32
	Gender   string
}

func (User) TableName() string {
	//user对象和simple_user表绑定
	return "simple_user"
}

func GetUserByID(userID int) (target User) {
	var user User
	db.Conn.Where("id = ?", userID).First(&user)
	return user
}
