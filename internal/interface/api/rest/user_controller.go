package rest

import "C"
import (
	"Or/Library/internal/application/user"
	userDto "Or/Library/internal/interface/api/dtos/user"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type UserController struct {
	userService *user.UserService
}

func NewUserController(userService *user.UserService) *UserController {
	return &UserController{userService}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDTO userDto.UserDTO
	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	u := userDto.MapToDomainUserFromDTO(&userDTO)
	err = c.userService.CreateUser(r.Context(), u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	u, err := c.userService.GetUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	userDTO := userDto.MapToUserDTO(u)
	_ = json.NewEncoder(w).Encode(userDTO)
}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var userDTO userDto.UserDTO
	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	u := userDto.MapToDomainUserFromDTO(&userDTO)
	u.ID = id
	err = c.userService.UpdateUser(r.Context(), u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := c.userService.DeleteUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *UserController) GetUserLoans(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	loans, err := c.userService.GetUserLoans(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(loans)
}
