#!/usr/bin/env python3
"""
Merge Intervals Pattern Examples

The Merge Intervals pattern is used to deal with overlapping intervals.
It involves sorting intervals and then merging or manipulating them based on overlap.

Time Complexity: O(n log n) due to sorting
Space Complexity: O(n) for output
"""

from typing import List


def merge_intervals(intervals: List[List[int]]) -> List[List[int]]:
    """
    Merge overlapping intervals.
    
    Args:
        intervals: List of intervals [start, end]
        
    Returns:
        List of merged intervals
        
    Time: O(n log n), Space: O(n)
    """
    if not intervals:
        return []
    
    # Sort intervals by start time
    intervals.sort(key=lambda x: x[0])
    merged = [intervals[0]]
    
    for current in intervals[1:]:
        last_merged = merged[-1]
        
        # If current interval overlaps with last merged interval
        if current[0] <= last_merged[1]:
            # Merge intervals by updating end time
            last_merged[1] = max(last_merged[1], current[1])
        else:
            # No overlap, add current interval
            merged.append(current)
    
    return merged


def insert_interval(intervals: List[List[int]], new_interval: List[int]) -> List[List[int]]:
    """
    Insert a new interval into a sorted list of non-overlapping intervals.
    
    Args:
        intervals: Sorted list of non-overlapping intervals
        new_interval: New interval to insert
        
    Returns:
        List of intervals after insertion and merging
        
    Time: O(n), Space: O(n)
    """
    result = []
    i = 0
    
    # Add all intervals that end before new interval starts
    while i < len(intervals) and intervals[i][1] < new_interval[0]:
        result.append(intervals[i])
        i += 1
    
    # Merge overlapping intervals with new interval
    while i < len(intervals) and intervals[i][0] <= new_interval[1]:
        new_interval[0] = min(new_interval[0], intervals[i][0])
        new_interval[1] = max(new_interval[1], intervals[i][1])
        i += 1
    
    result.append(new_interval)
    
    # Add remaining intervals
    while i < len(intervals):
        result.append(intervals[i])
        i += 1
    
    return result


def interval_intersection(first_list: List[List[int]], second_list: List[List[int]]) -> List[List[int]]:
    """
    Find intersection of two lists of intervals.
    
    Args:
        first_list: First list of intervals
        second_list: Second list of intervals
        
    Returns:
        List of intersection intervals
        
    Time: O(m + n), Space: O(min(m, n))
    """
    result = []
    i = j = 0
    
    while i < len(first_list) and j < len(second_list):
        # Find intersection
        start = max(first_list[i][0], second_list[j][0])
        end = min(first_list[i][1], second_list[j][1])
        
        # If valid intersection exists
        if start <= end:
            result.append([start, end])
        
        # Move pointer of interval that ends first
        if first_list[i][1] < second_list[j][1]:
            i += 1
        else:
            j += 1
    
    return result


def can_attend_meetings(intervals: List[List[int]]) -> bool:
    """
    Check if a person can attend all meetings (no overlapping intervals).
    
    Args:
        intervals: List of meeting intervals [start, end]
        
    Returns:
        True if can attend all meetings, False otherwise
        
    Time: O(n log n), Space: O(1)
    """
    if not intervals:
        return True
    
    intervals.sort(key=lambda x: x[0])
    
    for i in range(1, len(intervals)):
        if intervals[i][0] < intervals[i-1][1]:
            return False
    
    return True


def min_meeting_rooms(intervals: List[List[int]]) -> int:
    """
    Find minimum number of meeting rooms required.
    
    Args:
        intervals: List of meeting intervals [start, end]
        
    Returns:
        Minimum number of rooms needed
        
    Time: O(n log n), Space: O(n)
    """
    if not intervals:
        return 0
    
    # Create separate lists for start and end times
    starts = sorted([interval[0] for interval in intervals])
    ends = sorted([interval[1] for interval in intervals])
    
    rooms_needed = 0
    max_rooms = 0
    start_ptr = end_ptr = 0
    
    while start_ptr < len(intervals):
        # If a meeting starts before another ends, we need a new room
        if starts[start_ptr] < ends[end_ptr]:
            rooms_needed += 1
            max_rooms = max(max_rooms, rooms_needed)
            start_ptr += 1
        else:
            # A meeting ended, free up a room
            rooms_needed -= 1
            end_ptr += 1
    
    return max_rooms


