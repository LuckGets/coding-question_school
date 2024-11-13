package main

func main() {
	words := []string{"abba", "baba", "bbaa", "cd", "cd"}
	_ = removeAnagrams(words)

}

func removeAnagrams(words []string) []string {
	i := 0

	for {
		if i == len(words)-1 {
			break
		}

		if isAnagram(words[i], words[i+1]) {
			slice := words[i+2:]
			words = words[0 : i+1]
			words = append(words, slice...)
		} else {
			i++
		}
	}

	return words
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charMap := make(map[byte]int)

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
