# Top K Elements Pattern

## Overview
The Top K Elements pattern uses heaps (priority queues) to efficiently find the k largest, k smallest, or k most frequent elements in a dataset. This pattern is fundamental for problems involving ranking, selection, and optimization where you need to maintain a collection of the "best" k elements based on some criteria.

## When to Use
- **Ranking Problems**: Finding top k largest/smallest elements
- **Frequency Analysis**: K most/least frequent elements
- **Streaming Data**: Maintaining top k in data streams
- **Selection Problems**: Selecting k elements based on criteria
- **Optimization**: Finding k best solutions

## Time/Space Complexity
- **Time**: O(n log k) for most operations
- **Space**: O(k) for maintaining heap of k elements

## Heap Types
- **Min Heap**: Smallest element at root (use for k largest elements)
- **Max Heap**: Largest element at root (use for k smallest elements)

## Pattern Variations

### 1. K Largest Elements (using Min Heap)
```python
import heapq

def find_k_largest(nums, k):
    heap = []
    
    for num in nums:
        if len(heap) < k:
            heapq.heappush(heap, num)
        elif num > heap[0]:
            heapq.heappop(heap)
            heapq.heappush(heap, num)
    
    return list(heap)
```

### 2. K Smallest Elements (using Max Heap)
```python
import heapq

def find_k_smallest(nums, k):
    heap = []
    
    for num in nums:
        if len(heap) < k:
            heapq.heappush(heap, -num)  # Negate for max heap
        elif num < -heap[0]:
            heapq.heappop(heap)
            heapq.heappush(heap, -num)
    
    return [-x for x in heap]
```

## Common Problem Patterns

### Pattern 1: Kth Largest Element
**Problem**: Find the kth largest element in an unsorted array.

```python
import heapq

def find_kth_largest(nums, k):
    # Use min heap of size k
    heap = []
    
    for num in nums:
        if len(heap) < k:
            heapq.heappush(heap, num)
        elif num > heap[0]:
            heapq.heappop(heap)
            heapq.heappush(heap, num)
    
    return heap[0]

# Alternative: Using all elements then popping
def find_kth_largest_v2(nums, k):
    heapq.heapify(nums)  # Convert to min heap
    
    # Remove n-k smallest elements
    for _ in range(len(nums) - k):
        heapq.heappop(nums)
    
    return nums[0]

# Example usage
nums = [3, 2, 1, 5, 6, 4]
k = 2
print(find_kth_largest(nums, k))  # Output: 5
```

### Pattern 2: K Closest Points to Origin
**Problem**: Find k closest points to origin in 2D plane.

```python
import heapq

def k_closest_points(points, k):
    # Use max heap to keep k closest points
    heap = []
    
    for point in points:
        x, y = point
        distance = x*x + y*y  # Squared distance (no need for sqrt)
        
        if len(heap) < k:
            heapq.heappush(heap, (-distance, point))
        elif distance < -heap[0][0]:
            heapq.heappop(heap)
            heapq.heappush(heap, (-distance, point))
    
    return [point for _, point in heap]

# Example usage
points = [[1,1],[2,2],[3,3]]
k = 2
print(k_closest_points(points, k))  # Output: [[1,1],[2,2]]
```

### Pattern 3: Top K Frequent Elements
**Problem**: Find k most frequent elements in array.

```python
import heapq
from collections import Counter

def top_k_frequent(nums, k):
    # Count frequencies
    count = Counter(nums)
    
    # Use min heap of size k (frequency, element)
    heap = []
    
    for num, freq in count.items():
        if len(heap) < k:
            heapq.heappush(heap, (freq, num))
        elif freq > heap[0][0]:
            heapq.heappop(heap)
            heapq.heappush(heap, (freq, num))
    
    return [num for freq, num in heap]

# Alternative: Using max heap with all elements
def top_k_frequent_v2(nums, k):
    count = Counter(nums)
    
    # Create max heap with all elements
    heap = [(-freq, num) for num, freq in count.items()]
    heapq.heapify(heap)
    
    result = []
    for _ in range(k):
        freq, num = heapq.heappop(heap)
        result.append(num)
    
    return result

# Example usage
nums = [1,1,1,2,2,3]
k = 2
print(top_k_frequent(nums, k))  # Output: [1, 2]
```

### Pattern 4: Kth Smallest Element in Sorted Matrix
**Problem**: Find kth smallest element in row and column wise sorted matrix.

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
        val, row, col = heapq.heappop(heap)
        count += 1
        
        if count == k:
            return val
        
        # Add next element from same row if exists
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

### Pattern 5: K Closest Elements
**Problem**: Find k elements closest to target value in sorted array.

