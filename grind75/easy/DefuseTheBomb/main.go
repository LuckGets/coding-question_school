package main

func main() {

}

func decryptSolByLeetCodeChad(code []int, k int) []int {
	// We start declare counter first, this will use to track how many
	// increment or decrement we have done.
	// Not to let it over than k
	cnt := 1
	n := len(code)

	// If k is minus int, then the counter have to turn backward.
	if k < 0 {
		cnt = -1
	}
	// declare the current sum value
	sum := 0

	// We start by finding the first sum of the first index
	// to initiate the sliding window.
	if k > 0 {
		// if k is positive int,
		// we will the sum the value of current sum with the next index
		// and increment the counter
		for cnt <= k {
			sum += code[cnt%n]
			cnt++
		}
	} else if k < 0 {
		// if k is negative int,
		// we will minus the k before index
		// to reach the first index
		for cnt >= k {
			sum += code[(cnt+n)%n]
			cnt--
		}
	}

	// start the sliding window by declare the slice with
	// the answer sum of the index 0
	ans := []int{sum}
	for i := 1; i < n; i++ {
		// If k is a positive int,
		// we will minus the value of current index
		// because the sum before already have
		// the value of current index.
		// So, we subtract it and put in the next i + k index to the window.
		// to make new window
		if k > 0 {
			sum = sum - code[i] + code[(i+k)%n]
		} else if k < 0 {
			// If k is a negative int,
			// we will subtract the index next to the curr index cause it already have,
			// and plus the value of the before index
			sum = sum + code[i-1] - code[(i+k+n-1)%n]
		}
		// put the answer to the slice
		ans = append(ans, sum)
	}
	return ans
}

func decryptSolByNeetCodeIO(code []int, k int) []int {
	n := len(code)
	result := make([]int, n, n)
	l := 0
	cur := 0
	absK := int(math.Abs(float64(k)))

	for r := 0; r <= n+absK; r++ {
		cur += code[r%n]

		if r-l+1 > absK {
			cur -= code[l%n]
			l = (l + 1) % n
		}

		if r-l+1 == absK {
			var ind int
			if k > 0 {
				if l-1 < 0 {
					ind = n - 1
				} else {
					ind = (l - 1) % n
				}
			} else if k < 0 {
				ind = (r + 1) % n
			}
			result[ind] = cur
		}
	}
	return result
}
