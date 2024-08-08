function maxSubArrayByBruteForce(nums: number[]): number {
  let maxSum = -Infinity;
  if (nums.length <= 1) return nums[0];
  for (let i = 0; i < nums.length; i++) {
    for (let j = nums.length - 1; j >= i; j--) {
      const newSum = nums.slice(i, j + 1).reduce((prev, curr) => prev + curr);
      if (newSum <= maxSum) {
        continue;
      } else {
        maxSum = newSum;
      }
    }
  }
  return maxSum;
}

function maxSubArrayByKadaneAlgo(nums: number[]): number {
  // Declare the variable to store the maximum value of subarray
  let bestSum: number = -Infinity;
  // Declare the variable to store the current maximum value
  let currMaxSum: number = 0;

  // Loop through each number of array
  for (let item of nums) {
    // Variable for current maximum and item
    const currSum = currMaxSum + item;

    // if the new item + curr max greater than item, We will continue this subarray.
    // If not, we will start the new sub array by changing the curr max to the item.
    // It's like we sliding through each number in array.
    // If the last slide doesn't greater than new number than we start the slide by changing the start
    currMaxSum = currSum > item ? currSum : item;

    // Then we check if current max value is greater than best sum or no,
    // By first item, currMaxSum is going to greater than -Infinity ofc.
    // But the next item is going to change. So we have to check if the currMaxSum is greater than last best sum or not.
    // If not, best sum stay the same.
    bestSum = currMaxSum > bestSum ? currMaxSum : bestSum;
  }

  // returning maximum sum of sub array
  return bestSum;
}
