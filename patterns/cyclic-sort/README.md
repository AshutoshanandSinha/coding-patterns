# Cyclic Sort Pattern

## Overview
Cyclic Sort is an in-place sorting algorithm that works by placing each element at its correct position based on its value. This pattern is particularly useful when dealing with arrays containing numbers in a given range, typically from 1 to n or 0 to n-1. It's extremely efficient for finding missing numbers, duplicates, or the first missing positive integer.

## When to Use
- **Range-Bound Numbers**: Arrays with numbers in range 1 to n or 0 to n-1
- **Find Missing Numbers**: Identify missing elements in sequences
- **Find Duplicates**: Detect duplicate numbers efficiently
- **In-Place Operations**: When you need O(1) extra space
- **Unsorted Input**: When input array is unsorted but has known range

## Time/Space Complexity
- **Time**: O(n) - Each element is moved at most once to its correct position
- **Space**: O(1) - In-place sorting algorithm

## Pattern Variations

### 1. Basic Cyclic Sort (1 to n)
```python
def cyclic_sort(nums):
    i = 0
    while i < len(nums):
        correct_index = nums[i] - 1  # For numbers 1 to n
        
        if nums[i] != nums[correct_index]:
            nums[i], nums[correct_index] = nums[correct_index], nums[i]
        else:
            i += 1
    
    return nums
```

### 2. Cyclic Sort (0 to n-1)
```python
def cyclic_sort_zero_based(nums):
    i = 0
    while i < len(nums):
        correct_index = nums[i]  # For numbers 0 to n-1
        
        if nums[i] != nums[correct_index]:
            nums[i], nums[correct_index] = nums[correct_index], nums[i]
        else:
            i += 1
    
    return nums
```

## Common Problem Patterns

### Pattern 1: Find Missing Number
**Problem**: Find the missing number in an array containing n distinct numbers from 0 to n.

```python
def missing_number(nums):
    i = 0
    n = len(nums)
    
    # Cyclic sort: place each number at its correct index
    while i < n:
        correct_index = nums[i]
        
        # If number is in valid range and not in correct position
        if 0 <= nums[i] < n and nums[i] != nums[correct_index]:
            nums[i], nums[correct_index] = nums[correct_index], nums[i]
        else:
            i += 1
    
    # Find the missing number
    for i in range(n):
        if nums[i] != i:
            return i
    
    return n  # If all numbers 0 to n-1 are present, missing number is n

# Example usage
nums = [4, 0, 3, 1]
print(missing_number(nums))  # Output: 2
```

### Pattern 2: Find All Missing Numbers
**Problem**: Find all missing numbers in an array containing numbers from 1 to n.

```python
def find_missing_numbers(nums):
    i = 0
    
    # Cyclic sort: place each number at index (number - 1)
    while i < len(nums):
        correct_index = nums[i] - 1
        
        if 1 <= nums[i] <= len(nums) and nums[i] != nums[correct_index]:
            nums[i], nums[correct_index] = nums[correct_index], nums[i]
        else:
            i += 1
    
    # Find missing numbers
    missing = []
    for i in range(len(nums)):
        if nums[i] != i + 1:
            missing.append(i + 1)
    
    return missing

# Example usage
nums = [2, 3, 1, 8, 2, 3, 5, 1]
print(find_missing_numbers(nums))  # Output: [4, 6, 7]
```

### Pattern 3: Find Duplicate Number
**Problem**: Find the duplicate number in an array containing numbers from 1 to n.

```python
def find_duplicate(nums):
    i = 0
    
    while i < len(nums):
        correct_index = nums[i] - 1
        
        if nums[i] != nums[correct_index]:
            nums[i], nums[correct_index] = nums[correct_index], nums[i]
        else:
            i += 1
    
    # Find the duplicate
    for i in range(len(nums)):
        if nums[i] != i + 1:
            return nums[i]
    
    return -1

# Example usage
nums = [1, 4, 4, 3, 2]
print(find_duplicate(nums))  # Output: 4
```

### Pattern 4: Find All Duplicates
**Problem**: Find all duplicates in an array containing numbers from 1 to n.

```python
def find_duplicates(nums):
    i = 0
    
    # Cyclic sort
    while i < len(nums):
        correct_index = nums[i] - 1
        
        if nums[i] != nums[correct_index]:
            nums[i], nums[correct_index] = nums[correct_index], nums[i]
        else:
            i += 1
    
    # Find duplicates
    duplicates = []
    for i in range(len(nums)):
        if nums[i] != i + 1:
            duplicates.append(nums[i])
    
    return duplicates

# Example usage
nums = [4, 3, 2, 7, 8, 2, 3, 1]
print(find_duplicates(nums))  # Output: [2, 3]
```

