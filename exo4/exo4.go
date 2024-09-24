package main

import (
	"fmt"
)

func Ft_profit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}

		potentialProfit := price - minPrice
		if potentialProfit > maxProfit {
			maxProfit = potentialProfit
		}
	}

	return maxProfit
}

func main() {
	result := Ft_profit([]int{7, 1, 5, 3, 6, 4})
	fmt.Println(result)
}
