# [2064. Minimized Maximum of Products Distributed to Any Store](https://leetcode.com/problems/minimized-maximum-of-products-distributed-to-any-store/)
You are given an integer  `n`  indicating there are  `n`  specialty retail stores. There are  `m`  product types of varying amounts, which are given as a  **0-indexed**  integer array  `quantities`, where  `quantities[i]`  represents the number of products of the  `ith`  product type.

You need to distribute  **all products**  to the retail stores following these rules:

-   A store can only be given  **at most one product type**  but can be given  **any**  amount of it.
-   After distribution, each store will have been given some number of products (possibly  `0`). Let  `x`  represent the maximum number of products given to any store. You want  `x`  to be as small as possible, i.e., you want to  **minimize**  the  **maximum**  number of products that are given to any store.

Return  _the minimum possible_  `x`.

**Example 1:**

**Input:** n = 6, quantities = [11,6]
**Output:** 3
**Explanation:** One optimal way is:
- The 11 products of type 0 are distributed to the first four stores in these amounts: 2, 3, 3, 3
- The 6 products of type 1 are distributed to the other two stores in these amounts: 3, 3
The maximum number of products given to any store is max(2, 3, 3, 3, 3, 3) = 3.

**Example 2:**

**Input:** n = 7, quantities = [15,10,10]
**Output:** 5
**Explanation:** One optimal way is:
- The 15 products of type 0 are distributed to the first three stores in these amounts: 5, 5, 5
- The 10 products of type 1 are distributed to the next two stores in these amounts: 5, 5
- The 10 products of type 2 are distributed to the last two stores in these amounts: 5, 5
The maximum number of products given to any store is max(5, 5, 5, 5, 5, 5, 5) = 5.

**Example 3:**

**Input:** n = 1, quantities = [100000]
**Output:** 100000
**Explanation:** The only optimal way is:
- The 100000 products of type 0 are distributed to the only store.
The maximum number of products given to any store is max(100000) = 100000.

**Constraints:**

-   `m == quantities.length`
-   `1 <= m <= n <= 105`
-   `1 <= quantities[i] <= 105`

### Solution

#### My Solution
We need to find the maximum minimum number of product which can be distributed to satisfies the store. Each store can only take one type of the product.

```go
func minimizedMaximum(n int, quantities []int) int {
	slices.Sort(quantities)
	min := 1
	max := quantities[len(quantities)-1]

	if n == len(quantities) {
		return max
	}

	canDistribute := func(k int) int {
		var sum int
		for _, quantity := range quantities {
			divide := float64(quantity) / float64(k)
			dist := int(math.Ceil(divide))
			sum += dist
		}
		return sum
	}
	var ans int

	for min <= max {
		mid := min + (max-min)/2

		sum := canDistribute(mid)
		if sum > n {
			min = mid + 1
		} else if sum <= n {
			ans = mid
			max = mid - 1
		}
	}
	return ans

}
```

#### [Leetcode guy solution.](https://leetcode.com/problems/minimized-maximum-of-products-distributed-to-any-store/)
```go
// Binary search approach
func minimizedMaximumByLeetCodeGuy(n int, quantities []int) int {
	var maxn int

	for _, num := range quantities {
		maxn = max(maxn, num)
	}

	l, r := 1, maxn
	var res int

	for l <= r {
		mid := l + (r-l)/2
		if canDistribute(n, mid, quantities) {
			res = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}
func canDistribute(n, q int, quantities []int) bool {
	var cnt int
	for _, quan := range quantities {
		cnt = quan / q
		if quan%q > 0 {
			cnt++
		}
		n -= cnt
		if n < 0 {
			return false
		}
	}
	return true
}
```

###### Heap approach by Leetcode guy.

```go
type IntHeap [][]int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i][0] > h[j][0] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        {}
func (h *IntHeap) Pop() any          { return 0 }

func heapApproach(n int, q []int) int {
	buf := make([][]int, len(q))

	for i := range q {
		buf[i] = []int{q[i], q[i], 1, i}
	}

	temp := IntHeap(buf)
	h := &temp
	heap.Init(h)

	n -= len(q)

	var total, cnt, nextTry, newCnt, d int
	var top []int

	for n > 0 {
		top = (*h)[0]

		if top[0] == 1 {
			return 1
		}

		nextTry, total, cnt = top[0]-1, top[1], top[2]
		newCnt = total / nextTry

		if total%nextTry > 0 {
			newCnt++
		}

		d = newCnt - cnt

		if n < d {
			return top[0]
		}

		top[0], top[2] = total/newCnt, newCnt

		if total%newCnt > 0 {
			top[0]++
		}

		heap.Fix(h, 0)

		n -= d
	}
	return (*h)[0][0]
}
```
