package models

import (
	"os"

	"github.com/ValGrace/projects-showcase/src/pkg/config"
	"gorm.io/gorm"
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Project struct {
	gorm.Model
	Title       string `json:"title"`
	Livelink    string `json:"livelink"`
	Gitlink     string `json:"gitlink"`
	Description string `json:"description"`
	Photo       string `json:"photo"`
	Problem     string `json:"problem"`
	Solution    string `json:"solution"`
	Tech        string `json:"tech"`
	UserID      uint
}

func init() {
	config.Connect()
	db = config.GetDB()
	println(db)
	db.AutoMigrate(&Project{})
	db.AutoMigrate(&User{})
	var user = []User{{Name: os.Getenv("ADMIN_USERNAME"), Email: os.Getenv("ADMIN_EMAIL"), Password: os.Getenv("ADMIN_PASSWORD"), RoleID: 1, Projects: []Project{}}}
	var roles = []Role{{Name: "Admin", Description: "Admin has all the access"}, {Name: "User", Description: "User can only view the projects"}, {Name: "Anonymous", Description: "Unregistered user can only view the projects"}}
	db.Save(&roles)
	db.Save(&user)

}

func (b *Project) CreateProject() *Project {
	// db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllProjects() []Project {
	var Projects []Project
	db.Find(&Projects)
	return Projects
}

func GetProjectById(id int64) (*Project, *gorm.DB) {
	var getProject Project
	db := db.Where("ID=?", id).Find(&getProject)
	return &getProject, db
}

func DeleteProject(ID int64) Project {
	var project Project
	db.Where("ID=?", ID).Delete(project)
	return project
}
