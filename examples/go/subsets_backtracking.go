// Subsets and Backtracking Pattern Examples
//
// This package demonstrates backtracking patterns for generating subsets,
// permutations, combinations, and solving constraint satisfaction problems.

package main

import (
	"fmt"
	"sort"
)

// SubsetGeneration contains subset generation patterns
type SubsetGeneration struct{}

// Subsets generates all possible subsets
func (sg *SubsetGeneration) Subsets(nums []int) [][]int {
	result := [][]int{}

	var backtrack func(int, []int)
	backtrack = func(start int, currentSubset []int) {
		// Add current subset to result (make a copy)
		subset := make([]int, len(currentSubset))
		copy(subset, currentSubset)
		result = append(result, subset)

		// Try adding each remaining element
		for i := start; i < len(nums); i++ {
			currentSubset = append(currentSubset, nums[i])
			backtrack(i+1, currentSubset)
			currentSubset = currentSubset[:len(currentSubset)-1] // Backtrack
		}
	}

	backtrack(0, []int{})
	return result
}

// SubsetsWithDuplicates generates subsets with duplicate elements
func (sg *SubsetGeneration) SubsetsWithDuplicates(nums []int) [][]int {
	sort.Ints(nums) // Sort to handle duplicates
	result := [][]int{}

	var backtrack func(int, []int)
	backtrack = func(start int, currentSubset []int) {
		subset := make([]int, len(currentSubset))
		copy(subset, currentSubset)
		result = append(result, subset)

		for i := start; i < len(nums); i++ {
			// Skip duplicates
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			currentSubset = append(currentSubset, nums[i])
			backtrack(i+1, currentSubset)
			currentSubset = currentSubset[:len(currentSubset)-1]
		}
	}

	backtrack(0, []int{})
	return result
}

// CombinationSum finds combinations that sum to target (with repetition)
func (sg *SubsetGeneration) CombinationSum(candidates []int, target int) [][]int {
	result := [][]int{}

	var backtrack func(int, []int, int)
	backtrack = func(start int, currentCombination []int, currentSum int) {
		if currentSum == target {
			combination := make([]int, len(currentCombination))
			copy(combination, currentCombination)
			result = append(result, combination)
			return
		}

		if currentSum > target {
			return
		}

		for i := start; i < len(candidates); i++ {
			currentCombination = append(currentCombination, candidates[i])
			// Can reuse same element, so pass i (not i + 1)
			backtrack(i, currentCombination, currentSum+candidates[i])
			currentCombination = currentCombination[:len(currentCombination)-1]
		}
	}

	backtrack(0, []int{}, 0)
	return result
}

// CombinationSumII finds combinations that sum to target (without repetition)
func (sg *SubsetGeneration) CombinationSumII(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}

	var backtrack func(int, []int, int)
	backtrack = func(start int, currentCombination []int, currentSum int) {
		if currentSum == target {
			combination := make([]int, len(currentCombination))
			copy(combination, currentCombination)
			result = append(result, combination)
			return
		}

		if currentSum > target {
			return
		}

		for i := start; i < len(candidates); i++ {
			// Skip duplicates
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			currentCombination = append(currentCombination, candidates[i])
			backtrack(i+1, currentCombination, currentSum+candidates[i])
			currentCombination = currentCombination[:len(currentCombination)-1]
		}
	}

	backtrack(0, []int{}, 0)
	return result
}

// PermutationGeneration contains permutation generation patterns
type PermutationGeneration struct{}

// Permute generates all permutations
func (pg *PermutationGeneration) Permute(nums []int) [][]int {
	result := [][]int{}

	var backtrack func([]int)
	backtrack = func(currentPermutation []int) {
		if len(currentPermutation) == len(nums) {
			permutation := make([]int, len(currentPermutation))
			copy(permutation, currentPermutation)
			result = append(result, permutation)
			return
		}

		for _, num := range nums {
			// Check if num is already used
			used := false
			for _, used_num := range currentPermutation {
				if used_num == num {
					used = true
					break
				}
			}

			if !used {
				currentPermutation = append(currentPermutation, num)
				backtrack(currentPermutation)
				currentPermutation = currentPermutation[:len(currentPermutation)-1]
			}
		}
	}

	backtrack([]int{})
	return result
}

