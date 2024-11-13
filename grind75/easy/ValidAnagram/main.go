package main

func main() {
	ogStr := "anagram"
	testStr := "nagaram"

	isAnagram(ogStr, testStr)
}

func isAnagram(s string, t string) bool {
	charMap := make(map[byte]int)

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		if v, ok := charMap[s[i]]; ok != false {
			charMap[s[i]] = v + 1
		} else {
			charMap[s[i]] = 1
		}

		if v, ok := charMap[t[i]]; ok != false {
			charMap[t[i]] = v - 1
		} else {
			charMap[t[i]] = -1
		}
	}

	for _, value := range charMap {
		if value > 0 {
			return false
		}
	}
	return true
}
