package question_2

type CourseScoreMetrics struct {
	// Name of the course.
	CourseName string

	// List of assignments in the course, with metrics of students' performances.
	Assignments []AssignmentScoreMetrics

	// List of students in the course, with metrics of students' individual performance.
	// Sorted by the students' weightedPercentage, from highest to lowest.
	SortedStudentScores []StudentScoreMetrics

	// Highest weighted percentage achieved by a student for the course.
	// Should be the highest calculated `students.weightedPercentage`.
	MaxWeightedPercentage float32

	// Lowest weighted percentage achieved by a student for the course.
	// Should be the lowest calculated `students.weightedPercentage`.
	MinWeightedPercentage float32

	// Mean weighted percentage achieved by a student for the course.
	// Should be the sum of all calculated `students.weightedPercentage` divided by the int  of students (i.e. average).
	MeanWeightedPercentage float32

	// Median weighted percentage achieved by a student for the course.
	// Should be the middle `students.weightedPercentage` when ordered from lowest to highest.
	MedianWeightedPercentage float32

	// Median per course weighted percentage achieved by any student for the course.
	// Should be the sum of each assignment's ((median score / max possible score) * weight).
	PerAssignmentMedianWeightedPercentage float32
}

type AssignmentScoreMetrics struct {
	// Name of the assignment.
	AssignmentName string

	// Highest score achieved by a student for the assignment.
	MaxScore float32

	// Lowest score achieved by a student for the assignment.
	MinScore float32

	// Mean score achieved by a student for the assignment.
	// Should be the sum of all students' assignment scores divided by the int  of students (i.e. average).
	MeanScore float32

	// Median score achieved by a student for the assignment.
	// Should be the middle `score` when ordered from lowest to highest.
	MedianScore float32
}

type StudentScoreMetrics struct {
	// Name of the student.
	StudentName string

	// Total score achieved by the student over all assignments.
	TotalScore float32

	// Weighted percentage achieved by the student for the course.
	// Should be the sum of ((student assignment score / assignment max possible score) * assignment weight).
	WeightedPercentage float32
}
