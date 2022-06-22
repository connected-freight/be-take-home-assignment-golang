package question_1

import "time"

type UserInput struct {
	Name         string
	PhoneNumber  string
	EmailAddress string
	DateOfBirth  *string
}

type User struct {
	ID           string
	Name         string
	PhoneNumber  string
	EmailAddress string
	DateOfBirth  time.Time
	AgeToday     int
}
