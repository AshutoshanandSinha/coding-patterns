# Tree Depth-First Search (DFS) Patterns

## Table of Contents
1. [Overview](#overview)
2. [Core Concepts](#core-concepts)
3. [DFS Traversal Types](#dfs-traversal-types)
4. [Time and Space Complexity](#time-and-space-complexity)
5. [Recursive vs Iterative Approaches](#recursive-vs-iterative-approaches)
6. [Template Code](#template-code)
7. [Common Problem Patterns](#common-problem-patterns)
8. [Practice Problems](#practice-problems)
9. [Tips and Common Mistakes](#tips-and-common-mistakes)

## Overview

Tree Depth-First Search (DFS) is a fundamental algorithm for traversing tree data structures. It explores as far as possible along each branch before backtracking. DFS is particularly useful for problems involving path finding, tree structure analysis, and recursive solutions.

### When to Use Tree DFS
- Finding paths from root to leaf
- Calculating tree properties (depth, diameter, sum)
- Tree serialization/deserialization
- Validating tree properties
- Tree structure modifications

## Core Concepts

### Binary Tree Node Definition
```python
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
```

### Key Principles
1. **Recursive Nature**: Most tree problems can be solved recursively
2. **Base Case**: Handle null/leaf nodes appropriately
3. **Divide and Conquer**: Break problem into smaller subproblems
4. **State Management**: Pass information down (parameters) or up (return values)

## DFS Traversal Types

### 1. Pre-order Traversal
- **Order**: Root → Left → Right
- **Use Case**: Tree copying, prefix expressions, tree serialization
```python
def preorder(root):
    if not root:
        return []
    return [root.val] + preorder(root.left) + preorder(root.right)
```

### 2. In-order Traversal
- **Order**: Left → Root → Right
- **Use Case**: BST sorted output, expression evaluation
```python
def inorder(root):
    if not root:
        return []
    return inorder(root.left) + [root.val] + inorder(root.right)
```

### 3. Post-order Traversal
- **Order**: Left → Right → Root
- **Use Case**: Tree deletion, calculating tree properties, postfix expressions
```python
def postorder(root):
    if not root:
        return []
    return postorder(root.left) + postorder(root.right) + [root.val]
```

## Time and Space Complexity

### Time Complexity
- **All DFS traversals**: O(n) where n is the number of nodes
- Each node is visited exactly once

### Space Complexity
- **Best Case (balanced tree)**: O(log n) - recursion depth
- **Worst Case (skewed tree)**: O(n) - recursion depth
- **Iterative approach**: O(h) where h is tree height

## Recursive vs Iterative Approaches

### Recursive Approach
**Pros:**
- Clean and intuitive code
- Natural for tree problems
- Easy to implement

**Cons:**
- Stack overflow risk for deep trees
- Higher memory overhead

### Iterative Approach
**Pros:**
- No stack overflow risk
- Better space control
- More efficient for very deep trees

**Cons:**
- More complex implementation
- Less intuitive for some problems

## Template Code

### Recursive DFS Template
```python
def dfs_recursive(root):
    # Base case
    if not root:
        return  # or appropriate base value
    
    # Pre-order processing
    # process(root.val)
    
    # Recursive calls
    left_result = dfs_recursive(root.left)
    
    # In-order processing
    # process(root.val)
    
    right_result = dfs_recursive(root.right)
    
    # Post-order processing
    # Combine results or process current node
    return combined_result
```

### Iterative DFS Template
```python
def dfs_iterative(root):
    if not root:
        return
    
    stack = [root]
    while stack:
        node = stack.pop()
        
        # Process current node
        process(node.val)
        
        # Add children to stack (right first for left-to-right processing)
        if node.right:
            stack.append(node.right)
        if node.left:
            stack.append(node.left)
```

## Common Problem Patterns

### 1. Path Sum (Root to Leaf)
**Problem**: Check if there exists a root-to-leaf path with a given sum.

```python
def hasPathSum(root, targetSum):
    """
    Time: O(n), Space: O(h)
    """
    if not root:
        return False
    
    # Leaf node check
    if not root.left and not root.right:
        return root.val == targetSum
    
    # Check left and right subtrees with updated sum
    remaining_sum = targetSum - root.val
    return (hasPathSum(root.left, remaining_sum) or 
            hasPathSum(root.right, remaining_sum))

# Example usage:
# Tree: [5,4,8,11,null,13,4,7,2,null,null,null,1]
# Target: 22
# Expected: True (path: 5->4->11->2)
```

### 2. All Root-to-Leaf Paths
**Problem**: Return all root-to-leaf paths in a binary tree.

```python
def binaryTreePaths(root):
    """
    Time: O(n), Space: O(n)
    """
    def dfs(node, path, result):
        if not node:
            return
        
        path.append(str(node.val))
        
        # Leaf node - add path to result
        if not node.left and not node.right:
            result.append("->".join(path))
        else:
            # Continue DFS
            dfs(node.left, path, result)
            dfs(node.right, path, result)
        
        # Backtrack
        path.pop()
    
    result = []
    dfs(root, [], result)
    return result

# Example usage:
# Tree: [1,2,3,null,5]
# Expected: ["1->2->5", "1->3"]
```

### 3. Path Sum II (Specific Sum Paths)
**Problem**: Find all root-to-leaf paths where each path's sum equals the target sum.

```python
def pathSum(root, targetSum):
    """
    Time: O(n), Space: O(n)
    """
    def dfs(node, remaining_sum, current_path, all_paths):
        if not node:
            return
        
        current_path.append(node.val)
        
        # Check if it's a leaf and sum matches
        if not node.left and not node.right and remaining_sum == node.val:
            all_paths.append(current_path[:])  # Make a copy
        else:
            # Continue searching in subtrees
            new_remaining = remaining_sum - node.val
            dfs(node.left, new_remaining, current_path, all_paths)
            dfs(node.right, new_remaining, current_path, all_paths)
        
        # Backtrack
        current_path.pop()
    
    result = []
    dfs(root, targetSum, [], result)
    return result

# Example usage:
# Tree: [5,4,8,11,null,13,4,7,2,null,null,5,1]
# Target: 22
# Expected: [[5,4,11,2], [5,8,4,5]]
```

### 4. Tree Diameter
**Problem**: Find the diameter of a binary tree (longest path between any two nodes).

```python
def diameterOfBinaryTree(root):
    """
    Time: O(n), Space: O(h)
    """
    def dfs(node):
        if not node:
            return 0
        
        # Get depths of left and right subtrees
        left_depth = dfs(node.left)
        right_depth = dfs(node.right)
        
        # Update diameter (path through current node)
        self.diameter = max(self.diameter, left_depth + right_depth)
        
        # Return depth of current subtree
        return max(left_depth, right_depth) + 1
    
    self.diameter = 0
    dfs(root)
    return self.diameter

# Alternative implementation using nonlocal
def diameterOfBinaryTree(root):
    diameter = 0
    
    def dfs(node):
        nonlocal diameter
        if not node:
            return 0
        
        left_depth = dfs(node.left)
        right_depth = dfs(node.right)
        
        diameter = max(diameter, left_depth + right_depth)
        return max(left_depth, right_depth) + 1
    
    dfs(root)
    return diameter
```

### 5. Maximum Path Sum (Binary Tree)
**Problem**: Find the maximum sum of any path in a binary tree.

```python
def maxPathSum(root):
    """
    Time: O(n), Space: O(h)
    """
    def dfs(node):
        nonlocal max_sum
        if not node:
            return 0
        
        # Get max contribution from left and right subtrees
        # Use max with 0 to ignore negative paths
        left_gain = max(dfs(node.left), 0)
        right_gain = max(dfs(node.right), 0)
        
        # Path sum through current node
        current_max = node.val + left_gain + right_gain
        max_sum = max(max_sum, current_max)
        
        # Return max gain for parent (can only use one side)
        return node.val + max(left_gain, right_gain)
    
    max_sum = float('-inf')
    dfs(root)
    return max_sum

# Example usage:
# Tree: [1,2,3]
# Expected: 6 (path: 2->1->3)
```

### 6. Serialize and Deserialize Binary Tree
**Problem**: Design an algorithm to serialize and deserialize a binary tree.

```python
class Codec:
    def serialize(self, root):
        """
        Time: O(n), Space: O(n)
        """
        def dfs(node):
            if not node:
                return "null"
            return str(node.val) + "," + dfs(node.left) + "," + dfs(node.right)
        
        return dfs(root)
    
    def deserialize(self, data):
        """
        Time: O(n), Space: O(n)
        """
        def dfs():
            val = next(values)
            if val == "null":
                return None
            
            node = TreeNode(int(val))
            node.left = dfs()
            node.right = dfs()
            return node
        
        values = iter(data.split(","))
        return dfs()

# Example usage:
# Tree: [1,2,3,null,null,4,5]
# Serialized: "1,2,null,null,3,4,null,null,5,null,null"
```

### 7. Sum of Path Numbers
**Problem**: Find sum of all numbers formed by root-to-leaf paths.

```python
def sumNumbers(root):
    """
    Time: O(n), Space: O(h)
    """
    def dfs(node, current_number):
        if not node:
            return 0
        
        current_number = current_number * 10 + node.val
        
        # Leaf node
        if not node.left and not node.right:
            return current_number
        
        # Sum from both subtrees
        return dfs(node.left, current_number) + dfs(node.right, current_number)
    
    return dfs(root, 0)

# Example usage:
# Tree: [1,2,3]
# Paths: 1->2 (12), 1->3 (13)
# Expected: 25
```

### 8. Path with Given Sequence
**Problem**: Check if there exists a root-to-leaf path that matches a given sequence.

```python
def isValidSequence(root, arr):
    """
    Time: O(n), Space: O(h)
    """
    def dfs(node, index):
        if not node or index >= len(arr):
            return False
        
        # Check if current node matches sequence
        if node.val != arr[index]:
            return False
        
        # If leaf node, check if we've consumed entire sequence
        if not node.left and not node.right:
            return index == len(arr) - 1
        
        # Continue with next index
        return dfs(node.left, index + 1) or dfs(node.right, index + 1)
    
    return dfs(root, 0)

# Example usage:
# Tree: [1,2,3,4,5,6,7]
# Sequence: [1,2,5]
# Expected: True
```

## Practice Problems

### Easy
1. **Maximum Depth of Binary Tree** - Find the maximum depth of a binary tree
2. **Same Tree** - Check if two binary trees are the same
3. **Symmetric Tree** - Check if a tree is symmetric around its center
4. **Path Sum** - Check if root-to-leaf path sum equals target
5. **Minimum Depth of Binary Tree** - Find minimum depth to a leaf node

### Medium
1. **Binary Tree Paths** - Find all root-to-leaf paths
2. **Path Sum II** - Find all root-to-leaf paths with target sum
3. **Sum Root to Leaf Numbers** - Sum all numbers formed by paths
4. **Binary Tree Maximum Path Sum** - Find maximum path sum
5. **Diameter of Binary Tree** - Find the longest path between any nodes
6. **Count Univalue Subtrees** - Count subtrees with same values
7. **Path Sum III** - Count paths (not necessarily root-to-leaf) with target sum

### Hard
1. **Serialize and Deserialize Binary Tree** - Convert tree to string and back
2. **Binary Tree Maximum Path Sum** (with negative values)
3. **Recover Binary Search Tree** - Fix two swapped nodes in BST
4. **Tree Diameter** - Multiple variations and optimizations

## Tips and Common Mistakes

### Tips
1. **Think Recursively**: Break problems into smaller subproblems
2. **Handle Base Cases**: Always handle null nodes and leaf nodes
3. **State Management**: Decide whether to pass state down or return up
4. **Backtracking**: Remember to clean up when exploring multiple paths
5. **Edge Cases**: Consider empty trees, single nodes, skewed trees

### Common Mistakes
1. **Forgetting Base Cases**: Not handling null nodes properly
2. **Stack Overflow**: Deep recursion without considering iterative solutions
3. **Incorrect Leaf Detection**: `not node.left and not node.right`
4. **Path Copying**: Forgetting to copy paths when storing results
5. **State Pollution**: Not properly cleaning up shared state between recursive calls
6. **Integer Overflow**: Not considering large path sums

### Debugging Strategies
1. **Trace Execution**: Walk through small examples by hand
2. **Print Debug Info**: Add prints for node values and current state
3. **Test Edge Cases**: Empty tree, single node, skewed trees
4. **Verify Base Cases**: Ensure proper handling of null/leaf nodes

### Optimization Techniques
1. **Early Termination**: Return early when answer is found
2. **Memoization**: Cache results for overlapping subproblems
3. **Iterative Solutions**: Use stack to avoid recursion overhead
4. **Space Optimization**: Reuse data structures when possible

---

## Summary

Tree DFS is a powerful pattern for solving tree-related problems in coding interviews. The key is to:

1. **Identify the traversal type** needed (pre/in/post-order)
2. **Choose recursive vs iterative** based on constraints
3. **Handle base cases** properly
4. **Manage state** effectively (parameters vs return values)
5. **Consider edge cases** and optimization opportunities

Master these patterns and you'll be well-prepared for tree problems in coding interviews!