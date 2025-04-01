// https://www.hackerrank.com/challenges/apple-and-orange/problem?isFullScreen=true
package easy

import "fmt"

/*
 * Complete the 'countApplesAndOranges' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER s
 *  2. INTEGER t
 *  3. INTEGER a
 *  4. INTEGER b
 *  5. INTEGER_ARRAY apples
 *  6. INTEGER_ARRAY oranges
 */

func CountApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	// Write your code here
	var applesFallOnHouse, orangesFallOnHouse int32
	for _, apple := range apples {
		if a+apple >= s && a+apple <= t {
			applesFallOnHouse += 1
		}
	}
	fmt.Println(applesFallOnHouse)

	for _, orange := range oranges {
		if b+orange >= s && b+orange <= t {
			orangesFallOnHouse += 1
		}
	}

	fmt.Println(orangesFallOnHouse)
}
