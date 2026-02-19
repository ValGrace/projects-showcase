package models

import (
	// "github.com/ValGrace/portfolio-server/src/pkg/config"
	// "github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	UserID   uint   `gorm:"primary_key"`
	Name     string `gorm:"size:255"`
	Email    string `gorm:"size:255"`
	Level    string `gorm:"size:255"`
	Password string `gorm:"size:255"`
	RoleID   uint
	Projects []Project
}

func (usr *User) CreateUser() *User {

	db.Create(&usr)
	return usr
}

func (usr *User) HashPass(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	usr.Password = string(passwordHash)
	return nil
}

func (usr *User) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(password))
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

func UpdateUser(User *User) {
	db.Omit("password").Updates(User)

	return
}

func GetUserByUsername(uname string) User {
	var user User
	db.Where("name=?", uname).Find(&user)

	return user
}
