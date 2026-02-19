package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ValGrace/portfolio-server/src/models"
	"github.com/ValGrace/portfolio-server/src/utils"
	"github.com/gorilla/mux"
)

// var newProject models.Project

// var newSkill models.Skillset

func GetUser(w http.ResponseWriter, r *http.Request) {
	newUsers := models.GetAllUsers()
	res, _ := json.Marshal(newUsers)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	userDetails, _ := models.GetUserById(ID)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}

	utils.Authenticate(w, r, CreateUser)

	b := CreateUser.CreateUser()
	fmt.Print(b)
	res, err := json.Marshal(b)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Print(http.StatusCreated)
	w.Write(res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("success --------")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	user := models.DeleteUser(ID)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser = &models.User{}
	// var updateSkill = &models.Skillset{}
	utils.ParseBody(r, updateUser)
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userDetails, db := models.GetUserById(ID)

	if updateUser.Name != "" {
		userDetails.Name = updateUser.Name
	}
	if updateUser.Email != "" {
		userDetails.Email = updateUser.Email
	}
	if updateUser.Level != "" {
		userDetails.Level = updateUser.Level
	}
	// if updateUser.Photo != []byte {
	// 	userDetails.Photo = updateUser.Photo
	// }
	// if updateUser.Tech != []string {
	// 	userDetails.Tech = updateUser.Tech
	// }

	db.Save(&userDetails)
	res, _ := json.Marshal(userDetails)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
