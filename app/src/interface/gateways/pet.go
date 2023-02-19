package gateways

import (
	"github.com/yogotaku/schema-driven-app/app/src/data"
)

func FindPets(tags []string, limit *int) []data.Pet {
	ts := make(map[string]struct{}, len(tags))
	for _, t := range tags {
		ts[t] = struct{}{}
	}

	var pets []data.Pet

	if len(ts) == 0 {
		pets = data.Pets
	} else {
		for i := range data.Pets {
			pet := data.Pets[i]
			if _, ok := ts[pet.Tag]; ok {
				pets = append(pets, pet)
			}
		}
	}

	if limit != nil {
		pets = pets[:*limit]
	}
	return pets
}
