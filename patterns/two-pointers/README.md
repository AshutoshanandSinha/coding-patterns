# Two Pointers Pattern

## Overview
The Two Pointers technique uses two pointers to iterate through a data structure in tandem until one or both pointers hit a certain condition. This pattern is highly effective for solving problems involving arrays or linked lists where you need to find pairs of elements, reverse elements, or search for triplets.

## When to Use
- **Sorted Arrays/Lists**: When dealing with sorted arrays or lists
- **Pair Problems**: Finding pairs that meet certain criteria (sum, difference, etc.)
- **Palindrome Problems**: Checking if a string/array is a palindrome
- **Reverse Operations**: Reversing arrays or strings in-place
- **Sliding Window Alternative**: When you need to shrink/expand a window based on conditions

## Time/Space Complexity
- **Time**: O(n) - Single pass through the array
- **Space**: O(1) - Only two pointers used

## Pattern Variations

### 1. Opposite Direction (Most Common)
```python
def two_pointers_opposite(arr):
    left, right = 0, len(arr) - 1
    
    while left < right:
        # Process arr[left] and arr[right]
        if condition_met:
            # Found solution
            return result
        elif need_smaller_sum:
            right -= 1
        else:
            left += 1
    
    return not_found
```

### 2. Same Direction
```python
def two_pointers_same_direction(arr):
    slow, fast = 0, 0
    
    while fast < len(arr):
        # Process elements
        if condition:
            # Move slow pointer
            slow += 1
        fast += 1
    
    return result
```

## Common Problem Patterns

### Pattern 1: Target Sum (Two Sum in Sorted Array)
**Problem**: Find two numbers that add up to a target sum in a sorted array.

```python
def two_sum_sorted(arr, target):
    left, right = 0, len(arr) - 1
    
    while left < right:
        current_sum = arr[left] + arr[right]
        if current_sum == target:
            return [left, right]
        elif current_sum < target:
            left += 1
        else:
            right -= 1
    
    return [-1, -1]

# Example usage
arr = [1, 2, 3, 4, 6]
target = 6
print(two_sum_sorted(arr, target))  # Output: [1, 3]
```

### Pattern 2: Remove Duplicates
**Problem**: Remove duplicates from a sorted array in-place.

```python
def remove_duplicates(arr):
    if not arr:
        return 0
    
    slow = 0  # Position to place next unique element
    
    for fast in range(1, len(arr)):
        if arr[fast] != arr[slow]:
            slow += 1
            arr[slow] = arr[fast]
    
    return slow + 1

# Example usage
arr = [1, 1, 2, 2, 3, 3, 4]
length = remove_duplicates(arr)
print(arr[:length])  # Output: [1, 2, 3, 4]
```

### Pattern 3: Palindrome Check
**Problem**: Check if a string is a palindrome.

```python
def is_palindrome(s):
    left, right = 0, len(s) - 1
    
    while left < right:
        if s[left] != s[right]:
            return False
        left += 1
        right -= 1
    
    return True

# Example usage
print(is_palindrome("racecar"))  # Output: True
print(is_palindrome("hello"))    # Output: False
```

### Pattern 4: Three Sum
**Problem**: Find all unique triplets that sum to zero.

```python
def three_sum(arr):
    arr.sort()
    result = []
    
    for i in range(len(arr) - 2):
        # Skip duplicates for first element
        if i > 0 and arr[i] == arr[i - 1]:
            continue
            
        left, right = i + 1, len(arr) - 1
        
        while left < right:
            current_sum = arr[i] + arr[left] + arr[right]
            
            if current_sum == 0:
                result.append([arr[i], arr[left], arr[right]])
                
                # Skip duplicates
                while left < right and arr[left] == arr[left + 1]:
                    left += 1
                while left < right and arr[right] == arr[right - 1]:
                    right -= 1
                
                left += 1
                right -= 1
            elif current_sum < 0:
                left += 1
            else:
                right -= 1
    
    return result

# Example usage
arr = [-1, 0, 1, 2, -1, -4]
print(three_sum(arr))  # Output: [[-1, -1, 2], [-1, 0, 1]]
```

### Pattern 5: Container With Most Water
**Problem**: Find two lines that together with x-axis forms a container that holds the most water.

```python
def max_area(height):
    left, right = 0, len(height) - 1
    max_water = 0
    
    while left < right:
        # Calculate current area
        width = right - left
        current_area = min(height[left], height[right]) * width
        max_water = max(max_water, current_area)
        
        # Move pointer with smaller height
        if height[left] < height[right]:
            left += 1
        else:
            right -= 1
    
    return max_water

# Example usage
height = [1, 8, 6, 2, 5, 4, 8, 3, 7]
print(max_area(height))  # Output: 49
```

## Practice Problems

### Easy
1. **Two Sum II** - Input array is sorted
2. **Valid Palindrome** - Check if string is palindrome
3. **Remove Duplicates from Sorted Array**
4. **Move Zeros** - Move all zeros to end

### Medium  
1. **3Sum** - Find triplets that sum to zero
2. **Container With Most Water** - Maximum water container
3. **Sort Colors** - Dutch flag problem
4. **Remove Nth Node From End** - LinkedList problem

### Hard
1. **Trapping Rain Water** - Calculate trapped rainwater
2. **Minimum Window Substring** - Find minimum window containing all characters
3. **Longest Palindromic Substring** - Find longest palindrome
4. **4Sum** - Find quadruplets with target sum

## Tips and Tricks

1. **Sort First**: For sum problems, always consider sorting the array first
2. **Skip Duplicates**: When finding unique combinations, skip duplicate elements
3. **Boundary Checks**: Always check array boundaries to avoid index errors
4. **Multiple Pointers**: Some problems may need more than two pointers
5. **Direction Matters**: Choose pointer directions based on the problem requirements

## Common Mistakes

1. **Infinite Loops**: Forgetting to move pointers in certain conditions
2. **Index Bounds**: Not checking if pointers are within array bounds
3. **Duplicate Handling**: Not properly skipping duplicates in sum problems
4. **Early Termination**: Missing early termination conditions

## Related Patterns

- **Sliding Window**: Similar concept but maintains a window of elements
- **Fast & Slow Pointers**: Used for cycle detection in linked lists
- **Binary Search**: Uses two pointers (left/right) to narrow search space

## Implementation Languages

The pattern works similarly across languages:
- **Python**: Use list indexing
- **Java**: Array indexing with int pointers  
- **JavaScript**: Array indexing
- **C++**: Vector/array with iterator or index
- **Go**: Slice indexing