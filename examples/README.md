# Code Examples

This directory contains working implementations of coding patterns in multiple programming languages.

## Structure

```
examples/
├── python/           # Python3 implementations
│   ├── two_pointers.py
│   ├── sliding_window.py
│   ├── fast_slow_pointers.py
│   ├── merge_intervals.py
│   └── dynamic_programming.py
├── go/              # Go implementations
│   ├── two_pointers.go
│   ├── sliding_window.go
│   ├── fast_slow_pointers.go
│   ├── merge_intervals.go
│   └── dynamic_programming.go
└── README.md        # This file
```

## How to Run

### Python Examples

Prerequisites: Python 3.6+

```bash
# Run Two Pointers examples
python3 examples/python/two_pointers.py

# Run Sliding Window examples  
python3 examples/python/sliding_window.py

# Run Fast & Slow Pointers examples
python3 examples/python/fast_slow_pointers.py

# Run Merge Intervals examples
python3 examples/python/merge_intervals.py

# Run Dynamic Programming examples
python3 examples/python/dynamic_programming.py
```

### Go Examples

Prerequisites: Go 1.16+

```bash
# Run Two Pointers examples
go run examples/go/two_pointers.go

# Run Sliding Window examples
go run examples/go/sliding_window.go

# Run Fast & Slow Pointers examples
go run examples/go/fast_slow_pointers.go

# Run Merge Intervals examples
go run examples/go/merge_intervals.go

# Run Dynamic Programming examples
go run examples/go/dynamic_programming.go
```

## Available Patterns

### Two Pointers
- ✅ Two Sum (Sorted Array)
- ✅ Remove Duplicates
- ✅ Palindrome Check
- ✅ Three Sum
- ✅ Container With Most Water
- ✅ Sort Colors (Dutch Flag)
- ✅ Trapping Rain Water

### Sliding Window
- ✅ Maximum Sum Subarray of Size K
- ✅ Longest Substring Without Repeating Characters
- ✅ Minimum Window Substring
- ✅ Longest Substring with K Distinct Characters
- ✅ Subarray Product Less Than K
- ✅ Fruits Into Baskets
- ✅ Smallest Subarray Sum
- ✅ Find All Anagrams
- ✅ Character Replacement

### Fast & Slow Pointers
- ✅ Cycle Detection (Floyd's Algorithm)
- ✅ Find Cycle Start
- ✅ Find Middle of Linked List
- ✅ Palindrome Linked List
- ✅ Happy Number
- ✅ Circular Array Loop

### Merge Intervals
- ✅ Merge Overlapping Intervals
- ✅ Insert Interval
- ✅ Interval Intersection
- ✅ Meeting Rooms I & II
- ✅ Employee Free Time
- ✅ Intervals with Names

### Dynamic Programming
- ✅ Fibonacci Numbers
- ✅ Climbing Stairs
- ✅ House Robber
- ✅ Coin Change
- ✅ Longest Increasing Subsequence
- ✅ 0/1 Knapsack
- ✅ Longest Common Subsequence
- ✅ Edit Distance
- ✅ Word Break
- ✅ Maximum Product Subarray

## Testing

All implementations include comprehensive test suites that verify correctness with example inputs and expected outputs.

## Contributing

When adding new patterns:

1. Follow the existing code structure and naming conventions
2. Include comprehensive test cases
3. Add proper documentation and time/space complexity analysis
4. Update this README with the new pattern