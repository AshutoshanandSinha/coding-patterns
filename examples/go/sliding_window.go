// Package main demonstrates Sliding Window pattern implementations in Go
//
// Sliding Window is used to find a subarray or substring that satisfies
// specific conditions, optimizing from O(n²) to O(n) time complexity.
package main

import (
	"fmt"
	"math"
)

// SlidingWindow contains implementations of the Sliding Window pattern
type SlidingWindow struct{}

// MaxSumSubarray finds maximum sum of any contiguous subarray of size k
// Time: O(n), Space: O(1)
func (sw *SlidingWindow) MaxSumSubarray(arr []int, k int) int {
	if len(arr) < k {
		return -1
	}
	
	// Calculate sum of first window
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}
	maxSum := windowSum
	
	// Slide the window
	for i := k; i < len(arr); i++ {
		windowSum += arr[i] - arr[i-k]
		maxSum = max(maxSum, windowSum)
	}
	
	return maxSum
}

// LongestSubstringWithoutRepeating finds length of longest substring without repeating characters
// Time: O(n), Space: O(min(m,n)) where m is charset size
func (sw *SlidingWindow) LongestSubstringWithoutRepeating(s string) int {
	charMap := make(map[byte]int)
	left := 0
	maxLength := 0
	
	for right := 0; right < len(s); right++ {
		// If character already in window, move left pointer
		if lastIndex, exists := charMap[s[right]]; exists && lastIndex >= left {
			left = lastIndex + 1
		}
		
		charMap[s[right]] = right
		maxLength = max(maxLength, right-left+1)
	}
	
	return maxLength
}

// MinWindowSubstring finds minimum window in s that contains all characters of t
// Time: O(|s| + |t|), Space: O(|s| + |t|)
func (sw *SlidingWindow) MinWindowSubstring(s, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	
	// Character frequency in t
	required := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		required[t[i]]++
	}
	
	left, right := 0, 0
	formed := 0 // Number of unique chars with desired frequency
	windowCounts := make(map[byte]int)
	
	// (window length, left, right)
	ans := []int{int(math.Inf(1)), 0, 0}
	
	for right < len(s) {
		// Add character from right to window
		char := s[right]
		windowCounts[char]++
		
		// Check if current character's frequency matches required
		if requiredCount, exists := required[char]; exists && windowCounts[char] == requiredCount {
			formed++
		}
		
		// Try to contract window
		for left <= right && formed == len(required) {
			char = s[left]
			
			// Save the smallest window
			if right-left+1 < ans[0] {
				ans[0] = right - left + 1
				ans[1] = left
				ans[2] = right
			}
			
			// Remove from left of window
			windowCounts[char]--
			if requiredCount, exists := required[char]; exists && windowCounts[char] < requiredCount {
				formed--
			}
			
			left++
		}
		
		right++
	}
	
	if ans[0] == int(math.Inf(1)) {
		return ""
	}
	return s[ans[1] : ans[2]+1]
}

// LongestSubstringKDistinct finds length of longest substring with at most k distinct characters
// Time: O(n), Space: O(k)
func (sw *SlidingWindow) LongestSubstringKDistinct(s string, k int) int {
	if k == 0 {
		return 0
	}
	
	charFrequency := make(map[byte]int)
	left := 0
	maxLength := 0
	
	for right := 0; right < len(s); right++ {
		// Add character to frequency map
		charFrequency[s[right]]++
		
		// Shrink window if more than k distinct characters
		for len(charFrequency) > k {
			charFrequency[s[left]]--
			if charFrequency[s[left]] == 0 {
				delete(charFrequency, s[left])
			}
			left++
		}
		
		maxLength = max(maxLength, right-left+1)
	}
	
	return maxLength
}

