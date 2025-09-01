# Backtracking Pattern - Complete Interview Guide

## Table of Contents
1. [Overview](#overview)
2. [Core Concepts](#core-concepts)
3. [Backtracking Template](#backtracking-template)
4. [Problem Types](#problem-types)
5. [Key Problems with Solutions](#key-problems-with-solutions)
6. [Time/Space Complexity Analysis](#timespace-complexity-analysis)
7. [Optimization Techniques](#optimization-techniques)
8. [Practice Problems by Difficulty](#practice-problems-by-difficulty)
9. [Common Mistakes & Debugging Tips](#common-mistakes--debugging-tips)

## Overview

Backtracking is a algorithmic technique that considers searching every possible combination in order to solve computational problems. It builds candidates to the solution incrementally and abandons candidates ("backtracks") as soon as it determines that they cannot possibly be completed to a valid solution.

### When to Use Backtracking

- **Generate all possible solutions**: subsets, permutations, combinations
- **Find one valid solution**: N-Queens, Sudoku solving
- **Optimization problems**: finding the best path with constraints
- **Constraint satisfaction**: placing items with specific rules

### Key Characteristics
- Explores solution space using DFS (Depth-First Search)
- Makes choices and undoes them (backtrack) if they lead to invalid solutions
- Uses recursion to explore all possibilities
- Often involves choosing/not choosing elements or trying different positions

## Core Concepts

### 1. Decision Tree
Every backtracking problem can be visualized as a decision tree where:
- Each node represents a partial solution
- Each edge represents a decision/choice
- Leaves represent complete solutions (valid or invalid)

### 2. State Space
- **Current state**: partial solution being built
- **Valid state**: meets all constraints so far
- **Goal state**: complete valid solution

### 3. Pruning
Early termination of branches that cannot lead to valid solutions, improving efficiency.

## Backtracking Template

```python
def backtrack(path, choices, result):
    # Base case: if path is a complete solution
    if is_complete(path):
        result.append(path.copy())  # or path[:]
        return
    
    # Try each possible choice
    for choice in choices:
        # Check if choice is valid
        if is_valid(path, choice):
            # Make choice
            path.append(choice)
            
            # Recursively explore
            backtrack(path, get_next_choices(choices, choice), result)
            
            # Backtrack: undo the choice
            path.pop()

# Usage
def solve_problem(input_data):
    result = []
    backtrack([], get_initial_choices(input_data), result)
    return result
```

### Alternative Template (Index-based)

```python
def backtrack(nums, start, path, result):
    # Add current path to result (for subsets)
    result.append(path[:])
    
    # Try each choice from start index
    for i in range(start, len(nums)):
        # Make choice
        path.append(nums[i])
        
        # Recurse with next index
        backtrack(nums, i + 1, path, result)
        
        # Backtrack
        path.pop()
```

## Problem Types

### 1. Subsets/Combinations Generation
- Generate all possible subsets of a set
- Generate combinations of specific size
- Handle duplicates in input

### 2. Permutations
- Generate all arrangements of elements
- Handle duplicates and constraints
- Partial permutations

### 3. Board Problems
- N-Queens: place N queens on N×N board
- Sudoku solving
- Word search on grid
- Knight's tour

### 4. Path Finding with Constraints
- Find paths in maze with obstacles
- Word search with specific patterns
- Generate valid combinations with rules

## Key Problems with Solutions

### 1. Generate All Subsets (Leetcode 78)

**Problem**: Given an integer array nums of unique elements, return all possible subsets.

```python
def subsets(nums):
    """
    Generate all possible subsets using backtracking.
    
    Time: O(2^n * n) - 2^n subsets, each takes O(n) to copy
    Space: O(2^n * n) - storing all subsets
    """
    result = []
    
    def backtrack(start, path):
        # Add current subset to result
        result.append(path[:])
        
        # Try adding each remaining element
        for i in range(start, len(nums)):
            path.append(nums[i])
            backtrack(i + 1, path)
            path.pop()
    
    backtrack(0, [])
    return result

# Alternative: Choice-based approach
def subsets_choice(nums):
    result = []
    
    def backtrack(index, path):
        if index == len(nums):
            result.append(path[:])
            return
        
        # Choice 1: don't include current element
        backtrack(index + 1, path)
        
        # Choice 2: include current element
        path.append(nums[index])
        backtrack(index + 1, path)
        path.pop()
    
    backtrack(0, [])
    return result
```

### 2. Combinations (Leetcode 77)

**Problem**: Given two integers n and k, return all possible combinations of k numbers out of 1 to n.

```python
def combine(n, k):
    """
    Generate all combinations of k elements from 1 to n.
    
    Time: O(C(n,k) * k) - C(n,k) combinations, each takes O(k) to copy
    Space: O(C(n,k) * k) - storing all combinations
    """
    result = []
    
    def backtrack(start, path):
        # Base case: we have k elements
        if len(path) == k:
            result.append(path[:])
            return
        
        # Optimization: if we can't get enough elements, stop
        need = k - len(path)
        available = n - start + 1
        if available < need:
            return
        
        # Try each number from start to n
        for i in range(start, n + 1):
            path.append(i)
            backtrack(i + 1, path)
            path.pop()
    
    backtrack(1, [])
    return result
```

### 3. Permutations (Leetcode 46)

**Problem**: Given an array nums of distinct integers, return all possible permutations.

```python
def permute(nums):
    """
    Generate all permutations of distinct elements.
    
    Time: O(n! * n) - n! permutations, each takes O(n) to copy
    Space: O(n! * n) - storing all permutations
    """
    result = []
    
    def backtrack(path):
        # Base case: permutation is complete
        if len(path) == len(nums):
            result.append(path[:])
            return
        
        # Try each unused number
        for num in nums:
            if num not in path:
                path.append(num)
                backtrack(path)
                path.pop()
    
    backtrack([])
    return result

# Optimized version using visited array
def permute_optimized(nums):
    result = []
    used = [False] * len(nums)
    
    def backtrack(path):
        if len(path) == len(nums):
            result.append(path[:])
            return
        
        for i in range(len(nums)):
            if not used[i]:
                path.append(nums[i])
                used[i] = True
                backtrack(path)
                path.pop()
                used[i] = False
    
    backtrack([])
    return result
```

### 4. Combination Sum (Leetcode 39)

**Problem**: Given an array of distinct integers and a target, return all unique combinations where the candidates sum to target.

```python
def combinationSum(candidates, target):
    """
    Find all combinations that sum to target (can reuse elements).
    
    Time: O(2^target) in worst case
    Space: O(target) for recursion stack
    """
    result = []
    
    def backtrack(start, path, current_sum):
        # Base cases
        if current_sum == target:
            result.append(path[:])
            return
        if current_sum > target:
            return
        
        # Try each candidate starting from 'start' index
        for i in range(start, len(candidates)):
            path.append(candidates[i])
            # Can reuse same element, so pass 'i' not 'i+1'
            backtrack(i, path, current_sum + candidates[i])
            path.pop()
    
    backtrack(0, [], 0)
    return result

# Alternative with early pruning
def combinationSum_pruned(candidates, target):
    # Sort for better pruning
    candidates.sort()
    result = []
    
    def backtrack(start, path, remaining):
        if remaining == 0:
            result.append(path[:])
            return
        
        for i in range(start, len(candidates)):
            # Early pruning: if current candidate > remaining, stop
            if candidates[i] > remaining:
                break
            
            path.append(candidates[i])
            backtrack(i, path, remaining - candidates[i])
            path.pop()
    
    backtrack(0, [], target)
    return result
```

### 5. Generate Parentheses (Leetcode 22)

**Problem**: Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

```python
def generateParenthesis(n):
    """
    Generate all valid combinations of n pairs of parentheses.
    
    Time: O(4^n / sqrt(n)) - Catalan number
    Space: O(4^n / sqrt(n)) - storing all combinations
    """
    result = []
    
    def backtrack(path, open_count, close_count):
        # Base case: used all n pairs
        if len(path) == 2 * n:
            result.append(path)
            return
        
        # Can add opening bracket if we haven't used all n
        if open_count < n:
            backtrack(path + '(', open_count + 1, close_count)
        
        # Can add closing bracket if it won't exceed opening brackets
        if close_count < open_count:
            backtrack(path + ')', open_count, close_count + 1)
    
    backtrack('', 0, 0)
    return result

# Alternative using list for path
def generateParenthesis_list(n):
    result = []
    
    def backtrack(path, open_count, close_count):
        if len(path) == 2 * n:
            result.append(''.join(path))
            return
        
        if open_count < n:
            path.append('(')
            backtrack(path, open_count + 1, close_count)
            path.pop()
        
        if close_count < open_count:
            path.append(')')
            backtrack(path, open_count, close_count + 1)
            path.pop()
    
    backtrack([], 0, 0)
    return result
```

### 6. Word Search (Leetcode 79)

**Problem**: Given a 2D board and a word, find if the word exists in the grid.

```python
def exist(board, word):
    """
    Search for word in 2D board using backtracking.
    
    Time: O(m * n * 4^L) where L is length of word
    Space: O(L) for recursion stack
    """
    if not board or not board[0]:
        return False
    
    rows, cols = len(board), len(board[0])
    
    def backtrack(row, col, index):
        # Base case: found entire word
        if index == len(word):
            return True
        
        # Check bounds and character match
        if (row < 0 or row >= rows or col < 0 or col >= cols or 
            board[row][col] != word[index]):
            return False
        
        # Mark cell as visited
        temp = board[row][col]
        board[row][col] = '#'
        
        # Explore all 4 directions
        found = (backtrack(row + 1, col, index + 1) or
                backtrack(row - 1, col, index + 1) or
                backtrack(row, col + 1, index + 1) or
                backtrack(row, col - 1, index + 1))
        
        # Backtrack: restore cell
        board[row][col] = temp
        
        return found
    
    # Try starting from each cell
    for i in range(rows):
        for j in range(cols):
            if backtrack(i, j, 0):
                return True
    
    return False
```

### 7. N-Queens (Leetcode 51)

**Problem**: Place n queens on an n×n chessboard such that no two queens attack each other.

```python
def solveNQueens(n):
    """
    Solve N-Queens problem using backtracking.
    
    Time: O(n!) - pruning reduces this significantly
    Space: O(n^2) for the board
    """
    result = []
    board = [['.' for _ in range(n)] for _ in range(n)]
    
    def is_safe(row, col):
        # Check column
        for i in range(row):
            if board[i][col] == 'Q':
                return False
        
        # Check diagonal (top-left to bottom-right)
        i, j = row - 1, col - 1
        while i >= 0 and j >= 0:
            if board[i][j] == 'Q':
                return False
            i -= 1
            j -= 1
        
        # Check anti-diagonal (top-right to bottom-left)
        i, j = row - 1, col + 1
        while i >= 0 and j < n:
            if board[i][j] == 'Q':
                return False
            i -= 1
            j += 1
        
        return True
    
    def backtrack(row):
        # Base case: placed all queens
        if row == n:
            result.append([''.join(row) for row in board])
            return
        
        # Try placing queen in each column of current row
        for col in range(n):
            if is_safe(row, col):
                board[row][col] = 'Q'
                backtrack(row + 1)
                board[row][col] = '.'
    
    backtrack(0)
    return result

# Optimized version using sets for O(1) conflict checking
def solveNQueens_optimized(n):
    result = []
    
    def backtrack(row, cols, diag1, diag2, path):
        if row == n:
            result.append(path[:])
            return
        
        for col in range(n):
            if col in cols or (row - col) in diag1 or (row + col) in diag2:
                continue
            
            cols.add(col)
            diag1.add(row - col)
            diag2.add(row + col)
            path.append('.' * col + 'Q' + '.' * (n - col - 1))
            
            backtrack(row + 1, cols, diag1, diag2, path)
            
            cols.remove(col)
            diag1.remove(row - col)
            diag2.remove(row + col)
            path.pop()
    
    backtrack(0, set(), set(), set(), [])
    return result
```

### 8. Sudoku Solver (Leetcode 37)

**Problem**: Write a program to solve a Sudoku puzzle by filling the empty cells.

```python
def solveSudoku(board):
    """
    Solve Sudoku puzzle using backtracking.
    
    Time: O(9^(n*n)) in worst case, but pruning helps significantly
    Space: O(n*n) for recursion stack
    """
    def is_valid(row, col, num):
        # Check row
        for j in range(9):
            if board[row][j] == num:
                return False
        
        # Check column
        for i in range(9):
            if board[i][col] == num:
                return False
        
        # Check 3x3 box
        start_row, start_col = 3 * (row // 3), 3 * (col // 3)
        for i in range(start_row, start_row + 3):
            for j in range(start_col, start_col + 3):
                if board[i][j] == num:
                    return False
        
        return True
    
    def backtrack():
        for i in range(9):
            for j in range(9):
                if board[i][j] == '.':
                    # Try digits 1-9
                    for num in '123456789':
                        if is_valid(i, j, num):
                            board[i][j] = num
                            if backtrack():
                                return True
                            board[i][j] = '.'
                    return False
        return True
    
    backtrack()

# Optimized version with constraint propagation
def solveSudoku_optimized(board):
    def get_candidates(row, col):
        used = set()
        
        # Add numbers from row
        for j in range(9):
            if board[row][j] != '.':
                used.add(board[row][j])
        
        # Add numbers from column
        for i in range(9):
            if board[i][col] != '.':
                used.add(board[i][col])
        
        # Add numbers from 3x3 box
        start_row, start_col = 3 * (row // 3), 3 * (col // 3)
        for i in range(start_row, start_row + 3):
            for j in range(start_col, start_col + 3):
                if board[i][j] != '.':
                    used.add(board[i][j])
        
        return [str(i) for i in range(1, 10) if str(i) not in used]
    
    def find_best_cell():
        min_candidates = 10
        best_cell = None
        
        for i in range(9):
            for j in range(9):
                if board[i][j] == '.':
                    candidates = get_candidates(i, j)
                    if len(candidates) < min_candidates:
                        min_candidates = len(candidates)
                        best_cell = (i, j, candidates)
                        if min_candidates == 0:
                            return best_cell
        
        return best_cell
    
    def backtrack():
        cell = find_best_cell()
        if not cell:
            return True  # No empty cells left
        
        row, col, candidates = cell
        
        if not candidates:
            return False  # No valid candidates
        
        for num in candidates:
            board[row][col] = num
            if backtrack():
                return True
            board[row][col] = '.'
        
        return False
    
    backtrack()
```

## Time/Space Complexity Analysis

### Time Complexity Patterns

| Problem Type | Time Complexity | Explanation |
|--------------|----------------|-------------|
| Subsets | O(2^n * n) | 2^n subsets, O(n) to copy each |
| Permutations | O(n! * n) | n! permutations, O(n) to copy each |
| Combinations C(n,k) | O(C(n,k) * k) | C(n,k) combinations, O(k) to copy |
| N-Queens | O(n!) | Place n queens with constraints |
| Sudoku | O(9^(empty_cells)) | Try 1-9 for each empty cell |

### Space Complexity Patterns

| Component | Space Complexity | Description |
|-----------|------------------|-------------|
| Recursion Stack | O(depth) | Maximum depth of recursion |
| Path Storage | O(path_length) | Current solution being built |
| Result Storage | O(total_solutions * solution_size) | All solutions |
| Visited Tracking | O(n) | For permutations, word search |

## Optimization Techniques

### 1. Early Pruning
```python
# Stop exploring if remaining choices can't complete solution
if remaining_slots < needed_elements:
    return

# Stop if current sum exceeds target
if current_sum > target:
    return
```

### 2. Constraint Propagation
```python
# For Sudoku: choose cell with fewest valid candidates
def find_best_cell():
    min_candidates = 10
    best_cell = None
    for each empty cell:
        candidates = get_valid_numbers(cell)
        if len(candidates) < min_candidates:
            min_candidates = len(candidates)
            best_cell = cell
    return best_cell
```

### 3. Symmetry Breaking
```python
# For N-Queens: only try first half of first row
for col in range(n // 2 + n % 2):
    # place first queen and solve
```

### 4. Sorting for Better Pruning
```python
# Sort candidates in ascending order for combination sum
candidates.sort()
for candidate in candidates:
    if candidate > remaining_target:
        break  # All larger candidates will also exceed
```

### 5. Bit Manipulation for State
```python
# Use bits to track visited states efficiently
def backtrack(row, cols_mask, diag1_mask, diag2_mask):
    if row == n:
        return 1
    
    available_positions = ((1 << n) - 1) & (~(cols_mask | diag1_mask | diag2_mask))
    count = 0
    
    while available_positions:
        position = available_positions & (-available_positions)
        available_positions ^= position
        
        count += backtrack(row + 1,
                          cols_mask | position,
                          (diag1_mask | position) << 1,
                          (diag2_mask | position) >> 1)
    
    return count
```

## Practice Problems by Difficulty

### Easy
1. **Binary Tree Paths** (Leetcode 257)
   - Find all root-to-leaf paths
   - Time: O(n * h), Space: O(h)

2. **Letter Case Permutation** (Leetcode 784)
   - Generate all possible strings by changing case
   - Time: O(2^n * n), Space: O(2^n * n)

### Medium
1. **Subsets II** (Leetcode 90)
   - Generate subsets with duplicates
   - Key: Sort and skip duplicates

2. **Permutations II** (Leetcode 47)
   - Generate permutations with duplicates
   - Key: Sort and skip used duplicates

3. **Combination Sum II** (Leetcode 40)
   - Combination sum with each number used once
   - Key: Sort and avoid duplicate combinations

4. **Palindrome Partitioning** (Leetcode 131)
   - Partition string into palindromes
   - Key: Check palindrome validity

5. **Word Search II** (Leetcode 212)
   - Find multiple words in grid
   - Key: Use Trie for efficient searching

6. **Beautiful Arrangement** (Leetcode 526)
   - Count arrangements where nums[i] % i == 0 or i % nums[i] == 0
   - Key: Pruning and constraint checking

### Hard
1. **N-Queens II** (Leetcode 52)
   - Count total N-Queens solutions
   - Optimization: bit manipulation

2. **Sudoku Solver** (Leetcode 37)
   - Solve 9x9 Sudoku puzzle
   - Key: Constraint propagation

3. **Expression Add Operators** (Leetcode 282)
   - Insert +, -, * between digits to reach target
   - Key: Handle operator precedence

4. **Remove Invalid Parentheses** (Leetcode 301)
   - Remove minimum parentheses to make valid
   - Key: BFS + validation

## Common Mistakes & Debugging Tips

### 1. Forgetting to Backtrack
```python
# ❌ Wrong: forgetting to remove choice
def backtrack(path):
    for choice in choices:
        path.append(choice)
        backtrack(path)
        # Missing: path.pop()

# ✅ Correct: always backtrack
def backtrack(path):
    for choice in choices:
        path.append(choice)
        backtrack(path)
        path.pop()  # Backtrack
```

### 2. Shallow vs Deep Copy
```python
# ❌ Wrong: shallow copy - all results point to same list
result.append(path)

# ✅ Correct: deep copy
result.append(path[:])  # or path.copy() or list(path)
```

### 3. Index Management in Combinations
```python
# ❌ Wrong: can revisit same element
for i in range(len(nums)):
    backtrack(i)

# ✅ Correct: prevent revisiting
for i in range(start, len(nums)):
    backtrack(i + 1)
```

### 4. Handling Duplicates
```python
# For input with duplicates, sort first and skip
nums.sort()
for i in range(start, len(nums)):
    # Skip duplicates at same level
    if i > start and nums[i] == nums[i-1]:
        continue
    # Process nums[i]
```

### 5. State Restoration in Grid Problems
```python
# ❌ Wrong: modifying without restoration
board[row][col] = 'X'
result = backtrack(...)
return result

# ✅ Correct: save and restore state
original = board[row][col]
board[row][col] = 'X'
result = backtrack(...)
board[row][col] = original  # Restore
return result
```

### Debugging Strategies

1. **Add Print Statements**
```python
def backtrack(path, level=0):
    print("  " * level + f"Exploring: {path}")
    # ... backtracking logic
```

2. **Validate Intermediate States**
```python
def backtrack(path):
    assert is_valid_state(path), f"Invalid state: {path}"
    # ... continue
```

3. **Test with Small Inputs**
   - Start with n=1, n=2 to understand the pattern
   - Verify base cases work correctly

4. **Trace Execution**
   - Draw the recursion tree for small examples
   - Verify pruning conditions are correct

5. **Check Edge Cases**
   - Empty input
   - Single element
   - All elements the same
   - Maximum constraints

### Performance Tips

1. **Early Termination**
   - Add pruning conditions as early as possible
   - Check impossible cases before recursion

2. **Choose Right Data Structures**
   - Use sets for O(1) membership testing
   - Use lists for ordered sequences
   - Consider bit manipulation for flags

3. **Minimize Object Creation**
   - Reuse data structures when possible
   - Pass indices instead of creating sublists

4. **Sort When Beneficial**
   - Enables early pruning in many problems
   - Groups duplicates for easier handling

This comprehensive guide covers all major backtracking patterns and problems commonly seen in coding interviews. Practice these templates and patterns to master backtracking!