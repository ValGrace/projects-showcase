package models

import (
	"gorm.io/gorm"
)

var db *gorm.DB

type Project struct {
	gorm.Model
	Title       string `gorm:"title"`
	Livelink    string `gorm:"livelink"`
	Gitlink     string `gorm:"gitlink"`
	Description string `gorm:"description"`
	Photo       string `gorm:"photo"`
	Problem     string `gorm:"problem"`
	Solution    string `gorm:"solution"`
	Tech        string `gorm:"tech"`
	UserID      uint
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
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