// PermuteUnique generates unique permutations with duplicates
func (pg *PermutationGeneration) PermuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	used := make([]bool, len(nums))

	var backtrack func([]int)
	backtrack = func(currentPermutation []int) {
		if len(currentPermutation) == len(nums) {
			permutation := make([]int, len(currentPermutation))
			copy(permutation, currentPermutation)
			result = append(result, permutation)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}

			// Skip duplicates: if current element equals previous and
			// previous is not used, skip current
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			used[i] = true
			currentPermutation = append(currentPermutation, nums[i])
			backtrack(currentPermutation)
			currentPermutation = currentPermutation[:len(currentPermutation)-1]
			used[i] = false
		}
	}

	backtrack([]int{})
	return result
}

// NextPermutation finds next lexicographically greater permutation in-place
func (pg *PermutationGeneration) NextPermutation(nums []int) {
	// Find the largest index i such that nums[i] < nums[i + 1]
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		// Find the largest index j such that nums[i] < nums[j]
		j := len(nums) - 1
		for nums[j] <= nums[i] {
			j--
		}

		// Swap nums[i] and nums[j]
		nums[i], nums[j] = nums[j], nums[i]
	}

	// Reverse the suffix starting at nums[i + 1]
	left, right := i+1, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

// ConstraintSatisfaction contains constraint satisfaction problems using backtracking
type ConstraintSatisfaction struct{}

// SolveNQueens solves N-Queens problem
func (cs *ConstraintSatisfaction) SolveNQueens(n int) [][]string {
	result := [][]string{}
	board := make([]string, n)
	for i := range board {
		board[i] = ""
		for j := 0; j < n; j++ {
			board[i] += "."
		}
	}

	isSafe := func(row, col int) bool {
		// Check column
		for i := 0; i < row; i++ {
			if board[i][col] == 'Q' {
				return false
			}
		}

		// Check upper-left diagonal
		for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if board[i][j] == 'Q' {
				return false
			}
		}

		// Check upper-right diagonal
		for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
			if board[i][j] == 'Q' {
				return false
			}
		}

		return true
	}

	var backtrack func(int)
	backtrack = func(row int) {
		if row == n {
			solution := make([]string, len(board))
			copy(solution, board)
			result = append(result, solution)
			return
		}

		for col := 0; col < n; col++ {
			if isSafe(row, col) {
				// Place queen
				board[row] = board[row][:col] + "Q" + board[row][col+1:]
				backtrack(row + 1)
				// Remove queen (backtrack)
				board[row] = board[row][:col] + "." + board[row][col+1:]
			}
		}
	}

	backtrack(0)
	return result
}

// SolveSudoku solves Sudoku puzzle in-place
func (cs *ConstraintSatisfaction) SolveSudoku(board [][]byte) {
	isValid := func(row, col int, num byte) bool {
		// Check row
		for j := 0; j < 9; j++ {
			if board[row][j] == num {
				return false
			}
		}

		// Check column
		for i := 0; i < 9; i++ {
			if board[i][col] == num {
				return false
			}
		}

		// Check 3x3 box
		startRow, startCol := 3*(row/3), 3*(col/3)
		for i := startRow; i < startRow+3; i++ {
			for j := startCol; j < startCol+3; j++ {
				if board[i][j] == num {
					return false
				}
			}
		}

		return true
	}

	var backtrack func() bool
	backtrack = func() bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] == '.' {
					for num := byte('1'); num <= byte('9'); num++ {
						if isValid(i, j, num) {
							board[i][j] = num
							if backtrack() {
								return true
							}
							board[i][j] = '.' // Backtrack
						}
					}
					return false
				}
			}
		}
		return true
	}

	backtrack()
}

// WordProblems contains word-related backtracking problems
type WordProblems struct{}

