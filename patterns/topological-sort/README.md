# Topological Sort Pattern

## Overview
Topological Sort is a linear ordering of vertices in a directed acyclic graph (DAG) where for every directed edge (u, v), vertex u comes before vertex v in the ordering. This pattern is essential for solving dependency resolution problems, scheduling tasks, and determining the order of execution when there are prerequisite relationships.

## When to Use
- **Dependency Resolution**: Resolving dependencies between tasks/modules
- **Course Scheduling**: Ordering courses based on prerequisites
- **Build Systems**: Determining compilation order
- **Task Scheduling**: Scheduling tasks with dependencies
- **Cycle Detection**: Detecting cycles in directed graphs

## Time/Space Complexity
- **Time**: O(V + E) where V is vertices and E is edges
- **Space**: O(V + E) for adjacency list and auxiliary structures

## Key Concepts
- **DAG**: Directed Acyclic Graph (required for topological sort)
- **In-degree**: Number of incoming edges to a vertex
- **Prerequisites**: Dependencies that must be satisfied before processing

## Pattern Variations

### 1. Kahn's Algorithm (BFS-based)
```python
from collections import deque, defaultdict

def topological_sort_bfs(vertices, edges):
    # Build adjacency list and in-degree count
    graph = defaultdict(list)
    in_degree = [0] * vertices
    
    for src, dest in edges:
        graph[src].append(dest)
        in_degree[dest] += 1
    
    # Initialize queue with vertices having 0 in-degree
    queue = deque()
    for i in range(vertices):
        if in_degree[i] == 0:
            queue.append(i)
    
    result = []
    
    while queue:
        vertex = queue.popleft()
        result.append(vertex)
        
        # Reduce in-degree of adjacent vertices
        for neighbor in graph[vertex]:
            in_degree[neighbor] -= 1
            if in_degree[neighbor] == 0:
                queue.append(neighbor)
    
    return result if len(result) == vertices else []
```

### 2. DFS-based Approach
```python
def topological_sort_dfs(vertices, edges):
    graph = defaultdict(list)
    
    for src, dest in edges:
        graph[src].append(dest)
    
    visited = [False] * vertices
    stack = []
    
    def dfs(vertex):
        visited[vertex] = True
        
        for neighbor in graph[vertex]:
            if not visited[neighbor]:
                dfs(neighbor)
        
        stack.append(vertex)
    
    # Visit all vertices
    for i in range(vertices):
        if not visited[i]:
            dfs(i)
    
    return stack[::-1]  # Reverse to get correct order
```

## Common Problem Patterns

### Pattern 1: Course Schedule
**Problem**: Determine if it's possible to finish all courses given prerequisites.

```python
from collections import deque, defaultdict

def can_finish(num_courses, prerequisites):
    graph = defaultdict(list)
    in_degree = [0] * num_courses
    
    # Build graph and in-degree array
    for course, prereq in prerequisites:
        graph[prereq].append(course)
        in_degree[course] += 1
    
    # Find courses with no prerequisites
    queue = deque()
    for i in range(num_courses):
        if in_degree[i] == 0:
            queue.append(i)
    
    completed_courses = 0
    
    while queue:
        course = queue.popleft()
        completed_courses += 1
        
        # Remove this course as prerequisite for others
        for next_course in graph[course]:
            in_degree[next_course] -= 1
            if in_degree[next_course] == 0:
                queue.append(next_course)
    
    return completed_courses == num_courses

# Example usage
num_courses = 2
prerequisites = [[1, 0]]
print(can_finish(num_courses, prerequisites))  # Output: True
```

### Pattern 2: Course Schedule II
**Problem**: Return the order of courses to take to finish all courses.

```python
from collections import deque, defaultdict

def find_order(num_courses, prerequisites):
    graph = defaultdict(list)
    in_degree = [0] * num_courses
    
    for course, prereq in prerequisites:
        graph[prereq].append(course)
        in_degree[course] += 1
    
    queue = deque()
    for i in range(num_courses):
        if in_degree[i] == 0:
            queue.append(i)
    
    result = []
    
    while queue:
        course = queue.popleft()
        result.append(course)
        
        for next_course in graph[course]:
            in_degree[next_course] -= 1
            if in_degree[next_course] == 0:
                queue.append(next_course)
    
    return result if len(result) == num_courses else []

# Example usage
num_courses = 4
prerequisites = [[1,0],[2,0],[3,1],[3,2]]
print(find_order(num_courses, prerequisites))  # Output: [0,1,2,3] or [0,2,1,3]
```

