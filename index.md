---
layout: default
title: Home
---

# 🚀 Comprehensive Coding Interview Patterns

Welcome to your ultimate guide for mastering coding interviews through pattern recognition!

## 🎯 Why This Approach Works

Instead of memorizing hundreds of individual problems, learn **patterns** that help you solve entire categories of problems:

- ✅ **Faster Recognition** - Quickly identify problem types during interviews
- ✅ **Systematic Approach** - Apply proven templates to new problems  
- ✅ **Efficient Preparation** - Cover more ground in less time
- ✅ **Build Intuition** - Develop problem-solving instincts

## 📚 Pattern Categories

### 🎪 Array & String Fundamentals
<div class="pattern-grid">
  <div class="pattern-card">
    <h4><a href="{{ '/patterns/two-pointers/' | relative_url }}">Two Pointers</a></h4>
    <p>Pairs, triplets, palindromes</p>
    <span class="difficulty easy">Easy</span>
  </div>
  
  <div class="pattern-card">
    <h4><a href="{{ '/patterns/sliding-window/' | relative_url }}">Sliding Window</a></h4>
    <p>Subarrays, substrings</p>
    <span class="difficulty medium">Medium</span>
  </div>
  
  <div class="pattern-card">
    <h4><a href="{{ '/patterns/fast-slow-pointers/' | relative_url }}">Fast & Slow Pointers</a></h4>
    <p>Cycles, middle elements</p>
    <span class="difficulty medium">Medium</span>
  </div>
  
  <div class="pattern-card">
    <h4><a href="{{ '/patterns/merge-intervals/' | relative_url }}">Merge Intervals</a></h4>
    <p>Scheduling, time ranges</p>
    <span class="difficulty medium">Medium</span>
  </div>
</div>

### 🌳 Tree & Graph Patterns
<div class="pattern-grid">
  <div class="pattern-card">
    <h4><a href="{{ '/patterns/tree-bfs/' | relative_url }}">Tree BFS</a></h4>
    <p>Level-order traversal</p>
    <span class="difficulty easy">Easy</span>
  </div>
  
  <div class="pattern-card">
    <h4><a href="{{ '/patterns/tree-dfs/' | relative_url }}">Tree DFS</a></h4>
    <p>Path problems, tree depth</p>
    <span class="difficulty medium">Medium</span>
  </div>
</div>

### 🎯 Advanced Problem Solving
<div class="pattern-grid">
  <div class="pattern-card">
    <h4><a href="{{ '/patterns/dynamic-programming/' | relative_url }}">Dynamic Programming</a></h4>
    <p>Optimization problems</p>
    <span class="difficulty hard">Hard</span>
  </div>
  
  <div class="pattern-card">
    <h4><a href="{{ '/patterns/subsets-backtracking/' | relative_url }}">Backtracking</a></h4>
    <p>Combinations, permutations</p>
    <span class="difficulty medium">Medium</span>
  </div>
</div>

## 🚀 Quick Start Guide

### For Beginners
1. Start with **Two Pointers** - fundamental and intuitive
2. Move to **Sliding Window** - builds on two pointers concept
3. Practice 5 problems per pattern before advancing
4. Focus on understanding the pattern logic, not memorizing solutions

### For Experienced Developers
1. Review pattern summaries for quick refreshers
2. Jump to **Hard** difficulty problems
3. Practice under timed conditions
4. Focus on pattern recognition speed

## 📊 Your Progress

Track your journey through the patterns:

- [ ] **Foundation** (Two Pointers, Sliding Window, Fast & Slow Pointers)
- [ ] **Trees** (BFS, DFS traversals)
- [ ] **Advanced** (DP, Backtracking)
- [ ] **Expert** (Hard problems across all patterns)

## 🔗 Quick Links

- [**Study Plan**]({{ '/study-plan/' | relative_url }}) - Structured 8-week preparation
- [**Practice Problems**]({{ '/practice/' | relative_url }}) - Curated problem sets
- [**Templates**]({{ '/templates/' | relative_url }}) - Ready-to-use code patterns
- [**FAQ**]({{ '/faq/' | relative_url }}) - Common questions answered

---

## 💡 Pattern Recognition Tip

```
Keywords → Patterns

"pair", "two sum" → Two Pointers
"subarray", "window" → Sliding Window  
"cycle", "loop" → Fast & Slow Pointers
"intervals", "meeting" → Merge Intervals
"tree", "level" → Tree BFS/DFS
"optimization" → Dynamic Programming
"all combinations" → Backtracking
```

<style>
.pattern-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 1rem;
  margin: 1rem 0;
}

.pattern-card {
  border: 1px solid #e1e4e8;
  border-radius: 6px;
  padding: 1rem;
  background: #fff;
  box-shadow: 0 1px 3px rgba(0,0,0,0.12);
}

.pattern-card:hover {
  box-shadow: 0 4px 12px rgba(0,0,0,0.15);
  transform: translateY(-2px);
  transition: all 0.2s ease;
}

.pattern-card h4 {
  margin: 0 0 0.5rem 0;
  color: #0366d6;
}

.pattern-card p {
  margin: 0 0 0.5rem 0;
  color: #586069;
  font-size: 0.9rem;
}

.difficulty {
  font-size: 0.75rem;
  padding: 0.2rem 0.5rem;
  border-radius: 12px;
  font-weight: 500;
}

.difficulty.easy {
  background: #d4edda;
  color: #155724;
}

.difficulty.medium {
  background: #fff3cd;
  color: #856404;
}

.difficulty.hard {
  background: #f8d7da;
  color: #721c24;
}
</style>