### Pattern 5: Find Corrupt Pair
**Problem**: Find the corrupt pair (missing and duplicate numbers) in an array.

```python
def find_error_nums(nums):
    i = 0
    
    # Cyclic sort
    while i < len(nums):
        correct_index = nums[i] - 1
        
        if nums[i] != nums[correct_index]:
            nums[i], nums[correct_index] = nums[correct_index], nums[i]
        else:
            i += 1
    
    # Find the corrupt pair
    for i in range(len(nums)):
        if nums[i] != i + 1:
            return [nums[i], i + 1]  # [duplicate, missing]
    
    return [-1, -1]

# Example usage
nums = [1, 2, 2, 4]
print(find_error_nums(nums))  # Output: [2, 3] (2 is duplicate, 3 is missing)
```

### Pattern 6: First Missing Positive
**Problem**: Find the first missing positive integer in an unsorted array.

```python
def first_missing_positive(nums):
    n = len(nums)
    i = 0
    
    # Cyclic sort: place positive numbers 1 to n at correct positions
    while i < n:
        correct_index = nums[i] - 1
        
        # Only sort numbers in range [1, n] and not already in correct position
        if 1 <= nums[i] <= n and nums[i] != nums[correct_index]:
            nums[i], nums[correct_index] = nums[correct_index], nums[i]
        else:
            i += 1
    
    # Find first missing positive
    for i in range(n):
        if nums[i] != i + 1:
            return i + 1
    
    return n + 1  # All numbers 1 to n are present

# Example usage
nums = [3, 4, -1, 1]
print(first_missing_positive(nums))  # Output: 2
```

### Pattern 7: Find Smallest Missing Positive (K Missing)
**Problem**: Find the kth missing positive number.

```python
def find_kth_positive(arr, k):
    # First, let's find what positive numbers are missing
    i = 0
    n = len(arr)
    
    # Cyclic sort for positive numbers only
    while i < n:
        correct_index = arr[i] - 1
        
        if 0 < arr[i] <= n and arr[i] != arr[correct_index]:
            arr[i], arr[correct_index] = arr[correct_index], arr[i]
        else:
            i += 1
    
    # Count missing numbers and find the kth one
    missing_count = 0
    for i in range(n):
        if arr[i] != i + 1:
            missing_count += 1
            if missing_count == k:
                return i + 1
    
    # If we haven't found k missing numbers in range [1, n]
    # The answer is n + (k - missing_count)
    return n + (k - missing_count)

# Example usage
arr = [2, 3, 4, 7, 11]
k = 5
print(find_kth_positive(arr, k))  # Output: 9
```

## Practice Problems

### Easy
1. **Missing Number** - Find missing number in 0 to n
2. **Find All Numbers Disappeared in Array** - Missing numbers 1 to n
3. **Couples Holding Hands** - Arrange couples optimally

### Medium
1. **Find Duplicate Number** - Single duplicate in 1 to n
2. **Find All Duplicates in Array** - All duplicates in 1 to n
3. **Set Mismatch** - Find duplicate and missing number
4. **First Missing Positive** - First missing positive integer

### Hard
1. **Find the Duplicate Number** - Without modifying array (Floyd's algorithm)
2. **Missing Element in Sorted Array** - Binary search variation
3. **Smallest Missing Positive** - Complex constraints

## Tips and Tricks

1. **Range Check**: Always verify numbers are in expected range before sorting
2. **Index Calculation**: Be careful with 0-based vs 1-based indexing
3. **Boundary Conditions**: Handle edge cases like empty arrays
4. **Multiple Passes**: Sometimes need to iterate multiple times for complex problems
5. **Space Efficiency**: Remember this is an in-place algorithm

## Common Mistakes

1. **Index Out of Bounds**: Not checking if numbers are in valid range
2. **Infinite Loops**: Not handling duplicate numbers correctly
3. **Off-by-One Errors**: Confusion between 0-based and 1-based indexing
4. **Modifying vs Non-Modifying**: Some problems require preserving original array

## Related Patterns

- **Hash Tables**: Alternative approach for finding missing/duplicate elements
- **Binary Search**: Can be used for some missing number problems
- **Floyd's Cycle Detection**: Alternative for finding duplicates
- **Bit Manipulation**: XOR can solve some missing number problems

## Implementation Languages

The pattern works across all languages:
- **Python**: Use tuple unpacking for swaps: `a, b = b, a`
- **Java**: Traditional three-step swap or use helper method
- **JavaScript**: Array destructuring: `[a, b] = [b, a]`
- **C++**: Use `std::swap()` for cleaner code
- **Go**: Multiple assignment: `a, b = b, a`