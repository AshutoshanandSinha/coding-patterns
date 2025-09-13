"""
Sliding Window Pattern - Python3 Implementation
==============================================

This module contains comprehensive implementations of the Sliding Window pattern
with both fixed and variable window size variations.
"""

from typing import List, Dict, Optional
from collections import defaultdict


class SlidingWindow:
    """Sliding Window pattern implementations"""
    
    @staticmethod
    def max_sum_subarray(arr: List[int], k: int) -> int:
        """
        Find maximum sum of any contiguous subarray of size k.
        
        Time: O(n), Space: O(1)
        
        Args:
            arr: Array of integers
            k: Window size
            
        Returns:
            Maximum sum of k-sized subarray
        """
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
    
    @staticmethod
    def longest_substring_without_repeating(s: str) -> int:
        """
        Find length of longest substring without repeating characters.
        
        Time: O(n), Space: O(min(m,n)) where m is charset size
        
        Args:
            s: Input string
            
        Returns:
            Length of longest substring without repeating characters
        """
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
    
    @staticmethod
    def min_window_substring(s: str, t: str) -> str:
        """
        Find minimum window in s that contains all characters of t.
        
        Time: O(|s| + |t|), Space: O(|s| + |t|)
        
        Args:
            s: Source string
            t: Target string (pattern)
            
        Returns:
            Minimum window substring or empty string if not found
        """
        if not s or not t:
            return ""
        
        # Character frequency in t
        required = defaultdict(int)
        for char in t:
            required[char] += 1
        
        left = right = 0
        formed = 0  # Number of unique chars with desired frequency
        window_counts = defaultdict(int)
        
        # (window length, left, right)
        ans = float('inf'), None, None
        
        while right < len(s):
            # Add character from right to window
            char = s[right]
            window_counts[char] += 1
            
            # Check if current character's frequency matches required
            if char in required and window_counts[char] == required[char]:
                formed += 1
            
            # Try to contract window
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
    
    @staticmethod
    def longest_substring_k_distinct(s: str, k: int) -> int:
        """
        Find length of longest substring with at most k distinct characters.
        
        Time: O(n), Space: O(k)
        
        Args:
            s: Input string
            k: Maximum distinct characters allowed
            
        Returns:
            Length of longest valid substring
        """
        if k == 0:
            return 0
        
        char_frequency = defaultdict(int)
        left = 0
        max_length = 0
        
        for right in range(len(s)):
            # Add character to frequency map
            char_frequency[s[right]] += 1
            
            # Shrink window if more than k distinct characters
            while len(char_frequency) > k:
                char_frequency[s[left]] -= 1
                if char_frequency[s[left]] == 0:
                    del char_frequency[s[left]]
                left += 1
            
            max_length = max(max_length, right - left + 1)
        
        return max_length
    
    @staticmethod
    def subarray_product_less_than_k(nums: List[int], k: int) -> int:
        """
        Count subarrays where product of elements is less than k.
        
        Time: O(n), Space: O(1)
        
        Args:
            nums: Array of positive integers
            k: Target value
            
        Returns:
            Count of valid subarrays
        """
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
    
    @staticmethod
    def fruits_into_baskets(fruits: List[int]) -> int:
        """
        Pick maximum fruits with at most 2 different types.
        
        Time: O(n), Space: O(1)
        
        Args:
            fruits: Array representing fruit types
            
        Returns:
            Maximum number of fruits that can be picked
        """
        fruit_frequency = defaultdict(int)
        left = 0
        max_fruits = 0
        
        for right in range(len(fruits)):
            # Add fruit to basket
            fruit_frequency[fruits[right]] += 1
            
            # Shrink window if more than 2 fruit types
            while len(fruit_frequency) > 2:
                fruit_frequency[fruits[left]] -= 1
                if fruit_frequency[fruits[left]] == 0:
                    del fruit_frequency[fruits[left]]
                left += 1
            
            max_fruits = max(max_fruits, right - left + 1)
        
        return max_fruits
    
    @staticmethod
    def smallest_subarray_sum(arr: List[int], target: int) -> int:
        """
        Find length of smallest subarray with sum >= target.
        
        Time: O(n), Space: O(1)
        
        Args:
            arr: Array of positive integers
            target: Target sum
            
        Returns:
            Length of smallest subarray or 0 if not possible
        """
        left = 0
        window_sum = 0
        min_length = float('inf')
        
        for right in range(len(arr)):
            window_sum += arr[right]
            
            # Shrink window while sum >= target
            while window_sum >= target:
                min_length = min(min_length, right - left + 1)
                window_sum -= arr[left]
                left += 1
        
        return min_length if min_length != float('inf') else 0
    
    @staticmethod
    def find_anagrams(s: str, p: str) -> List[int]:
        """
        Find all start indices of anagrams of p in s.
        
        Time: O(|s| + |p|), Space: O(1) - at most 26 characters
        
        Args:
            s: Source string
            p: Pattern string
            
        Returns:
            List of starting indices where anagrams are found
        """
        if len(p) > len(s):
            return []
        
        # Character frequency in pattern
        p_freq = defaultdict(int)
        for char in p:
            p_freq[char] += 1
        
        window_freq = defaultdict(int)
        result = []
        left = 0
        
        for right in range(len(s)):
            # Add character to window
            window_freq[s[right]] += 1
            
            # Maintain window size
            if right - left + 1 > len(p):
                if window_freq[s[left]] == 1:
                    del window_freq[s[left]]
                else:
                    window_freq[s[left]] -= 1
                left += 1
            
            # Check if current window is anagram
            if right - left + 1 == len(p) and window_freq == p_freq:
                result.append(left)
        
        return result
    
    @staticmethod
    def character_replacement(s: str, k: int) -> int:
        """
        Find longest substring with same letters after k replacements.
        
        Time: O(n), Space: O(1) - at most 26 characters
        
        Args:
            s: Input string
            k: Number of allowed replacements
            
        Returns:
            Length of longest valid substring
        """
        char_frequency = defaultdict(int)
        left = 0
        max_length = 0
        max_count = 0
        
        for right in range(len(s)):
            char_frequency[s[right]] += 1
            max_count = max(max_count, char_frequency[s[right]])
            
            # If replacements needed > k, shrink window
            if (right - left + 1) - max_count > k:
                char_frequency[s[left]] -= 1
                left += 1
            
            max_length = max(max_length, right - left + 1)
        
        return max_length


