# Sliding Window Pattern

## Overview
The Sliding Window pattern is used to find a subarray or substring in an array or string that satisfies a specific condition. Instead of using nested loops (brute force), we maintain a 'window' of elements and slide it through the data structure, optimizing time complexity from O(n²) to O(n).

## When to Use
- **Subarray/Substring Problems**: Finding subarrays with specific properties
- **Fixed Window Size**: Problems asking for max/min in every window of size k
- **Variable Window Size**: Problems where window size changes based on conditions
- **Optimization Problems**: Finding optimal subarray (longest, shortest, maximum sum)
- **String Problems**: Pattern matching, anagrams, character frequency

## Time/Space Complexity
- **Time**: O(n) - Each element visited at most twice
- **Space**: O(1) for fixed window, O(k) for character frequency tracking

## Pattern Variations

### 1. Fixed Window Size
```python
def fixed_window(arr, k):
    if len(arr) < k:
        return []
    
    window_sum = sum(arr[:k])  # Initial window
    max_sum = window_sum
    
    # Slide the window
    for i in range(k, len(arr)):
        window_sum += arr[i] - arr[i - k]  # Add new, remove old
        max_sum = max(max_sum, window_sum)
    
    return max_sum
```

### 2. Variable Window Size
```python
def variable_window(arr, target):
    left = 0
    window_sum = 0
    min_length = float('inf')
    
    for right in range(len(arr)):
        window_sum += arr[right]
        
        # Shrink window while condition is met
        while window_sum >= target:
            min_length = min(min_length, right - left + 1)
            window_sum -= arr[left]
            left += 1
    
    return min_length if min_length != float('inf') else 0
```

## Common Problem Patterns

### Pattern 1: Maximum Sum Subarray of Size K
**Problem**: Find the maximum sum of any contiguous subarray of size k.

```python
def max_sum_subarray(arr, k):
    if len(arr) < k:
        return -1
    
    # Calculate sum of first window
    window_sum = sum(arr[:k])
    max_sum = window_sum
    
    # Slide the window
    for i in range(k, len(arr)):
        window_sum += arr[i] - arr[i - k]
        max_sum = max(max_sum, window_sum)
    
    return max_sum

# Example usage
arr = [2, 1, 5, 1, 3, 2]
k = 3
print(max_sum_subarray(arr, k))  # Output: 9 (5+1+3)
```

### Pattern 2: Longest Substring Without Repeating Characters
**Problem**: Find the length of the longest substring without repeating characters.

```python
def longest_substring_without_repeating(s):
    char_map = {}
    left = 0
    max_length = 0
    
    for right in range(len(s)):
        # If character already in window, move left pointer
        if s[right] in char_map and char_map[s[right]] >= left:
            left = char_map[s[right]] + 1
        
        char_map[s[right]] = right
        max_length = max(max_length, right - left + 1)
    
    return max_length

# Example usage
s = "abcabcbb"
print(longest_substring_without_repeating(s))  # Output: 3 ("abc")
```

### Pattern 3: Minimum Window Substring
**Problem**: Find the minimum window in string s that contains all characters of string t.

```python
def min_window_substring(s, t):
    if not s or not t:
        return ""
    
    # Character frequency in t
    required = {}
    for char in t:
        required[char] = required.get(char, 0) + 1
    
    left = right = 0
    formed = 0  # Number of unique chars in current window with desired frequency
    window_counts = {}
    
    # (window length, left, right)
    ans = float('inf'), None, None
    
    while right < len(s):
        # Add character from right to window
        char = s[right]
        window_counts[char] = window_counts.get(char, 0) + 1
        
        # Check if current character's frequency matches required frequency
        if char in required and window_counts[char] == required[char]:
            formed += 1
        
        # Try to contract window until it ceases to be 'desirable'
        while left <= right and formed == len(required):
            char = s[left]
            
            # Save the smallest window
            if right - left + 1 < ans[0]:
                ans = (right - left + 1, left, right)
            
            # Remove from left of window
            window_counts[char] -= 1
            if char in required and window_counts[char] < required[char]:
                formed -= 1
            
            left += 1
        
        right += 1
    
    return "" if ans[0] == float('inf') else s[ans[1]:ans[2] + 1]

# Example usage
s = "ADOBECODEBANC"
t = "ABC"
print(min_window_substring(s, t))  # Output: "BANC"
```

### Pattern 4: Longest Substring with At Most K Distinct Characters
**Problem**: Find the length of the longest substring with at most k distinct characters.

