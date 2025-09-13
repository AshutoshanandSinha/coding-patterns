# Two Heaps Pattern

## Overview
The Two Heaps pattern uses a combination of a min heap and a max heap to efficiently track and access the median or middle elements in a dataset. This pattern is particularly useful for problems involving streaming data, sliding window medians, and scenarios where you need to maintain balance between two halves of data.

## When to Use
- **Median Finding**: Finding median in streaming data or arrays
- **Sliding Window Median**: Median in sliding windows
- **Balance Maintenance**: Keeping two halves of data balanced
- **Priority-based Splitting**: Dividing data into two priority groups
- **Range Queries**: Problems requiring access to middle elements

## Time/Space Complexity
- **Time**: O(log n) for insertions, O(1) for median access
- **Space**: O(n) to store all elements in heaps

## Core Concept
- **Max Heap (Left)**: Stores smaller half of elements
- **Min Heap (Right)**: Stores larger half of elements
- **Balance**: Maintain size difference ≤ 1 between heaps
- **Median**: Root of larger heap (if unequal) or average of both roots

## Pattern Structure

### Basic Two Heaps Template
```python
import heapq

class MedianFinder:
    def __init__(self):
        self.max_heap = []  # Left half (smaller elements)
        self.min_heap = []  # Right half (larger elements)
    
    def add_num(self, num):
        # Add to max_heap first
        heapq.heappush(self.max_heap, -num)
        
        # Move largest from max_heap to min_heap
        heapq.heappush(self.min_heap, -heapq.heappop(self.max_heap))
        
        # Rebalance if min_heap is larger
        if len(self.min_heap) > len(self.max_heap):
            heapq.heappush(self.max_heap, -heapq.heappop(self.min_heap))
    
    def find_median(self):
        if len(self.max_heap) > len(self.min_heap):
            return -self.max_heap[0]
        else:
            return (-self.max_heap[0] + self.min_heap[0]) / 2.0
```

## Common Problem Patterns

### Pattern 1: Find Median from Data Stream
**Problem**: Design a data structure to find median from a stream of integers.

```python
import heapq

class MedianFinder:
    def __init__(self):
        self.small = []  # Max heap (use negative values)
        self.large = []  # Min heap
    
    def addNum(self, num):
        # Always add to small first
        heapq.heappush(self.small, -num)
        
        # Ensure all elements in small <= all elements in large
        if self.small and self.large and (-self.small[0] > self.large[0]):
            val = -heapq.heappop(self.small)
            heapq.heappush(self.large, val)
        
        # Balance the heaps
        if len(self.small) > len(self.large) + 1:
            val = -heapq.heappop(self.small)
            heapq.heappush(self.large, val)
        
        if len(self.large) > len(self.small) + 1:
            val = heapq.heappop(self.large)
            heapq.heappush(self.small, -val)
    
    def findMedian(self):
        if len(self.small) > len(self.large):
            return -self.small[0]
        elif len(self.large) > len(self.small):
            return self.large[0]
        else:
            return (-self.small[0] + self.large[0]) / 2.0

# Example usage
mf = MedianFinder()
mf.addNum(1)
mf.addNum(2)
print(mf.findMedian())  # Output: 1.5
mf.addNum(3)
print(mf.findMedian())  # Output: 2.0
```

### Pattern 2: Sliding Window Median
**Problem**: Find median in each sliding window of size k.

```python
import heapq
from collections import defaultdict

def median_sliding_window(nums, k):
    def get_median():
        if len(small) == len(large):
            return (small[0] * -1 + large[0]) / 2.0
        else:
            return small[0] * -1
    
    small = []  # Max heap
    large = []  # Min heap
    balance = 0  # balance = len(small) - len(large)
    
    result = []
    
    for i, num in enumerate(nums):
        # Add number to appropriate heap
        if not small or num <= small[0] * -1:
            heapq.heappush(small, -num)
            balance += 1
        else:
            heapq.heappush(large, num)
            balance -= 1
        
        # Rebalance heaps
        if balance < 0:
            heapq.heappush(small, -heapq.heappop(large))
            balance += 2
        elif balance > 1:
            heapq.heappush(large, -heapq.heappop(small))
            balance -= 2
        
        # Remove element going out of window
        if i >= k - 1:
            result.append(get_median())
            
            # Remove the element going out of window
            out_num = nums[i - k + 1]
            
            if out_num <= small[0] * -1:
                # Remove from small heap
                small.remove(-out_num)
                heapq.heapify(small)
                balance -= 1
            else:
                # Remove from large heap
                large.remove(out_num)
                heapq.heapify(large)
                balance += 1
            
            # Rebalance after removal
            if balance < 0:
                heapq.heappush(small, -heapq.heappop(large))
                balance += 2
            elif balance > 1:
                heapq.heappush(large, -heapq.heappop(small))
                balance -= 2
    
    return result

# Example usage
nums = [1, 3, -1, -3, 5, 3, 6, 7]
k = 3
print(median_sliding_window(nums, k))  # Output: [1.0, -1.0, -1.0, 3.0, 5.0, 6.0]
```