def run_tests():
    """Test all sliding window implementations"""
    sw = SlidingWindow()
    
    # Test max_sum_subarray
    print("Testing max_sum_subarray:")
    arr = [2, 1, 5, 1, 3, 2]
    k = 3
    result = sw.max_sum_subarray(arr, k)
    print(f"Array: {arr}, K: {k}, Max Sum: {result}")
    assert result == 9, f"Expected 9, got {result}"
    
    # Test longest_substring_without_repeating
    print("\nTesting longest_substring_without_repeating:")
    s = "abcabcbb"
    result = sw.longest_substring_without_repeating(s)
    print(f"String: '{s}', Longest Length: {result}")
    assert result == 3, f"Expected 3, got {result}"
    
    # Test min_window_substring
    print("\nTesting min_window_substring:")
    s = "ADOBECODEBANC"
    t = "ABC"
    result = sw.min_window_substring(s, t)
    print(f"Source: '{s}', Target: '{t}', Min Window: '{result}'")
    assert result == "BANC", f"Expected 'BANC', got '{result}'"
    
    # Test longest_substring_k_distinct
    print("\nTesting longest_substring_k_distinct:")
    s = "araaci"
    k = 2
    result = sw.longest_substring_k_distinct(s, k)
    print(f"String: '{s}', K: {k}, Longest Length: {result}")
    assert result == 4, f"Expected 4, got {result}"
    
    # Test subarray_product_less_than_k
    print("\nTesting subarray_product_less_than_k:")
    nums = [10, 5, 2, 6]
    k = 100
    result = sw.subarray_product_less_than_k(nums, k)
    print(f"Array: {nums}, K: {k}, Count: {result}")
    assert result == 8, f"Expected 8, got {result}"
    
    # Test fruits_into_baskets
    print("\nTesting fruits_into_baskets:")
    fruits = [1, 2, 1, 2, 3, 2, 2]
    result = sw.fruits_into_baskets(fruits)
    print(f"Fruits: {fruits}, Max Fruits: {result}")
    assert result == 4, f"Expected 4, got {result}"
    
    # Test smallest_subarray_sum
    print("\nTesting smallest_subarray_sum:")
    arr = [2, 3, 1, 2, 4, 3]
    target = 7
    result = sw.smallest_subarray_sum(arr, target)
    print(f"Array: {arr}, Target: {target}, Min Length: {result}")
    assert result == 2, f"Expected 2, got {result}"
    
    # Test find_anagrams
    print("\nTesting find_anagrams:")
    s = "abab"
    p = "ab"
    result = sw.find_anagrams(s, p)
    print(f"String: '{s}', Pattern: '{p}', Indices: {result}")
    assert result == [0, 1, 2], f"Expected [0, 1, 2], got {result}"
    
    # Test character_replacement
    print("\nTesting character_replacement:")
    s = "AABABBA"
    k = 1
    result = sw.character_replacement(s, k)
    print(f"String: '{s}', K: {k}, Max Length: {result}")
    assert result == 4, f"Expected 4, got {result}"
    
    print("\n✅ All tests passed!")


if __name__ == "__main__":
    run_tests()