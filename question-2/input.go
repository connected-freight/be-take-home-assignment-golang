package question_2

type CourseScores struct {
	// Name of the course.
	CourseName string

	// List of assignments in the course, with scores of every student for each assignment.
	AssignmentScores []AssignmentScores
}

type AssignmentScores struct {
	// Name of the assignment.
	AssignmentName string

	// Maximum possible score for the assignment.
	MaxPossibleScore uint

	// Weightage of the assignment scores out of the entire course.
	// All weights of every assignment should add up to 100 for the entire course.
	// Valid range 0 - 100.
	WeightInPercent uint

	// List of students and their scores for the assignment.
	StudentScores []StudentScore
}

type StudentScore struct {
	// Name of the student.
	StudentName string

	// Score that the student has achieved for the assignment.
	Score float32
}
