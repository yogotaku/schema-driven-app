package controllers

import (
	"fmt"
	"net/http"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateUser")
}

func (c *UserController) FindUserByID(w http.ResponseWriter, r *http.Request, userId int) {
	fmt.Println("FindUserByID")
}

func (c *UserController) UpdateUserByID(w http.ResponseWriter, r *http.Request, userId int) {
	fmt.Println("UpdateUserByID")
}
