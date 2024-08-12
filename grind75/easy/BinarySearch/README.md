# 704.Binary Search

Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

You must write an algorithm with O(log n) runtime complexity.

### Example

##### Example 1:

**Input**: nums = [-1,0,3,5,9,12], target = 9
**Output**: 4
**Explanation**: 9 exists in nums and its index is 4

##### Example 2:

**Input**: nums = [-1,0,3,5,9,12], target = 2
**Output**: -1
**Explanation**: 2 does not exist in nums so return -1

##### Constraints:

-     1 <= nums.length <= 104
-     -10^4 < nums[i], target < 10^4
-     All the integers in nums are unique.
-     nums is sorted in ascending order.

### Solution

##### My solution

I'm using two pointer technique by placing the two pointer at the half of the array (nums.length / 2) and move the pointer to the left and right distinctively by one loop.

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

##### Binary search

[AminiCK](https://leetcode.com/u/AminiCK/) has explained the binary search concept fairly easy and straightforward. You can check his profile or read the [AminiCK](https://leetcode.com/u/AminiCK/) explanation via my [github link](https://github.com/LuckGets/coding-question_school/tree/main/algorithms/basic/binarySearch) or [AminiCK leetcode solution page](https://leetcode.com/problems/binary-search/solutions/423162/binary-search-101/)

All credited to [AminiCK](https://leetcode.com/u/AminiCK/).

Salute 🫡🫡🫡

```javascript
const search = function (nums, target) {
  let lo = 0,
    hi = nums.length - 1;
  while (lo < hi) {
    let mid = lo + Math.floor((hi - lo + 1) / 2);
    if (target < nums[mid]) {
      hi = mid - 1;
    } else {
      lo = mid;
    }
  }
  return nums[lo] == target ? lo : -1;
};
```
