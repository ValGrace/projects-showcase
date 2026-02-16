package models

import (
	// "github.com/ValGrace/portfolio-server/src/pkg/config"
	// "github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	UserID   uint   `gorm:"primary_key"`
	Name     string `gorm:"size:255"`
	Email    string `gorm:"size:255"`
	Level    string `gorm:"size:255"`
	Password string `gorm:"size:255"`
	Projects []Project
}

func (usr *User) CreateUser() *User {
	db.NewRecord(usr)
	db.Create(&usr)
	return usr
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserProjects() ([]User, error) {
	var users []User
	err := db.Model(&User{}).Preload("Project").Find(&users).Error
	return users, err
}

func GetUserById(id int64) (*User, *gorm.DB) {
	var getUser User
	usrdb := db.Where("UserID=?", id).Find(&getUser)
	return &getUser, usrdb
}

func DeleteUser(ID int64) User {
	var user User
	db.Where("UserID=?", ID).Delete(user)
	return user
}
