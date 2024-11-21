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

func countUnguardedByDFS(m int, n int, guards [][]int, walls [][]int) int {
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
