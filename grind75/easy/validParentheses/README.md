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

##### My solution

```typescript
function isValid(str: string): boolean {
  const hashTable: HashTable = { "(": ")", "{": "}", "[": "]" };
  const strArr: Array<string> = str.split("");
  for (let i = 0; i < strArr.length; i++) {
    if (hashTable[strArr[i]] !== strArr[i + 1]) {
      return false;
    } else {
      i++;
    }
  }
  return true;
}
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
