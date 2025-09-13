"""
Binary Search Pattern Examples

This module demonstrates binary search patterns and variations
for efficient searching and optimization problems.
"""

from typing import List, Optional, Callable
import bisect


class BasicBinarySearch:
    """Basic binary search implementations."""
    
    @staticmethod
    def binary_search(arr: List[int], target: int) -> int:
        """Standard binary search returning index or -1."""
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
    
    @staticmethod
    def binary_search_recursive(arr: List[int], target: int, left: int = 0, right: int = None) -> int:
        """Recursive binary search."""
        if right is None:
            right = len(arr) - 1
        
        if left > right:
            return -1
        
        mid = left + (right - left) // 2
        
        if arr[mid] == target:
            return mid
        elif arr[mid] < target:
            return BasicBinarySearch.binary_search_recursive(arr, target, mid + 1, right)
        else:
            return BasicBinarySearch.binary_search_recursive(arr, target, left, mid - 1)
    
    @staticmethod
    def find_first_occurrence(arr: List[int], target: int) -> int:
        """Find first occurrence of target."""
        left, right = 0, len(arr) - 1
        result = -1
        
        while left <= right:
            mid = left + (right - left) // 2
            
            if arr[mid] == target:
                result = mid
                right = mid - 1  # Continue searching left
            elif arr[mid] < target:
                left = mid + 1
            else:
                right = mid - 1
        
        return result
    
    @staticmethod
    def find_last_occurrence(arr: List[int], target: int) -> int:
        """Find last occurrence of target."""
        left, right = 0, len(arr) - 1
        result = -1
        
        while left <= right:
            mid = left + (right - left) // 2
            
            if arr[mid] == target:
                result = mid
                left = mid + 1  # Continue searching right
            elif arr[mid] < target:
                left = mid + 1
            else:
                right = mid - 1
        
        return result


class BinarySearchVariations:
    """Advanced binary search patterns."""
    
    @staticmethod
    def search_insert_position(nums: List[int], target: int) -> int:
        """Find position where target should be inserted."""
        left, right = 0, len(nums)
        
        while left < right:
            mid = left + (right - left) // 2
            
            if nums[mid] < target:
                left = mid + 1
            else:
                right = mid
        
        return left
    
    @staticmethod
    def search_rotated_array(nums: List[int], target: int) -> int:
        """Search in rotated sorted array."""
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
    
    @staticmethod
    def find_minimum_rotated(nums: List[int]) -> int:
        """Find minimum in rotated sorted array."""
        left, right = 0, len(nums) - 1
        
        while left < right:
            mid = left + (right - left) // 2
            
            if nums[mid] > nums[right]:
                left = mid + 1
            else:
                right = mid
        
        return nums[left]
    
    @staticmethod
    def search_2d_matrix(matrix: List[List[int]], target: int) -> bool:
        """Search in row and column sorted matrix."""
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


class OptimizationProblems:
    """Problems using binary search for optimization."""
    
    @staticmethod
    def capacity_to_ship_packages(weights: List[int], D: int) -> int:
        """Minimum capacity to ship packages in D days."""
        def can_ship(capacity: int) -> bool:
            days, current_load = 1, 0
            
            for weight in weights:
                if current_load + weight > capacity:
                    days += 1
                    current_load = weight
                else:
                    current_load += weight
            
            return days <= D
        
        left, right = max(weights), sum(weights)
        
        while left < right:
            mid = left + (right - left) // 2
            
            if can_ship(mid):
                right = mid
            else:
                left = mid + 1
        
        return left
    
    @staticmethod
    def split_array_largest_sum(nums: List[int], m: int) -> int:
        """Split array to minimize largest sum."""
        def can_split(max_sum: int) -> bool:
            splits, current_sum = 1, 0
            
            for num in nums:
                if current_sum + num > max_sum:
                    splits += 1
                    current_sum = num
                else:
                    current_sum += num
            
            return splits <= m
        
        left, right = max(nums), sum(nums)
        
        while left < right:
            mid = left + (right - left) // 2
            
            if can_split(mid):
                right = mid
            else:
                left = mid + 1
        
        return left
    
    @staticmethod
    def smallest_divisor(nums: List[int], threshold: int) -> int:
        """Find smallest divisor such that sum <= threshold."""
        import math
        
        def get_sum(divisor: int) -> int:
            return sum(math.ceil(num / divisor) for num in nums)
        
        left, right = 1, max(nums)
        
        while left < right:
            mid = left + (right - left) // 2
            
            if get_sum(mid) <= threshold:
                right = mid
            else:
                left = mid + 1
        
        return left


