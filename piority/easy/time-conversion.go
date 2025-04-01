// https://www.hackerrank.com/challenges/one-week-preparation-kit-time-conversion/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=preparation-kits&playlist_slugs%5B%5D=one-week-preparation-kit&playlist_slugs%5B%5D=one-week-day-one
package easy

import (
	"fmt"
	"strconv"
)

func TimeConversion(s string) string {
	// Write your code here
	if s[8] == 'P' {
		if s[0:2] == "12" {
			return "12" + s[2:8]
		} else {
			h, _ := strconv.Atoi(s[0:2])
			h += 12
			return fmt.Sprintf("%02d", h) + s[2:8]
		}
	}

	if s[8] == 'A' {
		if s[0:2] == "12" {
			s = "00" + s[2:8]
		} else {
			return s[0:8]
		}
	}
	
	return s
}
