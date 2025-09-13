# Graph Traversal Pattern

## Overview
Graph traversal is a fundamental technique for visiting all vertices and edges in a graph systematically. The two primary traversal methods are Breadth-First Search (BFS) and Depth-First Search (DFS), each with specific use cases and properties that make them suitable for different types of problems.

## When to Use
- **Path Finding**: Finding paths between vertices
- **Connected Components**: Finding disconnected graph parts
- **Level-order Processing**: Processing nodes level by level
- **Cycle Detection**: Detecting cycles in graphs
- **Shortest Path**: Finding shortest unweighted paths (BFS)
- **Tree/Graph Validation**: Checking graph properties

## Time/Space Complexity
- **Time**: O(V + E) where V is vertices and E is edges
- **Space**: O(V) for visited tracking and auxiliary data structures

## Key Concepts
- **Visited Set**: Track processed vertices to avoid infinite loops
- **Queue (BFS)**: First-in-first-out for level-order traversal
- **Stack (DFS)**: Last-in-first-out for depth-first exploration
- **Adjacency Representation**: List, matrix, or edge list

## Pattern Variations

### 1. Breadth-First Search (BFS)
```python
from collections import deque

def bfs(graph, start):
    visited = set()
    queue = deque([start])
    result = []
    
    while queue:
        vertex = queue.popleft()
        if vertex not in visited:
            visited.add(vertex)
            result.append(vertex)
            
            # Add unvisited neighbors to queue
            for neighbor in graph[vertex]:
                if neighbor not in visited:
                    queue.append(neighbor)
    
    return result

# BFS with level tracking
def bfs_with_levels(graph, start):
    visited = set([start])
    queue = deque([(start, 0)])  # (node, level)
    levels = {}
    
    while queue:
        vertex, level = queue.popleft()
        levels[vertex] = level
        
        for neighbor in graph[vertex]:
            if neighbor not in visited:
                visited.add(neighbor)
                queue.append((neighbor, level + 1))
    
    return levels
```

### 2. Depth-First Search (DFS)
```python
def dfs_recursive(graph, start, visited=None):
    if visited is None:
        visited = set()
    
    visited.add(start)
    result = [start]
    
    for neighbor in graph[start]:
        if neighbor not in visited:
            result.extend(dfs_recursive(graph, neighbor, visited))
    
    return result

def dfs_iterative(graph, start):
    visited = set()
    stack = [start]
    result = []
    
    while stack:
        vertex = stack.pop()
        if vertex not in visited:
            visited.add(vertex)
            result.append(vertex)
            
            # Add neighbors to stack (reverse order for consistent results)
            for neighbor in reversed(graph[vertex]):
                if neighbor not in visited:
                    stack.append(neighbor)
    
    return result
```

## Common Problem Patterns

### Pattern 1: Number of Islands
**Problem**: Count connected components in a 2D grid.

```python
def num_islands(grid):
    if not grid or not grid[0]:
        return 0
    
    rows, cols = len(grid), len(grid[0])
    islands = 0
    
    def dfs(r, c):
        if (r < 0 or r >= rows or c < 0 or c >= cols or 
            grid[r][c] != '1'):
            return
        
        grid[r][c] = '0'  # Mark as visited
        
        # Explore all 4 directions
        dfs(r + 1, c)
        dfs(r - 1, c)
        dfs(r, c + 1)
        dfs(r, c - 1)
    
    for r in range(rows):
        for c in range(cols):
            if grid[r][c] == '1':
                islands += 1
                dfs(r, c)
    
    return islands

# Example usage
grid = [
    ["1","1","1","1","0"],
    ["1","1","0","1","0"],
    ["1","1","0","0","0"],
    ["0","0","0","0","0"]
]
print(num_islands(grid))  # Output: 1
```

### Pattern 2: Word Ladder
**Problem**: Find shortest transformation sequence between two words.

