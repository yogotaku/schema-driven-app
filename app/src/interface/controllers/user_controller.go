package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/types"
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
		DateOfBirth: body.DateOfBirth.Time,
	}

	u = gateways.CreateUser(u)

	res := schema.User{
		Id:            u.ID,
		FirstName:     u.FirstName,
		LastName:      u.LastName,
		Email:         types.Email(u.Email),
		DateOfBirth:   &types.Date{Time: u.DateOfBirth},
		EmailVerified: u.EmailVerified,
		CreateDate:    types.Date{Time: u.CreateDate},
	}

	schema.RenderJSONResponse(w, http.StatusCreated, res)
}

func (c *UserController) FindUserById(w http.ResponseWriter, r *http.Request, userId int) {
	u := gateways.FindUserByID(userId)
	res := schema.User{
		Id:            u.ID,
		FirstName:     u.FirstName,
		LastName:      u.LastName,
		Email:         types.Email(u.Email),
		DateOfBirth:   &types.Date{Time: u.DateOfBirth},
		EmailVerified: u.EmailVerified,
		CreateDate:    types.Date{Time: u.CreateDate},
	}

	schema.RenderJSONResponse(w, http.StatusOK, res)
}

func (c *UserController) UpdateUserByID(w http.ResponseWriter, r *http.Request, userId int) {
	var body schema.NewUser
	json.NewDecoder(r.Body).Decode(&body)

	u := gateways.FindUserByID(userId)
	u.FirstName = body.FirstName
	u.LastName = body.LastName
	u.Email = string(body.Email)
	u.DateOfBirth = body.DateOfBirth.Time

	gateways.UpdateUser(u)

	schema.RenderJSONResponse(w, http.StatusNoContent, nil)
}
