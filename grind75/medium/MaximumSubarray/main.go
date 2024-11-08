package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res := maxSubArray(nums)
	fmt.Printf("Result :: %v \n", res)
}

func maxSubArray(nums []int) int {
	carry := nums[0]
	bestSum := carry

	for i, num := range nums {
		if i == 0 {
			continue
		}
		sum := carry + num
		if num > sum {
			carry = num
		} else {
			carry = sum
		}

		if carry > bestSum {
			bestSum = carry
		}
	}

	return bestSum
}
