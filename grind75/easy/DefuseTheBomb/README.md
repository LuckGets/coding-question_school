# [1652. Defuse the Bomb](https://leetcode.com/problems/defuse-the-bomb/)

You have a bomb to defuse, and your time is running out! Your informer will provide you with a  **circular**  array  `code` of length of  `n` and a key  `k`.

To decrypt the code, you must replace every number. All the numbers are replaced  **simultaneously**.

-   If  `k > 0`, replace the  `ith`  number with the sum of the  **next**  `k`  numbers.
-   If  `k < 0`, replace the  `ith`  number with the sum of the  **previous**  `k`  numbers.
-   If  `k == 0`, replace the  `ith`  number with  `0`.

As  `code`  is circular, the next element of  `code[n-1]`  is  `code[0]`, and the previous element of  `code[0]`  is  `code[n-1]`.

Given the  **circular**  array  `code`  and an integer key  `k`, return  _the decrypted code to defuse the bomb_!

**Example 1:**

**Input:** code = [5,7,1,4], k = 3
**Output:** [12,10,16,13]
**Explanation:** Each number is replaced by the sum of the next 3 numbers. The decrypted code is [7+1+4, 1+4+5, 4+5+7, 5+7+1]. Notice that the numbers wrap around.

**Example 2:**

**Input:** code = [1,2,3,4], k = 0
**Output:** [0,0,0,0]
**Explanation:** When k is zero, the numbers are replaced by 0. 

**Example 3:**

**Input:** code = [2,4,9,3], k = -2
**Output:** [12,5,6,13]
**Explanation:** The decrypted code is [3+9, 2+3, 4+2, 9+4]. Notice that the numbers wrap around again. If k is negative, the sum is of the **previous** numbers.

**Constraints:**

-   `n == code.length`
-   `1 <= n <= 100`
-   `1 <= code[i] <= 100`
-   `-(n - 1) <= k <= n - 1`


### Solution

There is two approach to this solution. One with brute force which have BigO equal O(n * k) and it can pass this test cause Constraints state that there is not much test case which have many n. But, we will solve this with sliding window which give us O(n) time.

#### [Sliding window by NeetcodeIO](https://www.youtube.com/watch?v=c4oOIi5YTE4)

```go
func decryptSolByNeetCodeIO(code []int, k int) []int {
	n := len(code)
	result := make([]int, n, n)
	l := 0
	cur := 0
	absK := int(math.Abs(float64(k)))

	for r := 0; r <= n+absK; r++ {
		cur += code[r%n]

		if r-l+1 > absK {
			cur -= code[l%n]
			l = (l + 1) % n
		}

		if r-l+1 == absK {
			var ind int
			if k > 0 {
				if l-1 < 0 {
					ind = n - 1
				} else {
					ind = (l - 1) % n
				}
			} else if k < 0 {
				ind = (r + 1) % n
			}
			result[ind] = cur
		}
	}
	return result
}
```

and leetcode chad guy which I have more preference to this code due to the code which based on Golang.
```go
func decryptSolByLeetCodeChad(code []int, k int) []int {
	// We start declare counter first, this will use to track how many
	// increment or decrement we have done.
	// Not to let it over than k
	cnt := 1
	n := len(code)

	// If k is minus int, then the counter have to turn backward.
	if k < 0 {
		cnt = -1
	}
	// declare the current sum value
	sum := 0

	// We start by finding the first sum of the first index
	// to initiate the sliding window.
	if k > 0 {
		// if k is positive int,
		// we will the sum the value of current sum with the next index
		// and increment the counter
		for cnt <= k {
			sum += code[cnt%n]
			cnt++
		}
	} else if k < 0 {
		// if k is negative int,
		// we will minus the k before index
		// to reach the first index
		for cnt >= k {
			sum += code[(cnt+n)%n]
			cnt--
		}
	}

	// start the sliding window by declare the slice with
	// the answer sum of the index 0
	ans := []int{sum}
	for i := 1; i < n; i++ {
		// If k is a positive int,
		// we will minus the value of current index
		// because the sum before already have
		// the value of current index.
		// So, we subtract it and put in the next i + k index to the window.
		// to make new window
		if k > 0 {
			sum = sum - code[i] + code[(i+k)%n]
		} else if k < 0 {
			// If k is a negative int,
			// we will subtract the index next to the curr index cause it already have,
			// and plus the value of the before index
			sum = sum + code[i-1] - code[(i+k+n-1)%n]
		}
		// put the answer to the slice
		ans = append(ans, sum)
	}
	return ans
}
```

### This problem really not that hard if we using the brute force solution but it's a great question to learn about circular array which will have the same array if index is out of bound and that shows me the benefit of modulus operation which heavily need in this question.
