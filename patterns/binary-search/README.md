# Binary Search Pattern

## Overview
Binary Search is a highly efficient algorithm for finding elements in sorted data structures. It repeatedly divides the search space in half, eliminating half of the remaining elements with each iteration. This pattern is essential for optimization problems and searching in sorted arrays.

## When to Use
- **Sorted Arrays**: Searching in sorted arrays or lists
- **Search Space Reduction**: When you can eliminate half the possibilities each step
- **Optimization Problems**: Finding minimum/maximum values with a condition
- **Range Queries**: Finding first/last occurrence of elements
- **Matrix Search**: Searching in sorted 2D matrices

## Time/Space Complexity
- **Time**: O(log n) - Halves search space each iteration
- **Space**: O(1) for iterative, O(log n) for recursive

## Pattern Variations

### 1. Standard Binary Search
```python
def binary_search(arr, target):
    left, right = 0, len(arr) - 1
    
    while left <= right:
        mid = left + (right - left) // 2
        
        if arr[mid] == target:
            return mid
        elif arr[mid] < target:
            left = mid + 1
        else:
            right = mid - 1
    
    return -1
```

### 2. Binary Search for Insertion Point
```python
def search_insert_position(arr, target):
    left, right = 0, len(arr)
    
    while left < right:
        mid = left + (right - left) // 2
        
        if arr[mid] < target:
            left = mid + 1
        else:
            right = mid
    
    return left
```

## Common Problem Patterns

### Pattern 1: Find Element in Sorted Array
**Problem**: Find the index of a target element in a sorted array.

```python
def search(nums, target):
    left, right = 0, len(nums) - 1
    
    while left <= right:
        mid = left + (right - left) // 2
        
        if nums[mid] == target:
            return mid
        elif nums[mid] < target:
            left = mid + 1
        else:
            right = mid - 1
    
    return -1

# Example usage
nums = [1, 3, 5, 7, 9, 11]
target = 7
print(search(nums, target))  # Output: 3
```

### Pattern 2: Search in Rotated Sorted Array
**Problem**: Search for a target value in a rotated sorted array.

```python
def search_rotated(nums, target):
    left, right = 0, len(nums) - 1
    
    while left <= right:
        mid = left + (right - left) // 2
        
        if nums[mid] == target:
            return mid
        
        # Left half is sorted
        if nums[left] <= nums[mid]:
            if nums[left] <= target < nums[mid]:
                right = mid - 1
            else:
                left = mid + 1
        # Right half is sorted
        else:
            if nums[mid] < target <= nums[right]:
                left = mid + 1
            else:
                right = mid - 1
    
    return -1

# Example usage
nums = [4, 5, 6, 7, 0, 1, 2]
target = 0
print(search_rotated(nums, target))  # Output: 4
```

### Pattern 3: Find First and Last Position
**Problem**: Find the first and last position of a target element.

```python
def search_range(nums, target):
    def find_first():
        left, right = 0, len(nums) - 1
        result = -1
        
        while left <= right:
            mid = left + (right - left) // 2
            
            if nums[mid] == target:
                result = mid
                right = mid - 1  # Continue searching left
            elif nums[mid] < target:
                left = mid + 1
            else:
                right = mid - 1
        
        return result
    
    def find_last():
        left, right = 0, len(nums) - 1
        result = -1
        
        while left <= right:
            mid = left + (right - left) // 2
            
            if nums[mid] == target:
                result = mid
                left = mid + 1  # Continue searching right
            elif nums[mid] < target:
                left = mid + 1
            else:
                right = mid - 1
        
        return result
    
    first = find_first()
    if first == -1:
        return [-1, -1]
    
    last = find_last()
    return [first, last]

# Example usage
nums = [5, 7, 7, 8, 8, 10]
target = 8
print(search_range(nums, target))  # Output: [3, 4]
```

### Pattern 4: Search in 2D Matrix
**Problem**: Search for a target value in a 2D matrix where each row and column is sorted.

```python
def search_matrix(matrix, target):
    if not matrix or not matrix[0]:
        return False
    
    rows, cols = len(matrix), len(matrix[0])
    left, right = 0, rows * cols - 1
    
    while left <= right:
        mid = left + (right - left) // 2
        mid_value = matrix[mid // cols][mid % cols]
        
        if mid_value == target:
            return True
        elif mid_value < target:
            left = mid + 1
        else:
            right = mid - 1
    
    return False

# Example usage
matrix = [
    [1,  4,  7,  11],
    [2,  5,  8,  12],
    [3,  6,  9,  16]
]
target = 5
print(search_matrix(matrix, target))  # Output: True
```

### Pattern 5: Find Peak Element
**Problem**: Find a peak element in an array (element that is greater than its neighbors).

```python
def find_peak_element(nums):
    left, right = 0, len(nums) - 1
    
    while left < right:
        mid = left + (right - left) // 2
        
        if nums[mid] > nums[mid + 1]:
            # Peak is in left half (including mid)
            right = mid
        else:
            # Peak is in right half
            left = mid + 1
    
    return left

# Example usage
nums = [1, 2, 3, 1]
print(find_peak_element(nums))  # Output: 2 (element 3 is peak)
```

### Pattern 6: Square Root (Integer)
**Problem**: Find the integer square root of a number.

```python
def my_sqrt(x):
    if x < 2:
        return x
    
    left, right = 2, x // 2
    
    while left <= right:
        mid = left + (right - left) // 2
        num = mid * mid
        
        if num == x:
            return mid
        elif num < x:
            left = mid + 1
        else:
            right = mid - 1
    
    return right

# Example usage
print(my_sqrt(8))  # Output: 2
print(my_sqrt(16))  # Output: 4
```

## Practice Problems

### Easy
1. **Binary Search** - Find target in sorted array
2. **Search Insert Position** - Find insertion index
3. **First Bad Version** - Find first bad version
4. **Sqrt(x)** - Integer square root

### Medium
1. **Search in Rotated Sorted Array** - Rotated array search
2. **Find First and Last Position** - Range search
3. **Search a 2D Matrix** - 2D binary search
4. **Find Peak Element** - Peak finding
5. **Find Minimum in Rotated Sorted Array** - Minimum element

### Hard
1. **Median of Two Sorted Arrays** - Find median efficiently
2. **Search in Rotated Sorted Array II** - With duplicates
3. **Split Array Largest Sum** - Minimize maximum subarray sum
4. **Koko Eating Bananas** - Optimization problem

## Tips and Tricks

1. **Avoid Overflow**: Use `mid = left + (right - left) // 2` instead of `(left + right) // 2`
2. **Loop Invariants**: Maintain clear invariants about what left and right represent
3. **Boundary Conditions**: Be careful with `<=` vs `<` in while conditions
4. **Template Consistency**: Use consistent templates for different search types
5. **Integer Division**: Remember integer division behavior in your language

## Common Mistakes

1. **Infinite Loops**: Incorrect loop conditions causing infinite loops
2. **Off-by-One Errors**: Wrong boundary calculations
3. **Overflow Issues**: Integer overflow when calculating mid
4. **Template Confusion**: Mixing different binary search templates

## Related Patterns

- **Two Pointers**: Uses similar left/right pointer concept
- **Sliding Window**: Can be optimized using binary search
- **Divide and Conquer**: Binary search is a divide and conquer algorithm

## Implementation Languages

The pattern works across all languages:
- **Python**: Use integer division `//`
- **Java**: Watch for integer overflow with `left + (right - left) / 2`
- **JavaScript**: Use `Math.floor()` for integer division
- **C++**: Similar overflow considerations as Java
- **Go**: Integer division is built-in