```python
import heapq

def find_closest_elements(arr, k, x):
    # Max heap to store k closest elements
    heap = []
    
    for num in arr:
        distance = abs(num - x)
        
        if len(heap) < k:
            heapq.heappush(heap, (-distance, -num))
        elif distance < -heap[0][0]:
            heapq.heappop(heap)
            heapq.heappush(heap, (-distance, -num))
    
    result = sorted([-item[1] for item in heap])
    return result

# More efficient approach using two pointers
def find_closest_elements_v2(arr, k, x):
    left = 0
    right = len(arr) - k
    
    while left < right:
        mid = left + (right - left) // 2
        if x - arr[mid] > arr[mid + k] - x:
            left = mid + 1
        else:
            right = mid
    
    return arr[left:left + k]

# Example usage
arr = [1, 2, 3, 4, 5]
k = 4
x = 3
print(find_closest_elements(arr, k, x))  # Output: [1, 2, 3, 4]
```

### Pattern 6: Reorganize String
**Problem**: Reorganize string so no two adjacent characters are same.

```python
import heapq
from collections import Counter

def reorganize_string(s):
    count = Counter(s)
    
    # Create max heap based on frequency
    heap = [(-freq, char) for char, freq in count.items()]
    heapq.heapify(heap)
    
    result = []
    prev_freq, prev_char = 0, ''
    
    while heap:
        # Get most frequent character
        freq, char = heapq.heappop(heap)
        result.append(char)
        
        # Add back previous character if it still has frequency
        if prev_freq < 0:
            heapq.heappush(heap, (prev_freq, prev_char))
        
        # Update previous character info
        prev_freq, prev_char = freq + 1, char
    
    result_str = ''.join(result)
    return result_str if len(result_str) == len(s) else ""

# Example usage
s = "aab"
print(reorganize_string(s))  # Output: "aba"
```

### Pattern 7: Find K Pairs with Smallest Sums
**Problem**: Find k pairs with smallest sums from two sorted arrays.

```python
import heapq

def k_smallest_pairs(nums1, nums2, k):
    if not nums1 or not nums2:
        return []
    
    heap = []
    result = []
    
    # Initialize heap with pairs involving first element of nums1
    for i in range(min(k, len(nums1))):
        heapq.heappush(heap, (nums1[i] + nums2[0], i, 0))
    
    while heap and len(result) < k:
        sum_val, i, j = heapq.heappop(heap)
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

### Pattern 8: Task Scheduler
**Problem**: Schedule tasks with cooling period to minimize total time.

```python
import heapq
from collections import Counter

def least_interval(tasks, n):
    count = Counter(tasks)
    
    # Max heap based on frequency
    heap = [-freq for freq in count.values()]
    heapq.heapify(heap)
    
    time = 0
    
    while heap:
        cycle_tasks = []
        
        # Try to schedule n+1 tasks in current cycle
        for _ in range(n + 1):
            if heap:
                freq = heapq.heappop(heap)
                if freq < -1:  # Task still has remaining frequency
                    cycle_tasks.append(freq + 1)
        
        # Add back tasks that still need to be scheduled
        for freq in cycle_tasks:
            heapq.heappush(heap, freq)
        
        # Update time
        if heap:
            time += n + 1  # Full cycle used
        else:
            time += len(cycle_tasks)  # Only partial cycle needed
    
    return time

# Example usage
tasks = ["A","A","A","B","B","B"]
n = 2
print(least_interval(tasks, n))  # Output: 8
```

## Practice Problems

### Easy
1. **Kth Largest Element in Array** - Basic heap usage
2. **Last Stone Weight** - Max heap simulation
3. **K Closest Points to Origin** - Distance-based selection

### Medium
1. **Top K Frequent Elements** - Frequency analysis
2. **Kth Smallest Element in Sorted Matrix** - 2D array navigation
3. **Find K Closest Elements** - Target-based selection
4. **Reorganize String** - Greedy with heap
5. **K Pairs with Smallest Sums** - Two array combination

### Hard
1. **Merge k Sorted Lists** - Complex heap operations
2. **Task Scheduler** - Optimization with constraints
3. **Sliding Window Maximum** - Deque alternative to heap
4. **IPO (Maximum Capital)** - Dual heap problem

## Tips and Tricks

1. **Heap Size**: Keep heap size at k for efficiency
2. **Min vs Max Heap**: Choose based on what you're finding (k largest uses min heap)
3. **Custom Comparisons**: Use tuples for complex sorting criteria
4. **Frequency Maps**: Combine with Counter for frequency-based problems
5. **Streaming Data**: Heap is perfect for maintaining top k in streams

## Common Mistakes

1. **Wrong Heap Type**: Using max heap when min heap is needed
2. **Heap Size**: Not maintaining optimal heap size
3. **Comparison Keys**: Incorrect tuple ordering for complex comparisons
4. **Memory Issues**: Storing too much data in heap
5. **Edge Cases**: Not handling k larger than array size

## Related Patterns

- **Sliding Window**: Sometimes combined with heap for window maximum/minimum
- **Two Pointers**: Alternative approach for some closest element problems
- **Quick Select**: Alternative O(n) average case for kth element
- **Priority Queue**: Heap is implementation of priority queue

## Implementation Languages

Heap operations are well-supported across languages:
- **Python**: Use `heapq` module (min heap only, negate for max heap)
- **Java**: Use `PriorityQueue` class with custom comparators
- **JavaScript**: Implement custom heap or use libraries
- **C++**: Use `priority_queue` from STL (max heap by default)
- **Go**: Use `container/heap` package with interface implementation