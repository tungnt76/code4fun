// https://www.hackerrank.com/challenges/kangaroo/problem?isFullScreen=true
package easy

/*
 * Complete the 'kangaroo' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER x1
 *  2. INTEGER v1
 *  3. INTEGER x2
 *  4. INTEGER v2
 */

func Kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// Write your code here
	x := x1 - x2
	v := v2 - v1

	if (x == 0 && v != 0) || (v == 0 && x != 0) {
		return "NO"
	}

	if x%v == 0 && x/v > 0 {
		return "YES"
	}

	return "NO"
}
