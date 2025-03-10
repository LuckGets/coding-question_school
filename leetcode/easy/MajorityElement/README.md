# [169. Majority Element](https://leetcode.com/problems/majority-element/)

Given an array `nums` of size `n`, return _the majority element_.

The majority element is the element that appears more than `⌊n / 2⌋` times. You may assume that the majority element always exists in the array.

**Example 1:**

**Input:** nums = [3,2,3]
**Output:** 3

**Example 2:**

**Input:** nums = [2,2,1,1,1,2,2]
**Output:** 2

**Constraints:**

- `n == nums.length`
- `1 <= n <= 5 * 104`
- `-109 <= nums[i] <= 109`

**Follow-up:** Could you solve the problem in linear time and in `O(1)` space?

### Solution

##### My solution

The solution that pop into my head when I first read the question was hashmap approach which we will store the count and candidate inside some object or Map and we will check in map which candidate or number have the most count and return it as an answer.

```typescript
function majorityElementByHashMap(number: number[]): number {
  const countMap: Map<number, number> = new Map();

  let result: number = 0;
  let currCandidate: number = -Infinity;

  for (let num of number) {
    let currVal = countMap.get(num);
    if (currVal) {
      countMap.set(num, ++currVal);
    } else {
      countMap.set(num, 1);
    }
  }

  for (let value of countMap) {
    const [candidate, count] = value;
    if (count > result) {
      result = count;
      currCandidate = candidate;
    }
  }
  return currCandidate;
}
```

This function, even though there are two loops inside this function, the BigO in time will be `O(n)` as it will only iterate through the number of the input and nothing more but this will take the BigO of space as `O(n)` as we will need to have some record or any object to store the value and count of the input and WE NEED TO STORE EVERY INPUT. That's why this solution will take up `O(n)` of space.

##### Boyer-Moore's algorithm solution

```typescript
// Insane algorithm
function majorityElementByBoyerMoore(number: number[]): number {
  let candidate: number = -Infinity;
  let count: number = 0;

  for (let num of number) {
    if (count === 0) candidate = num;

    if (candidate === num) {
      count++;
    } else {
      count--;
    }
  }

  return candidate;
}
```

This function will iterate through every input one time and will take up only `O(1)` space as we will need two variable to store the data.

###### The intention behind this algorithm

This algorithm knows that the majority candidate will never going to have count reach in zero, if there is, then it will not be majority one and the other will be or there is none majority candidate.

**input**: [1,1,1,2,2]

`1` will be the majority here as it show up as the most count in this input. 2 will only have 2 count. If we thought about it as a number, 3 - 2 won't going to be zero and 1 will be the majority as it's the candidate which still have positive count.
