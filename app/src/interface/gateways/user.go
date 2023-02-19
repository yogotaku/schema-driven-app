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

func CreateUser(user data.User) data.User {
	users := data.Users
	lastUser := users[len(users)-1]
	user.ID = lastUser.ID + 1
	users = append(users, user)
	return user
}
