# Question 1

## Objectives
The assessment objectives of this question are as follows:
- Follows best practices.
- Code readability.
- Appropriate code structure.
- Appropriate naming of parameters, functions, types, and fields.
- Appropriate use of comments.

## Instructions
In this scenario, we have a user input that contains the details of a new user, and we would like to validate and map to our domain model. 
The `UserInput` model in `models.go` file describes the input structure, 
and the `User` model in the `models.go` file describes the structure of our domain model.

### Validation
We will need to validate the following, and return an error if any of the below conditions are not met:
1. The `Name` field is not null and is not an empty string.
2. Either `PhoneNumber` or `EmailAddress` must be specified. Both can be specified as well.
3. If specified, the `PhoneNumber` string must be exactly 10 characters, and can only contain numeric characters.
4. If specified, the `EmailAddress` string must be in a valid email address format.
5. The `DateOfBirth` string must be specified, must be in the ISO 8601 date format (`YYYY-MM-DD`), and must not be in the future.

### Mapping
The output fields should be mapped as follows:
- `ID`: Email address of the user if it exists, if not it should be the phone number
- `Name`: Name of the user
- `PhoneNumber`: Phone number of the user
- `emailAddress`: Email address of the user
- `DateOfBirth`: Date of birth of the user as a `time.Time` instance
- `AgeToday`: The current age of the user

### Additional Instructions
- You may use any additional Go packages as you deem fit, but you may not use any external API calls.
- You may create any additional files as necessary.
- You must not modify the signature of the `MapUserInputToDomainModel` function in `question_1.go`, you can only modify the body of the function.
- You must not modify the `question_1_test.go` file in any way.
- You may create additional unit tests in a separate file from `question_1_test.go`.
