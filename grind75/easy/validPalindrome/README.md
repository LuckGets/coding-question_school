# 125. Valid Palindrome

A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

#### Example

Example 1

> **Input**: s = "A man, a plan, a canal: Panama"
> **Output**: true
> **Explanation**: "amanaplanacanalpanama" is a palindrome.

Example 2

> **Input**: s = "race a car"
> **Output**: false
> **Explanation**: "raceacar" is not a palindrome.

Example 3

> **Input**: s = " "
> **Output**: true
> **Explanation**: s is an empty string "" after removing non-alphanumeric characters.
> Since an empty string reads the same forward and backward, it is a palindrome.

Constaints

-     1 <= s.length <= 2 * 105
-     s consists only of printable ASCII characters.

#### Solution

###### My solution

```typescript
function isPalindrome(s: string): boolean {
  // Check if string is " " or not, if it's then return true.
  if (s.trim() === "") {
    return true;
  }
  // Regex to check if the character is a string or num.
  const regex: RegExp = /^[A-Za-z0-9]+$/;

  // Declare stack which to push char in.
  const stack: Array<string> = [];

  // Loop to put char in stack
  for (let i = 0; i < s.length; i++) {
    // If special char, continue
    if (!regex.test(s[i])) {
      continue;
    }
    // push char to stack
    const strToPush = s[i].toLowerCase();
    stack.push(strToPush);
  }
  // Reverse the string in stack
  const reverStack = [...stack].reverse();

  // Test if it's palindrome or not.
  return stack.join("") === reverStack.join("") ? true : false;
}
```

Using regex to check whether the character is a string or not. If it's a special character or whitespace, We'll skip it. Then, loop through each char in array and using Array method: **reverse** to check if it's a palindrome.

###### Two-pointer solution

```typescript
function isPalindromeWithTwoPointer(s: string): boolean {
  // Check if string is " " or not, if it's then return true.
  if (s.trim() === "") {
    return true;
  }

  // Regex to check if the character is a string or num.
  const regex: RegExp = /^[A-Za-z0-9]+$/;

  // Declare two pointer which is first and last.
  let first: number = 0;
  let last: number = s.length - 1;

  // loop for two pointer
  while (first < last) {
    const currFirst = s.charAt(first);
    const currLast = s.charAt(last);

    if (!regex.test(currFirst)) {
      // If special char, continue
      first++;
    } else if (!regex.test(currLast)) {
      // If special char, continue
      last--;
    } else if (currFirst.toLowerCase() === currLast.toLowerCase()) {
      // If same char, continue
      first++;
      last--;
    } else {
      // If it's not the same char, then it's not palindrome. Return fasle.
      return false;
    }
  }
  // If the loop finished, then there is no err in loop. So it's palindrome
  return true;
}
```

Using two pointer which point to first and last char of the string. If each pointer is not the same, that mean it's not palindrome. But we have to check if it's string or special char first before check whether two pointer is the same or not.

###### BigO : O(n) because we have to loop through each character in array. So the worse possible is up to the length of array.
