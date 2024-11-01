package user

import (
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	UserService UserService
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// user struct
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	// decode the request body into user struct
	err := json.NewDecoder(r.Body).Decode(&user)
	// validation
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// call the service to create a user
	err = h.UserService.CreateUser(user.Username, user.Password, user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// return the response
	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.UserService.ListUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// return the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