def merge_intervals_with_names(intervals: List[List]) -> List[List]:
    """
    Merge intervals that also contain names/identifiers.
    
    Args:
        intervals: List of [start, end, name] intervals
        
    Returns:
        List of merged intervals with combined names
        
    Time: O(n log n), Space: O(n)
    """
    if not intervals:
        return []
    
    # Sort by start time
    intervals.sort(key=lambda x: x[0])
    merged = [intervals[0].copy()]
    
    for current in intervals[1:]:
        last_merged = merged[-1]
        
        # If current interval overlaps with last merged interval
        if current[0] <= last_merged[1]:
            # Merge intervals
            last_merged[1] = max(last_merged[1], current[1])
            # Combine names
            if len(current) > 2 and len(last_merged) > 2:
                last_merged[2] = f"{last_merged[2]},{current[2]}"
        else:
            # No overlap, add current interval
            merged.append(current.copy())
    
    return merged


def employee_free_time(schedule: List[List[List[int]]]) -> List[List[int]]:
    """
    Find common free time for all employees.
    
    Args:
        schedule: List of employee schedules, each containing intervals
        
    Returns:
        List of free time intervals common to all employees
        
    Time: O(n log n), Space: O(n)
    """
    # Flatten all intervals
    all_intervals = []
    for employee_schedule in schedule:
        all_intervals.extend(employee_schedule)
    
    # Merge all intervals
    merged = merge_intervals(all_intervals)
    
    # Find gaps between merged intervals
    free_time = []
    for i in range(1, len(merged)):
        if merged[i-1][1] < merged[i][0]:
            free_time.append([merged[i-1][1], merged[i][0]])
    
    return free_time


# Test functions
def test_merge_intervals():
    """Test merging overlapping intervals"""
    test_cases = [
        ([[1,3],[2,6],[8,10],[15,18]], [[1,6],[8,10],[15,18]]),
        ([[1,4],[4,5]], [[1,5]]),
        ([[1,4],[0,4]], [[0,4]]),
        ([[1,4],[2,3]], [[1,4]])
    ]
    
    for intervals, expected in test_cases:
        result = merge_intervals(intervals)
        assert result == expected, f"Expected {expected}, got {result}"
    
    print("✅ Merge intervals tests passed")


def test_insert_interval():
    """Test inserting interval"""
    test_cases = [
        ([[1,3],[6,9]], [2,5], [[1,5],[6,9]]),
        ([[1,2],[3,5],[6,7],[8,10],[12,16]], [4,8], [[1,2],[3,10],[12,16]]),
        ([], [5,7], [[5,7]]),
        ([[1,5]], [2,3], [[1,5]])
    ]
    
    for intervals, new_interval, expected in test_cases:
        result = insert_interval(intervals, new_interval)
        assert result == expected, f"Expected {expected}, got {result}"
    
    print("✅ Insert interval tests passed")


def test_interval_intersection():
    """Test interval intersection"""
    first = [[0,2],[5,10],[13,23],[24,25]]
    second = [[1,5],[8,12],[15,24],[25,26]]
    expected = [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
    
    result = interval_intersection(first, second)
    assert result == expected, f"Expected {expected}, got {result}"
    
    print("✅ Interval intersection tests passed")


def test_meeting_rooms():
    """Test meeting room problems"""
    # Can attend all meetings
    assert can_attend_meetings([[0,30],[5,10],[15,20]]) == False
    assert can_attend_meetings([[7,10],[2,4]]) == True
    
    # Minimum meeting rooms
    assert min_meeting_rooms([[0,30],[5,10],[15,20]]) == 2
    assert min_meeting_rooms([[7,10],[2,4]]) == 1
    assert min_meeting_rooms([[9,10],[4,9],[4,17]]) == 2
    
    print("✅ Meeting room tests passed")


def test_employee_free_time():
    """Test employee free time"""
    schedule = [[[1,3],[6,7]],[[2,4]],[[2,5],[9,12]]]
    expected = [[5,6],[7,9]]
    
    result = employee_free_time(schedule)
    assert result == expected, f"Expected {expected}, got {result}"
    
    print("✅ Employee free time tests passed")


if __name__ == "__main__":
    print("Testing Merge Intervals Pattern...")
    test_merge_intervals()
    test_insert_interval()
    test_interval_intersection()
    test_meeting_rooms()
    test_employee_free_time()
    print("\n🎉 All Merge Intervals tests passed!")