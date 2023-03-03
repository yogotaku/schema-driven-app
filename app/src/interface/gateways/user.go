package gateways

import (
	"time"

	"github.com/yogotaku/schema-driven-app/app/src/data"
)

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
	user.CreateDate = time.Now()
	users = append(users, user)
	return user
}

func UpdateUser(user data.User) data.User {
	users := data.Users
	for i := range users {
		if users[i].ID == user.ID {
			users[i] = user
			break
		}
	}
	return user
}
