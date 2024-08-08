# 242. Valid Anagram

Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

### Example

###### Example 1

**Input**: s = "anagram", t = "nagaram"
**Output**: true

###### Example 2:

**Input**: s = "rat", t = "car"
**Output**: false

###### Constraints:

-     1 <= s.length, t.length <= 5 * 104
-     s and t consist of lowercase English letters.

### Solution

##### My solution

###### intuition

This question has two parameter which use the same character but in different order. So, we check first if two parameter has the same length or no.

Second part is declare Object which reference for hash map which collect all of the string in parameter.

String is the same so we are gonna keep track of each character by collecting the value of character which is the key of hash map in number. Each time we met the character, we count it up in hash map.

So we loop through first parameter, collect the string and how many time we met these char in hash map.

Then, we loop through second parameter and each time we met the character, we minus the value of the key in hash map.

So, if the second parameter use the same character, all the value in hash map is going to equal 0.

<br>

```typescript
function isAnagram(s: string, t: string): boolean {
  if (t.length !== s.length) return false;

  const hashMap: { [key: string]: number } = {};

  for (let i = 0; i < s.length; i++) {
    hashMap[s[i]] = 1 + (hashMap[s[i]] || 0);
  }

  for (let j = 0; j < t.length; j++) {
    if (!hashMap[t[j]] || hashMap[t[j]] === 0) {
      return false;
    }
    hashMap[t[j]]--;
  }

  return true;
}
```
