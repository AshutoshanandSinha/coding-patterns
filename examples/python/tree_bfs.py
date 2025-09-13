"""
Tree BFS (Breadth-First Search) Pattern Examples

This module demonstrates BFS traversal patterns for tree problems,
including level-order traversal and tree manipulation using queues.
"""

from collections import deque
from typing import List, Optional


class TreeNode:
    """Binary tree node definition."""
    
    def __init__(self, val: int = 0, left: 'TreeNode' = None, right: 'TreeNode' = None):
        self.val = val
        self.left = left
        self.right = right


class LevelOrderTraversal:
    """Level-order traversal patterns."""
    
    @staticmethod
    def level_order(root: Optional[TreeNode]) -> List[List[int]]:
        """Basic level-order traversal."""
        if not root:
            return []
        
        result = []
        queue = deque([root])
        
        while queue:
            level_size = len(queue)
            level_nodes = []
            
            for _ in range(level_size):
                node = queue.popleft()
                level_nodes.append(node.val)
                
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)
            
            result.append(level_nodes)
        
        return result
    
    @staticmethod
    def level_order_bottom(root: Optional[TreeNode]) -> List[List[int]]:
        """Level-order traversal from bottom up."""
        if not root:
            return []
        
        result = []
        queue = deque([root])
        
        while queue:
            level_size = len(queue)
            level_nodes = []
            
            for _ in range(level_size):
                node = queue.popleft()
                level_nodes.append(node.val)
                
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)
            
            result.append(level_nodes)
        
        return result[::-1]  # Reverse the result
    
    @staticmethod
    def zigzag_level_order(root: Optional[TreeNode]) -> List[List[int]]:
        """Zigzag level-order traversal."""
        if not root:
            return []
        
        result = []
        queue = deque([root])
        left_to_right = True
        
        while queue:
            level_size = len(queue)
            level_nodes = []
            
            for _ in range(level_size):
                node = queue.popleft()
                level_nodes.append(node.val)
                
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)
            
            if not left_to_right:
                level_nodes.reverse()
            
            result.append(level_nodes)
            left_to_right = not left_to_right
        
        return result


class TreeLevelProblems:
    """Problems involving tree levels."""
    
    @staticmethod
    def right_side_view(root: Optional[TreeNode]) -> List[int]:
        """Get right side view of binary tree."""
        if not root:
            return []
        
        result = []
        queue = deque([root])
        
        while queue:
            level_size = len(queue)
            
            for i in range(level_size):
                node = queue.popleft()
                
                # Last node in level is rightmost
                if i == level_size - 1:
                    result.append(node.val)
                
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)
        
        return result
    
    @staticmethod
    def average_of_levels(root: Optional[TreeNode]) -> List[float]:
        """Calculate average of each level."""
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
    
    @staticmethod
    def level_with_minimum_nodes(root: Optional[TreeNode]) -> int:
        """Find level with minimum number of nodes."""
        if not root:
            return -1
        
        min_level = 0
        min_count = float('inf')
        queue = deque([root])
        level = 0
        
        while queue:
            level_size = len(queue)
            
            if level_size < min_count:
                min_count = level_size
                min_level = level
            
            for _ in range(level_size):
                node = queue.popleft()
                
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)
            
            level += 1
        
        return min_level


class TreeConstruction:
    """Tree construction using BFS."""
    
    @staticmethod
    def connect_next_right_pointers(root: Optional['Node']) -> Optional['Node']:
        """Connect nodes to their next right node."""
        if not root:
            return root
        
        queue = deque([root])
        
        while queue:
            level_size = len(queue)
            prev = None
            
            for i in range(level_size):
                node = queue.popleft()
                
                if prev:
                    prev.next = node
                prev = node
                
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)
        
        return root
    
    @staticmethod
    def populate_next_right_pointers_ii(root: Optional['Node']) -> Optional['Node']:
        """Connect next right pointers for any binary tree."""
        if not root:
            return root
        
        queue = deque([root])
        
        while queue:
            level_size = len(queue)
            
            for i in range(level_size):
                node = queue.popleft()
                
                # Connect to next node in level
                if i < level_size - 1:
                    node.next = queue[0]
                
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)
        
        return root


