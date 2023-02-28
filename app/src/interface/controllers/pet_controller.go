package controllers

// import (
// 	"net/http"

// 	"github.com/yogotaku/schema-driven-app/app/src/interface/gateways"
// 	"github.com/yogotaku/schema-driven-app/app/src/schema"
// )

// type PetController struct{}

// func NewPetController() *PetController {
// 	return &PetController{}
// }

// func (c *PetController) FindPets(w http.ResponseWriter, r *http.Request, params schema.FindPetsParams) {
// 	var tags []string

// 	if params.Tags != nil {
// 		tags = *params.Tags
// 	}

// 	ps := gateways.FindPets(tags, params.Limit)

// 	res := make([]schema.Pet, 0, len(ps))
// 	for i := range ps {
// 		p := ps[i]
// 		res = append(res, schema.Pet{
// 			Id:   p.ID,
// 			Name: p.Name,
// 			Tag:  p.Tag,
// 		})
// 	}

// 	schema.RenderJSONResponse(w, http.StatusOK, res)
// }
