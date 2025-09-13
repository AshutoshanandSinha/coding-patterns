package main

import "fmt"

/*
Dynamic Programming Pattern Examples

Dynamic Programming solves complex problems by breaking them down into simpler subproblems.
It stores the results of subproblems to avoid redundant calculations.

Common approaches:
1. Top-down (Memoization): Recursion + Cache
2. Bottom-up (Tabulation): Iterative + Table

Time Complexity: Usually O(n) to O(n²)
Space Complexity: O(n) to O(n²) for memoization/tables
*/

// Fibonacci calculates nth Fibonacci number using bottom-up DP
// Time: O(n), Space: O(1)
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	prev1, prev2 := 0, 1

	for i := 2; i <= n; i++ {
		current := prev1 + prev2
		prev1, prev2 = prev2, current
	}

	return prev2
}

// ClimbStairs counts ways to climb n stairs (1 or 2 steps at a time)
// Time: O(n), Space: O(1)
func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}

	prev1, prev2 := 1, 2

	for i := 3; i <= n; i++ {
		current := prev1 + prev2
		prev1, prev2 = prev2, current
	}

	return prev2
}

// HouseRobber robs houses to maximize money without robbing adjacent houses
// Time: O(n), Space: O(1)
func HouseRobber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	prev1, prev2 := 0, nums[0]

	for i := 1; i < len(nums); i++ {
		current := max(prev2, prev1+nums[i])
		prev1, prev2 = prev2, current
	}

	return prev2
}

// CoinChange finds minimum coins needed to make amount
// Time: O(amount * len(coins)), Space: O(amount)
func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1 // Initialize with impossible value
	}

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// LongestIncreasingSubsequence finds length of longest increasing subsequence
// Time: O(n²), Space: O(n)
func LongestIncreasingSubsequence(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	result := dp[0]
	for _, val := range dp {
		if val > result {
			result = val
		}
	}

	return result
}

// Knapsack01 solves 0/1 Knapsack problem
// Time: O(n * capacity), Space: O(capacity)
func Knapsack01(weights, values []int, capacity int) int {
	n := len(weights)
	dp := make([]int, capacity+1)

	for i := 0; i < n; i++ {
		// Traverse backwards to avoid using updated values
		for w := capacity; w >= weights[i]; w-- {
			dp[w] = max(dp[w], dp[w-weights[i]]+values[i])
		}
	}

	return dp[capacity]
}

// LongestCommonSubsequence finds length of longest common subsequence
// Time: O(m * n), Space: O(min(m, n))
func LongestCommonSubsequence(text1, text2 string) int {
	m, n := len(text1), len(text2)

	// Use space-optimized version
	prev := make([]int, n+1)

	for i := 1; i <= m; i++ {
		curr := make([]int, n+1)
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				curr[j] = prev[j-1] + 1
			} else {
				curr[j] = max(prev[j], curr[j-1])
			}
		}
		prev = curr
	}

	return prev[n]
}

// EditDistance finds minimum edit distance (insert, delete, replace)
// Time: O(m * n), Space: O(min(m, n))
func EditDistance(word1, word2 string) int {
	m, n := len(word1), len(word2)

	// Space-optimized version
	prev := make([]int, n+1)
	for j := 0; j <= n; j++ {
		prev[j] = j
	}

	for i := 1; i <= m; i++ {
		curr := make([]int, n+1)
		curr[0] = i
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				curr[j] = prev[j-1]
			} else {
				curr[j] = 1 + min(prev[j], min(curr[j-1], prev[j-1]))
			}
		}
		prev = curr
	}

	return prev[n]
}

// WordBreak checks if string can be segmented using dictionary words
// Time: O(n² * m), Space: O(n)
func WordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

// MaxProductSubarray finds maximum product of contiguous subarray
// Time: O(n), Space: O(1)
func MaxProductSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxProd, minProd, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]

		// Handle negative numbers by swapping max and min
		if num < 0 {
			maxProd, minProd = minProd, maxProd
		}

		maxProd = max(num, maxProd*num)
		minProd = min(num, minProd*num)

		result = max(result, maxProd)
	}

	return result
}

// UniquePathsWithObstacles counts unique paths in grid with obstacles
// Time: O(m * n), Space: O(n)
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 || obstacleGrid[0][0] == 1 {
		return 0
	}

	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([]int, n)
	dp[0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else if j > 0 {
				dp[j] += dp[j-1]
			}
		}
	}

	return dp[n-1]
}

// Helper functions
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Test functions
func testFibonacci() {
	fmt.Println("Testing Fibonacci calculation...")

	testCases := []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{10, 55},
		{15, 610},
	}

	for _, tc := range testCases {
		result := Fibonacci(tc.input)
		if result != tc.expected {
			panic(fmt.Sprintf("Fibonacci test failed for %d: expected %d, got %d",
				tc.input, tc.expected, result))
		}
	}

	fmt.Println("✅ Fibonacci tests passed")
}

