// https://www.hackerrank.com/challenges/the-birthday-bar/problem?isFullScreen=true
package easy

/*
 * Complete the 'birthday' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY s
 *  2. INTEGER d
 *  3. INTEGER m
 */

func Birthday(s []int32, d int32, m int32) int32 {
	if m == 0 || d == 0 {
		return 0
	}
	// 0 1 2 3 4

	// Write your code here
	var (
		sum          int32
		isEnough     = false
		numberOfWays int32
	)
	for i, num := range s {
		sum += num

		if i-int(m) == -1 {
			// only turn on once time
			isEnough = true
		} else if i-int(m) >= 0 {
			// auto re sum values
			sum -= s[i-int(m)]
		}

		if isEnough && sum == d {
			numberOfWays += 1
		}
	}

	return numberOfWays
}
