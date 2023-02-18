package gateways

import "github.com/yogotaku/schema-driven-app/app/src/data"

func FindUserByID(id int) data.User {
	var user data.User

	for i := range data.Users {
		u := data.Users[i]
		if id == u.ID {
			user = u
			break
		}
	}

	return user
}
