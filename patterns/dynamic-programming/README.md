# Dynamic Programming Patterns for Coding Interviews

## Table of Contents
1. [Overview of Dynamic Programming](#overview)
2. [When to Use Dynamic Programming](#when-to-use-dp)
3. [Main DP Patterns](#main-dp-patterns)
   - [1D DP](#1d-dp)
   - [2D DP](#2d-dp)
   - [Knapsack Problems](#knapsack-problems)
   - [String DP](#string-dp)
4. [DP Optimization Techniques](#optimization-techniques)
5. [Practice Problems by Difficulty](#practice-problems)
6. [Tips and Common Mistakes](#tips-and-mistakes)

## Overview {#overview}

Dynamic Programming (DP) is an algorithmic paradigm that solves complex problems by breaking them down into simpler subproblems. It stores the results of subproblems to avoid computing the same results again, making it an optimization technique over plain recursion.

### Core Principles
1. **Overlapping Subproblems**: The problem can be broken down into subproblems which are reused several times
2. **Optimal Substructure**: An optimal solution can be constructed from optimal solutions of its subproblems

## When to Use Dynamic Programming {#when-to-use-dp}

### Key Indicators
- Problem asks for **optimal** solution (minimum/maximum)
- Problem asks for **number of ways** to do something
- Problem involves making **choices** at each step
- Subproblems **overlap** (same calculation repeated)
- **Memoization** can improve brute force solution

### Problem Keywords
- "Find the minimum/maximum..."
- "Count the number of ways..."
- "Is it possible to..."
- "Find the longest/shortest..."

## Main DP Patterns {#main-dp-patterns}

## 1D DP {#1d-dp}

### When to Recognize
- Single sequence or array processing
- Decision at each element affects future decisions
- State can be represented with one dimension

### Template
```python
def dp_1d(arr):
    n = len(arr)
    dp = [0] * (n + 1)  # Initialize DP array
    
    # Base cases
    dp[0] = base_case_0
    dp[1] = base_case_1
    
    # Fill DP array
    for i in range(2, n + 1):
        dp[i] = transition_function(dp[i-1], dp[i-2], arr[i])
    
    return dp[n]
```

### Key Problems

#### 1. Fibonacci Sequence
**Problem**: Find the nth Fibonacci number.

```python
def fibonacci(n):
    """
    Time: O(n), Space: O(n)
    """
    if n <= 1:
        return n
    
    dp = [0] * (n + 1)
    dp[0], dp[1] = 0, 1
    
    for i in range(2, n + 1):
        dp[i] = dp[i-1] + dp[i-2]
    
    return dp[n]

# Space optimized version
def fibonacci_optimized(n):
    """
    Time: O(n), Space: O(1)
    """
    if n <= 1:
        return n
    
    prev2, prev1 = 0, 1
    for i in range(2, n + 1):
        current = prev1 + prev2
        prev2, prev1 = prev1, current
    
    return prev1
```

#### 2. Climbing Stairs
**Problem**: You can climb either 1 or 2 steps at a time. How many distinct ways to climb n stairs?

```python
def climb_stairs(n):
    """
    Time: O(n), Space: O(1)
    """
    if n <= 2:
        return n
    
    first, second = 1, 2
    for i in range(3, n + 1):
        third = first + second
        first, second = second, third
    
    return second
```

#### 3. House Robber
**Problem**: Rob houses in a line, can't rob adjacent houses. Maximize money robbed.

```python
def rob(nums):
    """
    Time: O(n), Space: O(1)
    """
    if not nums:
        return 0
    if len(nums) == 1:
        return nums[0]
    
    prev2, prev1 = 0, nums[0]
    
    for i in range(1, len(nums)):
        current = max(prev1, prev2 + nums[i])
        prev2, prev1 = prev1, current
    
    return prev1
```

## 2D DP {#2d-dp}

### When to Recognize
- Grid/matrix problems
- Two sequences comparison
- State depends on two variables
- Path counting problems

### Template
```python
def dp_2d(grid):
    m, n = len(grid), len(grid[0])
    dp = [[0] * n for _ in range(m)]
    
    # Initialize base cases
    for i in range(m):
        dp[i][0] = base_case_column
    for j in range(n):
        dp[0][j] = base_case_row
    
    # Fill DP table
    for i in range(1, m):
        for j in range(1, n):
            dp[i][j] = transition_function(dp[i-1][j], dp[i][j-1], grid[i][j])
    
    return dp[m-1][n-1]
```

### Key Problems

#### 1. Unique Paths
**Problem**: Count unique paths from top-left to bottom-right in a grid.

```python
def unique_paths(m, n):
    """
    Time: O(m*n), Space: O(m*n)
    """
    dp = [[1] * n for _ in range(m)]
    
    for i in range(1, m):
        for j in range(1, n):
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
    
    return dp[m-1][n-1]

# Space optimized version
def unique_paths_optimized(m, n):
    """
    Time: O(m*n), Space: O(n)
    """
    dp = [1] * n
    
    for i in range(1, m):
        for j in range(1, n):
            dp[j] += dp[j-1]
    
    return dp[n-1]
```

#### 2. Minimum Path Sum
**Problem**: Find path from top-left to bottom-right with minimum sum.

```python
def min_path_sum(grid):
    """
    Time: O(m*n), Space: O(1) - modifying input
    """
    if not grid or not grid[0]:
        return 0
    
    m, n = len(grid), len(grid[0])
    
    # Initialize first row
    for j in range(1, n):
        grid[0][j] += grid[0][j-1]
    
    # Initialize first column
    for i in range(1, m):
        grid[i][0] += grid[i-1][0]
    
    # Fill the grid
    for i in range(1, m):
        for j in range(1, n):
            grid[i][j] += min(grid[i-1][j], grid[i][j-1])
    
    return grid[m-1][n-1]
```

#### 3. Longest Common Subsequence
**Problem**: Find length of longest common subsequence between two strings.

```python
def longest_common_subsequence(text1, text2):
    """
    Time: O(m*n), Space: O(m*n)
    """
    m, n = len(text1), len(text2)
    dp = [[0] * (n + 1) for _ in range(m + 1)]
    
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if text1[i-1] == text2[j-1]:
                dp[i][j] = dp[i-1][j-1] + 1
            else:
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
    
    return dp[m][n]

# Space optimized version
def lcs_optimized(text1, text2):
    """
    Time: O(m*n), Space: O(min(m,n))
    """
    if len(text1) < len(text2):
        text1, text2 = text2, text1
    
    m, n = len(text1), len(text2)
    prev = [0] * (n + 1)
    curr = [0] * (n + 1)
    
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if text1[i-1] == text2[j-1]:
                curr[j] = prev[j-1] + 1
            else:
                curr[j] = max(prev[j], curr[j-1])
        prev, curr = curr, prev
    
    return prev[n]
```

## Knapsack Problems {#knapsack-problems}

### When to Recognize
- Given items with weights/values
- Capacity constraint
- Maximize/minimize total value
- Items can be taken once (0/1) or multiple times (unbounded)

### 0/1 Knapsack Template
```python
def knapsack_01(weights, values, capacity):
    n = len(weights)
    dp = [[0] * (capacity + 1) for _ in range(n + 1)]
    
    for i in range(1, n + 1):
        for w in range(1, capacity + 1):
            if weights[i-1] <= w:
                dp[i][w] = max(
                    dp[i-1][w],  # Don't take item
                    dp[i-1][w - weights[i-1]] + values[i-1]  # Take item
                )
            else:
                dp[i][w] = dp[i-1][w]
    
    return dp[n][capacity]
```

### Key Problems

#### 1. 0/1 Knapsack
**Problem**: Given items with weights and values, maximize value within weight capacity.

```python
def knapsack(weights, values, capacity):
    """
    Time: O(n*W), Space: O(W)
    """
    dp = [0] * (capacity + 1)
    
    for i in range(len(weights)):
        # Traverse backward to avoid using updated values
        for w in range(capacity, weights[i] - 1, -1):
            dp[w] = max(dp[w], dp[w - weights[i]] + values[i])
    
    return dp[capacity]
```

#### 2. Partition Equal Subset Sum
**Problem**: Can partition array into two subsets with equal sum?

```python
def can_partition(nums):
    """
    Time: O(n*sum), Space: O(sum)
    """
    total = sum(nums)
    if total % 2 != 0:
        return False
    
    target = total // 2
    dp = [False] * (target + 1)
    dp[0] = True
    
    for num in nums:
        for j in range(target, num - 1, -1):
            dp[j] = dp[j] or dp[j - num]
    
    return dp[target]
```

#### 3. Coin Change (Unbounded Knapsack)
**Problem**: Minimum coins needed to make amount (coins can be reused).

```python
def coin_change(coins, amount):
    """
    Time: O(amount * len(coins)), Space: O(amount)
    """
    dp = [float('inf')] * (amount + 1)
    dp[0] = 0
    
    for coin in coins:
        for j in range(coin, amount + 1):
            dp[j] = min(dp[j], dp[j - coin] + 1)
    
    return dp[amount] if dp[amount] != float('inf') else -1
```

## String DP {#string-dp}

### When to Recognize
- String manipulation problems
- Subsequence/substring problems
- String matching/transformation
- Palindrome problems

### Template
```python
def string_dp(s1, s2):
    m, n = len(s1), len(s2)
    dp = [[0] * (n + 1) for _ in range(m + 1)]
    
    # Initialize base cases
    for i in range(m + 1):
        dp[i][0] = base_case_for_empty_s2
    for j in range(n + 1):
        dp[0][j] = base_case_for_empty_s1
    
    # Fill DP table
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if s1[i-1] == s2[j-1]:
                dp[i][j] = dp[i-1][j-1] + match_cost
            else:
                dp[i][j] = min(
                    dp[i-1][j] + delete_cost,
                    dp[i][j-1] + insert_cost,
                    dp[i-1][j-1] + replace_cost
                )
    
    return dp[m][n]
```

### Key Problems

#### 1. Edit Distance
**Problem**: Minimum operations to convert one string to another.

```python
def min_distance(word1, word2):
    """
    Time: O(m*n), Space: O(m*n)
    """
    m, n = len(word1), len(word2)
    dp = [[0] * (n + 1) for _ in range(m + 1)]
    
    # Initialize base cases
    for i in range(m + 1):
        dp[i][0] = i
    for j in range(n + 1):
        dp[0][j] = j
    
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if word1[i-1] == word2[j-1]:
                dp[i][j] = dp[i-1][j-1]
            else:
                dp[i][j] = 1 + min(
                    dp[i-1][j],      # Delete
                    dp[i][j-1],      # Insert
                    dp[i-1][j-1]     # Replace
                )
    
    return dp[m][n]
```

#### 2. Longest Palindromic Subsequence
**Problem**: Find length of longest palindromic subsequence.

```python
def longest_palindromic_subsequence(s):
    """
    Time: O(n^2), Space: O(n^2)
    """
    n = len(s)
    dp = [[0] * n for _ in range(n)]
    
    # Every single character is a palindrome
    for i in range(n):
        dp[i][i] = 1
    
    # Fill for substrings of length 2 to n
    for length in range(2, n + 1):
        for i in range(n - length + 1):
            j = i + length - 1
            if s[i] == s[j]:
                dp[i][j] = dp[i+1][j-1] + 2
            else:
                dp[i][j] = max(dp[i+1][j], dp[i][j-1])
    
    return dp[0][n-1]
```

#### 3. Palindromic Substrings
**Problem**: Count number of palindromic substrings.

```python
def count_substrings(s):
    """
    Time: O(n^2), Space: O(1)
    """
    def expand_around_center(left, right):
        count = 0
        while left >= 0 and right < len(s) and s[left] == s[right]:
            count += 1
            left -= 1
            right += 1
        return count
    
    result = 0
    for i in range(len(s)):
        # Odd length palindromes
        result += expand_around_center(i, i)
        # Even length palindromes
        result += expand_around_center(i, i + 1)
    
    return result

# DP approach
def count_substrings_dp(s):
    """
    Time: O(n^2), Space: O(n^2)
    """
    n = len(s)
    dp = [[False] * n for _ in range(n)]
    count = 0
    
    # Single characters
    for i in range(n):
        dp[i][i] = True
        count += 1
    
    # Two characters
    for i in range(n - 1):
        if s[i] == s[i + 1]:
            dp[i][i + 1] = True
            count += 1
    
    # Three or more characters
    for length in range(3, n + 1):
        for i in range(n - length + 1):
            j = i + length - 1
            if s[i] == s[j] and dp[i + 1][j - 1]:
                dp[i][j] = True
                count += 1
    
    return count
```

## DP Optimization Techniques {#optimization-techniques}

### 1. Space Optimization

#### Rolling Array Technique
When DP state only depends on previous row/column:

```python
# Instead of 2D array
dp = [[0] * n for _ in range(m)]

# Use two 1D arrays
prev = [0] * n
curr = [0] * n

# Or single array when possible
dp = [0] * n
```

#### Example: Space-Optimized LCS
```python
def lcs_space_optimized(text1, text2):
    m, n = len(text1), len(text2)
    if m < n:
        text1, text2 = text2, text1
        m, n = n, m
    
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
```

### 2. Memoization vs Tabulation

#### Memoization (Top-Down)
```python
def fibonacci_memo(n, memo={}):
    if n in memo:
        return memo[n]
    if n <= 1:
        return n
    
    memo[n] = fibonacci_memo(n-1, memo) + fibonacci_memo(n-2, memo)
    return memo[n]
```

#### Tabulation (Bottom-Up)
```python
def fibonacci_tab(n):
    if n <= 1:
        return n
    
    dp = [0] * (n + 1)
    dp[0], dp[1] = 0, 1
    
    for i in range(2, n + 1):
        dp[i] = dp[i-1] + dp[i-2]
    
    return dp[n]
```

## Practice Problems by Difficulty {#practice-problems}

### Easy
1. **Climbing Stairs** - Classic 1D DP introduction
2. **Maximum Subarray** - Kadane's algorithm
3. **Best Time to Buy and Sell Stock** - State machine DP
4. **Range Sum Query - Immutable** - Prefix sum DP
5. **Counting Bits** - Bit manipulation DP

### Medium
1. **Unique Paths II** - 2D DP with obstacles
2. **Coin Change** - Unbounded knapsack
3. **Longest Increasing Subsequence** - 1D DP with binary search optimization
4. **Perfect Squares** - Mathematical DP
5. **Word Break** - String DP with dictionary
6. **House Robber II** - Circular array DP
7. **Decode Ways** - String parsing DP
8. **Target Sum** - Subset sum variation
9. **Palindromic Substrings** - Expand around center
10. **Maximum Product Subarray** - Modified Kadane's

### Hard
1. **Edit Distance** - Classic string DP
2. **Regular Expression Matching** - Complex string DP
3. **Longest Valid Parentheses** - Stack + DP combination
4. **Distinct Subsequences** - String matching DP
5. **Wildcard Matching** - Pattern matching DP
6. **Best Time to Buy and Sell Stock III** - Multi-transaction DP
7. **Interleaving String** - 2D string DP
8. **Scramble String** - Recursive DP with memoization

## Tips and Common Mistakes {#tips-and-mistakes}

### Identification Tips

1. **Look for keywords**: minimum, maximum, count, possible, optimal
2. **Check for choices**: At each step, can you make different decisions?
3. **Subproblem overlap**: Would brute force repeat calculations?
4. **Optimal substructure**: Can you build solution from smaller solutions?

### Common Patterns Recognition

| Pattern | Keywords | Examples |
|---------|----------|----------|
| 1D DP | Sequential decisions, array processing | Fibonacci, Climbing Stairs |
| 2D DP | Grid paths, two sequences | Unique Paths, LCS |
| Knapsack | Items with constraints, optimization | Subset Sum, Coin Change |
| String DP | String transformation, matching | Edit Distance, Palindromes |
| Tree DP | Tree traversal with state | Binary Tree Maximum Path Sum |

### State Definition Guidelines

1. **Identify variables**: What changes between subproblems?
2. **Define meaning clearly**: What does dp[i] or dp[i][j] represent?
3. **Consider all states**: Include all necessary information
4. **Minimize dimensions**: Use fewer dimensions when possible

### Transition Function Tips

1. **Base cases first**: Handle edge cases clearly
2. **Recurrence relation**: How does current state relate to previous states?
3. **All possibilities**: Consider all valid transitions
4. **Optimization**: Choose min/max or count appropriately

### Common Mistakes

1. **Wrong base cases**: Not handling empty inputs or single elements
2. **Off-by-one errors**: Index confusion in loops and arrays
3. **State definition**: Not capturing all necessary information
4. **Space optimization**: Breaking logic when reducing dimensions
5. **Initialization**: Forgetting to initialize DP array properly
6. **Order of computation**: Computing states before dependencies are ready

### Debugging Strategies

1. **Print DP table**: Visualize how values are filled
2. **Check base cases**: Verify initial conditions are correct
3. **Trace small examples**: Walk through algorithm step by step
4. **Verify recurrence**: Ensure transition logic is sound
5. **Test edge cases**: Empty inputs, single elements, boundary conditions

### Time/Space Complexity Analysis

| DP Type | Typical Time | Typical Space | Space Optimized |
|---------|-------------|---------------|-----------------|
| 1D DP | O(n) | O(n) | O(1) |
| 2D DP | O(n²) or O(m×n) | O(n²) or O(m×n) | O(n) or O(min(m,n)) |
| Knapsack | O(n×W) | O(n×W) | O(W) |
| String DP | O(m×n) | O(m×n) | O(min(m,n)) |

### Final Tips for Interviews

1. **Start with brute force**: Explain recursive approach first
2. **Identify optimization**: Point out overlapping subproblems
3. **Choose approach**: Memoization vs tabulation based on problem
4. **Optimize space**: Show space-optimized version if time permits
5. **Test thoroughly**: Walk through examples and edge cases
6. **Complexity analysis**: Always provide time and space complexity

Remember: DP is about finding the optimal way to break down problems. Focus on understanding the state, transitions, and base cases rather than memorizing solutions.