"""
Graph Traversal Pattern Examples

This module demonstrates common graph traversal patterns including
BFS, DFS, and their applications to solve various problems.
"""

from collections import deque, defaultdict
from typing import List, Set, Dict, Optional


class GraphTraversal:
    """Graph traversal algorithms and applications."""
    
    @staticmethod
    def bfs(graph: Dict[int, List[int]], start: int) -> List[int]:
        """Basic BFS traversal."""
        visited = set()
        queue = deque([start])
        result = []
        
        while queue:
            vertex = queue.popleft()
            if vertex not in visited:
                visited.add(vertex)
                result.append(vertex)
                
                for neighbor in graph[vertex]:
                    if neighbor not in visited:
                        queue.append(neighbor)
        
        return result
    
    @staticmethod
    def dfs_recursive(graph: Dict[int, List[int]], start: int, visited: Set[int] = None) -> List[int]:
        """Recursive DFS traversal."""
        if visited is None:
            visited = set()
        
        visited.add(start)
        result = [start]
        
        for neighbor in graph[start]:
            if neighbor not in visited:
                result.extend(GraphTraversal.dfs_recursive(graph, neighbor, visited))
        
        return result
    
    @staticmethod
    def dfs_iterative(graph: Dict[int, List[int]], start: int) -> List[int]:
        """Iterative DFS traversal."""
        visited = set()
        stack = [start]
        result = []
        
        while stack:
            vertex = stack.pop()
            if vertex not in visited:
                visited.add(vertex)
                result.append(vertex)
                
                for neighbor in reversed(graph[vertex]):
                    if neighbor not in visited:
                        stack.append(neighbor)
        
        return result


class IslandProblems:
    """Problems related to finding islands/connected components."""
    
    @staticmethod
    def num_islands(grid: List[List[str]]) -> int:
        """Count number of islands in a 2D grid."""
        if not grid or not grid[0]:
            return 0
        
        rows, cols = len(grid), len(grid[0])
        islands = 0
        
        def dfs(r: int, c: int) -> None:
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
    
    @staticmethod
    def max_area_of_island(grid: List[List[int]]) -> int:
        """Find the maximum area of an island."""
        if not grid or not grid[0]:
            return 0
        
        rows, cols = len(grid), len(grid[0])
        max_area = 0
        
        def dfs(r: int, c: int) -> int:
            if (r < 0 or r >= rows or c < 0 or c >= cols or 
                grid[r][c] != 1):
                return 0
            
            grid[r][c] = 0  # Mark as visited
            area = 1
            
            # Explore all 4 directions
            for dr, dc in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
                area += dfs(r + dr, c + dc)
            
            return area
        
        for r in range(rows):
            for c in range(cols):
                if grid[r][c] == 1:
                    max_area = max(max_area, dfs(r, c))
        
        return max_area


class PathFinding:
    """Path finding algorithms using graph traversal."""
    
    @staticmethod
    def word_ladder_length(begin_word: str, end_word: str, word_list: List[str]) -> int:
        """Find shortest transformation sequence length."""
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
    
    @staticmethod
    def shortest_bridge(grid: List[List[int]]) -> int:
        """Find shortest bridge between two islands."""
        rows, cols = len(grid), len(grid[0])
        
        def dfs(r: int, c: int, island: List[tuple]) -> None:
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


class CycleDetection:
    """Cycle detection algorithms."""
    
    @staticmethod
    def can_finish_courses(num_courses: int, prerequisites: List[List[int]]) -> bool:
        """Detect if course schedule has cycles."""
        graph = defaultdict(list)
        
        for course, prereq in prerequisites:
            graph[prereq].append(course)
        
        # 0: unvisited, 1: visiting, 2: visited
        color = [0] * num_courses
        
        def has_cycle(node: int) -> bool:
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


# Example usage and test cases
if __name__ == "__main__":
    # Test graph traversal
    graph = {
        0: [1, 2],
        1: [0, 3, 4],
        2: [0, 5, 6],
        3: [1],
        4: [1],
        5: [2],
        6: [2]
    }
    
    gt = GraphTraversal()
    print("BFS from 0:", gt.bfs(graph, 0))
    print("DFS (recursive) from 0:", gt.dfs_recursive(graph, 0))
    print("DFS (iterative) from 0:", gt.dfs_iterative(graph, 0))
    
    # Test island problems
    islands = IslandProblems()
    
    grid1 = [
        ["1","1","1","1","0"],
        ["1","1","0","1","0"],
        ["1","1","0","0","0"],
        ["0","0","0","0","0"]
    ]
    print("Number of islands:", islands.num_islands(grid1))
    
    grid2 = [
        [0,0,1,0,0,0,0,1,0,0,0,0,0],
        [0,0,0,0,0,0,0,1,1,1,0,0,0],
        [0,1,1,0,1,0,0,0,0,0,0,0,0],
        [0,1,0,0,1,1,0,0,1,0,1,0,0],
        [0,1,0,0,1,1,0,0,1,1,1,0,0],
        [0,0,0,0,0,0,0,0,0,0,1,0,0],
        [0,0,0,0,0,0,0,1,1,1,0,0,0],
        [0,0,0,0,0,0,0,1,1,0,0,0,0]
    ]
    print("Max area of island:", islands.max_area_of_island(grid2))
    
    # Test path finding
    pf = PathFinding()
    
    begin_word = "hit"
    end_word = "cog"
    word_list = ["hot","dot","dog","lot","log","cog"]
    print("Word ladder length:", pf.word_ladder_length(begin_word, end_word, word_list))
    
    bridge_grid = [[0,1],[1,0]]
    print("Shortest bridge:", pf.shortest_bridge(bridge_grid))
    
    # Test cycle detection
    cd = CycleDetection()
    
    num_courses = 2
    prerequisites = [[1, 0]]
    print("Can finish courses:", cd.can_finish_courses(num_courses, prerequisites))
    
    prerequisites_with_cycle = [[1, 0], [0, 1]]
    print("Can finish courses (with cycle):", cd.can_finish_courses(num_courses, prerequisites_with_cycle))