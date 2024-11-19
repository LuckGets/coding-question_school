# [2461. Maximum Sum of Distinct Subarrays With Length K](https://leetcode.com/problems/maximum-sum-of-distinct-subarrays-with-length-k/)

You are given an integer array `nums` and an integer `k`. Find the maximum subarray sum of all the subarrays of `nums` that meet the following conditions:

- The length of the subarray is `k`, and
- All the elements of the subarray are **distinct**.

Return _the maximum subarray sum of all the subarrays that meet the conditions\_\_._ If no subarray meets the conditions, return `0`.

_A **subarray** is a contiguous non-empty sequence of elements within an array._

**Example 1:**

**Input:** nums = [1,5,4,2,9,9,9], k = 3
**Output:** 15
**Explanation:** The subarrays of nums with length 3 are:

- [1,5,4] which meets the requirements and has a sum of 10.
- [5,4,2] which meets the requirements and has a sum of 11.
- [4,2,9] which meets the requirements and has a sum of 15.
- [2,9,9] which does not meet the requirements because the element 9 is repeated.
- [9,9,9] which does not meet the requirements because the element 9 is repeated.
  We return 15 because it is the maximum subarray sum of all the subarrays that meet the conditions

**Example 2:**

**Input:** nums = [4,4,4], k = 3
**Output:** 0
**Explanation:** The subarrays of nums with length 3 are:

- [4,4,4] which does not meet the requirements because the element 4 is repeated.
  We return 0 because no subarrays meet the conditions.

**Constraints:**

- `1 <= k <= nums.length <= 105`
- `1 <= nums[i] <= 105`

### Solution

#### First approach : By [NeetCodeIO](https://www.youtube.com/watch?v=pT-lOE1on3M) and me.

```go
/*
*
It totally obvious that this question need to use sliding window.
You can brute force through it but it will give you BigO : O(n * k) time.
If using sliding window, it will only be O(n) time
First approach will be using hash map to remember the number
we have been passing through and check if the length of
hash key or value is equal to K yet. If equal, we will update the
best sum, if not, we will cruise through again. And by remembering the number,
we will count each time we have seen the number if the hash length does not equal to K.
if the window length is greater than K now, we have to update the left pointer to move closer
to current pointer and subtract the number of the left pointer value in hash map, if the
time we have seen that number is equal to 0, we will remove it from the hash map.
*/
func maximumSubarraySumByHashMap(nums []int, k int) int64 {
	// declare a map to store the int
	h := make(map[int]int)
	n := len(nums)
	// declare the best_sum which is the ans
	best_sum := 0
	l, curr_sum := 0, 0
	r := l
	for r < n {
		h[nums[r]] += 1
		curr_sum += nums[r]
		// if window length is greater than k, we move the left pointer to the
		// next one, and subtract the value of the left pointer out of the hash
		if r-l+1 > k {
			h[nums[l]] -= 1
			if h[nums[l]] == 0 {
				delete(h, nums[l])
			}
			// subtract the value out of the current sum to match the window length
			curr_sum -= nums[l]
			l++
		}

		// if hash length is equal to k, which mean we have different number in this subarray
		// so we caompare the best sum to current sum, if current one is greater, we update the best sum.
		if len(h) == k && r-l+1 == k {
			if best_sum < curr_sum {
				best_sum = curr_sum
			}
		}
		r++
	}
	return int64(best_sum)
}
```

#### Second approach : By [NeetCodeIO](https://www.youtube.com/watch?v=pT-lOE1on3M)

```go
/*
*
This solution is a what I tried to achieve but did not know that it need to use while loop to
reset the left pointer but whatever.
This solution is the optimized version from the above approach
imagine the test case or array we want to find the max sum is looking like this
[9,9,9,9,1,2] --> if we use the above approach, we have to iterate through value of 9
k times before we reach the diffrent value. So, we want to skip the duplicate value
and start the left pointer on the index we find the first duplicate value
*/
func maximumSubarraySumByHashMapButOpt(nums []int, k int) int64 {
	// we save the previous index we have iterate through
	prevInd := make(map[int]int)
	n := len(nums)
	// declare the answer, current value and left pointer
	bestSum, currSum, l := 0, 0, 0

	for r := 0; r < n; r++ {
		// we plus the current sum with the current value
		currSum += nums[r]
		// check if this value has ever exists
		i, exists := prevInd[nums[r]]
		// If not exists, we don't have to do anything,
		// so we set i to be lesser than left pointer
		if !exists {
			i = -1
		}

		// if the current value is a duplicate value or the window length is greater than k,
		// we move the left pointer to the point where the window is equal k first

		for l <= i || r-l+1 > k {
			currSum -= nums[l]
			l++
		}

		// if window length is equal to k, we update the best sum
		if r-l+1 == k {
			if currSum > bestSum {
				bestSum = currSum
			}
		}

		// remember or update the current index before moving on
		prevInd[nums[r]] = r
	}
	return int64(bestSum)
}
```

#### Third approach. [LeetCode guy: [Piotr Maminski]](https://leetcode.com/u/Piotr_Maminski/)

I think this is what the approach that i really want to make but I can't see how to use while loop to move the left pointer to not become the same duplicate value. Salute to this guy,

```go
func maximumSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	elements := make(map[int]bool)
	currentSum := int64(0)
	maxSum := int64(0)
	begin := 0

	for end := 0; end < n; end++ {
		if !elements[nums[end]] {
			currentSum += int64(nums[end])
			elements[nums[end]] = true

			if end-begin+1 == k {
				if currentSum > maxSum {
					maxSum = currentSum
				}
				currentSum -= int64(nums[begin])
				delete(elements, nums[begin])
				begin++
			}
		} else {
			for begin < end && nums[begin] != nums[end] {
				currentSum -= int64(nums[begin])
				delete(elements, nums[begin])
				begin++
			}
			begin++
		}
	}
	return maxSum
}
}
```
