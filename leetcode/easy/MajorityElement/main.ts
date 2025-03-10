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
