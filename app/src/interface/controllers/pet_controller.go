package controllers

import (
	"fmt"
	"net/http"

	"github.com/yogotaku/schema-driven-app/app/src/schema"
)

type PetController struct{}

func NewPetController() *PetController {
	return &PetController{}
}

func (c *PetController) FindPets(w http.ResponseWriter, r *http.Request, params schema.FindPetsParams) {
	fmt.Println("FindPets")
}
