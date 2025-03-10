# [1072. Flip Columns For Maximum Number of Equal Rows](https://leetcode.com/problems/flip-columns-for-maximum-number-of-equal-rows/)
You are given an  `m x n`  binary matrix  `matrix`.

You can choose any number of columns in the matrix and flip every cell in that column (i.e., Change the value of the cell from  `0`  to  `1`  or vice versa).

Return  _the maximum number of rows that have all values equal after some number of flips_.

**Example 1:**

**Input:** matrix = [[0,1],[1,1]]
**Output:** 1
**Explanation:** After flipping no values, 1 row has all values equal.

**Example 2:**

**Input:** matrix = [[0,1],[1,0]]
**Output:** 2
**Explanation:** After flipping values in the first column, both rows have equal values.

**Example 3:**

**Input:** matrix = [[0,0,0],[0,0,1],[1,1,0]]
**Output:** 2
**Explanation:** After flipping values in the first two columns, the last two rows have equal values.

**Constraints:**

-   `m == matrix.length`
-   `n == matrix[i].length`
-   `1 <= m, n <= 300`
-   `matrix[i][j]`  is either `0`  or  `1`.


### Solution

#### Intuition
This question is really hard if you don't know the approach on how to solve this problem. The question desc is quite confusing and hard to see for the pattern but once you know, it's actually not hard to solve due to brute force solution.

So, the real question is return the maximum number of the same pattern of the row. If we looking closely, if you have one row looking like this [0,0,0,1,1,0] and the another looking like this [1,1,1,0,0,1]. If you flip at the 4th and 5th column. the first row will be [0,0,0,0,0,0] and the second is [1,1,1,1,1,1]. so both row now have all equal value which is what we are looking for. Then, what are we actually looking for? it's not the number in the array but it's the pattern how the number order in the row. as the example the first array pattern looking like this [x,x,x,o,o,x] and second is [o,o,o,x,x,o] if you revert all of the value in each one, it will be the same value which mean if we can find the same pattern, after revert or no, it'll included in the answer.
but how do we return the maximum pattern? How about we collect it in hash map as the pattern is unique and it's what we looking for in each row so we can store as a key and make the frequency of that as a count.
Gotcha. we can now solve this problem. Totally great question.
```go
func maxEqualRowsAfterFlips(matrix [][]int) int {

	set := make(map[string]int)

	for i := range matrix {
		row := matrix[i]
		var pattern bytes.Buffer
		for _, num := range row {
			if num == row[0] {
				pattern.WriteString("T")
			} else {
				pattern.WriteString("F")
			}
		}
		set[pattern.String()]++
	}

	freq := 0
	for i := range set {
		freq = max(freq, set[i])
	}
	return freq
} 
```
