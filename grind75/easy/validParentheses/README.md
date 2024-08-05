# Valid Parentheses

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

1. Open brackets must be closed by the same type of brackets.

2. Open brackets must be closed in the correct order.

3. Every close bracket has a corresponding open bracket of the same type.

### Example

**Example1**

> **Input**: s = "()" <br> > **Output**: true

**Example2**

> **Input**: s = "()[]{}" <br> > **Output**: true

**Example3**

> **Input**: s = "(]" <br> > **Output**: false

### Constraints:

- 1 <= s.length <= 104
- s consists of parentheses only '()[]{}'.

### Solution

##### Leetcode guy's solution by using pure stack [Vikas-Pathak-123](https://leetcode.com/u/Vikas-Pathak-123/)

```typescript
const isValid = function (s) {
  let stack = []; // create an empty stack to store opening brackets
  for (let c of s) {
    // loop through each character in the string
    if (c === "(" || c === "{" || c === "[") {
      // if the character is an opening bracket
      stack.push(c); // push it onto the stack
    } else {
      // if the character is a closing bracket
      if (
        !stack.length || // if the stack is empty or
        (c === ")" && stack[stack.length - 1] !== "(") || // the closing bracket doesn't match the corresponding opening bracket at the top of the stack
        (c === "}" && stack[stack.length - 1] !== "{") ||
        (c === "]" && stack[stack.length - 1] !== "[")
      ) {
        return false; // the string is not valid, so return false
      }
      stack.pop(); // otherwise, pop the opening bracket from the stack
    }
  }
  return !stack.length; // if the stack is empty, all opening brackets have been matched with their corresponding closing brackets,
  // so the string is valid, otherwise, there are unmatched opening brackets, so return false
};
```

> This function use hash table to store desired pattern and answer. If the string match the key in hash table then the next index has to be the desired or ==closed parentheses==.
> and if the next index is the closed parentheses, we have to skip the next index and check the next open parentheses.

    BigO is O(n) because we only loop one time through each string

<br/>

##### The Leetcode guy's solution

From [nitts](https://leetcode.com/u/niits/)

He apply stacks to check if the input have a correctly order of string.

```javascript
const isValid = function (s) {
  const stack = [];
  const mapping = {
    ")": "(",
    "}": "{",
    "]": "[",
  };

  for (const c of s) {
    if (Object.values(mapping).includes(c)) {
      stack.push(c);
    } else if (mapping.hasOwnProperty(c)) {
      if (!stack.length || mapping[c] !== stack.pop()) {
        return false;
      }
    }
  }

  return stack.length === 0;
};
```

which is the same as my solution but more safe and to learn how stack work in data structure.