// LetterCombinations generates letter combinations for phone number digits
func (wp *WordProblems) LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	phoneMap := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl",
		'6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
	}

	result := []string{}

	var backtrack func(int, string)
	backtrack = func(index int, currentCombination string) {
		if index == len(digits) {
			result = append(result, currentCombination)
			return
		}

		for _, letter := range phoneMap[digits[index]] {
			backtrack(index+1, currentCombination+string(letter))
		}
	}

	backtrack(0, "")
	return result
}

// WordSearch searches for word in 2D grid
func (wp *WordProblems) WordSearch(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])

	var backtrack func(int, int, int) bool
	backtrack = func(row, col, index int) bool {
		if index == len(word) {
			return true
		}

		if row < 0 || row >= rows || col < 0 || col >= cols ||
			board[row][col] != word[index] {
			return false
		}

		// Mark cell as visited
		temp := board[row][col]
		board[row][col] = '#'

		// Search in all 4 directions
		found := backtrack(row+1, col, index+1) ||
			backtrack(row-1, col, index+1) ||
			backtrack(row, col+1, index+1) ||
			backtrack(row, col-1, index+1)

		// Restore cell
		board[row][col] = temp

		return found
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if backtrack(i, j, 0) {
				return true
			}
		}
	}

	return false
}

// TrieNode represents a node in the Trie
type TrieNode struct {
	Children map[byte]*TrieNode
	Word     string
}

// WordSearchII finds all words in 2D grid using Trie
func (wp *WordProblems) WordSearchII(board [][]byte, words []string) []string {
	// Build Trie
	root := &TrieNode{Children: make(map[byte]*TrieNode)}
	for _, word := range words {
		node := root
		for i := 0; i < len(word); i++ {
			char := word[i]
			if _, exists := node.Children[char]; !exists {
				node.Children[char] = &TrieNode{Children: make(map[byte]*TrieNode)}
			}
			node = node.Children[char]
		}
		node.Word = word
	}

	rows, cols := len(board), len(board[0])
	result := []string{}

	var backtrack func(int, int, *TrieNode)
	backtrack = func(row, col int, node *TrieNode) {
		if row < 0 || row >= rows || col < 0 || col >= cols {
			return
		}

		char := board[row][col]
		if child, exists := node.Children[char]; !exists {
			return
		} else {
			node = child
		}

		if node.Word != "" {
			result = append(result, node.Word)
			node.Word = "" // Avoid duplicates
		}

		// Mark as visited
		board[row][col] = '#'

		// Search in all 4 directions
		directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		for _, dir := range directions {
			backtrack(row+dir[0], col+dir[1], node)
		}

		// Restore cell
		board[row][col] = char
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			backtrack(i, j, root)
		}
	}

	return result
}

func main() {
	// Test subset generation
	sg := &SubsetGeneration{}

	nums := []int{1, 2, 3}
	fmt.Println("Subsets of [1, 2, 3]:", sg.Subsets(nums))

	numsWithDups := []int{1, 2, 2}
	fmt.Println("Subsets with duplicates [1, 2, 2]:", sg.SubsetsWithDuplicates(numsWithDups))

	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println("Combination sum (target=7):", sg.CombinationSum(candidates, target))

	// Test permutation generation
	pg := &PermutationGeneration{}

	nums = []int{1, 2, 3}
	fmt.Println("Permutations of [1, 2, 3]:", pg.Permute(nums))

	numsWithDups = []int{1, 1, 2}
	fmt.Println("Unique permutations of [1, 1, 2]:", pg.PermuteUnique(numsWithDups))

	// Test constraint satisfaction
	cs := &ConstraintSatisfaction{}

	fmt.Println("4-Queens solutions:")
	solutions := cs.SolveNQueens(4)
	for i, solution := range solutions {
		fmt.Printf("Solution %d:\n", i+1)
		for _, row := range solution {
			fmt.Println(row)
		}
		fmt.Println()
	}

	// Test word problems
	wp := &WordProblems{}

	digits := "23"
	fmt.Println("Letter combinations for '23':", wp.LetterCombinations(digits))

	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Printf("Word '%s' exists in board: %t\n", word, wp.WordSearch(board, word))

	words := []string{"oath", "pea", "eat", "rain"}
	board2 := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	fmt.Println("Words found in board:", wp.WordSearchII(board2, words))
}