class TreeValidation:
    """Tree validation using BFS."""
    
    @staticmethod
    def is_complete_tree(root: Optional[TreeNode]) -> bool:
        """Check if binary tree is complete."""
        if not root:
            return True
        
        queue = deque([root])
        null_found = False
        
        while queue:
            node = queue.popleft()
            
            if node is None:
                null_found = True
            else:
                if null_found:  # Found node after null
                    return False
                queue.append(node.left)
                queue.append(node.right)
        
        return True
    
    @staticmethod
    def is_symmetric(root: Optional[TreeNode]) -> bool:
        """Check if tree is symmetric."""
        if not root:
            return True
        
        queue = deque([root.left, root.right])
        
        while queue:
            left = queue.popleft()
            right = queue.popleft()
            
            if not left and not right:
                continue
            if not left or not right or left.val != right.val:
                return False
            
            queue.append(left.left)
            queue.append(right.right)
            queue.append(left.right)
            queue.append(right.left)
        
        return True


class TreeMeasurements:
    """Tree measurements using BFS."""
    
    @staticmethod
    def max_depth(root: Optional[TreeNode]) -> int:
        """Find maximum depth of tree."""
        if not root:
            return 0
        
        queue = deque([root])
        depth = 0
        
        while queue:
            depth += 1
            level_size = len(queue)
            
            for _ in range(level_size):
                node = queue.popleft()
                
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)
        
        return depth
    
    @staticmethod
    def min_depth(root: Optional[TreeNode]) -> int:
        """Find minimum depth to leaf node."""
        if not root:
            return 0
        
        queue = deque([root])
        depth = 0
        
        while queue:
            depth += 1
            level_size = len(queue)
            
            for _ in range(level_size):
                node = queue.popleft()
                
                # First leaf node found
                if not node.left and not node.right:
                    return depth
                
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)
        
        return depth
    
    @staticmethod
    def max_width(root: Optional[TreeNode]) -> int:
        """Find maximum width of tree."""
        if not root:
            return 0
        
        max_width = 1
        queue = deque([(root, 0)])  # (node, index)
        
        while queue:
            level_size = len(queue)
            _, first_index = queue[0]
            _, last_index = queue[-1]
            
            max_width = max(max_width, last_index - first_index + 1)
            
            for _ in range(level_size):
                node, index = queue.popleft()
                
                if node.left:
                    queue.append((node.left, 2 * index))
                if node.right:
                    queue.append((node.right, 2 * index + 1))
        
        return max_width


# Helper class for next right pointers problems
class Node:
    def __init__(self, val: int = 0, left: 'Node' = None, right: 'Node' = None, next: 'Node' = None):
        self.val = val
        self.left = left
        self.right = right
        self.next = next


# Example usage and test cases
if __name__ == "__main__":
    # Build a test tree:
    #       3
    #      / \
    #     9   20
    #        /  \
    #       15   7
    root = TreeNode(3)
    root.left = TreeNode(9)
    root.right = TreeNode(20)
    root.right.left = TreeNode(15)
    root.right.right = TreeNode(7)
    
    # Test level-order traversal
    lot = LevelOrderTraversal()
    print("Level order:", lot.level_order(root))
    print("Level order bottom up:", lot.level_order_bottom(root))
    print("Zigzag level order:", lot.zigzag_level_order(root))
    
    # Test tree level problems
    tlp = TreeLevelProblems()
    print("Right side view:", tlp.right_side_view(root))
    print("Average of levels:", tlp.average_of_levels(root))
    print("Level with min nodes:", tlp.level_with_minimum_nodes(root))
    
    # Test tree validation
    tv = TreeValidation()
    print("Is complete tree:", tv.is_complete_tree(root))
    print("Is symmetric:", tv.is_symmetric(root))
    
    # Test tree measurements
    tm = TreeMeasurements()
    print("Max depth:", tm.max_depth(root))
    print("Min depth:", tm.min_depth(root))
    print("Max width:", tm.max_width(root))
    
    # Test with symmetric tree
    symmetric_root = TreeNode(1)
    symmetric_root.left = TreeNode(2)
    symmetric_root.right = TreeNode(2)
    symmetric_root.left.left = TreeNode(3)
    symmetric_root.left.right = TreeNode(4)
    symmetric_root.right.left = TreeNode(4)
    symmetric_root.right.right = TreeNode(3)
    
    print("Is symmetric (symmetric tree):", tv.is_symmetric(symmetric_root))