func testClimbStairs() {
	fmt.Println("Testing climb stairs...")

	testCases := []struct {
		input    int
		expected int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
	}

	for _, tc := range testCases {
		result := ClimbStairs(tc.input)
		if result != tc.expected {
			panic(fmt.Sprintf("Climb stairs test failed for %d: expected %d, got %d",
				tc.input, tc.expected, result))
		}
	}

	fmt.Println("✅ Climb stairs tests passed")
}

func testHouseRobber() {
	fmt.Println("Testing house robber...")

	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
		{[]int{5}, 5},
	}

	for _, tc := range testCases {
		result := HouseRobber(tc.input)
		if result != tc.expected {
			panic(fmt.Sprintf("House robber test failed for %v: expected %d, got %d",
				tc.input, tc.expected, result))
		}
	}

	fmt.Println("✅ House robber tests passed")
}

func testCoinChange() {
	fmt.Println("Testing coin change...")

	testCases := []struct {
		coins    []int
		amount   int
		expected int
	}{
		{[]int{1, 3, 4}, 6, 2},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
	}

	for _, tc := range testCases {
		result := CoinChange(tc.coins, tc.amount)
		if result != tc.expected {
			panic(fmt.Sprintf("Coin change test failed for coins=%v, amount=%d: expected %d, got %d",
				tc.coins, tc.amount, tc.expected, result))
		}
	}

	fmt.Println("✅ Coin change tests passed")
}

func testLIS() {
	fmt.Println("Testing longest increasing subsequence...")

	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{7, 7, 7, 7, 7, 7, 7}, 1},
	}

	for _, tc := range testCases {
		result := LongestIncreasingSubsequence(tc.input)
		if result != tc.expected {
			panic(fmt.Sprintf("LIS test failed for %v: expected %d, got %d",
				tc.input, tc.expected, result))
		}
	}

	fmt.Println("✅ LIS tests passed")
}

func testKnapsack() {
	fmt.Println("Testing 0/1 knapsack...")

	weights := []int{1, 3, 4, 5}
	values := []int{1, 4, 5, 7}
	expected := 9

	result := Knapsack01(weights, values, 7)
	if result != expected {
		panic(fmt.Sprintf("Knapsack test failed: expected %d, got %d", expected, result))
	}

	fmt.Println("✅ Knapsack tests passed")
}

func testLCS() {
	fmt.Println("Testing longest common subsequence...")

	testCases := []struct {
		text1    string
		text2    string
		expected int
	}{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
	}

	for _, tc := range testCases {
		result := LongestCommonSubsequence(tc.text1, tc.text2)
		if result != tc.expected {
			panic(fmt.Sprintf("LCS test failed for '%s', '%s': expected %d, got %d",
				tc.text1, tc.text2, tc.expected, result))
		}
	}

	fmt.Println("✅ LCS tests passed")
}

func testEditDistance() {
	fmt.Println("Testing edit distance...")

	testCases := []struct {
		word1    string
		word2    string
		expected int
	}{
		{"horse", "ros", 3},
		{"intention", "execution", 5},
	}

	for _, tc := range testCases {
		result := EditDistance(tc.word1, tc.word2)
		if result != tc.expected {
			panic(fmt.Sprintf("Edit distance test failed for '%s', '%s': expected %d, got %d",
				tc.word1, tc.word2, tc.expected, result))
		}
	}

	fmt.Println("✅ Edit distance tests passed")
}

func testWordBreak() {
	fmt.Println("Testing word break...")

	testCases := []struct {
		s        string
		wordDict []string
		expected bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
		{"applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
	}

	for _, tc := range testCases {
		result := WordBreak(tc.s, tc.wordDict)
		if result != tc.expected {
			panic(fmt.Sprintf("Word break test failed for '%s': expected %v, got %v",
				tc.s, tc.expected, result))
		}
	}

	fmt.Println("✅ Word break tests passed")
}

func testMaxProduct() {
	fmt.Println("Testing max product subarray...")

	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
		{[]int{-2, 3, -4}, 24},
	}

	for _, tc := range testCases {
		result := MaxProductSubarray(tc.input)
		if result != tc.expected {
			panic(fmt.Sprintf("Max product test failed for %v: expected %d, got %d",
				tc.input, tc.expected, result))
		}
	}

	fmt.Println("✅ Max product tests passed")
}

func main() {
	fmt.Println("Testing Dynamic Programming Pattern...")
	testFibonacci()
	testClimbStairs()
	testHouseRobber()
	testCoinChange()
	testLIS()
	testKnapsack()
	testLCS()
	testEditDistance()
	testWordBreak()
	testMaxProduct()
	fmt.Println("\n🎉 All Dynamic Programming tests passed!")
}