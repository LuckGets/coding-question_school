interface DictTable {
  [key: number]: number;
}

const twoSumSlower = function (nums: number[], target: number) {
  for (let i: number = 0; i < nums.length; i++) {
    for (let j: number = i + 1; j < nums.length; j++) {
      if (nums[i] + nums[j] === target) {
        return [i, j];
      }
    }
  }
};

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
