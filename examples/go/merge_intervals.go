package main

import (
	"fmt"
	"sort"
)

/*
Merge Intervals Pattern Examples

The Merge Intervals pattern is used to deal with overlapping intervals.
It involves sorting intervals and then merging or manipulating them based on overlap.

Time Complexity: O(n log n) due to sorting
Space Complexity: O(n) for output
*/

// MergeIntervals merges overlapping intervals
// Time: O(n log n), Space: O(n)
func MergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// Sort intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		lastMerged := merged[len(merged)-1]

		// If current interval overlaps with last merged interval
		if current[0] <= lastMerged[1] {
			// Merge intervals by updating end time
			if current[1] > lastMerged[1] {
				lastMerged[1] = current[1]
			}
		} else {
			// No overlap, add current interval
			merged = append(merged, current)
		}
	}

	return merged
}

// InsertInterval inserts a new interval into a sorted list of non-overlapping intervals
// Time: O(n), Space: O(n)
func InsertInterval(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	i := 0

	// Add all intervals that end before new interval starts
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	// Merge overlapping intervals with new interval
	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		if intervals[i][0] < newInterval[0] {
			newInterval[0] = intervals[i][0]
		}
		if intervals[i][1] > newInterval[1] {
			newInterval[1] = intervals[i][1]
		}
		i++
	}

	result = append(result, newInterval)

	// Add remaining intervals
	for i < len(intervals) {
		result = append(result, intervals[i])
		i++
	}

	return result
}

// IntervalIntersection finds intersection of two lists of intervals
// Time: O(m + n), Space: O(min(m, n))
func IntervalIntersection(firstList, secondList [][]int) [][]int {
	result := [][]int{}
	i, j := 0, 0

	for i < len(firstList) && j < len(secondList) {
		// Find intersection
		start := max(firstList[i][0], secondList[j][0])
		end := min(firstList[i][1], secondList[j][1])

		// If valid intersection exists
		if start <= end {
			result = append(result, []int{start, end})
		}

		// Move pointer of interval that ends first
		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}

	return result
}

// CanAttendMeetings checks if a person can attend all meetings
// Time: O(n log n), Space: O(1)
func CanAttendMeetings(intervals [][]int) bool {
	if len(intervals) == 0 {
		return true
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}

	return true
}

// MinMeetingRooms finds minimum number of meeting rooms required
// Time: O(n log n), Space: O(n)
func MinMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// Create separate slices for start and end times
	starts := make([]int, len(intervals))
	ends := make([]int, len(intervals))

	for i, interval := range intervals {
		starts[i] = interval[0]
		ends[i] = interval[1]
	}

	sort.Ints(starts)
	sort.Ints(ends)

	roomsNeeded := 0
	maxRooms := 0
	startPtr, endPtr := 0, 0

	for startPtr < len(intervals) {
		// If a meeting starts before another ends, we need a new room
		if starts[startPtr] < ends[endPtr] {
			roomsNeeded++
			if roomsNeeded > maxRooms {
				maxRooms = roomsNeeded
			}
			startPtr++
		} else {
			// A meeting ended, free up a room
			roomsNeeded--
			endPtr++
		}
	}

	return maxRooms
}

// EmployeeFreeTime finds common free time for all employees
// Time: O(n log n), Space: O(n)
func EmployeeFreeTime(schedule [][][]int) [][]int {
	// Flatten all intervals
	allIntervals := [][]int{}
	for _, employeeSchedule := range schedule {
		allIntervals = append(allIntervals, employeeSchedule...)
	}

	// Merge all intervals
	merged := MergeIntervals(allIntervals)

	// Find gaps between merged intervals
	freeTime := [][]int{}
	for i := 1; i < len(merged); i++ {
		if merged[i-1][1] < merged[i][0] {
			freeTime = append(freeTime, []int{merged[i-1][1], merged[i][0]})
		}
	}

	return freeTime
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

// Helper function to check if two 2D slices are equal
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

// Test functions
func testMergeIntervals() {
	fmt.Println("Testing merge intervals...")

	testCases := []struct {
		input    [][]int
		expected [][]int
	}{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{[][]int{{1, 4}, {0, 4}}, [][]int{{0, 4}}},
		{[][]int{{1, 4}, {2, 3}}, [][]int{{1, 4}}},
	}

	for _, tc := range testCases {
		result := MergeIntervals(tc.input)
		if !equal2D(result, tc.expected) {
			panic(fmt.Sprintf("Expected %v, got %v", tc.expected, result))
		}
	}

	fmt.Println("✅ Merge intervals tests passed")
}

func testInsertInterval() {
	fmt.Println("Testing insert interval...")

	testCases := []struct {
		intervals   [][]int
		newInterval []int
		expected    [][]int
	}{
		{[][]int{{1, 3}, {6, 9}}, []int{2, 5}, [][]int{{1, 5}, {6, 9}}},
		{[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}, [][]int{{1, 2}, {3, 10}, {12, 16}}},
		{[][]int{}, []int{5, 7}, [][]int{{5, 7}}},
		{[][]int{{1, 5}}, []int{2, 3}, [][]int{{1, 5}}},
	}

	for _, tc := range testCases {
		result := InsertInterval(tc.intervals, tc.newInterval)
		if !equal2D(result, tc.expected) {
			panic(fmt.Sprintf("Expected %v, got %v", tc.expected, result))
		}
	}

	fmt.Println("✅ Insert interval tests passed")
}

func testIntervalIntersection() {
	fmt.Println("Testing interval intersection...")

	first := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
	second := [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
	expected := [][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}}

	result := IntervalIntersection(first, second)
	if !equal2D(result, expected) {
		panic(fmt.Sprintf("Expected %v, got %v", expected, result))
	}

	fmt.Println("✅ Interval intersection tests passed")
}

func testMeetingRooms() {
	fmt.Println("Testing meeting room problems...")

	// Can attend all meetings
	if CanAttendMeetings([][]int{{0, 30}, {5, 10}, {15, 20}}) != false {
		panic("Expected false for overlapping meetings")
	}
	if CanAttendMeetings([][]int{{7, 10}, {2, 4}}) != true {
		panic("Expected true for non-overlapping meetings")
	}

	// Minimum meeting rooms
	if MinMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}}) != 2 {
		panic("Expected 2 meeting rooms")
	}
	if MinMeetingRooms([][]int{{7, 10}, {2, 4}}) != 1 {
		panic("Expected 1 meeting room")
	}
	if MinMeetingRooms([][]int{{9, 10}, {4, 9}, {4, 17}}) != 2 {
		panic("Expected 2 meeting rooms")
	}

	fmt.Println("✅ Meeting room tests passed")
}

func testEmployeeFreeTime() {
	fmt.Println("Testing employee free time...")

	schedule := [][][]int{
		{{1, 3}, {6, 7}},
		{{2, 4}},
		{{2, 5}, {9, 12}},
	}
	expected := [][]int{{5, 6}, {7, 9}}

	result := EmployeeFreeTime(schedule)
	if !equal2D(result, expected) {
		panic(fmt.Sprintf("Expected %v, got %v", expected, result))
	}

	fmt.Println("✅ Employee free time tests passed")
}

func main() {
	fmt.Println("Testing Merge Intervals Pattern...")
	testMergeIntervals()
	testInsertInterval()
	testIntervalIntersection()
	testMeetingRooms()
	testEmployeeFreeTime()
	fmt.Println("\n🎉 All Merge Intervals tests passed!")
}