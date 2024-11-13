# [2563. Count the Number of Fair Pairs](https://leetcode.com/problems/count-the-number-of-fair-pairs/)

Given a **0-indexed** integer array `nums` of size `n` and two integers `lower` and `upper`, return _the number of fair pairs_.

A pair `(i, j)` is **fair** if:

- `0 <= i < j < n`, and
- `lower <= nums[i] + nums[j] <= upper`

**Example 1:**

**Input:** nums = [0,1,7,4,4,5], lower = 3, upper = 6
**Output:** 6
**Explanation:** There are 6 fair pairs: (0,3), (0,4), (0,5), (1,3), (1,4), and (1,5).

**Example 2:**

**Input:** nums = [1,7,9,2,5], lower = 11, upper = 11
**Output:** 1
**Explanation:** There is a single fair pair: (2,3).

**Constraints:**

- `1 <= nums.length <= 105`
- `nums.length == n`
- `-109 <= nums[i] <= 109`
- `-109 <= lower <= upper <= 109`

### Solution

#### Binary search approach

You can watch on an [Neetcode YT video](https://www.youtube.com/watch?v=TjthKf7Mc_8&t=5s). It's totally perfect.

```go
func countFairPairs(nums []int, lower int, upper int) int64 {
	ans := 0

	// iterate through each number of nums slices
	for i:= 0; i < len(nums); i++ {
	// find the index of largest number to define the upper bound which nums[i] + upperBoundNum to equal upper
		upperBoundInd := binSearch(nums, i + 1, upper - nums[i] + 1)
			// find the index of largest number to define the lower bound which nums[i] + lowerBoundNum to lesser than lower
		lowerBoundInd := binSearch(nums, i + 1, lower - nums[i])
		// after creating the boundary, all of the number in between can make a fair pair to this nums[i]. So, we minus the upperBoundInd with lowerBoundInd to find the fair pair
			ans += upperBoundInd - lowerBoundInd
	}
	return ans

}

func binSearch(nums []int, currInd int, target int) int {
	l := currInd
	h := len(nums) - 1
	for l <= h {
		m := l + (h - l) /2
		if nums[m] >= target {
			h = m - 1
		} else {
			l = m + 1
		}
	}
	return h
}
```

#### Two-pointer approach

```go
// This approach using the same concept of boundary between largest

// number minus lowest number to find the In-between number

// but using only two loop give us a lot better performance

func  countFairPairsWithTwoPointerApproach(nums  []int,  lower  int,  upper  int)  int64  {

	// Sorting the slices in ascending order first.

	slices.Sort(nums)



	// declare the function for calculate the pair in the bound

	calBound :=  func(value  int64)  int64  {

		result :=  int64(0)

		left :=  0

		right :=  len(nums)  -  1

		for left <= right {

			sum := nums[right]  + nums[left]

			if sum <=  int(value)  {

				result +=  int64(right - left)

				left++

			}  else  {

				right--

			}

		}

		return result

	}

	return  calBound(int64(upper))  -  calBound(int64(lower-1))

}
```
