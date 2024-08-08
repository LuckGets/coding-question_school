// function isAnagram(s, t) {
//   if (t.length !== s.length) return false;
//   var tHashMap = {};
//   var sHashMap = {};
//   for (var i = 0; i < s.length; i++) {
//     sHashMap[s[i]] = 1 + (sHashMap[s[i]] || 0);
//     tHashMap[t[i]] = 1 + (tHashMap[t[i]] || 0);
//   }
//   var sortedSMap = Object.entries(sHashMap).sort();
//   var sortedTMap = Object.entries(tHashMap).sort();
//   console.log("sMap", JSON.stringify(sortedSMap));
//   console.log("tMap", JSON.stringify(sortedTMap));
//   console.log(JSON.stringify(sortedTMap) === JSON.stringify(sortedSMap));
// }

// isAnagram("anagram", "nagaram");
