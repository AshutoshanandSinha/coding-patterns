"""
Subsets and Backtracking Pattern Examples

This module demonstrates backtracking patterns for generating subsets,
permutations, combinations, and solving constraint satisfaction problems.
"""

from typing import List, Set


class SubsetGeneration:
    """Subset generation patterns."""
    
    @staticmethod
    def subsets(nums: List[int]) -> List[List[int]]:
        """Generate all possible subsets."""
        result = []
        
        def backtrack(start: int, current_subset: List[int]) -> None:
            # Add current subset to result
            result.append(current_subset[:])
            
            # Try adding each remaining element
            for i in range(start, len(nums)):
                current_subset.append(nums[i])
                backtrack(i + 1, current_subset)
                current_subset.pop()  # Backtrack
        
        backtrack(0, [])
        return result
    
    @staticmethod
    def subsets_with_duplicates(nums: List[int]) -> List[List[int]]:
        """Generate subsets with duplicate elements."""
        nums.sort()  # Sort to handle duplicates
        result = []
        
        def backtrack(start: int, current_subset: List[int]) -> None:
            result.append(current_subset[:])
            
            for i in range(start, len(nums)):
                # Skip duplicates
                if i > start and nums[i] == nums[i - 1]:
                    continue
                
                current_subset.append(nums[i])
                backtrack(i + 1, current_subset)
                current_subset.pop()
        
        backtrack(0, [])
        return result
    
    @staticmethod
    def combination_sum(candidates: List[int], target: int) -> List[List[int]]:
        """Find combinations that sum to target (with repetition)."""
        result = []
        
        def backtrack(start: int, current_combination: List[int], current_sum: int) -> None:
            if current_sum == target:
                result.append(current_combination[:])
                return
            
            if current_sum > target:
                return
            
            for i in range(start, len(candidates)):
                current_combination.append(candidates[i])
                # Can reuse same element, so pass i (not i + 1)
                backtrack(i, current_combination, current_sum + candidates[i])
                current_combination.pop()
        
        backtrack(0, [], 0)
        return result
    
    @staticmethod
    def combination_sum_ii(candidates: List[int], target: int) -> List[List[int]]:
        """Find combinations that sum to target (without repetition)."""
        candidates.sort()
        result = []
        
        def backtrack(start: int, current_combination: List[int], current_sum: int) -> None:
            if current_sum == target:
                result.append(current_combination[:])
                return
            
            if current_sum > target:
                return
            
            for i in range(start, len(candidates)):
                # Skip duplicates
                if i > start and candidates[i] == candidates[i - 1]:
                    continue
                
                current_combination.append(candidates[i])
                backtrack(i + 1, current_combination, current_sum + candidates[i])
                current_combination.pop()
        
        backtrack(0, [], 0)
        return result


class PermutationGeneration:
    """Permutation generation patterns."""
    
    @staticmethod
    def permute(nums: List[int]) -> List[List[int]]:
        """Generate all permutations."""
        result = []
        
        def backtrack(current_permutation: List[int]) -> None:
            if len(current_permutation) == len(nums):
                result.append(current_permutation[:])
                return
            
            for num in nums:
                if num not in current_permutation:
                    current_permutation.append(num)
                    backtrack(current_permutation)
                    current_permutation.pop()
        
        backtrack([])
        return result
    
    @staticmethod
    def permute_unique(nums: List[int]) -> List[List[int]]:
        """Generate unique permutations with duplicates."""
        nums.sort()
        result = []
        used = [False] * len(nums)
        
        def backtrack(current_permutation: List[int]) -> None:
            if len(current_permutation) == len(nums):
                result.append(current_permutation[:])
                return
            
            for i in range(len(nums)):
                if used[i]:
                    continue
                
                # Skip duplicates: if current element equals previous and
                # previous is not used, skip current
                if i > 0 and nums[i] == nums[i - 1] and not used[i - 1]:
                    continue
                
                used[i] = True
                current_permutation.append(nums[i])
                backtrack(current_permutation)
                current_permutation.pop()
                used[i] = False
        
        backtrack([])
        return result
    
    @staticmethod
    def next_permutation(nums: List[int]) -> None:
        """Find next lexicographically greater permutation in-place."""
        # Find the largest index i such that nums[i] < nums[i + 1]
        i = len(nums) - 2
        while i >= 0 and nums[i] >= nums[i + 1]:
            i -= 1
        
        if i >= 0:
            # Find the largest index j such that nums[i] < nums[j]
            j = len(nums) - 1
            while nums[j] <= nums[i]:
                j -= 1
            
            # Swap nums[i] and nums[j]
            nums[i], nums[j] = nums[j], nums[i]
        
        # Reverse the suffix starting at nums[i + 1]
        nums[i + 1:] = reversed(nums[i + 1:])


