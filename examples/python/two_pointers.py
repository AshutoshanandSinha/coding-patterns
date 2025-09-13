"""
Two Pointers Pattern - Python3 Implementation
============================================

This module contains comprehensive implementations of the Two Pointers pattern
with different variations and common problems.
"""

from typing import List, Optional


class TwoPointers:
    """Two Pointers pattern implementations"""
    
    @staticmethod
    def two_sum_sorted(arr: List[int], target: int) -> List[int]:
        """
        Find two numbers that add up to target in sorted array.
        
        Time: O(n), Space: O(1)
        
        Args:
            arr: Sorted array of integers
            target: Target sum
            
        Returns:
            List of indices [left, right] or [-1, -1] if not found
        """
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
    
    @staticmethod
    def remove_duplicates(arr: List[int]) -> int:
        """
        Remove duplicates from sorted array in-place.
        
        Time: O(n), Space: O(1)
        
        Args:
            arr: Sorted array with duplicates
            
        Returns:
            Length of array after removing duplicates
        """
        if not arr:
            return 0
        
        slow = 0  # Position for next unique element
        
        for fast in range(1, len(arr)):
            if arr[fast] != arr[slow]:
                slow += 1
                arr[slow] = arr[fast]
        
        return slow + 1
    
    @staticmethod
    def is_palindrome(s: str) -> bool:
        """
        Check if string is palindrome (case-insensitive, alphanumeric only).
        
        Time: O(n), Space: O(1)
        
        Args:
            s: Input string
            
        Returns:
            True if palindrome, False otherwise
        """
        left, right = 0, len(s) - 1
        
        while left < right:
            # Skip non-alphanumeric characters
            while left < right and not s[left].isalnum():
                left += 1
            while left < right and not s[right].isalnum():
                right -= 1
            
            if s[left].lower() != s[right].lower():
                return False
            
            left += 1
            right -= 1
        
        return True
    
    @staticmethod
    def three_sum(arr: List[int]) -> List[List[int]]:
        """
        Find all unique triplets that sum to zero.
        
        Time: O(n²), Space: O(1) excluding output
        
        Args:
            arr: Array of integers
            
        Returns:
            List of triplets that sum to zero
        """
        arr.sort()
        result = []
        n = len(arr)
        
        for i in range(n - 2):
            # Skip duplicates for first element
            if i > 0 and arr[i] == arr[i - 1]:
                continue
            
            left, right = i + 1, n - 1
            
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
    
    @staticmethod
    def max_area(height: List[int]) -> int:
        """
        Container with most water problem.
        
        Time: O(n), Space: O(1)
        
        Args:
            height: Array of heights
            
        Returns:
            Maximum water that can be contained
        """
        left, right = 0, len(height) - 1
        max_water = 0
        
        while left < right:
            width = right - left
            current_area = min(height[left], height[right]) * width
            max_water = max(max_water, current_area)
            
            # Move pointer with smaller height
            if height[left] < height[right]:
                left += 1
            else:
                right -= 1
        
        return max_water
    
    @staticmethod
    def sort_colors(nums: List[int]) -> None:
        """
        Dutch flag problem - sort array of 0s, 1s, and 2s.
        
        Time: O(n), Space: O(1)
        
        Args:
            nums: Array containing only 0, 1, 2
        """
        left, right = 0, len(nums) - 1
        current = 0
        
        while current <= right:
            if nums[current] == 0:
                nums[left], nums[current] = nums[current], nums[left]
                left += 1
                current += 1
            elif nums[current] == 2:
                nums[current], nums[right] = nums[right], nums[current]
                right -= 1
                # Don't increment current as we need to check swapped element
            else:  # nums[current] == 1
                current += 1
    
    @staticmethod
    def trap_rain_water(height: List[int]) -> int:
        """
        Trapping rain water problem.
        
        Time: O(n), Space: O(1)
        
        Args:
            height: Array of heights representing elevation map
            
        Returns:
            Units of water that can be trapped
        """
        if not height:
            return 0
        
        left, right = 0, len(height) - 1
        left_max = right_max = 0
        water_trapped = 0
        
        while left < right:
            if height[left] < height[right]:
                if height[left] >= left_max:
                    left_max = height[left]
                else:
                    water_trapped += left_max - height[left]
                left += 1
            else:
                if height[right] >= right_max:
                    right_max = height[right]
                else:
                    water_trapped += right_max - height[right]
                right -= 1
        
        return water_trapped


def run_tests():
    """Test all two pointers implementations"""
    tp = TwoPointers()
    
    # Test two_sum_sorted
    print("Testing two_sum_sorted:")
    arr = [1, 2, 3, 4, 6]
    target = 6
    result = tp.two_sum_sorted(arr, target)
    print(f"Array: {arr}, Target: {target}, Result: {result}")
    assert result == [1, 3], f"Expected [1, 3], got {result}"
    
    # Test remove_duplicates
    print("\nTesting remove_duplicates:")
    arr = [1, 1, 2, 2, 3, 3, 4]
    original = arr.copy()
    length = tp.remove_duplicates(arr)
    print(f"Original: {original}, After: {arr[:length]}, Length: {length}")
    assert arr[:length] == [1, 2, 3, 4], f"Expected [1, 2, 3, 4], got {arr[:length]}"
    
    # Test is_palindrome
    print("\nTesting is_palindrome:")
    test_cases = [
        ("A man a plan a canal Panama", True),
        ("race a car", False),
        ("", True)
    ]
    for s, expected in test_cases:
        result = tp.is_palindrome(s)
        print(f"String: '{s}', Palindrome: {result}")
        assert result == expected, f"Expected {expected}, got {result}"
    
    # Test three_sum
    print("\nTesting three_sum:")
    arr = [-1, 0, 1, 2, -1, -4]
    result = tp.three_sum(arr)
    print(f"Array: {arr}, Triplets: {result}")
    expected = [[-1, -1, 2], [-1, 0, 1]]
    assert result == expected, f"Expected {expected}, got {result}"
    
    # Test max_area
    print("\nTesting max_area:")
    height = [1, 8, 6, 2, 5, 4, 8, 3, 7]
    result = tp.max_area(height)
    print(f"Heights: {height}, Max Area: {result}")
    assert result == 49, f"Expected 49, got {result}"
    
    # Test sort_colors
    print("\nTesting sort_colors:")
    nums = [2, 0, 2, 1, 1, 0]
    original = nums.copy()
    tp.sort_colors(nums)
    print(f"Original: {original}, Sorted: {nums}")
    assert nums == [0, 0, 1, 1, 2, 2], f"Expected [0, 0, 1, 1, 2, 2], got {nums}"
    
    # Test trap_rain_water
    print("\nTesting trap_rain_water:")
    height = [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]
    result = tp.trap_rain_water(height)
    print(f"Heights: {height}, Water Trapped: {result}")
    assert result == 6, f"Expected 6, got {result}"
    
    print("\n✅ All tests passed!")


if __name__ == "__main__":
    run_tests()