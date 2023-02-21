package data

type Pet struct {
	ID   int `json:"id"`
	Tag  string `json:"tag"`
	Name string `json:"name"`
}

var Pets []Pet = []Pet{
	{
		ID:   1,
		Tag:  "dog",
		Name: "Pochi",
	},
	{
		ID:   2,
		Tag:  "cat",
		Name: "Tama",
	},
	{
		ID:   3,
		Tag:  "dog",
		Name: "Shiro",
	},
	{
		ID:   4,
		Tag:  "hamster",
		Name: "Hamkichi",
	},
}
