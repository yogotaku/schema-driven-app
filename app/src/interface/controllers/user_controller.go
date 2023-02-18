package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yogotaku/schema-driven-app/app/src/interface/gateways"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateUser")
}

func (c *UserController) FindUserByID(w http.ResponseWriter, r *http.Request, userId int) {
	user := gateways.FindUserByID(userId)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (c *UserController) UpdateUserByID(w http.ResponseWriter, r *http.Request, userId int) {
	fmt.Println("UpdateUserByID")
}
