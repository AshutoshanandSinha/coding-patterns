# K-Way Merge Pattern

## Overview
The K-Way Merge pattern is used to merge k sorted arrays, lists, or streams efficiently. This pattern leverages a heap (priority queue) to keep track of the smallest elements from each array, allowing us to merge them in sorted order. It's essential for problems involving multiple sorted data sources.

## When to Use
- **Multiple Sorted Arrays**: Merging k sorted arrays or lists
- **Sorted Streams**: Processing multiple sorted data streams
- **Priority-Based Selection**: Selecting elements based on priority from multiple sources
- **Range Queries**: Finding elements in specific ranges across sorted data
- **Optimization Problems**: Finding optimal solutions across multiple sorted options

## Time/Space Complexity
- **Time**: O(n log k) where n is total elements and k is number of arrays
- **Space**: O(k) for the heap to store one element from each array

## Pattern Structure

### Basic K-Way Merge Template
```python
import heapq

def k_way_merge(lists):
    result = []
    heap = []
    
    # Initialize heap with first element from each list
    for i, lst in enumerate(lists):
        if lst:
            heapq.heappush(heap, (lst[0], i, 0))
    
    while heap:
        value, list_index, element_index = heapq.heappop(heap)
        result.append(value)
        
        # Add next element from the same list
        if element_index + 1 < len(lists[list_index]):
            next_value = lists[list_index][element_index + 1]
            heapq.heappush(heap, (next_value, list_index, element_index + 1))
    
    return result
```

## Common Problem Patterns

### Pattern 1: Merge K Sorted Lists
**Problem**: Merge k sorted linked lists into one sorted linked list.

```python
import heapq

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

def merge_k_lists(lists):
    heap = []
    
    # Add first node from each list to heap
    for i, head in enumerate(lists):
        if head:
            heapq.heappush(heap, (head.val, i, head))
    
    dummy = ListNode(0)
    current = dummy
    
    while heap:
        val, list_index, node = heapq.heappop(heap)
        
        current.next = node
        current = current.next
        
        # Add next node from same list
        if node.next:
            heapq.heappush(heap, (node.next.val, list_index, node.next))
    
    return dummy.next

# Example usage
# lists = [[1,4,5],[1,3,4],[2,6]]
# Result: 1->1->2->3->4->4->5->6
```

### Pattern 2: Smallest Range Covering Elements from K Lists
**Problem**: Find the smallest range that includes at least one element from each of k sorted lists.

```python
import heapq

def smallest_range(nums):
    heap = []
    current_max = float('-inf')
    
    # Initialize heap with first element from each list
    for i, lst in enumerate(nums):
        if lst:
            heapq.heappush(heap, (lst[0], i, 0))
            current_max = max(current_max, lst[0])
    
    range_start, range_end = 0, float('inf')
    
    while heap:
        current_min, list_index, element_index = heapq.heappop(heap)
        
        # Update smallest range if current range is smaller
        if current_max - current_min < range_end - range_start:
            range_start, range_end = current_min, current_max
        
        # Move to next element in the same list
        if element_index + 1 < len(nums[list_index]):
            next_val = nums[list_index][element_index + 1]
            heapq.heappush(heap, (next_val, list_index, element_index + 1))
            current_max = max(current_max, next_val)
        else:
            # If we can't advance any list, we're done
            break
    
    return [range_start, range_end]

# Example usage
nums = [[4,10,15,24,26], [0,9,12,20], [5,18,22,30]]
print(smallest_range(nums))  # Output: [20, 24]
```

### Pattern 3: K Pairs with Smallest Sums
**Problem**: Find k pairs with smallest sums from two sorted arrays.

```python
import heapq

def k_smallest_pairs(nums1, nums2, k):
    if not nums1 or not nums2 or k == 0:
        return []
    
    heap = []
    result = []
    
    # Initialize heap with pairs from first element of nums1
    for i in range(min(k, len(nums1))):
        heapq.heappush(heap, (nums1[i] + nums2[0], i, 0))
    
    while heap and len(result) < k:
        current_sum, i, j = heapq.heappop(heap)
        result.append([nums1[i], nums2[j]])
        
        # Add next pair from same row if available
        if j + 1 < len(nums2):
            heapq.heappush(heap, (nums1[i] + nums2[j + 1], i, j + 1))
    
    return result

# Example usage
nums1 = [1, 7, 11]
nums2 = [2, 4, 6]
k = 3
print(k_smallest_pairs(nums1, nums2, k))  # Output: [[1,2],[1,4],[1,6]]
```

### Pattern 4: Merge K Sorted Arrays
**Problem**: Merge k sorted arrays into one sorted array.

```python
import heapq

def merge_k_sorted_arrays(arrays):
    result = []
    heap = []
    
    # Initialize heap with first element from each array
    for i, arr in enumerate(arrays):
        if arr:
            heapq.heappush(heap, (arr[0], i, 0))
    
    while heap:
        value, array_index, element_index = heapq.heappop(heap)
        result.append(value)
        
        # Add next element from the same array
        if element_index + 1 < len(arrays[array_index]):
            next_value = arrays[array_index][element_index + 1]
            heapq.heappush(heap, (next_value, array_index, element_index + 1))
    
    return result

# Example usage
arrays = [[1, 4, 5], [1, 3, 4], [2, 6]]
print(merge_k_sorted_arrays(arrays))  # Output: [1, 1, 2, 3, 4, 4, 5, 6]
```

