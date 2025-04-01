// https://www.hackerrank.com/challenges/encryption/problem?isFullScreen=true
package medium

import (
	"math"
	"strings"
)

func Encryption(s string) string {
	// Write your code heres
	s = strings.ReplaceAll(s, " ", "")
	l := len(s)

	col := int(math.Ceil(math.Sqrt(float64(l)))) // 4
	row := int(col - 1)                          // 3

	// 0+4+8 1+5+9 ...

	output := ""
	for i := 0; i < col; i++ {
		for j := 0; j <= row; j++ {
			ic := j*col+i
			if ic >= l {
				break
			}
			output += string(s[ic])
		}

		output += " "
	}

	return output
}