class MedianProblems:
    """Problems finding medians using binary search."""
    
    @staticmethod
    def find_median_sorted_arrays(nums1: List[int], nums2: List[int]) -> float:
        """Find median of two sorted arrays."""
        if len(nums1) > len(nums2):
            nums1, nums2 = nums2, nums1
        
        m, n = len(nums1), len(nums2)
        total = m + n
        half = total // 2
        
        left, right = 0, m
        
        while left <= right:
            i = (left + right) // 2
            j = half - i
            
            nums1_left = nums1[i - 1] if i > 0 else float('-inf')
            nums1_right = nums1[i] if i < m else float('inf')
            
            nums2_left = nums2[j - 1] if j > 0 else float('-inf')
            nums2_right = nums2[j] if j < n else float('inf')
            
            if nums1_left <= nums2_right and nums2_left <= nums1_right:
                if total % 2 == 0:
                    return (max(nums1_left, nums2_left) + min(nums1_right, nums2_right)) / 2
                else:
                    return min(nums1_right, nums2_right)
            elif nums1_left > nums2_right:
                right = i - 1
            else:
                left = i + 1
        
        return 0.0


class PeakFinding:
    """Peak finding problems using binary search."""
    
    @staticmethod
    def find_peak_element(nums: List[int]) -> int:
        """Find any peak element index."""
        left, right = 0, len(nums) - 1
        
        while left < right:
            mid = left + (right - left) // 2
            
            if nums[mid] < nums[mid + 1]:
                left = mid + 1
            else:
                right = mid
        
        return left
    
    @staticmethod
    def find_peak_2d(matrix: List[List[int]]) -> List[int]:
        """Find peak in 2D matrix."""
        rows, cols = len(matrix), len(matrix[0])
        left, right = 0, cols - 1
        
        while left <= right:
            mid_col = left + (right - left) // 2
            max_row = 0
            
            # Find maximum element in middle column
            for i in range(rows):
                if matrix[i][mid_col] > matrix[max_row][mid_col]:
                    max_row = i
            
            # Check if it's a peak
            left_val = matrix[max_row][mid_col - 1] if mid_col > 0 else -1
            right_val = matrix[max_row][mid_col + 1] if mid_col < cols - 1 else -1
            
            if matrix[max_row][mid_col] >= left_val and matrix[max_row][mid_col] >= right_val:
                return [max_row, mid_col]
            elif left_val > matrix[max_row][mid_col]:
                right = mid_col - 1
            else:
                left = mid_col + 1
        
        return [-1, -1]


# Example usage and test cases
if __name__ == "__main__":
    # Test basic binary search
    bbs = BasicBinarySearch()
    arr = [1, 3, 5, 5, 5, 7, 9, 11]
    
    print("Basic binary search for 5:", bbs.binary_search(arr, 5))
    print("Recursive binary search for 5:", bbs.binary_search_recursive(arr, 5))
    print("First occurrence of 5:", bbs.find_first_occurrence(arr, 5))
    print("Last occurrence of 5:", bbs.find_last_occurrence(arr, 5))
    
    # Test binary search variations
    bsv = BinarySearchVariations()
    
    nums = [1, 3, 5, 6]
    print("Insert position for 5:", bsv.search_insert_position(nums, 5))
    print("Insert position for 2:", bsv.search_insert_position(nums, 2))
    
    rotated = [4, 5, 6, 7, 0, 1, 2]
    print("Search 0 in rotated array:", bsv.search_rotated_array(rotated, 0))
    print("Minimum in rotated array:", bsv.find_minimum_rotated(rotated))
    
    matrix = [[1, 4, 7, 11], [2, 5, 8, 12], [3, 6, 9, 16], [10, 13, 14, 17]]
    print("Search 5 in 2D matrix:", bsv.search_2d_matrix(matrix, 5))
    
    # Test optimization problems
    opt = OptimizationProblems()
    
    weights = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
    print("Ship capacity for 5 days:", opt.capacity_to_ship_packages(weights, 5))
    
    nums = [7, 2, 5, 10, 8]
    print("Split array largest sum (m=2):", opt.split_array_largest_sum(nums, 2))
    
    nums = [1, 2, 5, 9]
    print("Smallest divisor (threshold=6):", opt.smallest_divisor(nums, 6))
    
    # Test median problems
    mp = MedianProblems()
    
    nums1, nums2 = [1, 3], [2]
    print("Median of two arrays:", mp.find_median_sorted_arrays(nums1, nums2))
    
    # Test peak finding
    pf = PeakFinding()
    
    nums = [1, 2, 3, 1]
    print("Peak element index:", pf.find_peak_element(nums))
    
    matrix_2d = [[1, 4], [3, 2]]
    print("Peak in 2D matrix:", pf.find_peak_2d(matrix_2d))