### Pattern 5: Kth Smallest Element in Sorted Matrix
**Problem**: Find the kth smallest element in a sorted matrix.

```python
import heapq

def kth_smallest_in_matrix(matrix, k):
    n = len(matrix)
    heap = []
    
    # Add first element of each row
    for i in range(min(k, n)):
        heapq.heappush(heap, (matrix[i][0], i, 0))
    
    count = 0
    while heap:
        value, row, col = heapq.heappop(heap)
        count += 1
        
        if count == k:
            return value
        
        # Add next element from the same row
        if col + 1 < len(matrix[row]):
            heapq.heappush(heap, (matrix[row][col + 1], row, col + 1))
    
    return -1

# Example usage
matrix = [
    [1,  5,  9],
    [10, 11, 13],
    [12, 13, 15]
]
k = 8
print(kth_smallest_in_matrix(matrix, k))  # Output: 13
```

### Pattern 6: Find K Closest Elements
**Problem**: Find k closest elements to a target in a sorted array (using merge concept).

```python
import heapq

def find_closest_elements(arr, k, x):
    # Create two pointers approach combined with heap concept
    left = 0
    right = len(arr) - k
    
    # Binary search for the best starting position
    while left < right:
        mid = left + (right - left) // 2
        
        # Compare distances from x to arr[mid] and arr[mid + k]
        if x - arr[mid] > arr[mid + k] - x:
            left = mid + 1
        else:
            right = mid
    
    return arr[left:left + k]

# Alternative heap-based approach
def find_closest_elements_heap(arr, k, x):
    heap = []
    
    for num in arr:
        distance = abs(num - x)
        if len(heap) < k:
            heapq.heappush(heap, (-distance, -num))
        elif distance < -heap[0][0]:
            heapq.heappop(heap)
            heapq.heappush(heap, (-distance, -num))
    
    result = [-item[1] for item in heap]
    return sorted(result)

# Example usage
arr = [1, 2, 3, 4, 5]
k = 4
x = 3
print(find_closest_elements(arr, k, x))  # Output: [1, 2, 3, 4]
```

### Pattern 7: Merge Intervals from K Sorted Lists
**Problem**: Merge overlapping intervals from k sorted interval lists.

```python
import heapq

def merge_k_interval_lists(interval_lists):
    heap = []
    
    # Initialize heap with first interval from each list
    for i, intervals in enumerate(interval_lists):
        if intervals:
            heapq.heappush(heap, (intervals[0][0], i, 0, intervals[0]))
    
    merged = []
    
    while heap:
        start, list_idx, interval_idx, interval = heapq.heappop(heap)
        
        # Merge with previous interval if overlapping
        if merged and start <= merged[-1][1]:
            merged[-1][1] = max(merged[-1][1], interval[1])
        else:
            merged.append(interval)
        
        # Add next interval from the same list
        if interval_idx + 1 < len(interval_lists[list_idx]):
            next_interval = interval_lists[list_idx][interval_idx + 1]
            heapq.heappush(heap, (next_interval[0], list_idx, interval_idx + 1, next_interval))
    
    return merged

# Example usage
interval_lists = [
    [[1, 3], [5, 8]],
    [[2, 4], [6, 8]],
    [[3, 7]]
]
print(merge_k_interval_lists(interval_lists))  # Output: [[1, 8]]
```

## Practice Problems

### Easy
1. **Merge Two Sorted Lists** - Basic merge operation
2. **Merge Sorted Array** - In-place merge
3. **Intersection of Two Arrays** - Find common elements

### Medium
1. **Merge k Sorted Lists** - Classic k-way merge
2. **K Pairs with Smallest Sums** - Heap-based selection
3. **Find K Closest Elements** - Distance-based selection
4. **Smallest Range Covering Elements** - Range optimization

### Hard
1. **Merge k Sorted Arrays** - Multiple array merge
2. **Kth Smallest in Sorted Matrix** - 2D matrix navigation
3. **Smallest Range in K Lists** - Complex range finding
4. **Super Ugly Number** - Multiple factor streams

## Tips and Tricks

1. **Heap Initialization**: Always initialize heap with first elements from each source
2. **Index Tracking**: Keep track of current position in each array/list
3. **Boundary Checks**: Always check if next element exists before adding to heap
4. **Custom Comparisons**: Use tuples for complex sorting criteria
5. **Memory Efficiency**: Only store necessary information in heap

## Common Mistakes

1. **Heap Overflow**: Not checking if arrays are empty before accessing elements
2. **Index Errors**: Going out of bounds when accessing next elements
3. **Infinite Loops**: Not properly advancing through arrays
4. **Memory Issues**: Storing too much information in heap
5. **Comparison Errors**: Incorrect tuple ordering for heap comparisons

## Related Patterns

- **Two Pointers**: Used within individual array processing
- **Binary Search**: For optimization in some k-way merge problems
- **Divide and Conquer**: Merge sort is a simpler version of k-way merge
- **Priority Queue**: Fundamental data structure for this pattern

## Implementation Languages

The pattern adapts well across languages:
- **Python**: Use `heapq` module for heap operations
- **Java**: Use `PriorityQueue` class
- **JavaScript**: Implement custom heap or use libraries
- **C++**: Use `priority_queue` from STL
- **Go**: Use `container/heap` package