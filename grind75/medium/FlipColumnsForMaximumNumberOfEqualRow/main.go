func maxEqualRowsAfterFlips(matrix [][]int) int {

	set := make(map[string]int)

	for i := range matrix {
		row := matrix[i]
		var pattern bytes.Buffer
		for _, num := range row {
			if num == row[0] {
				pattern.WriteString("T")
			} else {
				pattern.WriteString("F")
			}
		}
		set[pattern.String()]++
	}

	freq := 0
	for i := range set {
		freq = max(freq, set[i])
	}
	return freq
}