### Pattern 3: IPO (Initial Public Offering)
**Problem**: Maximize capital by choosing at most k projects with given initial capital.

```python
import heapq

def find_maximized_capital(k, w, profits, capital):
    # Min heap for available projects (sorted by capital required)
    min_capital_heap = []
    # Max heap for affordable projects (sorted by profit)
    max_profit_heap = []
    
    # Add all projects to min_capital_heap
    for i in range(len(profits)):
        heapq.heappush(min_capital_heap, (capital[i], profits[i]))
    
    for _ in range(k):
        # Move all affordable projects to max_profit_heap
        while min_capital_heap and min_capital_heap[0][0] <= w:
            cap, profit = heapq.heappop(min_capital_heap)
            heapq.heappush(max_profit_heap, -profit)
        
        # If no affordable projects, break
        if not max_profit_heap:
            break
        
        # Select the most profitable project
        w += -heapq.heappop(max_profit_heap)
    
    return w

# Example usage
k = 2
w = 0
profits = [1, 2, 3]
capital = [0, 1, 1]
print(find_maximized_capital(k, w, profits, capital))  # Output: 4
```

### Pattern 4: Schedule Tasks to Minimize Interval
**Problem**: Find minimum interval to finish all tasks with given durations.

```python
import heapq

def minimum_effort_path(heights):
    rows, cols = len(heights), len(heights[0])
    
    # Two heaps to maintain efforts
    # In this context, we use one heap but demonstrate two-heap thinking
    min_heap = [(0, 0, 0)]  # (effort, row, col)
    efforts = [[float('inf')] * cols for _ in range(rows)]
    efforts[0][0] = 0
    
    directions = [(0, 1), (1, 0), (0, -1), (-1, 0)]
    
    while min_heap:
        current_effort, row, col = heapq.heappop(min_heap)
        
        if row == rows - 1 and col == cols - 1:
            return current_effort
        
        if current_effort > efforts[row][col]:
            continue
        
        for dr, dc in directions:
            new_row, new_col = row + dr, col + dc
            
            if 0 <= new_row < rows and 0 <= new_col < cols:
                new_effort = max(current_effort, 
                               abs(heights[new_row][new_col] - heights[row][col]))
                
                if new_effort < efforts[new_row][new_col]:
                    efforts[new_row][new_col] = new_effort
                    heapq.heappush(min_heap, (new_effort, new_row, new_col))
    
    return efforts[rows-1][cols-1]
```

### Pattern 5: Next Interval
**Problem**: Find the next interval for each interval in the array.

```python
import heapq

def find_right_interval(intervals):
    # Max heap for intervals sorted by start time
    start_heap = []
    # Min heap for intervals sorted by end time  
    end_heap = []
    
    result = [-1] * len(intervals)
    
    # Add all intervals to start_heap with their indices
    for i, interval in enumerate(intervals):
        heapq.heappush(start_heap, (-interval[0], i))
    
    # Process intervals in order of end time
    intervals_with_index = [(interval[1], interval[0], i) for i, interval in enumerate(intervals)]
    intervals_with_index.sort()
    
    for end_time, start_time, original_index in intervals_with_index:
        # Remove intervals that start before current end time
        while start_heap and -start_heap[0][0] < end_time:
            heapq.heappop(start_heap)
        
        # The interval with minimum start time >= end_time is the answer
        if start_heap:
            result[original_index] = start_heap[0][1]
    
    return result

# Simpler approach using sorting and binary search
def find_right_interval_v2(intervals):
    # Create array of (start_time, original_index)
    starts = sorted((interval[0], i) for i, interval in enumerate(intervals))
    
    result = []
    for interval in intervals:
        end_time = interval[1]
        
        # Binary search for smallest start >= end_time
        left, right = 0, len(starts)
        while left < right:
            mid = (left + right) // 2
            if starts[mid][0] >= end_time:
                right = mid
            else:
                left = mid + 1
        
        result.append(starts[left][1] if left < len(starts) else -1)
    
    return result

# Example usage
intervals = [[1,2], [2,3], [3,4]]
print(find_right_interval_v2(intervals))  # Output: [1, 2, -1]
```

