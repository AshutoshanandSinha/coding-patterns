# 🚀 Comprehensive Coding Interview Patterns

A comprehensive collection of coding patterns to master technical interviews. This repository contains detailed explanations, implementations, and practice problems for the most important algorithmic patterns you'll encounter in coding interviews.

## 📚 Table of Contents

- [About This Repository](#about-this-repository)
- [How to Use This Guide](#how-to-use-this-guide)  
- [Pattern Categories](#pattern-categories)
- [Quick Reference](#quick-reference)
- [Study Plan](#study-plan)
- [Contributing](#contributing)

## 🎯 About This Repository

Master coding interviews by learning **patterns** instead of memorizing individual problems. This approach helps you:

- ✅ **Recognize problem types** quickly during interviews
- ✅ **Apply systematic approaches** to solve new problems
- ✅ **Reduce preparation time** by focusing on core patterns
- ✅ **Build problem-solving intuition** through pattern recognition
- ✅ **Cover more ground** than solving random problems

## 🔧 How to Use This Guide

### For Beginners
1. Start with **Two Pointers** and **Sliding Window** patterns
2. Practice 3-5 problems per pattern before moving on
3. Focus on understanding the pattern logic rather than memorizing solutions
4. Use the provided templates as starting points

### For Experienced Programmers
1. Review pattern summaries for quick refreshers
2. Focus on **Hard** level problems and edge cases
3. Practice implementing patterns in your preferred language
4. Time yourself to simulate interview conditions

### Pattern Recognition Tips
- Read the problem statement carefully
- Look for keywords that hint at specific patterns:
  - "two sum", "pairs" → **Two Pointers**
  - "subarray", "substring" → **Sliding Window**
  - "cycle", "middle" → **Fast & Slow Pointers**
  - "intervals", "scheduling" → **Merge Intervals**
  - "all combinations", "subsets" → **Backtracking**

## 📊 Pattern Categories

### 🎪 Array & String Patterns

| Pattern | Difficulty | Key Problems | Time Complexity |
|---------|------------|--------------|-----------------|
| [Two Pointers](./patterns/two-pointers/) | 🟢 Easy | Two Sum, 3Sum, Container With Most Water | O(n) |
| [Sliding Window](./patterns/sliding-window/) | 🟡 Medium | Longest Substring, Min Window Substring | O(n) |
| [Fast & Slow Pointers](./patterns/fast-slow-pointers/) | 🟡 Medium | Cycle Detection, Find Middle, Happy Number | O(n) |
| [Merge Intervals](./patterns/merge-intervals/) | 🟡 Medium | Meeting Rooms, Insert Interval | O(n log n) |

### 🌳 Tree & Graph Patterns

| Pattern | Difficulty | Key Problems | Time Complexity |
|---------|------------|--------------|-----------------|
| [Tree BFS](./patterns/tree-bfs/) | 🟢 Easy | Level Order, Right Side View | O(n) |
| [Tree DFS](./patterns/tree-dfs/) | 🟡 Medium | Path Sum, Diameter, Serialize Tree | O(n) |
| [Graph Traversal](./patterns/graph-traversal/) | 🟡 Medium | Clone Graph, Course Schedule | O(V + E) |
| [Topological Sort](./patterns/topological-sort/) | 🔴 Hard | Task Scheduling, Alien Dictionary | O(V + E) |

### 🔍 Search & Sort Patterns

| Pattern | Difficulty | Key Problems | Time Complexity |
|---------|------------|--------------|-----------------|
| [Binary Search](./patterns/binary-search/) | 🟡 Medium | Search in Rotated Array, First Bad Version | O(log n) |
| [Cyclic Sort](./patterns/cyclic-sort/) | 🟢 Easy | Missing Number, Find Duplicates | O(n) |
| [Top K Elements](./patterns/top-k-elements/) | 🟡 Medium | Kth Largest, Top K Frequent | O(n log k) |
| [K-way Merge](./patterns/k-way-merge/) | 🔴 Hard | Merge K Sorted Lists, Smallest Range | O(n log k) |

### 🎯 Dynamic Programming Patterns

| Pattern | Difficulty | Key Problems | Time Complexity |
|---------|------------|--------------|-----------------|
| [0/1 Knapsack](./patterns/dynamic-programming/knapsack/) | 🟡 Medium | Subset Sum, Partition Equal Sum | O(n*m) |
| [Unbounded Knapsack](./patterns/dynamic-programming/unbounded-knapsack/) | 🟡 Medium | Coin Change, Rod Cutting | O(n*m) |
| [Fibonacci Numbers](./patterns/dynamic-programming/fibonacci/) | 🟢 Easy | Climbing Stairs, House Robber | O(n) |
| [Palindromic Subsequence](./patterns/dynamic-programming/palindromic/) | 🔴 Hard | Longest Palindromic Substring | O(n²) |

### 🔄 Backtracking & Recursion

| Pattern | Difficulty | Key Problems | Time Complexity |
|---------|------------|--------------|-----------------|
| [Subsets](./patterns/subsets-backtracking/) | 🟡 Medium | Generate Subsets, Permutations | O(2ⁿ) |
| [Combinations](./patterns/combinations/) | 🟡 Medium | Combination Sum, Generate Parentheses | O(2ⁿ) |
| [N-Queens](./patterns/n-queens/) | 🔴 Hard | N-Queens, Sudoku Solver | O(n!) |

### ⚡ Advanced Patterns

| Pattern | Difficulty | Key Problems | Time Complexity |
|---------|------------|--------------|-----------------|
| [Bitwise XOR](./patterns/bitwise-xor/) | 🟢 Easy | Single Number, Missing Number | O(n) |
| [Two Heaps](./patterns/two-heaps/) | 🔴 Hard | Median from Data Stream | O(log n) |
| [Modified Binary Search](./patterns/binary-search/) | 🔴 Hard | Search in Rotated Array | O(log n) |

## ⚡ Quick Reference

### Pattern Recognition Cheat Sheet

```
🔍 KEYWORDS → PATTERNS

"two sum", "pair", "triplet" → Two Pointers
"substring", "subarray", "window" → Sliding Window  
"cycle", "circular", "loop" → Fast & Slow Pointers
"intervals", "meeting", "schedule" → Merge Intervals
"tree", "binary tree" → Tree DFS/BFS
"graph", "connected", "path" → Graph Traversal
"sorted array", "search" → Binary Search
"missing number", "duplicate" → Cyclic Sort
"k largest", "k smallest" → Top K Elements
"merge", "k sorted" → K-way Merge
"subset", "combination", "permutation" → Backtracking
"optimization", "maximum", "minimum" → Dynamic Programming
```

### Time Complexity Quick Guide

| Complexity | Typical Patterns | Example Operations |
|------------|------------------|-------------------|
| O(1) | Hash Map Access | Array access, hash lookup |
| O(log n) | Binary Search | Search in sorted array |
| O(n) | Two Pointers, Sliding Window | Single pass through array |
| O(n log n) | Merge Intervals | Sorting + linear scan |
| O(n²) | Nested loops, DP | Matrix traversal, some DP |
| O(2ⁿ) | Backtracking | Subsets, permutations |

## 📅 Study Plan

### Week 1-2: Foundation Patterns
- **Day 1-3**: Two Pointers
- **Day 4-7**: Sliding Window
- **Day 8-10**: Fast & Slow Pointers
- **Day 11-14**: Merge Intervals

### Week 3-4: Tree & Graph Patterns  
- **Day 15-18**: Tree BFS
- **Day 19-22**: Tree DFS
- **Day 23-26**: Basic Graph Traversal
- **Day 27-28**: Review and Practice

### Week 5-6: Advanced Patterns
- **Day 29-32**: Binary Search & Variations
- **Day 33-36**: Top K Elements
- **Day 37-40**: Dynamic Programming (Basic)
- **Day 41-42**: Review and Mock Interviews

### Week 7-8: Expert Level
- **Day 43-46**: Backtracking & Subsets
- **Day 47-50**: Advanced DP Patterns
- **Day 51-54**: Hard Problems Mix
- **Day 55-56**: Final Review & Timed Practice

## 📈 Progress Tracking

Use this checklist to track your progress:

### Beginner Patterns ✅
- [ ] Two Pointers (15 problems)
- [ ] Sliding Window (12 problems)  
- [ ] Fast & Slow Pointers (8 problems)
- [ ] Merge Intervals (10 problems)

### Intermediate Patterns 🟡
- [ ] Tree BFS (10 problems)
- [ ] Tree DFS (15 problems)
- [ ] Binary Search (12 problems)
- [ ] Top K Elements (10 problems)

### Advanced Patterns 🔴
- [ ] Backtracking (15 problems)
- [ ] Dynamic Programming (20 problems)
- [ ] Graph Advanced (8 problems)
- [ ] Hard Mixed (10 problems)

## 🛠️ Setup Instructions

### Local Development
```bash
# Clone the repository
git clone https://github.com/your-username/coding-patterns.git
cd coding-patterns

# Choose your preferred language directory
cd examples/python  # or java, javascript, cpp, etc.

# Run example solutions
python two_pointers_examples.py
```

### Online Practice
- **LeetCode**: Practice problems organized by pattern
- **HackerRank**: Additional problems for each pattern
- **CodeSignal**: Timed practice sessions

## 🤝 Contributing

We welcome contributions! Here's how you can help:

1. **Add New Patterns**: Found a pattern we missed?
2. **Improve Explanations**: Make existing content clearer
3. **Add Solutions**: Contribute solutions in different languages  
4. **Report Issues**: Found bugs or unclear explanations?
5. **Suggest Problems**: Know of great practice problems?

### Contribution Guidelines
- Follow the existing pattern structure
- Include time/space complexity analysis
- Add 3-5 example problems per pattern
- Test all code examples
- Update the main README

## 📖 Additional Resources

### Books
- "Cracking the Coding Interview" by Gayle McDowell
- "Elements of Programming Interviews" by Aziz, Prakash & Lee
- "Algorithm Design Manual" by Steven Skiena

### Online Courses
- Grokking the Coding Interview (Educative)
- AlgoExpert
- InterviewBit

### Practice Platforms
- [LeetCode](https://leetcode.com/) - Most popular platform
- [HackerRank](https://www.hackerrank.com/) - Diverse problem set
- [CodeSignal](https://codesignal.com/) - Interview simulation

## ⭐ Star History

If this repository helps you in your interview preparation, please consider giving it a star! ⭐

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Thanks to all contributors who help improve this resource
- Inspired by the "Grokking the Coding Interview" methodology
- Problem examples adapted from LeetCode and other platforms

---

**Happy Coding and Good Luck with Your Interviews! 🚀**

*Remember: Consistent practice and pattern recognition are the keys to interview success.*