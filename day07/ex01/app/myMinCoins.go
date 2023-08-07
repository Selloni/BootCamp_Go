// This is Day07 ex02.
//
// I need to Generate html with godoc.

package app

import "sort"

// sorts the passed slice
// starting from the highest element fills the buffer
// until it gets the value
// minCouns does not sort the received data
func MinCoins2(val int, coins []int) []int {
	sort.Ints(coins)
	res := make([]int, 0)
	i := len(coins) - 1
	for i >= 0 {
		for val >= coins[i] {
			val -= coins[i]
			res = append(res, coins[i])
		}
		i -= 1
	}
	return res
}
