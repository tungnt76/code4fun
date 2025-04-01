// https://www.hackerrank.com/challenges/day-of-the-programmer/problem?isFullScreen=true
package easy

import "fmt"

func DayOfProgrammer(year int32) string {
	// Write your code here
	if year == 1918 {
		return "26.09.1918"
	}

	if isLeafYear(year) {
		return fmt.Sprintf("13.09.%d", year)
	}
	return fmt.Sprintf("12.09.%d", year)
}

func isLeafYear(year int32) bool {
	if year <= 1917 {
		return year%4 == 0
	}
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}
