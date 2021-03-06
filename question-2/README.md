# Question 2

## Objectives
The assessment objectives of this question are as follows:
- Knowledge of time complexity.
- Follows best practices.
- Code readability.
- Appropriate code structure.
- Appropriate naming of parameters, functions, types, and fields.
- Appropriate use of functional programming concepts.
- Appropriate use of comments.

## Instructions
In this scenario, we have a user input that contains the assignments and raw scores of students for a particular course, 
and we would like to calculate metrics to give an overview of the performance of the course. 
The `CourseScores` model in `input.go` file describes the input structure, 
and the `CourseScoreMetrics` interface in the `output.go` file describes the structure of our metrics output. 
You are to write the code with the lowest time complexity possible, while maintaining readability, 
and document cases where you made a tradeoff between readability and performance. 
Both input and output interfaces are documented to provide information on how to calculate the metrics.

### Additional Instructions
- You may use any additional Go packages as you deem fit, but you may not use any external API calls.
- It is highly recommended to use functional programming concepts to improve the readability of your code.
- You may create any additional files as necessary.
- You must not modify the signature of the `GetCourseScoreMetrics` function in `question_2.go`, you can only modify the body of the function.
- You must not modify the `input.go` and `output.go` files in any way.
- You must not modify the `question_2_test.go` file in any way.
- You may create additional unit tests in a separate file from `question_2_test.go`.
