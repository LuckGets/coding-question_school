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

function twoSum(nums: number[], target: number): number[] {
  const dict: { [key: number]: { idx: number } } = {};
  for (let i = 0; i < nums.length; i++) {
    if (dict[target - nums[i]]) {
      return [dict[target - nums[i]].idx, i];
    }
    if (!dict[nums[i]]) {
      dict[nums[i]] = { idx: i };
    }
  }
  throw new Error("No solution or invalid input");
}
