function search(nums: number[], target: number): number {
  let left = Math.trunc(nums.length /2)
  let right = Math.trunc(nums.length /2) + 1
  while (left >= 0 && right <= nums.length) {
      if (nums[left] !== target && nums[right] !== target) {
          if (left === 0 && right === nums.length) {
              break;
          }
          if (left > 0) {
          left--
          }
          if (right < nums.length) {
          right++
          }
      } else if (nums[left] === target) {
          return left
      } else if (nums[right] === target) {
          return right
      }
    }
    return -1
};