// https://www.hackerrank.com/challenges/the-birthday-bar/problem?isFullScreen=true
package easy

func GradingStudents(grades []int32) []int32 {
	// Write your code here
	rs := make([]int32, len(grades))
	for i, grade := range grades {
		q := grade % 5

		if grade < 38 || q <= 2 {
			rs[i] = int32(grade)
		} else if q == 3 {
			rs[i] = grade + 2
		} else {
			rs[i] = grade + 1
		}
	}

	return rs
}
