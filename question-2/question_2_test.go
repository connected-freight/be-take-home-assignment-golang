package question_2_test

import (
	"reflect"
	"testing"

	. "github.com/connected-freight/be-toke-home-assignment-golang/question-2"
)

func TestQuestion2(t *testing.T) {
	tests := []struct {
		it    string
		input CourseScores
		want  CourseScoreMetrics
	}{
		{
			it: "Should calculate course score metrics from raw course scores",
			input: CourseScores{
				CourseName: "Defense Against the Dark Arts",
				AssignmentScores: []AssignmentScores{
					{
						AssignmentName:   "Verdimillious Charm",
						MaxPossibleScore: 10,
						WeightInPercent:  20,
						StudentScores: []StudentScore{
							{StudentName: "Harry Potter", Score: 9},
							{StudentName: "Hermione Granger", Score: 10},
							{StudentName: "Ron Weasley", Score: 6},
							{StudentName: "Ginny Weasley", Score: 8},
							{StudentName: "Draco Malfoy", Score: 4},
						},
					},
					{
						AssignmentName:   "Homorphus Charm",
						MaxPossibleScore: 10,
						WeightInPercent:  30,
						StudentScores: []StudentScore{
							{StudentName: "Ginny Weasley", Score: 7},
							{StudentName: "Draco Malfoy", Score: 2},
							{StudentName: "Hermione Granger", Score: 9.5},
							{StudentName: "Ron Weasley", Score: 8},
							{StudentName: "Harry Potter", Score: 7},
						},
					},
					{
						AssignmentName:   "Full Body-Bind Curse",
						MaxPossibleScore: 20,
						WeightInPercent:  50,
						StudentScores: []StudentScore{
							{StudentName: "Ron Weasley", Score: 12},
							{StudentName: "Hermione Granger", Score: 19},
							{StudentName: "Harry Potter", Score: 0},
							{StudentName: "Draco Malfoy", Score: 8},
							{StudentName: "Ginny Weasley", Score: 16},
						},
					},
				},
			},
			want: CourseScoreMetrics{
				CourseName: "Defense Against the Dark Arts",
				Assignments: []AssignmentScoreMetrics{
					{AssignmentName: "Verdimillious Charm", MaxScore: 10, MinScore: 4, MeanScore: 7.4, MedianScore: 8},
					{AssignmentName: "Homorphus Charm", MaxScore: 9.5, MinScore: 2, MeanScore: 6.7, MedianScore: 7},
					{AssignmentName: "Full Body-Bind Curse", MaxScore: 19, MinScore: 0, MeanScore: 11, MedianScore: 12},
				},
				SortedStudentScores: []StudentScoreMetrics{
					{StudentName: "Hermione Granger", TotalScore: 38.5, WeightedPercentage: 96},
					{StudentName: "Ginny Weasley", TotalScore: 31, WeightedPercentage: 77},
					{StudentName: "Ron Weasley", TotalScore: 26, WeightedPercentage: 66},
					{StudentName: "Harry Potter", TotalScore: 16, WeightedPercentage: 39},
					{StudentName: "Draco Malfoy", TotalScore: 14, WeightedPercentage: 34},
				},
				MaxWeightedPercentage:                 96,
				MinWeightedPercentage:                 34,
				MeanWeightedPercentage:                62.4,
				MedianWeightedPercentage:              66,
				PerAssignmentMedianWeightedPercentage: 67,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.it, func(t *testing.T) {
			got := GetCourseScoreMetrics(tt.input)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
