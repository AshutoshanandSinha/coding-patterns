# Merge Intervals Pattern

## Overview
The Merge Intervals pattern is used to deal with overlapping intervals. It's particularly useful for problems involving scheduling, time management, or any scenario where you need to merge, insert, or manipulate overlapping ranges.

## When to Use
- **Overlapping Intervals**: When dealing with overlapping time periods or ranges
- **Scheduling Problems**: Meeting rooms, calendar conflicts, resource allocation
- **Range Queries**: Merging ranges, finding gaps, interval intersections
- **Timeline Problems**: Event scheduling, availability checking
- **Data Compression**: Merging consecutive or overlapping data ranges

## Time/Space Complexity
- **Time**: O(n log n) for sorting + O(n) for merging = O(n log n)
- **Space**: O(1) if modifying in-place, O(n) for creating new result

## Core Concepts

### Interval Representation
```python
# Common interval representations
# 1. List/Array: [start, end]
interval = [1, 3]

# 2. Named tuple or class
from collections import namedtuple
Interval = namedtuple('Interval', ['start', 'end'])
interval = Interval(1, 3)

# 3. Custom class
class Interval:
    def __init__(self, start, end):
        self.start = start
        self.end = end
```

### Basic Overlap Detection
```python
def is_overlap(interval1, interval2):
    # Two intervals overlap if:
    # interval1.start <= interval2.end and interval2.start <= interval1.end
    return interval1[0] <= interval2[1] and interval2[0] <= interval1[1]

def merge_two_intervals(interval1, interval2):
    # Merge two overlapping intervals
    return [min(interval1[0], interval2[0]), max(interval1[1], interval2[1])]
```

## Common Problem Patterns

### Pattern 1: Merge Overlapping Intervals
**Problem**: Given a collection of intervals, merge all overlapping intervals.

```python
def merge_intervals(intervals):
    if not intervals:
        return []
    
    # Sort intervals by start time
    intervals.sort(key=lambda x: x[0])
    
    merged = [intervals[0]]
    
    for current in intervals[1:]:
        last_merged = merged[-1]
        
        # Check if current interval overlaps with last merged interval
        if current[0] <= last_merged[1]:
            # Merge intervals
            merged[-1] = [last_merged[0], max(last_merged[1], current[1])]
        else:
            # No overlap, add current interval
            merged.append(current)
    
    return merged

# Example usage
intervals = [[1, 3], [2, 6], [8, 10], [15, 18]]
print(merge_intervals(intervals))  # Output: [[1, 6], [8, 10], [15, 18]]
```

### Pattern 2: Insert Interval
**Problem**: Insert a new interval into a list of sorted non-overlapping intervals.

```python
def insert_interval(intervals, new_interval):
    result = []
    i = 0
    
    # Add all intervals that end before new interval starts
    while i < len(intervals) and intervals[i][1] < new_interval[0]:
        result.append(intervals[i])
        i += 1
    
    # Merge overlapping intervals
    while i < len(intervals) and intervals[i][0] <= new_interval[1]:
        # Update new_interval to merge with current interval
        new_interval = [
            min(new_interval[0], intervals[i][0]),
            max(new_interval[1], intervals[i][1])
        ]
        i += 1
    
    result.append(new_interval)
    
    # Add remaining intervals
    while i < len(intervals):
        result.append(intervals[i])
        i += 1
    
    return result

# Example usage
intervals = [[1, 3], [6, 9]]
new_interval = [2, 5]
print(insert_interval(intervals, new_interval))  # Output: [[1, 5], [6, 9]]
```

### Pattern 3: Interval List Intersections
**Problem**: Find intersection of two sorted interval lists.

```python
def interval_intersection(first_list, second_list):
    result = []
    i = j = 0
    
    while i < len(first_list) and j < len(second_list):
        # Find intersection
        start = max(first_list[i][0], second_list[j][0])
        end = min(first_list[i][1], second_list[j][1])
        
        # If there's an intersection
        if start <= end:
            result.append([start, end])
        
        # Move the pointer of interval that ends first
        if first_list[i][1] < second_list[j][1]:
            i += 1
        else:
            j += 1
    
    return result

# Example usage
A = [[0, 2], [5, 10], [13, 23], [24, 25]]
B = [[1, 5], [8, 12], [15, 24], [25, 26]]
print(interval_intersection(A, B))  
# Output: [[1, 2], [5, 5], [8, 10], [15, 23], [24, 24], [25, 25]]
```

