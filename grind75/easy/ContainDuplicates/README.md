# 217. Contains Duplicate

Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

### Example

##### Example 1:

**Input**: nums = [1,2,3,1]

**Output**: true

##### Example 2:

**Input**: nums = [1,2,3,4]

**Output**: false

##### Example 3:

**Input**: nums = [1,1,1,3,3,4,3,2,4,2]

**Output**: true

##### Constraints:

-     1 <= nums.length <= 105
-     -109 <= nums[i] <= 109

### Solution

###### My solution

```typescript
function containsDuplicate(nums: number[]): boolean {
  const hash: { [key: number]: boolean } = {};
  for (let i = 0; i < nums.length; i++) {
    if (hash[nums[i]]) {
      return true;
    } else {
      hash[nums[i]] = true;
    }
  }
  return false;
}
```

Using Hash map to help collect the data we've seen.

> BigO : O(n)
