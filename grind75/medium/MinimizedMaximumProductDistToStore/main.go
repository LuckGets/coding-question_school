package main

import (
	"math"
	"slices"
)

func main() {
}

func minimizedMaximum(n int, quantities []int) int {
	slices.Sort(quantities)
	min := 1
	max := quantities[len(quantities)-1]

	if n == len(quantities) {
		return max
	}

	canDistribute := func(k int) int {
		var sum int
		for _, quantity := range quantities {
			divide := float64(quantity) / float64(k)
			dist := int(math.Ceil(divide))
			sum += dist
		}
		return sum
	}
	var ans int

	for min <= max {
		mid := min + (max-min)/2

		sum := canDistribute(mid)
		if sum > n {
			min = mid + 1
		} else if sum <= n {
			ans = mid
			max = mid - 1
		}
	}
	return ans

}

func minimizedMaximumByLeetCodeGuy(n int, quantities []int) int {
	var maxn int

	for _, num := range quantities {
		maxn = max(maxn, num)
	}

	l, r := 1, maxn
	var res int

	for l <= r {
		mid := l + (r-l)/2
		if canDistribute(n, mid, quantities) {
			res = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}
func canDistribute(n, q int, quantities []int) bool {
	var cnt int
	for _, quan := range quantities {
		cnt = quan / q
		if quan%q > 0 {
			cnt++
		}
		n -= cnt
		if n < 0 {
			return false
		}
	}
	return true
}
