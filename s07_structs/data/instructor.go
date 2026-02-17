package data

// capital name so we can export
type Instructor struct {
	Id        int
	FirstName string // IMPORTANT : note when we have export to another package every value inside the struct we need to write in capital first letter.
	LastName  string
	Score     int
}

// Note: avoid making a type public until needed to avoid exposing unnecessary package details.

// Files in the same package can access each other's types and methods, regardless of which file they are in. The package is the key organizational unit, not individual files. Imports use the module name and package path to reference types from other packages.
