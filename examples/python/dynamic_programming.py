#!/usr/bin/env python3
"""
Dynamic Programming Pattern Examples

Dynamic Programming solves complex problems by breaking them down into simpler subproblems.
It stores the results of subproblems to avoid redundant calculations.

Common approaches:
1. Top-down (Memoization): Recursion + Cache
2. Bottom-up (Tabulation): Iterative + Table

Time Complexity: Usually O(n) to O(n²)
Space Complexity: O(n) to O(n²) for memoization/tables
"""

from typing import List, Dict
from functools import lru_cache


def fibonacci(n: int) -> int:
    """
    Calculate nth Fibonacci number using bottom-up DP.
    
    Args:
        n: The position in Fibonacci sequence
        
    Returns:
        The nth Fibonacci number
        
    Time: O(n), Space: O(1)
    """
    if n <= 1:
        return n
    
    prev1, prev2 = 0, 1
    
    for i in range(2, n + 1):
        current = prev1 + prev2
        prev1, prev2 = prev2, current
    
    return prev2


def climb_stairs(n: int) -> int:
    """
    Count ways to climb n stairs (1 or 2 steps at a time).
    
    Args:
        n: Number of stairs
        
    Returns:
        Number of ways to climb stairs
        
    Time: O(n), Space: O(1)
    """
    if n <= 2:
        return n
    
    prev1, prev2 = 1, 2
    
    for i in range(3, n + 1):
        current = prev1 + prev2
        prev1, prev2 = prev2, current
    
    return prev2


def house_robber(nums: List[int]) -> int:
    """
    Rob houses to maximize money without robbing adjacent houses.
    
    Args:
        nums: Money in each house
        
    Returns:
        Maximum money that can be robbed
        
    Time: O(n), Space: O(1)
    """
    if not nums:
        return 0
    if len(nums) == 1:
        return nums[0]
    
    prev1, prev2 = 0, nums[0]
    
    for i in range(1, len(nums)):
        current = max(prev2, prev1 + nums[i])
        prev1, prev2 = prev2, current
    
    return prev2


def coin_change(coins: List[int], amount: int) -> int:
    """
    Find minimum coins needed to make amount.
    
    Args:
        coins: Available coin denominations
        amount: Target amount
        
    Returns:
        Minimum coins needed, or -1 if impossible
        
    Time: O(amount * len(coins)), Space: O(amount)
    """
    dp = [float('inf')] * (amount + 1)
    dp[0] = 0
    
    for i in range(1, amount + 1):
        for coin in coins:
            if coin <= i:
                dp[i] = min(dp[i], dp[i - coin] + 1)
    
    return dp[amount] if dp[amount] != float('inf') else -1


def longest_increasing_subsequence(nums: List[int]) -> int:
    """
    Find length of longest increasing subsequence.
    
    Args:
        nums: Input array
        
    Returns:
        Length of LIS
        
    Time: O(n²), Space: O(n)
    """
    if not nums:
        return 0
    
    n = len(nums)
    dp = [1] * n
    
    for i in range(1, n):
        for j in range(i):
            if nums[j] < nums[i]:
                dp[i] = max(dp[i], dp[j] + 1)
    
    return max(dp)


def knapsack_01(weights: List[int], values: List[int], capacity: int) -> int:
    """
    0/1 Knapsack problem - maximize value within weight capacity.
    
    Args:
        weights: Weight of each item
        values: Value of each item
        capacity: Maximum weight capacity
        
    Returns:
        Maximum value achievable
        
    Time: O(n * capacity), Space: O(capacity)
    """
    n = len(weights)
    dp = [0] * (capacity + 1)
    
    for i in range(n):
        # Traverse backwards to avoid using updated values
        for w in range(capacity, weights[i] - 1, -1):
            dp[w] = max(dp[w], dp[w - weights[i]] + values[i])
    
    return dp[capacity]


def longest_common_subsequence(text1: str, text2: str) -> int:
    """
    Find length of longest common subsequence.
    
    Args:
        text1: First string
        text2: Second string
        
    Returns:
        Length of LCS
        
    Time: O(m * n), Space: O(min(m, n))
    """
    m, n = len(text1), len(text2)
    
    # Use space-optimized version
    prev = [0] * (n + 1)
    
    for i in range(1, m + 1):
        curr = [0] * (n + 1)
        for j in range(1, n + 1):
            if text1[i-1] == text2[j-1]:
                curr[j] = prev[j-1] + 1
            else:
                curr[j] = max(prev[j], curr[j-1])
        prev = curr
    
    return prev[n]


def edit_distance(word1: str, word2: str) -> int:
    """
    Find minimum edit distance (insert, delete, replace).
    
    Args:
        word1: First word
        word2: Second word
        
    Returns:
        Minimum edit distance
        
    Time: O(m * n), Space: O(min(m, n))
    """
    m, n = len(word1), len(word2)
    
    # Space-optimized version
    prev = list(range(n + 1))
    
    for i in range(1, m + 1):
        curr = [i]
        for j in range(1, n + 1):
            if word1[i-1] == word2[j-1]:
                curr.append(prev[j-1])
            else:
                curr.append(1 + min(prev[j], curr[j-1], prev[j-1]))
        prev = curr
    
    return prev[n]


