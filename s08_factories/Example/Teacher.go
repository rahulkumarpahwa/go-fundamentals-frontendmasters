package teacher

type Teacher struct {
	Id        int
	FirstName string // IMPORTANT : note when we have export to another package every value inside the struct we need to write in capital first letter.
	LastName  string
	Score     int
}

func New(firstName string, lastName string) Teacher {
	return Teacher{FirstName: firstName, LastName: lastName}
}