```python
from collections import deque

def ladder_length(begin_word, end_word, word_list):
    word_set = set(word_list)
    if end_word not in word_set:
        return 0
    
    queue = deque([(begin_word, 1)])
    visited = set([begin_word])
    
    while queue:
        word, level = queue.popleft()
        
        if word == end_word:
            return level
        
        # Try all possible one-character changes
        for i in range(len(word)):
            for char in 'abcdefghijklmnopqrstuvwxyz':
                new_word = word[:i] + char + word[i+1:]
                
                if new_word in word_set and new_word not in visited:
                    visited.add(new_word)
                    queue.append((new_word, level + 1))
    
    return 0

# Example usage
begin_word = "hit"
end_word = "cog"
word_list = ["hot","dot","dog","lot","log","cog"]
print(ladder_length(begin_word, end_word, word_list))  # Output: 5
```

### Pattern 3: Clone Graph
**Problem**: Create a deep copy of an undirected graph.

```python
from collections import deque

class Node:
    def __init__(self, val=0, neighbors=None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []

def clone_graph_bfs(node):
    if not node:
        return None
    
    visited = {}
    queue = deque([node])
    visited[node] = Node(node.val)
    
    while queue:
        current = queue.popleft()
        
        for neighbor in current.neighbors:
            if neighbor not in visited:
                visited[neighbor] = Node(neighbor.val)
                queue.append(neighbor)
            
            visited[current].neighbors.append(visited[neighbor])
    
    return visited[node]

def clone_graph_dfs(node):
    if not node:
        return None
    
    visited = {}
    
    def dfs(node):
        if node in visited:
            return visited[node]
        
        clone = Node(node.val)
        visited[node] = clone
        
        for neighbor in node.neighbors:
            clone.neighbors.append(dfs(neighbor))
        
        return clone
    
    return dfs(node)
```

### Pattern 4: Pacific Atlantic Water Flow
**Problem**: Find cells where water can flow to both oceans.

```python
def pacific_atlantic(heights):
    if not heights or not heights[0]:
        return []
    
    rows, cols = len(heights), len(heights[0])
    pacific = set()
    atlantic = set()
    
    def dfs(r, c, visited, prev_height):
        if (r < 0 or r >= rows or c < 0 or c >= cols or
            (r, c) in visited or heights[r][c] < prev_height):
            return
        
        visited.add((r, c))
        
        # Explore all 4 directions
        for dr, dc in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
            dfs(r + dr, c + dc, visited, heights[r][c])
    
    # Start DFS from all border cells
    for r in range(rows):
        dfs(r, 0, pacific, 0)        # Pacific (left border)
        dfs(r, cols - 1, atlantic, 0)  # Atlantic (right border)
    
    for c in range(cols):
        dfs(0, c, pacific, 0)        # Pacific (top border)
        dfs(rows - 1, c, atlantic, 0)  # Atlantic (bottom border)
    
    return [[r, c] for r, c in pacific & atlantic]

# Example usage
heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]
print(pacific_atlantic(heights))  # Output: [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]
```

### Pattern 5: Course Schedule (Cycle Detection)
**Problem**: Detect cycles using DFS with coloring.

```python
from collections import defaultdict

def can_finish_dfs(num_courses, prerequisites):
    graph = defaultdict(list)
    
    for course, prereq in prerequisites:
        graph[prereq].append(course)
    
    # 0: unvisited, 1: visiting, 2: visited
    color = [0] * num_courses
    
    def has_cycle(node):
        if color[node] == 1:  # Back edge found (cycle)
            return True
        if color[node] == 2:  # Already processed
            return False
        
        color[node] = 1  # Mark as visiting
        
        for neighbor in graph[node]:
            if has_cycle(neighbor):
                return True
        
        color[node] = 2  # Mark as visited
        return False
    
    for i in range(num_courses):
        if color[i] == 0 and has_cycle(i):
            return False
    
    return True

# Example usage
num_courses = 2
prerequisites = [[1, 0]]
print(can_finish_dfs(num_courses, prerequisites))  # Output: True
```

### Pattern 6: Shortest Bridge
**Problem**: Find shortest path between two islands.

