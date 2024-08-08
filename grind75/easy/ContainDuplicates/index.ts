function containsDuplicate(nums: number[]): boolean {
  const hash: {[key:number] : boolean} = {}
  for (let i = 0; i < nums.length; i++) {
      if(hash[nums[i]]) {
          return true
      } else {
      hash[nums[i]] = true
      }
  }
  return false
};