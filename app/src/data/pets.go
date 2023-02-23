package data

type Pet struct {
	ID   int
	Tag  string
	Name string
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
		Name: "Hamham",
	},
}
