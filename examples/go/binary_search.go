// Binary Search Pattern Examples
//
// This package demonstrates binary search patterns and variations
// for efficient searching and optimization problems.

package main

import (
	"fmt"
	"math"
)

// BasicBinarySearch contains basic binary search implementations
type BasicBinarySearch struct{}

// BinarySearch performs standard binary search returning index or -1
func (bbs *BasicBinarySearch) BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// BinarySearchRecursive performs recursive binary search
func (bbs *BasicBinarySearch) BinarySearchRecursive(arr []int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return bbs.BinarySearchRecursive(arr, target, mid+1, right)
	} else {
		return bbs.BinarySearchRecursive(arr, target, left, mid-1)
	}
}

// FindFirstOccurrence finds first occurrence of target
func (bbs *BasicBinarySearch) FindFirstOccurrence(arr []int, target int) int {
	left, right := 0, len(arr)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			result = mid
			right = mid - 1 // Continue searching left
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

// FindLastOccurrence finds last occurrence of target
func (bbs *BasicBinarySearch) FindLastOccurrence(arr []int, target int) int {
	left, right := 0, len(arr)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			result = mid
			left = mid + 1 // Continue searching right
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

// BinarySearchVariations contains advanced binary search patterns
type BinarySearchVariations struct{}

// SearchInsertPosition finds position where target should be inserted
func (bsv *BinarySearchVariations) SearchInsertPosition(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

// SearchRotatedArray searches in rotated sorted array
func (bsv *BinarySearchVariations) SearchRotatedArray(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		// Left half is sorted
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // Right half is sorted
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

// FindMinimumRotated finds minimum in rotated sorted array
func (bsv *BinarySearchVariations) FindMinimumRotated(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}

// Search2DMatrix searches in row and column sorted matrix
func (bsv *BinarySearchVariations) Search2DMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows, cols := len(matrix), len(matrix[0])
	left, right := 0, rows*cols-1

	for left <= right {
		mid := left + (right-left)/2
		midValue := matrix[mid/cols][mid%cols]

		if midValue == target {
			return true
		} else if midValue < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

// OptimizationProblems contains problems using binary search for optimization
type OptimizationProblems struct{}

// CapacityToShipPackages finds minimum capacity to ship packages in D days
func (opt *OptimizationProblems) CapacityToShipPackages(weights []int, D int) int {
	canShip := func(capacity int) bool {
		days, currentLoad := 1, 0

		for _, weight := range weights {
			if currentLoad+weight > capacity {
				days++
				currentLoad = weight
			} else {
				currentLoad += weight
			}
		}

		return days <= D
	}

	left, right := 0, 0
	for _, weight := range weights {
		if weight > left {
			left = weight
		}
		right += weight
	}

	for left < right {
		mid := left + (right-left)/2

		if canShip(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

// SplitArrayLargestSum splits array to minimize largest sum
func (opt *OptimizationProblems) SplitArrayLargestSum(nums []int, m int) int {
	canSplit := func(maxSum int) bool {
		splits, currentSum := 1, 0

		for _, num := range nums {
			if currentSum+num > maxSum {
				splits++
				currentSum = num
			} else {
				currentSum += num
			}
		}

		return splits <= m
	}

	left, right := 0, 0
	for _, num := range nums {
		if num > left {
			left = num
		}
		right += num
	}

	for left < right {
		mid := left + (right-left)/2

		if canSplit(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

// SmallestDivisor finds smallest divisor such that sum <= threshold
func (opt *OptimizationProblems) SmallestDivisor(nums []int, threshold int) int {
	getSum := func(divisor int) int {
		sum := 0
		for _, num := range nums {
			sum += int(math.Ceil(float64(num) / float64(divisor)))
		}
		return sum
	}

	left, right := 1, 0
	for _, num := range nums {
		if num > right {
			right = num
		}
	}

	for left < right {
		mid := left + (right-left)/2

		if getSum(mid) <= threshold {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

// MedianProblems contains problems finding medians using binary search
type MedianProblems struct{}

// FindMedianSortedArrays finds median of two sorted arrays
func (mp *MedianProblems) FindMedianSortedArrays(nums1, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	total := m + n
	half := total / 2

	left, right := 0, m

	for left <= right {
		i := (left + right) / 2
		j := half - i

		var nums1Left, nums1Right, nums2Left, nums2Right float64

		if i > 0 {
			nums1Left = float64(nums1[i-1])
		} else {
			nums1Left = math.Inf(-1)
		}

		if i < m {
			nums1Right = float64(nums1[i])
		} else {
			nums1Right = math.Inf(1)
		}

		if j > 0 {
			nums2Left = float64(nums2[j-1])
		} else {
			nums2Left = math.Inf(-1)
		}

		if j < n {
			nums2Right = float64(nums2[j])
		} else {
			nums2Right = math.Inf(1)
		}

		if nums1Left <= nums2Right && nums2Left <= nums1Right {
			if total%2 == 0 {
				return (math.Max(nums1Left, nums2Left) + math.Min(nums1Right, nums2Right)) / 2
			} else {
				return math.Min(nums1Right, nums2Right)
			}
		} else if nums1Left > nums2Right {
			right = i - 1
		} else {
			left = i + 1
		}
	}

	return 0.0
}

// PeakFinding contains peak finding problems using binary search
type PeakFinding struct{}

// FindPeakElement finds any peak element index
func (pf *PeakFinding) FindPeakElement(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

// FindPeak2D finds peak in 2D matrix
func (pf *PeakFinding) FindPeak2D(matrix [][]int) []int {
	rows, cols := len(matrix), len(matrix[0])
	left, right := 0, cols-1

	for left <= right {
		midCol := left + (right-left)/2
		maxRow := 0

		// Find maximum element in middle column
		for i := 0; i < rows; i++ {
			if matrix[i][midCol] > matrix[maxRow][midCol] {
				maxRow = i
			}
		}

		// Check if it's a peak
		leftVal := -1
		if midCol > 0 {
			leftVal = matrix[maxRow][midCol-1]
		}

		rightVal := -1
		if midCol < cols-1 {
			rightVal = matrix[maxRow][midCol+1]
		}

		if matrix[maxRow][midCol] >= leftVal && matrix[maxRow][midCol] >= rightVal {
			return []int{maxRow, midCol}
		} else if leftVal > matrix[maxRow][midCol] {
			right = midCol - 1
		} else {
			left = midCol + 1
		}
	}

	return []int{-1, -1}
}

func main() {
	// Test basic binary search
	bbs := &BasicBinarySearch{}
	arr := []int{1, 3, 5, 5, 5, 7, 9, 11}

	fmt.Println("Basic binary search for 5:", bbs.BinarySearch(arr, 5))
	fmt.Println("Recursive binary search for 5:", bbs.BinarySearchRecursive(arr, 5, 0, len(arr)-1))
	fmt.Println("First occurrence of 5:", bbs.FindFirstOccurrence(arr, 5))
	fmt.Println("Last occurrence of 5:", bbs.FindLastOccurrence(arr, 5))

	// Test binary search variations
	bsv := &BinarySearchVariations{}

	nums := []int{1, 3, 5, 6}
	fmt.Println("Insert position for 5:", bsv.SearchInsertPosition(nums, 5))
	fmt.Println("Insert position for 2:", bsv.SearchInsertPosition(nums, 2))

	rotated := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println("Search 0 in rotated array:", bsv.SearchRotatedArray(rotated, 0))
	fmt.Println("Minimum in rotated array:", bsv.FindMinimumRotated(rotated))

	matrix := [][]int{{1, 4, 7, 11}, {2, 5, 8, 12}, {3, 6, 9, 16}, {10, 13, 14, 17}}
	fmt.Println("Search 5 in 2D matrix:", bsv.Search2DMatrix(matrix, 5))

	// Test optimization problems
	opt := &OptimizationProblems{}

	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Ship capacity for 5 days:", opt.CapacityToShipPackages(weights, 5))

	nums2 := []int{7, 2, 5, 10, 8}
	fmt.Println("Split array largest sum (m=2):", opt.SplitArrayLargestSum(nums2, 2))

	nums3 := []int{1, 2, 5, 9}
	fmt.Println("Smallest divisor (threshold=6):", opt.SmallestDivisor(nums3, 6))

	// Test median problems
	mp := &MedianProblems{}

	nums1, nums2_med := []int{1, 3}, []int{2}
	fmt.Println("Median of two arrays:", mp.FindMedianSortedArrays(nums1, nums2_med))

	// Test peak finding
	pf := &PeakFinding{}

	nums4 := []int{1, 2, 3, 1}
	fmt.Println("Peak element index:", pf.FindPeakElement(nums4))

	matrix2D := [][]int{{1, 4}, {3, 2}}
	fmt.Println("Peak in 2D matrix:", pf.FindPeak2D(matrix2D))
}