// SubarrayProductLessThanK counts subarrays where product of elements is less than k
// Time: O(n), Space: O(1)
func (sw *SlidingWindow) SubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	
	left := 0
	product := 1
	count := 0
	
	for right := 0; right < len(nums); right++ {
		product *= nums[right]
		
		// Shrink window while product >= k
		for product >= k {
			product /= nums[left]
			left++
		}
		
		// Add all subarrays ending at right
		count += right - left + 1
	}
	
	return count
}

// FruitsIntoBaskets picks maximum fruits with at most 2 different types
// Time: O(n), Space: O(1)
func (sw *SlidingWindow) FruitsIntoBaskets(fruits []int) int {
	fruitFrequency := make(map[int]int)
	left := 0
	maxFruits := 0
	
	for right := 0; right < len(fruits); right++ {
		// Add fruit to basket
		fruitFrequency[fruits[right]]++
		
		// Shrink window if more than 2 fruit types
		for len(fruitFrequency) > 2 {
			fruitFrequency[fruits[left]]--
			if fruitFrequency[fruits[left]] == 0 {
				delete(fruitFrequency, fruits[left])
			}
			left++
		}
		
		maxFruits = max(maxFruits, right-left+1)
	}
	
	return maxFruits
}

// SmallestSubarraySum finds length of smallest subarray with sum >= target
// Time: O(n), Space: O(1)
func (sw *SlidingWindow) SmallestSubarraySum(arr []int, target int) int {
	left := 0
	windowSum := 0
	minLength := int(math.Inf(1))
	
	for right := 0; right < len(arr); right++ {
		windowSum += arr[right]
		
		// Shrink window while sum >= target
		for windowSum >= target {
			minLength = min(minLength, right-left+1)
			windowSum -= arr[left]
			left++
		}
	}
	
	if minLength == int(math.Inf(1)) {
		return 0
	}
	return minLength
}

// FindAnagrams finds all start indices of anagrams of p in s
// Time: O(|s| + |p|), Space: O(1) - at most 26 characters
func (sw *SlidingWindow) FindAnagrams(s, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}
	
	// Character frequency in pattern
	pFreq := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		pFreq[p[i]]++
	}
	
	windowFreq := make(map[byte]int)
	result := []int{}
	left := 0
	
	for right := 0; right < len(s); right++ {
		// Add character to window
		windowFreq[s[right]]++
		
		// Maintain window size
		if right-left+1 > len(p) {
			if windowFreq[s[left]] == 1 {
				delete(windowFreq, s[left])
			} else {
				windowFreq[s[left]]--
			}
			left++
		}
		
		// Check if current window is anagram
		if right-left+1 == len(p) && mapsEqual(windowFreq, pFreq) {
			result = append(result, left)
		}
	}
	
	return result
}

