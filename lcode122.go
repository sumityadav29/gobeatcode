package main

func maxProfit(prices []int) int {
	p := 0

	for i := range len(prices) {
		if i == 0 {
			continue
		}

		p += max(0, prices[i]-prices[i-1])
	}

	return p
}
