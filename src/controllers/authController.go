package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/LordRahl90/little_quiz_backend/src/models"
	"github.com/LordRahl90/little_quiz_backend/src/utils"
)

//CreateAccount - function to create a new account.
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid Request"))
		return
	}
	response := account.Create()
	utils.Respond(w, response)
}

//Login - function to log into the application
func Login(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid Request"))
		return
	}

	resp := models.Login(account.Email, account.Password)
	log.Println(resp)
	utils.Respond(w, resp)
}
