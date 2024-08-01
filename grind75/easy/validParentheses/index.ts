type HashTable = {
  [key: string]: string;
};

function isValid(str: string): boolean {
  const hashTable: HashTable = { "(": ")", "{": "}", "[": "]" };
  const strArr: Array<string> = str.split("");
  for (let i = 0; i < strArr.length; i++) {
    if (hashTable[strArr[i]] !== strArr[i + 1]) {
      return false;
    } else {
      i++;
    }
  }
  return true;
}
