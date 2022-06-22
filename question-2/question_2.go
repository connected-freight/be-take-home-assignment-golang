package question_2

func GetCourseScoreMetrics(input CourseScores) CourseScoreMetrics {
	// TODO: Implement and replace the return object
	return CourseScoreMetrics{
		CourseName: "Foo Course",
		Assignments: []AssignmentScoreMetrics{
			{AssignmentName: "Assignment 1", MaxScore: 0, MinScore: 0, MeanScore: 0, MedianScore: 0},
			{AssignmentName: "Assignment 2", MaxScore: 0, MinScore: 0, MeanScore: 0, MedianScore: 0},
		},
		SortedStudentScores: []StudentScoreMetrics{
			{StudentName: "Student 1", TotalScore: 0, WeightedPercentage: 0},
			{StudentName: "Student 2", TotalScore: 0, WeightedPercentage: 0},
		},
		MaxWeightedPercentage:                 0,
		MinWeightedPercentage:                 0,
		MeanWeightedPercentage:                0,
		MedianWeightedPercentage:              0,
		PerAssignmentMedianWeightedPercentage: 0,
	}
}
