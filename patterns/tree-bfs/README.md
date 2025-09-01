# Tree BFS (Breadth-First Search) Pattern

## Table of Contents
1. [Overview](#overview)
2. [When to Use Tree BFS](#when-to-use-tree-bfs)
3. [Time and Space Complexity](#time-and-space-complexity)
4. [Core Concepts](#core-concepts)
5. [Basic Template](#basic-template)
6. [Common Problem Patterns](#common-problem-patterns)
7. [Practice Problems](#practice-problems)
8. [Tips, Tricks and Common Mistakes](#tips-tricks-and-common-mistakes)
9. [Reusable Template Code](#reusable-template-code)

## Overview

Tree BFS (Breadth-First Search) is a tree traversal technique that explores nodes level by level, from left to right. It uses a queue data structure to maintain the order of nodes to be processed. This pattern is particularly useful when you need to process tree nodes level by level or when the problem requires understanding the tree's structure in terms of levels.

### Key Characteristics:
- **Level-by-level traversal**: Processes all nodes at depth `d` before any node at depth `d+1`
- **Queue-based**: Uses FIFO (First In, First Out) principle
- **Iterative approach**: Typically implemented using loops rather than recursion
- **Memory intensive**: Requires storing all nodes at the current level

## When to Use Tree BFS

Use Tree BFS when:
- **Level-order processing**: Need to process nodes level by level
- **Shortest path in unweighted trees**: Finding minimum depth or shortest path
- **Level-specific operations**: Need to perform operations on entire levels
- **Tree structure analysis**: Analyzing tree width, finding nodes at specific levels
- **Zigzag or reverse level traversal**: Special level-order requirements

### Problem Indicators:
- Keywords like "level", "depth", "layer", "minimum", "shortest"
- Need to process siblings before children
- Tree visualization or printing requirements
- Finding nodes at specific distances from root

## Time and Space Complexity

### Time Complexity: O(N)
- Must visit each node exactly once
- N = number of nodes in the tree

### Space Complexity: O(W)
- W = maximum width of the tree (maximum nodes at any level)
- In worst case (complete binary tree): O(N/2) = O(N)
- In best case (skewed tree): O(1)

## Core Concepts

### 1. Queue Operations
```python
from collections import deque

# Initialize queue with root
queue = deque([root])

# Process nodes level by level
while queue:
    node = queue.popleft()  # Remove from front
    # Process current node
    
    # Add children to queue
    if node.left:
        queue.append(node.left)
    if node.right:
        queue.append(node.right)
```

### 2. Level Tracking
```python
# Method 1: Process entire level at once
while queue:
    level_size = len(queue)
    current_level = []
    
    for _ in range(level_size):
        node = queue.popleft()
        current_level.append(node.val)
        
        if node.left:
            queue.append(node.left)
        if node.right:
            queue.append(node.right)
```

### 3. Level Counting
```python
# Method 2: Track level with tuple
queue = deque([(root, 0)])  # (node, level)

while queue:
    node, level = queue.popleft()
    # Process node at specific level
    
    if node.left:
        queue.append((node.left, level + 1))
    if node.right:
        queue.append((node.right, level + 1))
```

## Basic Template

```python
from collections import deque

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

def bfs_template(root):
    if not root:
        return []
    
    result = []
    queue = deque([root])
    
    while queue:
        level_size = len(queue)
        current_level = []
        
        for _ in range(level_size):
            node = queue.popleft()
            current_level.append(node.val)
            
            # Add children to queue
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)
        
        result.append(current_level)
    
    return result
```

## Common Problem Patterns

### 1. Level Order Traversal (Basic BFS)

**Problem**: Return the level order traversal of a binary tree's nodes' values.

```python
def levelOrder(root):
    """
    Time: O(N), Space: O(W) where W is max width
    """
    if not root:
        return []
    
    result = []
    queue = deque([root])
    
    while queue:
        level_size = len(queue)
        current_level = []
        
        for _ in range(level_size):
            node = queue.popleft()
            current_level.append(node.val)
            
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)
        
        result.append(current_level)
    
    return result

# Example:
#     3
#    / \
#   9  20
#     /  \
#    15   7
# Output: [[3], [9, 20], [15, 7]]
```

### 2. Binary Tree Right Side View

**Problem**: Return the values of the nodes you can see when looking at the tree from the right side.

```python
def rightSideView(root):
    """
    Key insight: The rightmost node at each level
    Time: O(N), Space: O(W)
    """
    if not root:
        return []
    
    result = []
    queue = deque([root])
    
    while queue:
        level_size = len(queue)
        
        for i in range(level_size):
            node = queue.popleft()
            
            # If it's the last node in the level, add to result
            if i == level_size - 1:
                result.append(node.val)
            
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)
    
    return result

# Example:
#     1
#    / \
#   2   3
#    \   \
#     5   4
# Output: [1, 3, 4]
```

### 3. Average of Levels in Binary Tree

**Problem**: Return the average value of the nodes on each level.

```python
def averageOfLevels(root):
    """
    Time: O(N), Space: O(W)
    """
    if not root:
        return []
    
    result = []
    queue = deque([root])
    
    while queue:
        level_size = len(queue)
        level_sum = 0
        
        for _ in range(level_size):
            node = queue.popleft()
            level_sum += node.val
            
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)
        
        result.append(level_sum / level_size)
    
    return result

# Example:
#     3
#    / \
#   9  20
#     /  \
#    15   7
# Output: [3.0, 14.5, 11.0]
```

### 4. Minimum Depth of Binary Tree

**Problem**: Find the minimum depth (shortest path from root to any leaf).

```python
def minDepth(root):
    """
    Key insight: Stop at first leaf node encountered
    Time: O(N) worst case, Space: O(W)
    """
    if not root:
        return 0
    
    queue = deque([(root, 1)])  # (node, depth)
    
    while queue:
        node, depth = queue.popleft()
        
        # If leaf node, return depth (first leaf = minimum depth)
        if not node.left and not node.right:
            return depth
        
        if node.left:
            queue.append((node.left, depth + 1))
        if node.right:
            queue.append((node.right, depth + 1))
    
    return 0

# Alternative approach tracking levels
def minDepthLevelByLevel(root):
    if not root:
        return 0
    
    queue = deque([root])
    depth = 0
    
    while queue:
        depth += 1
        level_size = len(queue)
        
        for _ in range(level_size):
            node = queue.popleft()
            
            # If leaf node, return current depth
            if not node.left and not node.right:
                return depth
            
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)
    
    return depth
```

### 5. Level Order Traversal II (Bottom-Up)

**Problem**: Return the bottom-up level order traversal of a binary tree.

```python
def levelOrderBottom(root):
    """
    Time: O(N), Space: O(W)
    """
    if not root:
        return []
    
    result = []
    queue = deque([root])
    
    while queue:
        level_size = len(queue)
        current_level = []
        
        for _ in range(level_size):
            node = queue.popleft()
            current_level.append(node.val)
            
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)
        
        result.append(current_level)
    
    # Reverse the result for bottom-up order
    return result[::-1]

# Example:
#     3
#    / \
#   9  20
#     /  \
#    15   7
# Output: [[15, 7], [9, 20], [3]]
```

### 6. Zigzag Level Order Traversal

**Problem**: Return the zigzag level order traversal (alternating left-to-right and right-to-left).

```python
def zigzagLevelOrder(root):
    """
    Time: O(N), Space: O(W)
    """
    if not root:
        return []
    
    result = []
    queue = deque([root])
    left_to_right = True
    
    while queue:
        level_size = len(queue)
        current_level = []
        
        for _ in range(level_size):
            node = queue.popleft()
            current_level.append(node.val)
            
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)
        
        # Reverse every other level
        if not left_to_right:
            current_level.reverse()
        
        result.append(current_level)
        left_to_right = not left_to_right
    
    return result

# Alternative using deque for efficient reversal
def zigzagLevelOrderDeque(root):
    if not root:
        return []
    
    result = []
    queue = deque([root])
    left_to_right = True
    
    while queue:
        level_size = len(queue)
        current_level = deque()
        
        for _ in range(level_size):
            node = queue.popleft()
            
            # Add to appropriate end based on direction
            if left_to_right:
                current_level.append(node.val)
            else:
                current_level.appendleft(node.val)
            
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)
        
        result.append(list(current_level))
        left_to_right = not left_to_right
    
    return result

# Example:
#     3
#    / \
#   9  20
#     /  \
#    15   7
# Output: [[3], [20, 9], [15, 7]]
```

### 7. Connect Level Order Siblings (Populating Next Right Pointers)

**Problem**: Connect each node to its next right node in the same level.

```python
class Node:
    def __init__(self, val=0, left=None, right=None, next=None):
        self.val = val
        self.left = left
        self.right = right
        self.next = next

def connect(root):
    """
    Time: O(N), Space: O(W)
    """
    if not root:
        return root
    
    queue = deque([root])
    
    while queue:
        level_size = len(queue)
        
        for i in range(level_size):
            node = queue.popleft()
            
            # Connect to next node in the same level
            if i < level_size - 1:
                node.next = queue[0]  # Peek at next node
            
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)
    
    return root

# Space-optimized approach for perfect binary tree
def connectPerfectBinaryTree(root):
    """
    Time: O(N), Space: O(1) for perfect binary tree
    """
    if not root:
        return root
    
    leftmost = root
    
    while leftmost.left:  # While not leaf level
        head = leftmost
        
        while head:
            # Connect children of current node
            head.left.next = head.right
            
            # Connect to next node's left child
            if head.next:
                head.right.next = head.next.left
            
            head = head.next
        
        leftmost = leftmost.left  # Move to next level
    
    return root
```

## Practice Problems

### Easy
- **Binary Tree Level Order Traversal** (LeetCode 102)
- **Average of Levels in Binary Tree** (LeetCode 637)
- **Minimum Depth of Binary Tree** (LeetCode 111)
- **Maximum Depth of Binary Tree** (LeetCode 104) - can be solved with BFS
- **Same Tree** (LeetCode 100) - level-by-level comparison

### Medium
- **Binary Tree Level Order Traversal II** (LeetCode 107)
- **Binary Tree Zigzag Level Order Traversal** (LeetCode 103)
- **Binary Tree Right Side View** (LeetCode 199)
- **Populating Next Right Pointers in Each Node** (LeetCode 116)
- **Populating Next Right Pointers in Each Node II** (LeetCode 117)
- **Find Largest Value in Each Tree Row** (LeetCode 515)
- **Add One Row to Tree** (LeetCode 623)

### Hard
- **Serialize and Deserialize Binary Tree** (LeetCode 297) - BFS approach
- **Vertical Order Traversal of a Binary Tree** (LeetCode 987)
- **Binary Tree Maximum Path Sum** (LeetCode 124) - can use BFS for level processing

### Pattern Variations
- **N-ary Tree Level Order Traversal** (LeetCode 429)
- **Find Bottom Left Tree Value** (LeetCode 513)
- **Cousins in Binary Tree** (LeetCode 993)
- **Even Odd Tree** (LeetCode 1609)

## Tips, Tricks and Common Mistakes

### Tips
1. **Use `collections.deque`**: More efficient than list for queue operations
2. **Track level size**: Process entire levels at once when needed
3. **Early termination**: For minimum depth problems, return as soon as you find a leaf
4. **Space optimization**: For perfect binary trees, consider O(1) space solutions
5. **Handle edge cases**: Always check for empty tree (`root is None`)

### Common Patterns Recognition
- **"Level by level"** → Standard BFS with level processing
- **"Minimum depth"** → BFS with early termination at first leaf
- **"Right side view"** → Take last element of each level
- **"Zigzag"** → Alternate direction for each level
- **"Connect siblings"** → Process level and connect adjacent nodes

### Common Mistakes
1. **Forgetting to check for None root**: Always handle empty tree case
2. **Queue operations confusion**: Use `popleft()` and `append()` correctly
3. **Level size calculation**: Calculate `len(queue)` before processing level
4. **Memory issues**: BFS can be memory-intensive for wide trees
5. **Modifying queue during iteration**: Calculate level size first
6. **Not handling leaf nodes properly**: Check both left and right children

### Debugging Tips
```python
# Add debug prints to track queue state
def debug_bfs(root):
    if not root:
        return []
    
    queue = deque([root])
    level = 0
    
    while queue:
        print(f"Level {level}: Queue size = {len(queue)}")
        level_size = len(queue)
        current_level = []
        
        for i in range(level_size):
            node = queue.popleft()
            print(f"  Processing: {node.val}")
            current_level.append(node.val)
            
            if node.left:
                queue.append(node.left)
                print(f"    Added left: {node.left.val}")
            if node.right:
                queue.append(node.right)
                print(f"    Added right: {node.right.val}")
        
        print(f"  Level result: {current_level}")
        level += 1
```

## Reusable Template Code

### 1. Basic Level-Order Template
```python
from collections import deque

def level_order_template(root):
    """
    Template for level-by-level processing
    """
    if not root:
        return []
    
    result = []
    queue = deque([root])
    
    while queue:
        level_size = len(queue)
        current_level = []
        
        for _ in range(level_size):
            node = queue.popleft()
            
            # Process current node
            current_level.append(node.val)
            
            # Add children
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)
        
        # Process complete level
        result.append(current_level)
    
    return result
```

### 2. Early Termination Template
```python
def early_termination_template(root, target_condition):
    """
    Template for problems requiring early termination
    """
    if not root:
        return None
    
    queue = deque([(root, 0)])  # (node, level)
    
    while queue:
        node, level = queue.popleft()
        
        # Check termination condition
        if target_condition(node, level):
            return level  # or node, or whatever needed
        
        if node.left:
            queue.append((node.left, level + 1))
        if node.right:
            queue.append((node.right, level + 1))
    
    return -1  # Not found
```

### 3. Level Processing Template
```python
def level_processing_template(root):
    """
    Template for operations on entire levels
    """
    if not root:
        return []
    
    queue = deque([root])
    result = []
    
    while queue:
        level_size = len(queue)
        level_values = []
        
        for i in range(level_size):
            node = queue.popleft()
            level_values.append(node.val)
            
            # Add children
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)
        
        # Process entire level
        level_result = process_level(level_values, i)
        result.append(level_result)
    
    return result

def process_level(values, level_index):
    """Override this function based on specific requirements"""
    return values
```

### 4. Directional BFS Template
```python
def directional_bfs_template(root):
    """
    Template for zigzag or directional traversals
    """
    if not root:
        return []
    
    result = []
    queue = deque([root])
    reverse = False
    
    while queue:
        level_size = len(queue)
        current_level = []
        
        for _ in range(level_size):
            node = queue.popleft()
            current_level.append(node.val)
            
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)
        
        # Apply direction
        if reverse:
            current_level.reverse()
        
        result.append(current_level)
        reverse = not reverse
    
    return result
```

Remember: Tree BFS is your go-to pattern when you need to process trees level by level. It's particularly powerful for shortest path problems, level-specific operations, and tree structure analysis. Master these templates and patterns, and you'll be well-prepared for most tree BFS interview questions!