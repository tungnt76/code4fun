// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
package array

import "testing"

func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}

func TestMaxProfit(t *testing.T) {
	prices := []int{2, 1, 2, 1, 0, 1, 2}
	expected := 2
	result := maxProfit(prices)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
