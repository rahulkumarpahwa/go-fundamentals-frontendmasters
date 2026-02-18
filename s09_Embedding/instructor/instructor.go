package instructor


type Instructor struct {
	Id        int
	FirstName string // IMPORTANT : note when we have export to another package every value inside the struct we need to write in capital first letter.
	LastName  string
	Score     int
}