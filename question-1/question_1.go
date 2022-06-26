package question_1

import "time"

func MapUserInputToDomainModel(input UserInput) (User, error) {
	// TODO: Implement and replace the return object
	return User{
		ID:           "user@example.com",
		Name:         "User 1",
		PhoneNumber:  "1234567890",
		EmailAddress: "user@example.com",
		DateOfBirth:  time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
	}, nil
}