### Pattern 4: Meeting Rooms (Can Attend All Meetings)
**Problem**: Determine if a person can attend all meetings.

```python
def can_attend_meetings(intervals):
    if not intervals:
        return True
    
    # Sort by start time
    intervals.sort(key=lambda x: x[0])
    
    for i in range(1, len(intervals)):
        # Check if current meeting starts before previous ends
        if intervals[i][0] < intervals[i - 1][1]:
            return False
    
    return True

# Example usage
meetings = [[0, 30], [5, 10], [15, 20]]
print(can_attend_meetings(meetings))  # Output: False (overlap between [0,30] and [5,10])
```

### Pattern 5: Meeting Rooms II (Minimum Meeting Rooms)
**Problem**: Find minimum number of meeting rooms required.

```python
import heapq

def min_meeting_rooms(intervals):
    if not intervals:
        return 0
    
    # Sort meetings by start time
    intervals.sort(key=lambda x: x[0])
    
    # Min heap to track end times of ongoing meetings
    min_heap = []
    
    for meeting in intervals:
        # If current meeting starts after earliest ending meeting
        if min_heap and meeting[0] >= min_heap[0]:
            heapq.heappop(min_heap)  # Remove ended meeting
        
        # Add current meeting's end time
        heapq.heappush(min_heap, meeting[1])
    
    return len(min_heap)

# Example usage
meetings = [[0, 30], [5, 10], [15, 20]]
print(min_meeting_rooms(meetings))  # Output: 2
```

### Pattern 6: Employee Free Time
**Problem**: Find common free time for all employees.

```python
def employee_free_time(schedule):
    # Flatten all intervals and sort by start time
    all_intervals = []
    for employee_schedule in schedule:
        all_intervals.extend(employee_schedule)
    
    all_intervals.sort(key=lambda x: x[0])
    
    # Merge overlapping intervals (working hours)
    merged_working = []
    for interval in all_intervals:
        if not merged_working or interval[0] > merged_working[-1][1]:
            merged_working.append(interval)
        else:
            merged_working[-1] = [
                merged_working[-1][0], 
                max(merged_working[-1][1], interval[1])
            ]
    
    # Find gaps between merged intervals (free time)
    free_time = []
    for i in range(1, len(merged_working)):
        if merged_working[i - 1][1] < merged_working[i][0]:
            free_time.append([merged_working[i - 1][1], merged_working[i][0]])
    
    return free_time

# Example usage
schedule = [
    [[1, 3], [6, 7]],      # Employee 1
    [[2, 4]],              # Employee 2  
    [[2, 5], [9, 12]]      # Employee 3
]
print(employee_free_time(schedule))  # Output: [[5, 6], [7, 9]]
```

### Pattern 7: Non-overlapping Intervals
**Problem**: Find minimum number of intervals to remove to make rest non-overlapping.

```python
def erase_overlap_intervals(intervals):
    if not intervals:
        return 0
    
    # Sort by end time (greedy approach)
    intervals.sort(key=lambda x: x[1])
    
    count = 0
    end = intervals[0][1]
    
    for i in range(1, len(intervals)):
        if intervals[i][0] < end:
            # Overlapping interval found, remove it
            count += 1
        else:
            # Update end to current interval's end
            end = intervals[i][1]
    
    return count

# Example usage
intervals = [[1, 2], [2, 3], [3, 4], [1, 3]]
print(erase_overlap_intervals(intervals))  # Output: 1 (remove [1,3])
```

### Pattern 8: Range Addition
**Problem**: Apply range updates efficiently.

