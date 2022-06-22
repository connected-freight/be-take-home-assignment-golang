package question_1_test

import (
	"reflect"
	"testing"
	"time"

	. "github.com/connected-freight/be-toke-home-assignment-golang/question-1"
)

func TestQuestion1(t *testing.T) {
	tests := []struct {
		it             string
		input          UserInput
		want           User
		wantAgeOfToday int
		wantErr        bool
	}{
		{
			it: "Should map emailAddress to id if both phoneNumber and emailAddress are available",
			input: UserInput{
				Name:         "John Doe",
				PhoneNumber:  "1234567890",
				EmailAddress: "john@example.com",
				DateOfBirth:  strPtr("1990-02-21"),
			},
			want: User{
				ID:           "john@example.com",
				Name:         "John Doe",
				PhoneNumber:  "1234567890",
				EmailAddress: "john@example.com",
				DateOfBirth:  time.Date(1990, 1, 21, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			it: "Should return correct ageToday based on the dateOfBirth",
			input: UserInput{
				Name:         "John Doe",
				PhoneNumber:  "1234567890",
				EmailAddress: "john@example.com",
				DateOfBirth:  strPtr("1990-02-21"),
			},
			wantAgeOfToday: 32,
		},
		{
			it: "Should map emailAddress to id if only emailAddress is available",
			input: UserInput{
				Name:         "John Doe",
				PhoneNumber:  "",
				EmailAddress: "john@example.com",
				DateOfBirth:  strPtr("1990-02-21"),
			},
			want: User{
				ID:           "john@example.com",
				Name:         "John Doe",
				PhoneNumber:  "",
				EmailAddress: "john@example.com",
				DateOfBirth:  time.Date(1990, 1, 21, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			it: "Should map phoneNumber to id if only phoneNumber is available",
			input: UserInput{
				Name:         "John Doe",
				PhoneNumber:  "1234567890",
				EmailAddress: "",
				DateOfBirth:  strPtr("1990-02-21"),
			},
			want: User{
				ID:           "1234567890",
				Name:         "John Doe",
				PhoneNumber:  "1234567890",
				EmailAddress: "",
				DateOfBirth:  time.Date(1990, 1, 21, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			it: "Should return error if name is an empty string",
			input: UserInput{
				Name:         "",
				PhoneNumber:  "1234567890",
				EmailAddress: "john@example.com",
				DateOfBirth:  strPtr("1990-02-21"),
			},
			wantErr: true,
		},
		{
			it: "Should return error if name is null",
			input: UserInput{
				Name:         "",
				PhoneNumber:  "1234567890",
				EmailAddress: "john@example.com",
				DateOfBirth:  strPtr("1990-02-21"),
			},
			wantErr: true,
		},
		{
			it: "Should return error if both phoneNumber and emailAddress are empty",
			input: UserInput{
				Name:         "John Doe",
				PhoneNumber:  "",
				EmailAddress: "",
				DateOfBirth:  strPtr("1990-02-21"),
			},
			wantErr: true,
		},
		{
			it: "Should return error if phoneNumber is not 10 characters long",
			input: UserInput{
				Name:         "John Doe",
				PhoneNumber:  "12345678",
				EmailAddress: "john@example.com",
				DateOfBirth:  strPtr("1990-02-21"),
			},
			wantErr: true,
		},
		{
			it: "Should return error if phoneNumber consists of non-numeric characters",
			input: UserInput{
				Name:         "John Doe",
				PhoneNumber:  "123456abcd",
				EmailAddress: "",
				DateOfBirth:  strPtr("1990-02-21"),
			},
			wantErr: true,
		},
		{
			it: "Should return error if emailAddress is not in a valid email address format",
			input: UserInput{
				Name:         "John Doe",
				PhoneNumber:  "1234567890",
				EmailAddress: "john@example@com",
				DateOfBirth:  strPtr("1990-02-21"),
			},
			wantErr: true,
		},
		{
			it: "Should return error if dateOfBirth is nil",
			input: UserInput{
				Name:         "John Doe",
				PhoneNumber:  "1234567890",
				EmailAddress: "john@example.com",
				DateOfBirth:  nil,
			},
			wantErr: true,
		},
		{
			it: "Should return error if dateOfBirth is empty",
			input: UserInput{
				Name:         "John Doe",
				PhoneNumber:  "1234567890",
				EmailAddress: "john@example.com",
				DateOfBirth:  strPtr(""),
			},
			wantErr: true,
		},
		{
			it: "Should return error if dateOfBirth is not in the ISO 8601 date format",
			input: UserInput{
				Name:         "John Doe",
				PhoneNumber:  "1234567890",
				EmailAddress: "john@example.com",
				DateOfBirth:  strPtr("01-01-1990"),
			},
			wantErr: true,
		},
		{
			it: "Should return error if dateOfBirth is in the future",
			input: UserInput{
				Name:         "John Doe",
				PhoneNumber:  "1234567890",
				EmailAddress: "john@example.com",
				DateOfBirth:  strPtr("2050-01-01"),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.it, func(t *testing.T) {
			got, err := MapUserInputToDomainModel(tt.input)

			if err != nil && !tt.wantErr {
				t.Fatalf("got unexpected error: %v", err)
			}
			if err == nil && tt.wantErr {
				t.Fatalf("got: %v, want error", err)
			}

			if tt.wantAgeOfToday > 0 {
				if got.AgeToday != tt.wantAgeOfToday {
					t.Errorf("want age today: %d, got: %d", tt.wantAgeOfToday, got.AgeToday)
				}
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func strPtr(s string) *string {
	return &s
}
