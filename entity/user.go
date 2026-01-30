package entity

type User struct {
	ID          uint
	PhoneNumber string
	Name        string
	// password always keep hashed password.
	Password string
}