// CharacterReplacement finds longest substring with same letters after k replacements
// Time: O(n), Space: O(1) - at most 26 characters
func (sw *SlidingWindow) CharacterReplacement(s string, k int) int {
	charFrequency := make(map[byte]int)
	left := 0
	maxLength := 0
	maxCount := 0
	
	for right := 0; right < len(s); right++ {
		charFrequency[s[right]]++
		maxCount = max(maxCount, charFrequency[s[right]])
		
		// If replacements needed > k, shrink window
		if (right-left+1)-maxCount > k {
			charFrequency[s[left]]--
			left++
		}
		
		maxLength = max(maxLength, right-left+1)
	}
	
	return maxLength
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

func mapsEqual(a, b map[byte]int) bool {
	if len(a) != len(b) {
		return false
	}
	
	for key, valueA := range a {
		if valueB, exists := b[key]; !exists || valueA != valueB {
			return false
		}
	}
	
	return true
}

// Test functions

func runTests() {
	sw := &SlidingWindow{}
	
	// Test MaxSumSubarray
	fmt.Println("Testing MaxSumSubarray:")
	arr := []int{2, 1, 5, 1, 3, 2}
	k := 3
	result := sw.MaxSumSubarray(arr, k)
	fmt.Printf("Array: %v, K: %d, Max Sum: %d\n", arr, k, result)
	if result != 9 {
		panic(fmt.Sprintf("Expected 9, got %d", result))
	}
	
	// Test LongestSubstringWithoutRepeating
	fmt.Println("\nTesting LongestSubstringWithoutRepeating:")
	s := "abcabcbb"
	resultInt := sw.LongestSubstringWithoutRepeating(s)
	fmt.Printf("String: '%s', Longest Length: %d\n", s, resultInt)
	if resultInt != 3 {
		panic(fmt.Sprintf("Expected 3, got %d", resultInt))
	}
	
	// Test MinWindowSubstring
	fmt.Println("\nTesting MinWindowSubstring:")
	s = "ADOBECODEBANC"
	t := "ABC"
	resultStr := sw.MinWindowSubstring(s, t)
	fmt.Printf("Source: '%s', Target: '%s', Min Window: '%s'\n", s, t, resultStr)
	if resultStr != "BANC" {
		panic(fmt.Sprintf("Expected 'BANC', got '%s'", resultStr))
	}
	
	// Test LongestSubstringKDistinct
	fmt.Println("\nTesting LongestSubstringKDistinct:")
	s = "araaci"
	k = 2
	resultInt = sw.LongestSubstringKDistinct(s, k)
	fmt.Printf("String: '%s', K: %d, Longest Length: %d\n", s, k, resultInt)
	if resultInt != 4 {
		panic(fmt.Sprintf("Expected 4, got %d", resultInt))
	}
	
	// Test SubarrayProductLessThanK
	fmt.Println("\nTesting SubarrayProductLessThanK:")
	nums := []int{10, 5, 2, 6}
	k = 100
	resultInt = sw.SubarrayProductLessThanK(nums, k)
	fmt.Printf("Array: %v, K: %d, Count: %d\n", nums, k, resultInt)
	if resultInt != 8 {
		panic(fmt.Sprintf("Expected 8, got %d", resultInt))
	}
	
	// Test FruitsIntoBaskets
	fmt.Println("\nTesting FruitsIntoBaskets:")
	fruits := []int{1, 2, 1, 2, 3, 2, 2}
	resultInt = sw.FruitsIntoBaskets(fruits)
	fmt.Printf("Fruits: %v, Max Fruits: %d\n", fruits, resultInt)
	if resultInt != 5 {
		panic(fmt.Sprintf("Expected 5, got %d", resultInt))
	}
	
	// Test SmallestSubarraySum
	fmt.Println("\nTesting SmallestSubarraySum:")
	arr = []int{2, 3, 1, 2, 4, 3}
	target := 7
	resultInt = sw.SmallestSubarraySum(arr, target)
	fmt.Printf("Array: %v, Target: %d, Min Length: %d\n", arr, target, resultInt)
	if resultInt != 2 {
		panic(fmt.Sprintf("Expected 2, got %d", resultInt))
	}
	
	// Test FindAnagrams
	fmt.Println("\nTesting FindAnagrams:")
	s = "abab"
	p := "ab"
	resultSlice := sw.FindAnagrams(s, p)
	fmt.Printf("String: '%s', Pattern: '%s', Indices: %v\n", s, p, resultSlice)
	expected := []int{0, 2}
	if !equalSlices(resultSlice, expected) {
		panic(fmt.Sprintf("Expected %v, got %v", expected, resultSlice))
	}
	
	// Test CharacterReplacement
	fmt.Println("\nTesting CharacterReplacement:")
	s = "AABABBA"
	k = 1
	resultInt = sw.CharacterReplacement(s, k)
	fmt.Printf("String: '%s', K: %d, Max Length: %d\n", s, k, resultInt)
	if resultInt != 4 {
		panic(fmt.Sprintf("Expected 4, got %d", resultInt))
	}
	
	fmt.Println("\n✅ All tests passed!")
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	
	return true
}

func main() {
	runTests()
}