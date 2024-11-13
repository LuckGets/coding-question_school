package main

import (
	"slices"
)

func main() {}

func countFairPairsWithBinSearch(nums []int, lower int, upper int) int {
	slices.Sort(nums)
	ans := 0

	// iterate through each number of nums slices
	for i := 0; i < len(nums); i++ {
		// find the index of largest number to define the upper bound
		// which nums[i] + upperBoundNum to equal upper
		upperBoundInd := binSearch(nums, i+1, upper-nums[i]+1)

		// find the index of largest number to define the lower bound
		// which nums[i] + lowerBoundNum to lesser than lower
		lowerBoundInd := binSearch(nums, i+1, lower-nums[i])

		// after creating the boundary,
		// all of the number in between can make a fair pair to this nums[i].
		// So, we minus the upperBoundInd with lowerBoundInd
		//to find the fair pair
		ans += upperBoundInd - lowerBoundInd
	}

	return ans
}

func binSearch(nums []int, currInd int, target int) int {
	// declare the left boundary
	l := currInd
	// declare the right boundary
	r := len(nums) - 1

	// doing the binary search
	for l <= r {
		m := l + (r-l)/2
		if nums[m] >= target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	// We will return the right boundary because
	// the while loops will end when left equal or greater than right
	// which can miscorrect if we return left when left is greater than right
	// because in this case we want to find the GREATEST nums
	// to lesser than target
	return r
}

// This approach using the same concept of boundary between largest
// number minus lowest number to find the In-between number
// but using only two loop give us a lot better performance
func countFairPairsWithTwoPointerApproach(nums []int, lower int, upper int) int64 {
	// Sorting the slices in ascending order first.
	slices.Sort(nums)

	// declare the function for calculate the pair in the bound
	calBound := func(value int64) int64 {
		result := int64(0)
		left := 0
		right := len(nums) - 1
		for left <= right {
			sum := nums[right] + nums[left]
			if sum <= int(value) {
				result += int64(right - left)
				left++
			} else {
				right--
			}
		}
		return result
	}

	return calBound(int64(upper)) - calBound(int64(lower-1))
}
