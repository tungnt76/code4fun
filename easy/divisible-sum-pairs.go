// https://www.hackerrank.com/challenges/divisible-sum-pairs/problem?isFullScreen=true&h_r=next-challenge&h_v=zen
package easy

/*
 * Complete the 'divisibleSumPairs' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER k
 *  3. INTEGER_ARRAY ar
 */

func DivisibleSumPairs(n int32, k int32, ar []int32) int32 {
	// Write your code here
	var counter, i, j int32
	for i = 0; i < n-1; i++ {
		for j = i + 1; j < n; j++ {
			if (ar[i]+ar[j])%k == 0 {
				counter += 1
			}
		}
	}
	return counter
}
