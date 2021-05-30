package main

import (
	"math"
)

func coinChange(coins []int, amount int) int {

	// no coins are required to make change for an amount less than 1

	if amount < 1 {
		return 0
	}
	// initialize "count" array
	// count[i] stores optimal solution for amount i+1

	return change(coins, amount, make([]int, amount))

}

func change(coins []int, rem int, count []int) int {

	// degenerate case

	if rem < 0 {
		return -1
	}

	if rem == 0 {
		return 0
	}

	// it is possible to re-encounter an already computed optimal value
	// that is the point of the cache
	// note that rem-1 is the index of the optimal solution for rem

	if count[rem-1] != 0 {
		return count[rem-1]
	}

	// initialize optimal solution as "infinity"

	min := math.MaxInt64

	// for a given value of rem (remainder)
	// compute minimum (optimal) solution
	// subtract denomination value from current case
	// core recursion here
	// as per recurrence relation discussed above
	// this is at the rem level

	for _, coin := range coins {

		res := change(coins, rem-coin, count)

		if res >= 0 && res < min {
			min = 1 + res
		}
	}

	// update cache accordingly

	if min == math.MaxInt64 {
		count[rem-1] = -1
	} else {
		count[rem-1] = min
	}

	return count[rem-1]

}
