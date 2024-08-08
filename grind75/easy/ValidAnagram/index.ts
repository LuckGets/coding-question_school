function isAnagramSortWithJSON(s: string, t: string): void | boolean {
  if (t.length !== s.length) return false;
  const tHashMap: { [key: string]: number } = {};
  const sHashMap: { [key: string]: number } = {};
  for (let i = 0; i < s.length; i++) {
    sHashMap[s[i]] = 1 + (sHashMap[s[i]] || 0);
    tHashMap[t[i]] = 1 + (tHashMap[t[i]] || 0);
  }

  const sortedSMap = Object.entries(sHashMap);
  const sortedTMap = Object.entries(tHashMap);

  return JSON.stringify(sortedSMap) === JSON.stringify(sortedTMap);
}

function isAnagram(s: string, t: string): boolean {
  if (t.length !== s.length) return false;

  const hashMap: { [key: string]: number } = {};

  for (let i = 0; i < s.length; i++) {
    hashMap[s[i]] = 1 + (hashMap[s[i]] || 0);
  }

  for (let j = 0; j < t.length; j++) {
    if (!hashMap[t[j]] || hashMap[t[j]] === 0) {
      return false;
    }
    hashMap[t[j]]--;
  }

  return true;
}