```python
def range_addition(length, updates):
    # Use difference array technique
    diff = [0] * (length + 1)
    
    # Apply updates to difference array
    for start, end, inc in updates:
        diff[start] += inc
        if end + 1 < len(diff):
            diff[end + 1] -= inc
    
    # Convert difference array to actual values
    result = [0] * length
    result[0] = diff[0]
    
    for i in range(1, length):
        result[i] = result[i - 1] + diff[i]
    
    return result

# Example usage
updates = [[1, 3, 2], [2, 4, 3], [0, 2, -2]]
print(range_addition(5, updates))  # Output: [-2, 0, 3, 5, 3]
```

## Practice Problems

### Easy
1. **Merge Intervals** - Basic interval merging
2. **Meeting Rooms** - Check if all meetings can be attended
3. **Summary Ranges** - Find consecutive ranges in sorted array

### Medium
1. **Insert Interval** - Insert new interval and merge
2. **Meeting Rooms II** - Minimum meeting rooms needed
3. **Interval List Intersections** - Find intersections of two lists
4. **Non-overlapping Intervals** - Minimum removals to make non-overlapping

### Hard
1. **Employee Free Time** - Find common free time across all employees
2. **Data Stream as Disjoint Intervals** - Maintain disjoint intervals from stream
3. **Range Module** - Track and query ranges efficiently
4. **My Calendar III** - Maximum simultaneous bookings

## Key Strategies

### 1. Sorting Strategy
```python
# Sort by start time (most common)
intervals.sort(key=lambda x: x[0])

# Sort by end time (for greedy algorithms)
intervals.sort(key=lambda x: x[1])

# Custom sorting for specific needs
intervals.sort(key=lambda x: (x[0], x[1]))
```

### 2. Merging Strategy
```python
def merge_strategy(intervals):
    merged = [intervals[0]]
    
    for current in intervals[1:]:
        if current[0] <= merged[-1][1]:  # Overlap condition
            # Merge: extend the end time
            merged[-1][1] = max(merged[-1][1], current[1])
        else:
            # No overlap: add new interval
            merged.append(current)
    
    return merged
```

### 3. Two-Pointer Strategy
```python
def two_pointer_intersection(list1, list2):
    result = []
    i = j = 0
    
    while i < len(list1) and j < len(list2):
        # Process intersection logic
        # Move appropriate pointer based on end times
        pass
    
    return result
```

## Tips and Tricks

1. **Always Sort First**: Most interval problems require sorting by start or end time
2. **Overlap Condition**: `interval1.start <= interval2.end and interval2.start <= interval1.end`
3. **Merge Condition**: `current.start <= previous.end`
4. **Use Min-Heap**: For problems requiring tracking of multiple ending times
5. **Greedy Approach**: Sort by end time for minimum removal problems
6. **Edge Cases**: Empty intervals, single interval, all overlapping

## Common Mistakes

1. **Wrong Sort Order**: Sorting by wrong parameter
2. **Boundary Conditions**: Missing edge cases for interval boundaries
3. **Off-by-One Errors**: Incorrect overlap detection (`<` vs `<=`)
4. **Mutating Input**: Modifying original input when not required
5. **Infinite Loops**: Not advancing pointers properly

## Template Code

```python
def interval_template(intervals):
    if not intervals:
        return handle_empty_case()
    
    # Sort intervals (usually by start time)
    intervals.sort(key=lambda x: x[0])
    
    result = []
    
    for interval in intervals:
        if not result or no_overlap_condition(result[-1], interval):
            # No overlap, add interval
            result.append(interval)
        else:
            # Overlap detected, merge or handle accordingly
            result[-1] = merge_intervals(result[-1], interval)
    
    return result
```

## Related Patterns

- **Greedy Algorithm**: For optimization problems (minimum rooms, removals)
- **Heap/Priority Queue**: For tracking multiple simultaneous events
- **Two Pointers**: For intersection and comparison problems
- **Line Sweep**: For more complex interval processing

## Implementation Notes

- **Python**: Use `lambda` for custom sorting, `heapq` for priority queues
- **Java**: Implement `Comparator` for sorting, `PriorityQueue` for heaps
- **JavaScript**: Use custom compare functions, array methods
- **C++**: Use `std::sort` with custom comparators, `priority_queue`