package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/yogotaku/schema-driven-app/app/src/interface/gateways"
	"github.com/yogotaku/schema-driven-app/app/src/schema"
)

type PetController struct{}

func NewPetController() *PetController {
	return &PetController{}
}

func (c *PetController) FindPets(w http.ResponseWriter, r *http.Request, params schema.FindPetsParams) {
	var tags []string

	if params.Tags != nil {
		tags = *params.Tags
	}

	pets := gateways.FindPets(tags, params.Limit)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pets)
}