```python
def longest_substring_k_distinct(s, k):
    if k == 0:
        return 0
    
    char_frequency = {}
    left = 0
    max_length = 0
    
    for right in range(len(s)):
        # Add character to frequency map
        char_frequency[s[right]] = char_frequency.get(s[right], 0) + 1
        
        # Shrink window if we have more than k distinct characters
        while len(char_frequency) > k:
            char_frequency[s[left]] -= 1
            if char_frequency[s[left]] == 0:
                del char_frequency[s[left]]
            left += 1
        
        max_length = max(max_length, right - left + 1)
    
    return max_length

# Example usage
s = "araaci"
k = 2
print(longest_substring_k_distinct(s, k))  # Output: 4 ("araa")
```

### Pattern 5: Subarray Product Less Than K
**Problem**: Count number of subarrays where product of elements is less than k.

```python
def subarray_product_less_than_k(nums, k):
    if k <= 1:
        return 0
    
    left = 0
    product = 1
    count = 0
    
    for right in range(len(nums)):
        product *= nums[right]
        
        # Shrink window while product >= k
        while product >= k:
            product //= nums[left]
            left += 1
        
        # Add all subarrays ending at right
        count += right - left + 1
    
    return count

# Example usage
nums = [10, 5, 2, 6]
k = 100
print(subarray_product_less_than_k(nums, k))  # Output: 8
```

### Pattern 6: Fruits Into Baskets (Max 2 Types)
**Problem**: Pick maximum fruits with at most 2 different types.

```python
def total_fruits(fruits):
    fruit_frequency = {}
    left = 0
    max_fruits = 0
    
    for right in range(len(fruits)):
        # Add fruit to basket
        fruit_frequency[fruits[right]] = fruit_frequency.get(fruits[right], 0) + 1
        
        # Shrink window if more than 2 fruit types
        while len(fruit_frequency) > 2:
            fruit_frequency[fruits[left]] -= 1
            if fruit_frequency[fruits[left]] == 0:
                del fruit_frequency[fruits[left]]
            left += 1
        
        max_fruits = max(max_fruits, right - left + 1)
    
    return max_fruits

# Example usage
fruits = [1, 2, 1, 2, 3, 2, 2]
print(total_fruits(fruits))  # Output: 5 ([2, 1, 2, 3, 2] -> [2, 1, 2, 2])
```

## Practice Problems

### Easy
1. **Maximum Sum Subarray of Size K** - Find max sum of k-sized window
2. **Average of All Subarrays of Size K** - Calculate averages
3. **Smallest Subarray with Sum >= S** - Find minimum length subarray

### Medium
1. **Longest Substring Without Repeating Characters** - No duplicate chars
2. **Longest Substring with K Distinct Characters** - At most k distinct
3. **String Anagrams** - Find all anagram substrings
4. **Subarray Product Less Than K** - Count valid subarrays

### Hard
1. **Minimum Window Substring** - Smallest window containing all characters
2. **Substring with Concatenation of All Words** - Complex pattern matching
3. **Longest Substring with Same Letters after Replacement** - With k replacements
4. **Shortest Subarray with Sum at Least K** - Using deque optimization

## Tips and Tricks

1. **Expand First, Then Contract**: Generally expand right pointer first, then contract left
2. **Use Hash Maps**: For character/element frequency tracking
3. **Multiple Conditions**: Handle multiple shrinking conditions carefully
4. **Edge Cases**: Empty arrays, single elements, impossible conditions
5. **Template Approach**: Use consistent template for different variations

## Common Sliding Window Template

```python
def sliding_window_template(s, pattern):
    # Initialize frequency map for pattern
    freq_map = {}
    for char in pattern:
        freq_map[char] = freq_map.get(char, 0) + 1
    
    left = 0
    matched = 0
    result = []
    
    for right in range(len(s)):
        right_char = s[right]
        
        # Expand window
        if right_char in freq_map:
            freq_map[right_char] -= 1
            if freq_map[right_char] == 0:
                matched += 1
        
        # Check if window condition is met
        while matched == len(freq_map):
            # Process current window
            result.append(left)  # or other processing
            
            # Contract window
            left_char = s[left]
            if left_char in freq_map:
                if freq_map[left_char] == 0:
                    matched -= 1
                freq_map[left_char] += 1
            left += 1
    
    return result
```

## Common Mistakes

1. **Infinite Loops**: Not moving pointers properly
2. **Off-by-One Errors**: Incorrect window size calculations
3. **Hash Map Cleanup**: Not removing zero-frequency entries
4. **Edge Cases**: Not handling empty inputs or impossible conditions
5. **Overflow**: Not considering integer overflow in product problems

## Related Patterns

- **Two Pointers**: Similar window concept but different movement
- **Fast & Slow Pointers**: Different pointer speeds
- **Dynamic Programming**: Some sliding window problems can be DP alternatives

## Implementation Notes

- **Python**: Use dictionaries for frequency maps
- **Java**: HashMap for frequency, StringBuilder for string operations
- **JavaScript**: Map or Object for frequency tracking
- **C++**: unordered_map for frequency, careful with iterators