package arrays

import (
	"fmt"
)

func maxProfit(prices []int) int {
	var buyPrice = -1
	var profit = 0
	var next = 1
	for next < len(prices) {
		if buyPrice == -1 && prices[next] > prices[next-1] {
			// Tomorrow stock goes up, lets buy today
			buyPrice = prices[next-1]
		} else if buyPrice != -1 && prices[next] < prices[next-1] {
			//tomorrow price will go down, let's sell today
			profit += prices[next-1] - buyPrice
			buyPrice = -1
		}
		next++
	}
	if buyPrice != -1 {
		profit += prices[len(prices)-1] - buyPrice
	}

	return profit
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit([]int{1, 2, 4, 3, 1}))
	fmt.Println(maxProfit([]int{2, 1, 2, 0, 1}))
}
