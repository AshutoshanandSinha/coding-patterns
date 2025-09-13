// Graph Traversal Pattern Examples
//
// This package demonstrates common graph traversal patterns including
// BFS, DFS, and their applications to solve various problems.

package main

import (
	"fmt"
)

// GraphTraversal contains graph traversal algorithms
type GraphTraversal struct{}

// BFS performs breadth-first search traversal
func (gt *GraphTraversal) BFS(graph map[int][]int, start int) []int {
	visited := make(map[int]bool)
	queue := []int{start}
	result := []int{}

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]

		if !visited[vertex] {
			visited[vertex] = true
			result = append(result, vertex)

			for _, neighbor := range graph[vertex] {
				if !visited[neighbor] {
					queue = append(queue, neighbor)
				}
			}
		}
	}

	return result
}

// DFSRecursive performs recursive depth-first search
func (gt *GraphTraversal) DFSRecursive(graph map[int][]int, start int, visited map[int]bool) []int {
	if visited == nil {
		visited = make(map[int]bool)
	}

	visited[start] = true
	result := []int{start}

	for _, neighbor := range graph[start] {
		if !visited[neighbor] {
			result = append(result, gt.DFSRecursive(graph, neighbor, visited)...)
		}
	}

	return result
}

// DFSIterative performs iterative depth-first search
func (gt *GraphTraversal) DFSIterative(graph map[int][]int, start int) []int {
	visited := make(map[int]bool)
	stack := []int{start}
	result := []int{}

	for len(stack) > 0 {
		vertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[vertex] {
			visited[vertex] = true
			result = append(result, vertex)

			// Add neighbors in reverse order to maintain consistent results
			neighbors := graph[vertex]
			for i := len(neighbors) - 1; i >= 0; i-- {
				if !visited[neighbors[i]] {
					stack = append(stack, neighbors[i])
				}
			}
		}
	}

	return result
}

// IslandProblems contains island-related algorithms
type IslandProblems struct{}

// NumIslands counts connected components in a 2D grid
func (ip *IslandProblems) NumIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	islands := 0

	var dfs func(int, int)
	dfs = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != '1' {
			return
		}

		grid[r][c] = '0' // Mark as visited

		// Explore all 4 directions
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				islands++
				dfs(r, c)
			}
		}
	}

	return islands
}

// MaxAreaOfIsland finds the maximum area of an island
func (ip *IslandProblems) MaxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	maxArea := 0

	var dfs func(int, int) int
	dfs = func(r, c int) int {
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != 1 {
			return 0
		}

		grid[r][c] = 0 // Mark as visited
		area := 1

		// Explore all 4 directions
		directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		for _, dir := range directions {
			area += dfs(r+dir[0], c+dir[1])
		}

		return area
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				area := dfs(r, c)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}

// PathFinding contains path finding algorithms
type PathFinding struct{}

// Point represents a coordinate
type Point struct {
	R, C int
}

// WordLadderLength finds shortest transformation sequence length
func (pf *PathFinding) WordLadderLength(beginWord, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}

	if !wordSet[endWord] {
		return 0
	}

	type QueueItem struct {
		Word  string
		Level int
	}

	queue := []QueueItem{{beginWord, 1}}
	visited := map[string]bool{beginWord: true}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		if item.Word == endWord {
			return item.Level
		}

		// Try all possible one-character changes
		word := []rune(item.Word)
		for i := 0; i < len(word); i++ {
			original := word[i]
			for ch := 'a'; ch <= 'z'; ch++ {
				word[i] = ch
				newWord := string(word)

				if wordSet[newWord] && !visited[newWord] {
					visited[newWord] = true
					queue = append(queue, QueueItem{newWord, item.Level + 1})
				}
			}
			word[i] = original
		}
	}

	return 0
}

// ShortestBridge finds shortest bridge between two islands
func (pf *PathFinding) ShortestBridge(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	var firstIsland []Point
	var dfs func(int, int)
	dfs = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != 1 {
			return
		}

		grid[r][c] = 2 // Mark as part of first island
		firstIsland = append(firstIsland, Point{r, c})

		for _, dir := range directions {
			dfs(r+dir[0], c+dir[1])
		}
	}

	// Find first island and mark it
	found := false
	for r := 0; r < rows && !found; r++ {
		for c := 0; c < cols && !found; c++ {
			if grid[r][c] == 1 {
				dfs(r, c)
				found = true
			}
		}
	}

	// BFS to find shortest path to second island
	type QueueItem struct {
		R, C, Dist int
	}

	queue := []QueueItem{}
	for _, point := range firstIsland {
		queue = append(queue, QueueItem{point.R, point.C, 0})
	}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			nr, nc := item.R+dir[0], item.C+dir[1]

			if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
				if grid[nr][nc] == 1 { // Found second island
					return item.Dist
				} else if grid[nr][nc] == 0 { // Water
					grid[nr][nc] = 2 // Mark as visited
					queue = append(queue, QueueItem{nr, nc, item.Dist + 1})
				}
			}
		}
	}

	return -1
}

// CycleDetection contains cycle detection algorithms
type CycleDetection struct{}

// CanFinishCourses detects if course schedule has cycles
func (cd *CycleDetection) CanFinishCourses(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)

	for _, prereq := range prerequisites {
		course, pre := prereq[0], prereq[1]
		graph[pre] = append(graph[pre], course)
	}

	// 0: unvisited, 1: visiting, 2: visited
	color := make([]int, numCourses)

	var hasCycle func(int) bool
	hasCycle = func(node int) bool {
		if color[node] == 1 { // Back edge found (cycle)
			return true
		}
		if color[node] == 2 { // Already processed
			return false
		}

		color[node] = 1 // Mark as visiting

		for _, neighbor := range graph[node] {
			if hasCycle(neighbor) {
				return true
			}
		}

		color[node] = 2 // Mark as visited
		return false
	}

	for i := 0; i < numCourses; i++ {
		if color[i] == 0 && hasCycle(i) {
			return false
		}
	}

	return true
}

func main() {
	// Test graph traversal
	graph := map[int][]int{
		0: {1, 2},
		1: {0, 3, 4},
		2: {0, 5, 6},
		3: {1},
		4: {1},
		5: {2},
		6: {2},
	}

	gt := &GraphTraversal{}
	fmt.Println("BFS from 0:", gt.BFS(graph, 0))
	fmt.Println("DFS (recursive) from 0:", gt.DFSRecursive(graph, 0, nil))
	fmt.Println("DFS (iterative) from 0:", gt.DFSIterative(graph, 0))

	// Test island problems
	islands := &IslandProblems{}

	grid1 := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println("Number of islands:", islands.NumIslands(grid1))

	grid2 := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	fmt.Println("Max area of island:", islands.MaxAreaOfIsland(grid2))

	// Test path finding
	pf := &PathFinding{}

	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println("Word ladder length:", pf.WordLadderLength(beginWord, endWord, wordList))

	bridgeGrid := [][]int{{0, 1}, {1, 0}}
	fmt.Println("Shortest bridge:", pf.ShortestBridge(bridgeGrid))

	// Test cycle detection
	cd := &CycleDetection{}

	numCourses := 2
	prerequisites := [][]int{{1, 0}}
	fmt.Println("Can finish courses:", cd.CanFinishCourses(numCourses, prerequisites))

	prerequisitesWithCycle := [][]int{{1, 0}, {0, 1}}
	fmt.Println("Can finish courses (with cycle):", cd.CanFinishCourses(numCourses, prerequisitesWithCycle))
}