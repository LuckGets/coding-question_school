# [2516. Take K of Each Character From Left and Right](https://leetcode.com/problems/take-k-of-each-character-from-left-and-right/)

You are given a string  `s`  consisting of the characters  `'a'`,  `'b'`, and  `'c'`  and a non-negative integer  `k`. Each minute, you may take either the  **leftmost**  character of  `s`, or the  **rightmost**  character of  `s`.

Return _the  **minimum**  number of minutes needed for you to take  **at least**_ `k` _of each character, or return_ `-1` _if it is not possible to take_ `k` _of each character._

**Example 1:**

**Input:** s = "aabaaaacaabc", k = 2
**Output:** 8
**Explanation:** 
Take three characters from the left of s. You now have two 'a' characters, and one 'b' character.
Take five characters from the right of s. You now have four 'a' characters, two 'b' characters, and two 'c' characters.
A total of 3 + 5 = 8 minutes is needed.
It can be proven that 8 is the minimum number of minutes needed.

**Example 2:**

**Input:** s = "a", k = 1
**Output:** -1
**Explanation:** It is not possible to take one 'b' or 'c' so return -1.

**Constraints:**

-   `1 <= s.length <= 105`
-   `s`  consists of only the letters  `'a'`,  `'b'`, and  `'c'`.
-   `0 <= k <= s.length`


### Solution
```go

/**
*  This is actually hard sliding window question for me in term of intuition.
*  You can actually using two pointer combining with hashSet that have key as
*  a,b and c and iterate through array to find the pre and suf.
*  but we will invert thinking for this question.
*  the result is finding the minimum number of minutes to take each
*  char from left and right. So, the invert is the maximum char we can
*  remove from the string to have the required amount of char in the result string
*  -------------------------------------------
*  ## Than how to?
*  The intuition is we will find the maximum window which can remove
*  and the outer window still have the amount of char same as the required amount
*  So, we will find the count of all character for two reason.
*  One, to check does the amount of char in string match the required amount or not, if not,
*  we will return -1
*  Two, to have the references for the sliding window, by using sliding window,
*  we will using two pointer to keep track of window. We will increment the right pointer while we
*  iterate through string and minus the count of character from the outside window.
*  then, we will check if the outside window have the amount of char lower than the required amount or no,
*  if yes, than we will move the left pointer untill it match the right pointer or the amount of char in
*  outside window is equal to required amount
*	 once, the window is valid which mean the outside window is still contain all char with required amount
*	 we will update the maximum window.
*	 after finish the iteration, we will have the maximum window of the string that can remove and
*	 the outside window is still valid.
*	 So, the result of this question will be the length of string minus with maximum window
 */

func takeCharacters(s string, k int) int {
	// we declare the array to collect the number of each char
	countAllChar := []int{0, 0, 0}
	n := len(s)
	// we count how many amount of char in the string
	for c := range s {
		ind := int(s[c]) - int('a')
		countAllChar[ind]++
	}

	// if the string is invalid, we will return -1
	if countAllChar[0] < k || countAllChar[1] < k || countAllChar[2] < k {
		return -1
	}

	// we declare left pointer and maximum window
	l, maxWindow := 0, 0
	// Iterate through each char
	for r := range s {
		// we will compare the index of countArray by find the ascii value
		ind := int(s[r]) - int('a')
		// subtract the count of the char
		countAllChar[ind]--

		// than we will check if current window is valid or no, by compare each count of char to
		// the required amount.
		// if window is not valid, we will shrink the window by moving the left pointer
		// to right pointer and return the char amount to the Counting array.
		for l <= r && (countAllChar[0] < k || countAllChar[1] < k || countAllChar[2] < k) {
			lInd := int(s[l]) - int('a')
			countAllChar[lInd]++
			l++
		}

		// Update the maximumWindow by compare the old maxwindow to current window
		maxWindow = max(maxWindow, (r - l + 1))
	}
	// The result is the subtract of length of string and maximum window
	return n - maxWindow
}

```
