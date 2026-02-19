package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/ValGrace/projects-showcase/src/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func CreateRole(w http.ResponseWriter, r *http.Request) {
	CreateRoles := &models.Role{}

	json.NewDecoder(r.Body).Decode(CreateRoles)

	models.CreateRole(CreateRoles)
	res, _ := json.Marshal(CreateRoles)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func GetRoles(w http.ResponseWriter, r *http.Request) {
	newRoles := []models.Role{}
	models.GetAllRoles(&newRoles)
	res, _ := json.Marshal(newRoles)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetRole(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	role := models.Role{}
	err := models.GetRoleByID(uint(id), &role)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Role not found"))
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error retrieving role: " + err.Error()))
		return
	}
	res, _ := json.Marshal(role)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateRole(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	role := models.Role{}
	err := models.GetRoleByID(uint(id), &role)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Role not found"))
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error retrieving role: " + err.Error()))
		return
	}
	json.NewDecoder(r.Body).Decode(&role)
	models.UpdateRole(&role)
	res, _ := json.Marshal(role)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
