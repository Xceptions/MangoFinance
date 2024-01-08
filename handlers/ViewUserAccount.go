package handlers

import (
	"encoding/json"
	"errors"

	"net/http"

	"gorm.io/gorm"
	"mangofinance.com/bank-backend/helpers"
	"mangofinance.com/bank-backend/models"
)

// Receives: response and request writers
// function receives input from the user containing the access
// token. We then check if the token is authorized. Retrieve
// the username from it, retrieve the user id from the
// username, and use that to retrieve the account detail
// Returns: json containing account detail
func (h handler) ViewUserAccount(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	userToken := r.Header.Get("x-access-token")

	var user models.User
	var account models.Account

	userName := helpers.IsAuthorized(w, userToken)

	findUser := h.DB.Where("Username = ?", userName).First(&user)
	if errors.Is(findUser.Error, gorm.ErrRecordNotFound) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not found")
	}

	getUserAccount := h.DB.Where("user_id = ?", user.ID).First(&account)
	if errors.Is(getUserAccount.Error, gorm.ErrRecordNotFound) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Account not found")
	}

	// Send a 201 retrieved response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}
