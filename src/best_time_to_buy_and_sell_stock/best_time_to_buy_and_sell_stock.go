// Package best_time_to_buy_and_sell_stock
//
// You are given an array prices where prices[i] is the price of a given stock on the ith day.
//
// You want to maximize your profit by choosing a single day to buy one stock
// and choosing a different day in the future to sell that stock.
//
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
package best_time_to_buy_and_sell_stock

// bruteForce simples approach O(n^2), too long
//
// Iterate over the array on each value, calculate
//func bruteForce(prices []int) int {
//	if len(prices) <= 1 {
//		return 0
//	}
//
//	maxProfit := 0
//	for i, buyPrice := range prices[:len(prices)-1] {
//		for _, sellPrice := range prices[i+1:] {
//			if sellPrice <= buyPrice {
//				continue
//			}
//			profit := sellPrice - buyPrice
//			if profit > maxProfit {
//				maxProfit = profit
//			}
//		}
//	}
//
//	return maxProfit
//}

type sellOption struct {
	buyPrice  int
	sellPrice int
}

// minMax we iterate over the array once, O(n)
//
// For each minimum price we are looking for a maximum price.
// If we found new minimum, we save previous known price pair, and continue to look for a maximum.
//
// Then we iterate over the pair list, for each maximum we iterate over previous known pairs,
// in case we saw any less price
func minMax(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	options := make([]sellOption, 0, len(prices)/2)
	option := sellOption{buyPrice: prices[0]}
	for _, price := range prices[1:] {
		if price > option.buyPrice && price > option.sellPrice {
			option.sellPrice = price
			continue
		}
		if price < option.buyPrice {
			if option.sellPrice > option.buyPrice {
				options = append(options, option)
			}
			option.buyPrice = price
			option.sellPrice = price
		}
	}
	options = append(options, option)
	maxProfit := 0
	for i, option := range options {
		profit := option.sellPrice - option.buyPrice
		for _, prevOption := range options[:i] {
			if prevOption.buyPrice < option.buyPrice {
				profit = option.sellPrice - prevOption.buyPrice
			}
		}
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}
