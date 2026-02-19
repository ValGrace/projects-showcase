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
	return d
}

func GetDB() *gorm.DB {
	db = Connect()
	return db
}

func ConnectDatabase() {
	Connect()
	db = GetDB()
	println(db)
	db.AutoMigrate(&models.Project{}, &models.User{}, &models.Role{})
	Seeding()
}

func Seeding() {
	var user = []models.User{{Name: os.Getenv("ADMIN_USERNAME"), Email: os.Getenv("ADMIN_EMAIL"), Password: os.Getenv("ADMIN_PASSWORD"), RoleID: 1, Projects: []models.Project{}}}
	var roles = []models.Role{{Name: "Admin", Description: "Admin has all the access"}, {Name: "User", Description: "User can only view the projects"}, {Name: "Anonymous", Description: "Unregistered user can only view the projects"}}
	db.Save(&roles)
	db.Save(&user)
}