class ConstraintSatisfaction:
    """Constraint satisfaction problems using backtracking."""
    
    @staticmethod
    def solve_n_queens(n: int) -> List[List[str]]:
        """Solve N-Queens problem."""
        result = []
        board = ['.' * n for _ in range(n)]
        
        def is_safe(row: int, col: int) -> bool:
            # Check column
            for i in range(row):
                if board[i][col] == 'Q':
                    return False
            
            # Check upper-left diagonal
            i, j = row - 1, col - 1
            while i >= 0 and j >= 0:
                if board[i][j] == 'Q':
                    return False
                i, j = i - 1, j - 1
            
            # Check upper-right diagonal
            i, j = row - 1, col + 1
            while i >= 0 and j < n:
                if board[i][j] == 'Q':
                    return False
                i, j = i - 1, j + 1
            
            return True
        
        def backtrack(row: int) -> None:
            if row == n:
                result.append(board[:])
                return
            
            for col in range(n):
                if is_safe(row, col):
                    # Place queen
                    board[row] = board[row][:col] + 'Q' + board[row][col + 1:]
                    backtrack(row + 1)
                    # Remove queen (backtrack)
                    board[row] = board[row][:col] + '.' + board[row][col + 1:]
        
        backtrack(0)
        return result
    
    @staticmethod
    def solve_sudoku(board: List[List[str]]) -> None:
        """Solve Sudoku puzzle in-place."""
        def is_valid(row: int, col: int, num: str) -> bool:
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
        
        def backtrack() -> bool:
            for i in range(9):
                for j in range(9):
                    if board[i][j] == '.':
                        for num in '123456789':
                            if is_valid(i, j, num):
                                board[i][j] = num
                                if backtrack():
                                    return True
                                board[i][j] = '.'  # Backtrack
                        return False
            return True
        
        backtrack()


class WordProblems:
    """Word-related backtracking problems."""
    
    @staticmethod
    def letter_combinations(digits: str) -> List[str]:
        """Generate letter combinations for phone number digits."""
        if not digits:
            return []
        
        phone_map = {
            '2': 'abc', '3': 'def', '4': 'ghi', '5': 'jkl',
            '6': 'mno', '7': 'pqrs', '8': 'tuv', '9': 'wxyz'
        }
        
        result = []
        
        def backtrack(index: int, current_combination: str) -> None:
            if index == len(digits):
                result.append(current_combination)
                return
            
            for letter in phone_map[digits[index]]:
                backtrack(index + 1, current_combination + letter)
        
        backtrack(0, "")
        return result
    
    @staticmethod
    def word_search(board: List[List[str]], word: str) -> bool:
        """Search for word in 2D grid."""
        rows, cols = len(board), len(board[0])
        
        def backtrack(row: int, col: int, index: int) -> bool:
            if index == len(word):
                return True
            
            if (row < 0 or row >= rows or col < 0 or col >= cols or
                board[row][col] != word[index]):
                return False
            
            # Mark cell as visited
            temp = board[row][col]
            board[row][col] = '#'
            
            # Search in all 4 directions
            found = (backtrack(row + 1, col, index + 1) or
                    backtrack(row - 1, col, index + 1) or
                    backtrack(row, col + 1, index + 1) or
                    backtrack(row, col - 1, index + 1))
            
            # Restore cell
            board[row][col] = temp
            
            return found
        
        for i in range(rows):
            for j in range(cols):
                if backtrack(i, j, 0):
                    return True
        
        return False
    
    @staticmethod
    def word_search_ii(board: List[List[str]], words: List[str]) -> List[str]:
        """Find all words in 2D grid using Trie."""
        class TrieNode:
            def __init__(self):
                self.children = {}
                self.word = None
        
        # Build Trie
        root = TrieNode()
        for word in words:
            node = root
            for char in word:
                if char not in node.children:
                    node.children[char] = TrieNode()
                node = node.children[char]
            node.word = word
        
        rows, cols = len(board), len(board[0])
        result = []
        
        def backtrack(row: int, col: int, node: TrieNode) -> None:
            if (row < 0 or row >= rows or col < 0 or col >= cols or
                board[row][col] not in node.children):
                return
            
            char = board[row][col]
            node = node.children[char]
            
            if node.word:
                result.append(node.word)
                node.word = None  # Avoid duplicates
            
            # Mark as visited
            board[row][col] = '#'
            
            # Search in all 4 directions
            for dr, dc in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
                backtrack(row + dr, col + dc, node)
            
            # Restore cell
            board[row][col] = char
        
        for i in range(rows):
            for j in range(cols):
                backtrack(i, j, root)
        
        return result


# Example usage and test cases
if __name__ == "__main__":
    # Test subset generation
    sg = SubsetGeneration()
    
    nums = [1, 2, 3]
    print("Subsets of [1, 2, 3]:", sg.subsets(nums))
    
    nums_with_dups = [1, 2, 2]
    print("Subsets with duplicates [1, 2, 2]:", sg.subsets_with_duplicates(nums_with_dups))
    
    candidates = [2, 3, 6, 7]
    target = 7
    print("Combination sum (target=7):", sg.combination_sum(candidates, target))
    
    # Test permutation generation
    pg = PermutationGeneration()
    
    nums = [1, 2, 3]
    print("Permutations of [1, 2, 3]:", pg.permute(nums))
    
    nums_with_dups = [1, 1, 2]
    print("Unique permutations of [1, 1, 2]:", pg.permute_unique(nums_with_dups))
    
    # Test constraint satisfaction
    cs = ConstraintSatisfaction()
    
    print("4-Queens solutions:")
    solutions = cs.solve_n_queens(4)
    for i, solution in enumerate(solutions):
        print(f"Solution {i + 1}:")
        for row in solution:
            print(row)
        print()
    
    # Test word problems
    wp = WordProblems()
    
    digits = "23"
    print("Letter combinations for '23':", wp.letter_combinations(digits))
    
    board = [
        ['A', 'B', 'C', 'E'],
        ['S', 'F', 'C', 'S'],
        ['A', 'D', 'E', 'E']
    ]
    word = "ABCCED"
    print(f"Word '{word}' exists in board:", wp.word_search(board, word))
    
    words = ["oath", "pea", "eat", "rain"]
    board2 = [
        ['o', 'a', 'a', 'n'],
        ['e', 't', 'a', 'e'],
        ['i', 'h', 'k', 'r'],
        ['i', 'f', 'l', 'v']
    ]
    print("Words found in board:", wp.word_search_ii(board2, words))