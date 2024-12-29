package controllers

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/Dostonlv/http3-go/models"
	"github.com/google/uuid"
	
)

var (
	users = make(map[string]models.User)
	mu    sync.Mutex
)

// RegisterUser handles user registration
// @Summary Register a new user
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 201 {object} models.User
// @Failure 400 {string} string "Invalid input"
// @Router /register [post]
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	mu.Lock()
	defer mu.Unlock()

	user.ID = uuid.New().String()
	users[user.ID] = user

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUsers handles fetching all users
// @Summary Get all users
// @Tags User
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var userList []models.User
	for _, user := range users {
		userList = append(userList, user)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userList)
}
