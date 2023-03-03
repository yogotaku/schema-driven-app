package data

import "time"

type User struct {
	ID            int
	FirstName     string
	LastName      string
	Email         string
	DateOfBirth   time.Time
	EmailVerified bool
	CreateDate    time.Time
}

var Users []User = []User{
	{
		ID:            1,
		FirstName:     "Ichiro",
		LastName:      "Ichinose",
		Email:         "ichiro.ichinose@example.com",
		DateOfBirth:   time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
		EmailVerified: true,
		CreateDate:    time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		ID:            2,
		FirstName:     "Jiro",
		LastName:      "Ninomiya",
		Email:         "jiro.ninomiya@example.com",
		DateOfBirth:   time.Date(2002, 2, 2, 0, 0, 0, 0, time.UTC),
		EmailVerified: true,
		CreateDate:    time.Date(2002, 2, 2, 0, 0, 0, 0, time.UTC),
	},
	{
		ID:            3,
		FirstName:     "Saburo",
		LastName:      "Sanbonmatsu",
		Email:         "saburo.sanbonmatsu@example.com",
		DateOfBirth:   time.Date(2003, 3, 3, 0, 0, 0, 0, time.UTC),
		EmailVerified: false,
		CreateDate:    time.Date(2003, 3, 3, 0, 0, 0, 0, time.UTC),
	},
}
