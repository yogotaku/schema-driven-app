package data

type User struct {
	ID            int
	FirstName     string
	LastName      string
	Email         string
	DateOfBirth   string
	EmailVerified bool
	CreatedDate   string
}

var Users []User = []User{
	{
		ID:            1,
		FirstName:     "Ichiro",
		LastName:      "Ichinose",
		Email:         "ichiro.ichinose@example.com",
		DateOfBirth:   "2001-01-01",
		EmailVerified: true,
		CreatedDate:   "2021-01-01",
	},
	{
		ID:            2,
		FirstName:     "Jiro",
		LastName:      "Ninomiya",
		Email:         "jiro.ninomiya@example.com",
		DateOfBirth:   "2002-02-02",
		EmailVerified: true,
		CreatedDate:   "2022-02-02",
	},
	{
		ID:            3,
		FirstName:     "Saburo",
		LastName:      "Sanbonmatsu",
		Email:         "saburo.sanbonmatsu@example.com",
		DateOfBirth:   "2003-03-03",
		EmailVerified: false,
		CreatedDate:   "",
	},
}
