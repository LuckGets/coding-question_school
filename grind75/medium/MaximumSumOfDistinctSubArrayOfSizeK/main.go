package main

func main() {
}

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
