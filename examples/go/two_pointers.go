// Package main demonstrates Two Pointers pattern implementations in Go
//
// Two Pointers is a technique that uses two pointers to iterate through
// a data structure in tandem until one or both pointers hit a condition.
package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

// TwoPointers contains implementations of the Two Pointers pattern
type TwoPointers struct{}

// TwoSumSorted finds two numbers that add up to target in sorted array
// Time: O(n), Space: O(1)
func (tp *TwoPointers) TwoSumSorted(arr []int, target int) []int {
	left, right := 0, len(arr)-1
	
	for left < right {
		currentSum := arr[left] + arr[right]
		if currentSum == target {
			return []int{left, right}
		} else if currentSum < target {
			left++
		} else {
			right--
		}
	}
	
	return []int{-1, -1}
}

// RemoveDuplicates removes duplicates from sorted array in-place
// Time: O(n), Space: O(1)
func (tp *TwoPointers) RemoveDuplicates(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	
	slow := 0 // Position for next unique element
	
	for fast := 1; fast < len(arr); fast++ {
		if arr[fast] != arr[slow] {
			slow++
			arr[slow] = arr[fast]
		}
	}
	
	return slow + 1
}

// IsPalindrome checks if string is palindrome (case-insensitive, alphanumeric only)
// Time: O(n), Space: O(1)
func (tp *TwoPointers) IsPalindrome(s string) bool {
	left, right := 0, len(s)-1
	
	for left < right {
		// Skip non-alphanumeric characters
		for left < right && !isAlphanumeric(s[left]) {
			left++
		}
		for left < right && !isAlphanumeric(s[right]) {
			right--
		}
		
		if strings.ToLower(string(s[left])) != strings.ToLower(string(s[right])) {
			return false
		}
		
		left++
		right--
	}
	
	return true
}

// ThreeSum finds all unique triplets that sum to zero
// Time: O(n²), Space: O(1) excluding output
func (tp *TwoPointers) ThreeSum(arr []int) [][]int {
	sort.Ints(arr)
	result := [][]int{}
	n := len(arr)
	
	for i := 0; i < n-2; i++ {
		// Skip duplicates for first element
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		
		left, right := i+1, n-1
		
		for left < right {
			currentSum := arr[i] + arr[left] + arr[right]
			
			if currentSum == 0 {
				result = append(result, []int{arr[i], arr[left], arr[right]})
				
				// Skip duplicates
				for left < right && arr[left] == arr[left+1] {
					left++
				}
				for left < right && arr[right] == arr[right-1] {
					right--
				}
				
				left++
				right--
			} else if currentSum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	
	return result
}

// MaxArea solves container with most water problem
// Time: O(n), Space: O(1)
func (tp *TwoPointers) MaxArea(height []int) int {
	left, right := 0, len(height)-1
	maxWater := 0
	
	for left < right {
		width := right - left
		currentArea := min(height[left], height[right]) * width
		maxWater = max(maxWater, currentArea)
		
		// Move pointer with smaller height
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	
	return maxWater
}

// SortColors solves Dutch flag problem - sort array of 0s, 1s, and 2s
// Time: O(n), Space: O(1)
func (tp *TwoPointers) SortColors(nums []int) {
	left, right := 0, len(nums)-1
	current := 0
	
	for current <= right {
		if nums[current] == 0 {
			nums[left], nums[current] = nums[current], nums[left]
			left++
			current++
		} else if nums[current] == 2 {
			nums[current], nums[right] = nums[right], nums[current]
			right--
			// Don't increment current as we need to check swapped element
		} else { // nums[current] == 1
			current++
		}
	}
}

// TrapRainWater solves trapping rain water problem
// Time: O(n), Space: O(1)
func (tp *TwoPointers) TrapRainWater(height []int) int {
	if len(height) == 0 {
		return 0
	}
	
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	waterTrapped := 0
	
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				waterTrapped += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				waterTrapped += rightMax - height[right]
			}
			right--
		}
	}
	
	return waterTrapped
}

// Helper functions

func isAlphanumeric(c byte) bool {
	return unicode.IsLetter(rune(c)) || unicode.IsDigit(rune(c))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Test functions

func runTests() {
	tp := &TwoPointers{}
	
	// Test TwoSumSorted
	fmt.Println("Testing TwoSumSorted:")
	arr := []int{1, 2, 3, 4, 6}
	target := 6
	result := tp.TwoSumSorted(arr, target)
	fmt.Printf("Array: %v, Target: %d, Result: %v\n", arr, target, result)
	assertEqual(result, []int{1, 3}, "TwoSumSorted")
	
	// Test RemoveDuplicates
	fmt.Println("\nTesting RemoveDuplicates:")
	arr = []int{1, 1, 2, 2, 3, 3, 4}
	original := make([]int, len(arr))
	copy(original, arr)
	length := tp.RemoveDuplicates(arr)
	fmt.Printf("Original: %v, After: %v, Length: %d\n", original, arr[:length], length)
	assertEqual(arr[:length], []int{1, 2, 3, 4}, "RemoveDuplicates")
	
	// Test IsPalindrome
	fmt.Println("\nTesting IsPalindrome:")
	testCases := []struct {
		input    string
		expected bool
	}{
		{"A man a plan a canal Panama", true},
		{"race a car", false},
		{"", true},
	}
	
	for _, tc := range testCases {
		result := tp.IsPalindrome(tc.input)
		fmt.Printf("String: '%s', Palindrome: %t\n", tc.input, result)
		if result != tc.expected {
			panic(fmt.Sprintf("Expected %t, got %t for '%s'", tc.expected, result, tc.input))
		}
	}
	
	// Test ThreeSum
	fmt.Println("\nTesting ThreeSum:")
	arr = []int{-1, 0, 1, 2, -1, -4}
	result2D := tp.ThreeSum(arr)
	fmt.Printf("Array: %v, Triplets: %v\n", arr, result2D)
	expected2D := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	if !equal2D(result2D, expected2D) {
		panic(fmt.Sprintf("Expected %v, got %v", expected2D, result2D))
	}
	
	// Test MaxArea
	fmt.Println("\nTesting MaxArea:")
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	resultInt := tp.MaxArea(height)
	fmt.Printf("Heights: %v, Max Area: %d\n", height, resultInt)
	if resultInt != 49 {
		panic(fmt.Sprintf("Expected 49, got %d", resultInt))
	}
	
	// Test SortColors
	fmt.Println("\nTesting SortColors:")
	nums := []int{2, 0, 2, 1, 1, 0}
	original = make([]int, len(nums))
	copy(original, nums)
	tp.SortColors(nums)
	fmt.Printf("Original: %v, Sorted: %v\n", original, nums)
	assertEqual(nums, []int{0, 0, 1, 1, 2, 2}, "SortColors")
	
	// Test TrapRainWater
	fmt.Println("\nTesting TrapRainWater:")
	height = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	resultInt = tp.TrapRainWater(height)
	fmt.Printf("Heights: %v, Water Trapped: %d\n", height, resultInt)
	if resultInt != 6 {
		panic(fmt.Sprintf("Expected 6, got %d", resultInt))
	}
	
	fmt.Println("\n✅ All tests passed!")
}

// Helper test functions

func assertEqual(actual, expected []int, testName string) {
	if len(actual) != len(expected) {
		panic(fmt.Sprintf("%s: Expected length %d, got %d", testName, len(expected), len(actual)))
	}
	
	for i := range actual {
		if actual[i] != expected[i] {
			panic(fmt.Sprintf("%s: Expected %v, got %v", testName, expected, actual))
		}
	}
}

func equal2D(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	
	return true
}

func main() {
	runTests()
}