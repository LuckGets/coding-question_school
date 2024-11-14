package firstlettertoappeartwice

// We assign the type of data to assign in List to be
// the smallest data type which is uint8

type Bits uint8

// assign the enum as 0 and 1
const (
	NotFound Bits = iota
	Found
)

func repeatedCharacter(s string) byte {
	bitSlice := make([]Bits, 26, 26)
	var res byte
	for _, rune := range s {
		charInd := rune - 'a'
		if bitSlice[charInd] != Found {
			bitSlice[charInd] = Found
		} else {
			res = byte(rune)
			break
		}
	}
	return res
}

func repeatedCharacterByHashSet(s string) byte {
	var res byte
	charMap := make(map[rune]uint8)

	for _, rune := range s {
		if charMap[rune] != 1 {
			res = byte(rune)
		} else {
			charMap[rune] = 1
		}
	}
	return res
}
