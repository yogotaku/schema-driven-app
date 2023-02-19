package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/yogotaku/schema-driven-app/app/src/data"
	"github.com/yogotaku/schema-driven-app/app/src/interface/gateways"
	"github.com/yogotaku/schema-driven-app/app/src/schema"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var body schema.NewUser
	json.NewDecoder(r.Body).Decode(&body)

	u := data.User{
		FirstName:   body.FirstName,
		LastName:    body.LastName,
		Email:       string(body.Email),
		DateOfBirth: body.DateOfBirth.Format(time.DateOnly),
	}

	user := gateways.CreateUser(u)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (c *UserController) FindUserByID(w http.ResponseWriter, r *http.Request, userId int) {
	user := gateways.FindUserByID(userId)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (c *UserController) UpdateUserByID(w http.ResponseWriter, r *http.Request, userId int) {
	var body schema.NewUser
	json.NewDecoder(r.Body).Decode(&body)

	u := gateways.FindUserByID(userId)
	u.FirstName = body.FirstName
	u.LastName = body.LastName
	u.Email = string(body.Email)
	u.DateOfBirth = body.DateOfBirth.Format(time.DateOnly)

	gateways.UpdateUser(u)

	w.WriteHeader(http.StatusNoContent)
}
