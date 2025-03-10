# [2351. First Letter to Appear Twice](https://leetcode.com/problems/first-letter-to-appear-twice/)
Given a string  `s`  consisting of lowercase English letters, return  _the first letter to appear  **twice**_.

**Note**:

-   A letter  `a`  appears twice before another letter  `b`  if the  **second**  occurrence of  `a`  is before the  **second**  occurrence of  `b`.
-   `s`  will contain at least one letter that appears twice.

**Example 1:**

**Input:** s = "abccbaacz"
**Output:** "c"
**Explanation:**
The letter 'a' appears on the indexes 0, 5 and 6.
The letter 'b' appears on the indexes 1 and 4.
The letter 'c' appears on the indexes 2, 3 and 7.
The letter 'z' appears on the index 8.
The letter 'c' is the first letter to appear twice, because out of all the letters the index of its second occurrence is the smallest.

**Example 2:**

**Input:** s = "abcdd"
**Output:** "d"
**Explanation:**
The only letter that appears twice is 'd' so we return 'd'.

**Constraints:**

-   `2 <= s.length <= 100`
-   `s`  consists of lowercase English letters.
-   `s`  has at least one repeated letter.

### Solution

##### Hash set approach
We are using the hash set to memorize the letter which we have visited and add the value to 1. If we met the same letter again, which appeared to be the key of the set, the value will be one and we will return that letter.
```go
func repeatedCharacterByHashSet(s string) byte {
	var res byte
	charMap := make(map[rune]uint8)

	for _, rune := range s {
		if charMap[rune] != 1 {
			res = byte(rune)
		} else {
			charMap[rune] = 1
		}
	}
	return res
}
```

#### List approach
This is almost like the hash set approach but we use list instead, we know that all of the letter of s will be lower case english which have `ASCII` value. So, we assume that we won't meet the letter outside of english and we will apply each Index of the list as the letter `ASCII` value. 

We will assigned the bit value which is 2 value integer --> `0` and `1`. We are doing this for memory optimization because collecting the bit in List will using the memory less than the hash set. If we know all the letter of the word will be the consistent thing, we can do this approach. 

```go
func repeatedCharacter(s string) byte {
	bitSlice := make([]Bits, 26, 26)
	var res byte
	for _, rune := range s {
		charInd := rune - 'a'
		if bitSlice[charInd] != Found {
			bitSlice[charInd] = Found
		} else {
			res = byte(rune)
			break
		}
	}
	return res
}

```

