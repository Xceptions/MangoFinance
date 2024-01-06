package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"mangofinance.com/bank-backend/helpers"
	"mangofinance.com/bank-backend/models"
)

// function receives input from the user containing username, email
// and password. We then use these values to populate both the
// user table and account table. The account is generated from the
// user detail.
// The return value is a string
func (h handler) CreateUserAccount(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	var user models.User
	json.Unmarshal(body, &user)

	generatePassword := helpers.HashAndSalt([]byte(user.Password))
	user.Password = generatePassword
	h.DB.Create(&user)

	account := models.Account{
		Type:    "Daily Account",
		Name:    string(user.Username + "'s" + " account"),
		Balance: uint(10000),
		UserID:  user.ID}
	h.DB.Create(&account)

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
