package data

type User struct {
	ID            int    `json:"id"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Email         string `json:"email"`
	DateOfBirth   string `json:"dateOfBirth"`
	EmailVerified bool   `json:"emailVerified"`
	CreateDate    string `json:"createDate"`
}

var Users []User = []User{
	{
		ID:            1,
		FirstName:     "Ichiro",
		LastName:      "Ichinose",
		Email:         "ichiro.ichinose@example.com",
		DateOfBirth:   "2001-01-01",
		EmailVerified: true,
		CreateDate:    "2021-01-01",
	},
	{
		ID:            2,
		FirstName:     "Jiro",
		LastName:      "Ninomiya",
		Email:         "jiro.ninomiya@example.com",
		DateOfBirth:   "2002-02-02",
		EmailVerified: true,
		CreateDate:    "2022-02-02",
	},
	{
		ID:            3,
		FirstName:     "Saburo",
		LastName:      "Sanbonmatsu",
		Email:         "saburo.sanbonmatsu@example.com",
		DateOfBirth:   "2003-03-03",
		EmailVerified: false,
		CreateDate:    "",
	},
}
