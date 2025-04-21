package hashmap

import (
	"fmt"
	"testing"
)

func groupAnagrams(strs []string) [][]string {
	groups := map[string][]string{}

	for _, s := range strs {
		seen := make([]int, 26)
		for _, char := range s {
			seen[char-'a']++
		}

		key := ""
		for i, count := range seen {
			if count > 0 {
				key += string(rune('a'+i)) + fmt.Sprintf("%d", count)
			}
		}

		if _, ok := groups[key]; !ok {
			groups[key] = []string{}
		}

		groups[key] = append(groups[key], s)
	}

	results := [][]string{}
	for _, group := range groups {
		results = append(results, group)
	}

	return results
}

func TestGroupAnagrams(t *testing.T) {
	scenarios := []struct {
		strs     []string
		expected [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{[]string{""}, [][]string{{""}}},
	}

	for _, scenario := range scenarios {
		result := groupAnagrams(scenario.strs)
		if len(result) != len(scenario.expected) {
			t.Errorf("Expected %v, got %v", scenario.expected, result)
		}

		for i := range result {
			if len(result[i]) != len(scenario.expected[i]) {
				t.Errorf("Expected %v, got %v", scenario.expected, result)
			}
		}
	}
}