```python
from collections import deque

def shortest_bridge(grid):
    rows, cols = len(grid), len(grid[0])
    
    def dfs(r, c, island):
        if (r < 0 or r >= rows or c < 0 or c >= cols or 
            grid[r][c] != 1):
            return
        
        grid[r][c] = 2  # Mark as part of first island
        island.append((r, c))
        
        for dr, dc in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
            dfs(r + dr, c + dc, island)
    
    # Find first island and mark it
    first_island = []
    found = False
    for r in range(rows):
        if found:
            break
        for c in range(cols):
            if grid[r][c] == 1:
                dfs(r, c, first_island)
                found = True
                break
    
    # BFS to find shortest path to second island
    queue = deque([(r, c, 0) for r, c in first_island])
    
    while queue:
        r, c, dist = queue.popleft()
        
        for dr, dc in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
            nr, nc = r + dr, c + dc
            
            if (0 <= nr < rows and 0 <= nc < cols):
                if grid[nr][nc] == 1:  # Found second island
                    return dist
                elif grid[nr][nc] == 0:  # Water
                    grid[nr][nc] = 2  # Mark as visited
                    queue.append((nr, nc, dist + 1))
    
    return -1

# Example usage
grid = [[0,1],[1,0]]
print(shortest_bridge(grid))  # Output: 1
```

### Pattern 7: Surrounded Regions
**Problem**: Capture regions surrounded by 'X'.

```python
def solve(board):
    if not board or not board[0]:
        return
    
    rows, cols = len(board), len(board[0])
    
    def dfs(r, c):
        if (r < 0 or r >= rows or c < 0 or c >= cols or 
            board[r][c] != 'O'):
            return
        
        board[r][c] = 'T'  # Temporary mark
        
        for dr, dc in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
            dfs(r + dr, c + dc)
    
    # Mark all 'O's connected to borders as 'T'
    for r in range(rows):
        if board[r][0] == 'O':
            dfs(r, 0)
        if board[r][cols - 1] == 'O':
            dfs(r, cols - 1)
    
    for c in range(cols):
        if board[0][c] == 'O':
            dfs(0, c)
        if board[rows - 1][c] == 'O':
            dfs(rows - 1, c)
    
    # Convert remaining 'O's to 'X' and 'T's back to 'O'
    for r in range(rows):
        for c in range(cols):
            if board[r][c] == 'O':
                board[r][c] = 'X'
            elif board[r][c] == 'T':
                board[r][c] = 'O'

# Example usage
board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
solve(board)
print(board)  # Modified in-place
```

## Practice Problems

### Easy
1. **Number of Islands** - Connected components with DFS/BFS
2. **Max Area of Island** - Find largest connected component
3. **Flood Fill** - Basic DFS/BFS application
4. **Employee Importance** - Tree traversal

### Medium
1. **Word Ladder** - Shortest path with BFS
2. **Clone Graph** - Deep copy with traversal
3. **Pacific Atlantic Water Flow** - Multi-source traversal
4. **Surrounded Regions** - Border-connected regions
5. **Course Schedule** - Cycle detection with DFS
6. **Number of Provinces** - Connected components

### Hard
1. **Word Ladder II** - All shortest paths
2. **Shortest Bridge** - BFS for shortest path
3. **Remove Invalid Parentheses** - BFS with state space
4. **Alien Dictionary** - Topological sort with DFS

## Tips and Tricks

1. **Choose Right Approach**: BFS for shortest paths, DFS for exhaustive search
2. **Visited Tracking**: Use sets, arrays, or modify input for space efficiency
3. **Level Processing**: Use queue size to process nodes level by level
4. **State Representation**: Include additional state in queue/stack when needed
5. **Early Termination**: Return immediately when target is found

## Common Mistakes

1. **Infinite Loops**: Forgetting to mark nodes as visited
2. **Boundary Checks**: Not validating coordinates in grid problems
3. **State Modification**: Modifying shared state incorrectly in recursion
4. **Queue vs Stack**: Using wrong data structure for the problem type
5. **Memory Issues**: Not optimizing space for large graphs

## Related Patterns

- **Topological Sort**: DFS-based ordering of vertices
- **Union Find**: Alternative for connectivity problems
- **Shortest Path**: Dijkstra's and Bellman-Ford algorithms
- **Tree Traversal**: Specialized graph traversal for trees

## Implementation Languages

Graph traversal implementations across languages:
- **Python**: Use `collections.deque` for BFS, recursion/stack for DFS
- **Java**: Use `Queue`, `Stack`, and proper collection interfaces
- **JavaScript**: Use arrays as stacks/queues with proper methods
- **C++**: Use `queue`, `stack`, and STL containers
- **Go**: Use slices and channels for concurrent traversal