### Pattern 3: Alien Dictionary
**Problem**: Derive the order of letters in an alien language from sorted words.

```python
from collections import deque, defaultdict

def alien_order(words):
    # Build initial graph with all characters
    graph = defaultdict(set)
    in_degree = {}
    
    # Initialize in_degree for all characters
    for word in words:
        for char in word:
            in_degree[char] = 0
    
    # Build graph by comparing adjacent words
    for i in range(len(words) - 1):
        first, second = words[i], words[i + 1]
        min_len = min(len(first), len(second))
        
        # Find first differing character
        for j in range(min_len):
            if first[j] != second[j]:
                if second[j] not in graph[first[j]]:
                    graph[first[j]].add(second[j])
                    in_degree[second[j]] += 1
                break
        else:
            # If first word is longer and no difference found, invalid order
            if len(first) > len(second):
                return ""
    
    # Topological sort
    queue = deque()
    for char in in_degree:
        if in_degree[char] == 0:
            queue.append(char)
    
    result = []
    
    while queue:
        char = queue.popleft()
        result.append(char)
        
        for neighbor in graph[char]:
            in_degree[neighbor] -= 1
            if in_degree[neighbor] == 0:
                queue.append(neighbor)
    
    return ''.join(result) if len(result) == len(in_degree) else ""

# Example usage
words = ["wrt", "wrf", "er", "ett", "rftt"]
print(alien_order(words))  # Output: "wertf"
```

### Pattern 4: Sequence Reconstruction
**Problem**: Check if there's only one topological sort order.

```python
from collections import deque, defaultdict

def sequence_reconstruction(original, sequences):
    graph = defaultdict(list)
    in_degree = defaultdict(int)
    
    # Build graph from sequences
    for seq in sequences:
        for i in range(len(seq)):
            in_degree[seq[i]] = 0  # Initialize all nodes
        
        for i in range(len(seq) - 1):
            from_node, to_node = seq[i], seq[i + 1]
            if to_node not in graph[from_node]:
                graph[from_node].append(to_node)
                in_degree[to_node] += 1
    
    # Check if original sequence is valid
    if len(original) != len(in_degree):
        return False
    
    queue = deque()
    for num in in_degree:
        if in_degree[num] == 0:
            queue.append(num)
    
    result = []
    
    while queue:
        if len(queue) > 1:  # More than one choice, not unique
            return False
        
        num = queue.popleft()
        result.append(num)
        
        for neighbor in graph[num]:
            in_degree[neighbor] -= 1
            if in_degree[neighbor] == 0:
                queue.append(neighbor)
    
    return result == original

# Example usage
original = [1, 2, 3, 4]
sequences = [[1, 2], [1, 3], [2, 3], [3, 4]]
print(sequence_reconstruction(original, sequences))  # Output: True
```

### Pattern 5: Minimum Height Trees
**Problem**: Find all possible roots that result in minimum height trees.

```python
from collections import defaultdict, deque

def find_min_height_trees(n, edges):
    if n <= 2:
        return list(range(n))
    
    # Build adjacency list
    graph = defaultdict(list)
    for a, b in edges:
        graph[a].append(b)
        graph[b].append(a)
    
    # Initialize leaves (nodes with degree 1)
    leaves = deque()
    for i in range(n):
        if len(graph[i]) == 1:
            leaves.append(i)
    
    remaining_nodes = n
    
    # Remove leaves layer by layer until 1-2 nodes remain
    while remaining_nodes > 2:
        leaves_count = len(leaves)
        remaining_nodes -= leaves_count
        
        for _ in range(leaves_count):
            leaf = leaves.popleft()
            
            # Remove leaf from its neighbor
            neighbor = graph[leaf][0]
            graph[neighbor].remove(leaf)
            
            # If neighbor becomes leaf, add it to queue
            if len(graph[neighbor]) == 1:
                leaves.append(neighbor)
    
    return list(leaves)

# Example usage
n = 6
edges = [[3,0],[3,1],[3,2],[3,4],[5,4]]
print(find_min_height_trees(n, edges))  # Output: [3,4]
```

