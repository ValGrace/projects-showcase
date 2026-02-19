package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:255;not null;unique" json:"name"`
	Description string `gorm:"size:255;not null" json:"description"`
}

func CreateRole(role *Role) {
	db.Create(role)
}

func GetAllRoles(roles *[]Role) {
	db.Find(roles)
}

func GetRoleByID(id uint, role *Role) (err error) {
	err = db.Where("id = ?", id).First(role).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateRole(role *Role) {
	db.Save(role)
}