def word_break(s: str, word_dict: List[str]) -> bool:
    """
    Check if string can be segmented using dictionary words.
    
    Args:
        s: Input string
        word_dict: Dictionary of valid words
        
    Returns:
        True if string can be segmented
        
    Time: O(n² * m), Space: O(n)
    """
    word_set = set(word_dict)
    dp = [False] * (len(s) + 1)
    dp[0] = True
    
    for i in range(1, len(s) + 1):
        for j in range(i):
            if dp[j] and s[j:i] in word_set:
                dp[i] = True
                break
    
    return dp[len(s)]


def max_product_subarray(nums: List[int]) -> int:
    """
    Find maximum product of contiguous subarray.
    
    Args:
        nums: Input array
        
    Returns:
        Maximum product
        
    Time: O(n), Space: O(1)
    """
    if not nums:
        return 0
    
    max_prod = min_prod = result = nums[0]
    
    for i in range(1, len(nums)):
        num = nums[i]
        
        # Handle negative numbers by swapping max and min
        if num < 0:
            max_prod, min_prod = min_prod, max_prod
        
        max_prod = max(num, max_prod * num)
        min_prod = min(num, min_prod * num)
        
        result = max(result, max_prod)
    
    return result


@lru_cache(maxsize=None)
def unique_paths_with_obstacles_memo(m: int, n: int, obstacles: tuple) -> int:
    """
    Count unique paths with obstacles using memoization.
    
    Args:
        m: Grid height
        n: Grid width
        obstacles: Tuple of obstacle positions
        
    Returns:
        Number of unique paths
        
    Time: O(m * n), Space: O(m * n)
    """
    obstacle_set = set(obstacles)
    
    def dfs(row: int, col: int) -> int:
        if row >= m or col >= n or (row, col) in obstacle_set:
            return 0
        if row == m - 1 and col == n - 1:
            return 1
        
        return dfs(row + 1, col) + dfs(row, col + 1)
    
    return dfs(0, 0)


# Test functions
def test_fibonacci():
    """Test Fibonacci calculation"""
    assert fibonacci(0) == 0
    assert fibonacci(1) == 1
    assert fibonacci(10) == 55
    assert fibonacci(15) == 610
    print("✅ Fibonacci tests passed")


def test_climb_stairs():
    """Test stair climbing"""
    assert climb_stairs(1) == 1
    assert climb_stairs(2) == 2
    assert climb_stairs(3) == 3
    assert climb_stairs(4) == 5
    print("✅ Climb stairs tests passed")


def test_house_robber():
    """Test house robber"""
    assert house_robber([1, 2, 3, 1]) == 4
    assert house_robber([2, 7, 9, 3, 1]) == 12
    assert house_robber([5]) == 5
    print("✅ House robber tests passed")


def test_coin_change():
    """Test coin change"""
    assert coin_change([1, 3, 4], 6) == 2
    assert coin_change([2], 3) == -1
    assert coin_change([1], 0) == 0
    print("✅ Coin change tests passed")


def test_lis():
    """Test longest increasing subsequence"""
    assert longest_increasing_subsequence([10, 9, 2, 5, 3, 7, 101, 18]) == 4
    assert longest_increasing_subsequence([0, 1, 0, 3, 2, 3]) == 4
    assert longest_increasing_subsequence([7, 7, 7, 7, 7, 7, 7]) == 1
    print("✅ LIS tests passed")


def test_knapsack():
    """Test 0/1 knapsack"""
    weights = [1, 3, 4, 5]
    values = [1, 4, 5, 7]
    assert knapsack_01(weights, values, 7) == 9
    print("✅ Knapsack tests passed")


def test_lcs():
    """Test longest common subsequence"""
    assert longest_common_subsequence("abcde", "ace") == 3
    assert longest_common_subsequence("abc", "abc") == 3
    assert longest_common_subsequence("abc", "def") == 0
    print("✅ LCS tests passed")


def test_edit_distance():
    """Test edit distance"""
    assert edit_distance("horse", "ros") == 3
    assert edit_distance("intention", "execution") == 5
    print("✅ Edit distance tests passed")


def test_word_break():
    """Test word break"""
    assert word_break("leetcode", ["leet", "code"]) == True
    assert word_break("applepenapple", ["apple", "pen"]) == True
    assert word_break("catsandog", ["cats", "dog", "sand", "and", "cat"]) == False
    print("✅ Word break tests passed")


def test_max_product():
    """Test max product subarray"""
    assert max_product_subarray([2, 3, -2, 4]) == 6
    assert max_product_subarray([-2, 0, -1]) == 0
    assert max_product_subarray([-2, 3, -4]) == 24
    print("✅ Max product tests passed")


if __name__ == "__main__":
    print("Testing Dynamic Programming Pattern...")
    test_fibonacci()
    test_climb_stairs()
    test_house_robber()
    test_coin_change()
    test_lis()
    test_knapsack()
    test_lcs()
    test_edit_distance()
    test_word_break()
    test_max_product()
    print("\n🎉 All Dynamic Programming tests passed!")