### Pattern 6: Build Order (Tasks with Dependencies)
**Problem**: Find valid build order given project dependencies.

```python
from collections import defaultdict, deque

def build_order(projects, dependencies):
    graph = defaultdict(list)
    in_degree = defaultdict(int)
    
    # Initialize in_degree for all projects
    for project in projects:
        in_degree[project] = 0
    
    # Build graph
    for dependency, project in dependencies:
        graph[dependency].append(project)
        in_degree[project] += 1
    
    # Find projects with no dependencies
    queue = deque()
    for project in projects:
        if in_degree[project] == 0:
            queue.append(project)
    
    result = []
    
    while queue:
        project = queue.popleft()
        result.append(project)
        
        # Remove dependency for dependent projects
        for dependent in graph[project]:
            in_degree[dependent] -= 1
            if in_degree[dependent] == 0:
                queue.append(dependent)
    
    return result if len(result) == len(projects) else []

# Example usage
projects = ['a', 'b', 'c', 'd', 'e', 'f']
dependencies = [('a', 'd'), ('f', 'b'), ('b', 'd'), ('f', 'a'), ('d', 'c')]
print(build_order(projects, dependencies))  # Output: ['e', 'f', 'a', 'b', 'd', 'c']
```

### Pattern 7: Parallel Course Scheduling
**Problem**: Find minimum semesters to complete all courses with prerequisites.

```python
from collections import defaultdict, deque

def minimum_semesters(n, relations):
    graph = defaultdict(list)
    in_degree = [0] * (n + 1)
    
    for prev_course, next_course in relations:
        graph[prev_course].append(next_course)
        in_degree[next_course] += 1
    
    queue = deque()
    for i in range(1, n + 1):
        if in_degree[i] == 0:
            queue.append(i)
    
    semesters = 0
    studied = 0
    
    while queue:
        semesters += 1
        courses_this_semester = len(queue)
        
        for _ in range(courses_this_semester):
            course = queue.popleft()
            studied += 1
            
            for next_course in graph[course]:
                in_degree[next_course] -= 1
                if in_degree[next_course] == 0:
                    queue.append(next_course)
    
    return semesters if studied == n else -1

# Example usage
n = 3
relations = [[1,3],[2,3]]
print(minimum_semesters(n, relations))  # Output: 2
```

## Practice Problems

### Easy
1. **Course Schedule** - Basic topological sort
2. **Find Judge** - Graph degree analysis
3. **Keys and Rooms** - Graph traversal

### Medium
1. **Course Schedule II** - Topological sort with order
2. **Alien Dictionary** - String comparison topological sort
3. **Minimum Height Trees** - Tree center finding
4. **Sequence Reconstruction** - Unique topological order
5. **Build Order** - Dependency resolution

### Hard
1. **Parallel Courses** - Level-wise topological sort
2. **Sort Items by Groups** - Multi-level topological sort
3. **Reconstruct Itinerary** - Eulerian path finding
4. **Critical Connections** - Bridge finding in graphs

## Tips and Tricks

1. **Cycle Detection**: If topological sort doesn't include all vertices, there's a cycle
2. **Multiple Solutions**: Some problems have multiple valid topological orders
3. **Level-wise Processing**: Use BFS for problems requiring level information
4. **In-degree Tracking**: Key concept for Kahn's algorithm
5. **Graph Representation**: Choose between adjacency list/matrix based on density

## Common Mistakes

1. **Forgetting Initialization**: Not initializing in-degree for all vertices
2. **Cycle Handling**: Not properly detecting cycles in directed graphs
3. **Multiple Components**: Not handling disconnected graph components
4. **Index Errors**: Off-by-one errors with vertex numbering
5. **Memory Issues**: Not optimizing space for large graphs

## Related Patterns

- **DFS**: Alternative approach using depth-first search
- **BFS**: Kahn's algorithm is BFS-based
- **Union Find**: For some graph connectivity problems
- **Shortest Path**: Topological sort can optimize shortest path in DAGs

## Implementation Languages

Topological sort implementations across languages:
- **Python**: Use `collections.deque` and `defaultdict`
- **Java**: Use `Queue`, `ArrayList`, and `HashMap`
- **JavaScript**: Use arrays and objects for graph representation
- **C++**: Use `queue`, `vector`, and `unordered_map`
- **Go**: Use slices and maps with proper graph structures