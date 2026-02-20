package models

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	ID       uint   `gorm:"primary_key"`
	Name     string `gorm:"size:255"`
	Email    string `gorm:"size:255"`
	Password string `gorm:"size:255"`
	RoleID   uint
	Role     Role      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Projects []Project `gorm:"foreignKey:UserID"`
}

func (usr *User) CreateUser() *User {
	ctx := context.Background()

	result := gorm.WithResult()
	gorm.G[User](db, result).Create(ctx, usr)

	fmt.Println("User creation result:", usr.ID)

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

}

func GetUserByUsername(uname string) (User, error) {
	var user User
	err := db.Where("name=?", uname).Find(&user).Error

	if err != nil {
		return User{}, err
	}
	return user, err
}
