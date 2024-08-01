# 1. Two sum

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

### Example

#### Example 1 :

**Input**: nums = [2,7,11,15], target = 9

**Output**: [0,1]

**Explanation**: Because nums[0] + nums[1] == 9, we return [0, 1].

#### Example 2:

**Input**: nums = [3,2,4], target = 6

**Output**: [1,2]

#### Example 3:

**Input**: nums = [3,3], target = 6

**Output**: [0,1]

#### Constraints:

- 2 <= nums.length <= 104

- -109 <= nums[i] <= 109

- -109 <= target <= 109

**Only one valid answer exists.**

**Follow-up:** Can you come up with an algorithm that is less than O(n2) time complexity?

### Solution

There is 2 solution to this question which separated by BigO notation of each solution

#### Solution 1 BigO : n^2^

```typescript
const twoSumSlower = function (nums: number[], target: number) {
  for (let i: number = 0; i < nums.length; i++) {
    for (let j: number = i + 1; j < nums.length; j++) {
      if (nums[i] + nums[j] === target) {
        return [i, j];
      }
    }
  }
};
```

> This solution is kinda brute force and will be super slow if there are so many content in nums.

---

#### Solution 2 BigO : `O(n2)`

```typescript
function twoSum(nums: number[], target: number): Array<number> | undefined {
  const dict: DictTable = {};
  for (let i: number = 0; i <= nums.length; i++) {
    const targetNumber: number = Math.abs(target - nums[i]);

    if (target in dict) {
      return [dict[targetNumber], i];
    }

    dict[nums[i]] = i;
  }
}
```

> This solution is faster because we only loop one time. By putting every seen index into dictionary table. We can easily find the target number when there is desired number.
