# 1047. Remove All Adjacent Duplicates In String

You are given a string s consisting of lowercase English letters. A duplicate removal consists of choosing two adjacent and equal letters and removing them.

We repeatedly make duplicate removals on s until we no longer can.

Return the final string after all such duplicate removals have been made. It can be proven that the answer is unique.

### Example

Example 1

> **Input**: s = "abbaca"
> **Output**: "ca"
> **Explanation**:
> For example, in "abbaca" we could remove "bb" since the letters are adjacent and equal, and this is the only possible move. The result of this move is that the string is "aaca", of which only "aa" is possible, so the final string is "ca".

Example 2

> **Input**: s = "azxxzy"
> **Output**: "ay"

Constraints

-     1 <= s.length <= 105
-     s consists of lowercase English letters.

### Solution

##### My solution

```typescript
function removeDuplicates(s: string): string {
  const stack: Array<string> = [];
  for (let i = 0; i < s.length; i++) {
    if (stack[stack.length - 1] === s.charAt(i) && i !== 0) {
      stack.pop();
    } else {
      stack.push(s.charAt(i));
    }
  }
  return stack.join("");
}
```

With this question, I using stack to track each character I've been looped through. Each character have been checked whether it was duplicated with the character(previous one) in stack or not. If it was actually duplicated, I pop the lastest character in the stack. If not, pushing the new non-dupe one to the stack and check again.

###### BigO : O(n)
