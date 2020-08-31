package main

var R, C int
var island [][]int

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isValid(A [][]int, i, j int) bool {
	return 0 <= i && i < R && 0 <= j && j < C && A[i][j] != -1
}

// dfs is used to find the islands
func dfs(A [][]int, i, j int) {
	if i < 0 || i == R || j < 0 || j == C || A[i][j] < 1 {
		return
	}

	// used to track visited
	A[i][j] = -1
	// 2d matrix of i,j coordinates
	island = append(island, []int{i, j})
	dfs(A, i-1, j)
	dfs(A, i+1, j)
	dfs(A, i, j-1)
	dfs(A, i, j+1)
}

func firstIsland(A [][]int) {
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if A[i][j] == 1 {
				dfs(A, i, j)
				return
			}
		}
	}
}

// bfs is used to grow an island
func bfs(A [][]int) int {
	moves := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	steps := 0

	for len(island) > 0 {
		next := [][]int{}
		for _, block := range island {
			for k := 0; k < 4; k++ {
				i, j := block[0]+moves[k][0], block[1]+moves[k][1]
				if isValid(A, i, j) {
					if A[i][j] == 1 {
						return steps
					} else {
						A[i][j] = -1
						next = append(next, []int{i, j})
					}
				}
			}
		}
		steps++
		island = next
	}
	return steps
}

func shortestBridge(A [][]int) int {
	R = len(A)
	C = len(A[0])

	island = nil
	firstIsland(A)

	return bfs(A)
}
