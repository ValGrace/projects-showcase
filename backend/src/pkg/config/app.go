package config

import (
	"os"

	"github.com/ValGrace/projects-showcase/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() *gorm.DB {
	dsn := "root:Cloud@WebDev23@tcp(127.0.0.1:3306)/projects_showcase?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	println("Connection to database successful!!!☁️")
	db = d
	return db
}

func GetDB() *gorm.DB {
	// db = Connect()
	return db
}

func ConnectDatabase() {
	Connect()
	db = GetDB()
	println(db)
	err := db.AutoMigrate(&models.Project{}, &models.User{}, &models.Role{})
	if err != nil {
		panic(err)
	}
	Seeding()
}

func Seeding() {
	var user = []models.User{{Name: os.Getenv("ADMIN_USERNAME"), Email: os.Getenv("ADMIN_EMAIL"), Password: os.Getenv("ADMIN_PASSWORD"), RoleID: 1, Projects: []models.Project{}}}
	var roles = []models.Role{{Name: "admin", Description: "Admin has all the access"}, {Name: "User", Description: "User can only view the projects"}, {Name: "Anonymous", Description: "Unregistered user can only view the projects"}}
	// models.CreateRole(&roles[0])
	// (*models.User).CreateUser(&user[0])
	db.Save(&roles)
	db.Save(&user)
}
