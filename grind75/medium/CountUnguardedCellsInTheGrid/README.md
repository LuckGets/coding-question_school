# [2257. Count Unguarded Cells in the Grid](https://leetcode.com/problems/count-unguarded-cells-in-the-grid/)
You are given two integers  `m`  and  `n`  representing a  **0-indexed**  `m x n`  grid. You are also given two 2D integer arrays  `guards`  and  `walls`  where  `guards[i] = [rowi, coli]`  and  `walls[j] = [rowj, colj]`  represent the positions of the  `ith`  guard and  `jth`  wall respectively.

A guard can see  **every**  cell in the four cardinal directions (north, east, south, or west) starting from their position unless  **obstructed**  by a wall or another guard. A cell is  **guarded**  if there is  **at least**  one guard that can see it.

Return _the number of unoccupied cells that are  **not**  **guarded**._

**Example 1:**

![](https://assets.leetcode.com/uploads/2022/03/10/example1drawio2.png)

**Input:** m = 4, n = 6, guards = [[0,0],[1,1],[2,3]], walls = [[0,1],[2,2],[1,4]]
**Output:** 7
**Explanation:** The guarded and unguarded cells are shown in red and green respectively in the above diagram.
There are a total of 7 unguarded cells, so we return 7.

**Example 2:**

![](https://assets.leetcode.com/uploads/2022/03/10/example2drawio.png)

**Input:** m = 3, n = 3, guards = [[1,1]], walls = [[0,1],[1,0],[2,1],[1,2]]
**Output:** 4
**Explanation:** The unguarded cells are shown in green in the above diagram.
There are a total of 4 unguarded cells, so we return 4.

**Constraints:**

-   `1 <= m, n <= 105`
-   `2 <= m * n <= 105`
-   `1 <= guards.length, walls.length <= 5 * 104`
-   `2 <= guards.length + walls.length <= m * n`
-   `guards[i].length == walls[j].length == 2`
-   `0 <= rowi, rowj < m`
-   `0 <= coli, colj < n`
-   All the positions in  `guards`  and  `walls`  are  **unique**.

### Solution

#### 1st approach, simple iterating
if we look at constraints of this question, we will see that the maximum border of m & n length is not much, so we can actually using simple for loop and it won't be LTE.

##### Intuition
We know the size of the grid and we know that guard and wall will occupied one grid. The question need us to answer the unoccupied but we can find the occupied first and after marking process finished, we will count the unoccupied and return it.

The problem in this answer which I don't know how to solve is grid. I am using hash set which won't be the problem but we can expand a 2D array to become 1D array and it will be much easier.
```go
func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	size := m * n
	// We will make the 1D grid by expand the 2D to be 1D by
	//  make the column as the index while the row is determined
	//  by index % column
	g := make([]uint8, size)

	// we need to set up which grid contain what
	// 0 == free, 1 == guardable, 2 == wall or guard

	// check the guards grid
	for _, pos := range guards {
		// we set row by number of column
		row := pos[0] * n
		col := pos[1]
		g[row+col] = 2
	}

	for _, pos := range walls {
		// we set row by number of column
		row := pos[0] * n
		col := pos[1]
		g[row+col] = 2
	}

	for _, guard := range guards {
		// we get the row and column of the guard first
		r, c := guard[0], guard[1]

		// move the row up one row
		// check if the grid is the grid occupied by wall or guard first
		// and we will mark the grid as guardable
		for r := r - 1; r >= 0 && g[r*n+c] != 2; r-- {
			g[r*n+c] = 1
		}
		// the same as above iteration, we move the row down one row
		for r := r + 1; r < m && g[r*n+c] != 2; r++ {
			g[r*n+c] = 1
		}
		// we move the column to the left to check
		for c := c - 1; c >= 0 && g[r*n+c] != 2; c-- {
			g[r*n+c] = 1
		}
		// move the column to the right
		for c := c + 1; c < n && g[r*n+c] != 2; c++ {
			g[r*n+c] = 1
		}
	}

	var res int

	// then we count how many unguarded or not occupied by wall or guard
	for i := 0; i < m*n; i++ {
		if g[i] != 1 && g[i] != 2 {
			res++
		}
	}
	// return the result
	return res
}

```

#### 2nd approach, DFS by [NeetCodeIO](https://www.youtube.com/watch?v=3WVHdSWHxxQ)
I will look to it later once I know what is it.
```go
func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
    grid := make([][]int, m)
    for i := range grid {
        grid[i] = make([]int, n)
    }
    // 0 = free, 1 = guard, 2 = wall, 3 = guardable
    
    for _, g := range guards {
        grid[g[0]][g[1]] = 1
    }
    for _, w := range walls {
        grid[w[0]][w[1]] = 2
    }
    
    markGuarded := func(r, c int) {
        for row := r + 1; row < m; row++ {
            if grid[row][c] == 1 || grid[row][c] == 2 {
                break
            }
            grid[row][c] = 3
        }
        for row := r - 1; row >= 0; row-- {
            if grid[row][c] == 1 || grid[row][c] == 2 {
                break
            }
            grid[row][c] = 3
        }
        for col := c + 1; col < n; col++ {
            if grid[r][col] == 1 || grid[r][col] == 2 {
                break
            }
            grid[r][col] = 3
        }
        for col := c - 1; col >= 0; col-- {
            if grid[r][col] == 1 || grid[r][col] == 2 {
                break
            }
            grid[r][col] = 3
        }
    }
    
    for _, g := range guards {
        markGuarded(g[0], g[1])
    }
    
    res := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 0 {
                res++
            }
        }
    }
    
    return res
}

```