### Pattern 6: Meeting Rooms with Priority
**Problem**: Find minimum meeting rooms needed with priority scheduling.

```python
import heapq

def min_meeting_rooms_priority(meetings):
    if not meetings:
        return 0
    
    # Sort meetings by start time
    meetings.sort(key=lambda x: x[0])
    
    # Min heap to track end times of ongoing meetings
    min_heap = []
    # Max heap to track priorities (if needed)
    max_rooms = 0
    
    for start, end in meetings:
        # Remove meetings that have ended
        while min_heap and min_heap[0] <= start:
            heapq.heappop(min_heap)
        
        # Add current meeting's end time
        heapq.heappush(min_heap, end)
        
        # Update maximum rooms needed
        max_rooms = max(max_rooms, len(min_heap))
    
    return max_rooms

# Example usage
meetings = [[0,30], [5,10], [15,20]]
print(min_meeting_rooms_priority(meetings))  # Output: 2
```

### Pattern 7: Maximize Sum with K Negations
**Problem**: Maximize sum after negating at most k elements.

```python
import heapq

def largest_sum_after_k_negations(nums, k):
    # Convert to min heap
    heapq.heapify(nums)
    
    # Negate k smallest elements
    for _ in range(k):
        val = heapq.heappop(nums)
        heapq.heappush(nums, -val)
    
    return sum(nums)

# Alternative approach using two heaps concept
def largest_sum_after_k_negations_v2(nums, k):
    # Separate positive and negative numbers
    negatives = []
    positives = []
    
    for num in nums:
        if num < 0:
            heapq.heappush(negatives, num)  # Min heap for negatives
        else:
            heapq.heappush(positives, num)  # Min heap for positives
    
    # First, negate as many negatives as possible
    while k > 0 and negatives:
        val = heapq.heappop(negatives)
        heapq.heappush(positives, -val)
        k -= 1
    
    # If k is still positive and odd, negate the smallest positive
    if k % 2 == 1 and positives:
        smallest = heapq.heappop(positives)
        heapq.heappush(positives, -smallest)
    
    return sum(positives) + sum(negatives)

# Example usage
nums = [4, 2, 3]
k = 1
print(largest_sum_after_k_negations(nums, k))  # Output: 5
```

## Practice Problems

### Easy
1. **Find Median from Data Stream** - Basic two heaps
2. **Last Stone Weight** - Max heap simulation
3. **Kth Largest Element** - Heap maintenance

### Medium
1. **Sliding Window Median** - Two heaps with window
2. **Meeting Rooms II** - Interval scheduling
3. **Top K Frequent Words** - Frequency-based heaps
4. **Find Right Interval** - Interval processing

### Hard
1. **IPO** - Capital and profit optimization
2. **Sliding Window Maximum** - Deque vs heap approaches
3. **Employee Free Time** - Complex interval merging
4. **Minimum Effort Path** - Path finding with effort

## Tips and Tricks

1. **Balance Maintenance**: Keep heap sizes balanced (difference ≤ 1)
2. **Invariant**: Max heap contains smaller half, min heap contains larger half
3. **Median Access**: O(1) time after proper balance maintenance
4. **Lazy Deletion**: Sometimes easier than removing from middle of heap
5. **Heap Properties**: Remember max heap uses negative values in Python

## Common Mistakes

1. **Unbalanced Heaps**: Not maintaining proper size balance
2. **Wrong Median Calculation**: Incorrect formula when heaps have equal size
3. **Heap Type Confusion**: Using min heap where max heap is needed
4. **Index Errors**: Off-by-one errors in sliding window problems
5. **Performance Issues**: Not optimizing heap operations

## Related Patterns

- **Top K Elements**: Uses single heap for k elements
- **Sliding Window**: Combined with heaps for window-based problems
- **Binary Search**: Alternative approach for some median problems
- **Quick Select**: Alternative O(n) approach for single median

## Implementation Languages

Two heaps pattern implementation across languages:
- **Python**: Use `heapq` with negative values for max heap
- **Java**: Use `PriorityQueue` with custom comparators
- **JavaScript**: Implement custom heap classes
- **C++**: Use `priority_queue` with greater<int> for min heap
- **Go**: Use `container/heap` with custom implementation