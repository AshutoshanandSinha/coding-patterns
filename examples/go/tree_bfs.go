// Tree BFS (Breadth-First Search) Pattern Examples
//
// This package demonstrates BFS traversal patterns for tree problems,
// including level-order traversal and tree manipulation using queues.

package main

import (
	"fmt"
	"math"
)

// TreeNode represents a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Node represents a tree node with next pointer
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// LevelOrderTraversal contains level-order traversal patterns
type LevelOrderTraversal struct{}

// LevelOrder performs basic level-order traversal
func (lot *LevelOrderTraversal) LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		levelNodes := []int{}

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			levelNodes = append(levelNodes, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, levelNodes)
	}

	return result
}

// LevelOrderBottom performs level-order traversal from bottom up
func (lot *LevelOrderTraversal) LevelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		levelNodes := []int{}

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			levelNodes = append(levelNodes, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, levelNodes)
	}

	// Reverse the result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

// ZigzagLevelOrder performs zigzag level-order traversal
func (lot *LevelOrderTraversal) ZigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{root}
	leftToRight := true

	for len(queue) > 0 {
		levelSize := len(queue)
		levelNodes := []int{}

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			levelNodes = append(levelNodes, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if !leftToRight {
			// Reverse level nodes
			for i, j := 0, len(levelNodes)-1; i < j; i, j = i+1, j-1 {
				levelNodes[i], levelNodes[j] = levelNodes[j], levelNodes[i]
			}
		}

		result = append(result, levelNodes)
		leftToRight = !leftToRight
	}

	return result
}

// TreeLevelProblems contains problems involving tree levels
type TreeLevelProblems struct{}

// RightSideView gets right side view of binary tree
func (tlp *TreeLevelProblems) RightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// Last node in level is rightmost
			if i == levelSize-1 {
				result = append(result, node.Val)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return result
}

// AverageOfLevels calculates average of each level
func (tlp *TreeLevelProblems) AverageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	result := []float64{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		levelSum := 0

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			levelSum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, float64(levelSum)/float64(levelSize))
	}

	return result
}

// LevelWithMinimumNodes finds level with minimum number of nodes
func (tlp *TreeLevelProblems) LevelWithMinimumNodes(root *TreeNode) int {
	if root == nil {
		return -1
	}

	minLevel := 0
	minCount := math.MaxInt32
	queue := []*TreeNode{root}
	level := 0

	for len(queue) > 0 {
		levelSize := len(queue)

		if levelSize < minCount {
			minCount = levelSize
			minLevel = level
		}

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		level++
	}

	return minLevel
}

// TreeConstruction contains tree construction using BFS
type TreeConstruction struct{}

// ConnectNextRightPointers connects nodes to their next right node
func (tc *TreeConstruction) ConnectNextRightPointers(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		var prev *Node

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if prev != nil {
				prev.Next = node
			}
			prev = node

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return root
}

// TreeValidation contains tree validation using BFS
type TreeValidation struct{}

// IsCompleteTree checks if binary tree is complete
func (tv *TreeValidation) IsCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*TreeNode{root}
	nullFound := false

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			nullFound = true
		} else {
			if nullFound { // Found node after null
				return false
			}
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}

	return true
}

// IsSymmetric checks if tree is symmetric
func (tv *TreeValidation) IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*TreeNode{root.Left, root.Right}

	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]

		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}

		queue = append(queue, left.Left)
		queue = append(queue, right.Right)
		queue = append(queue, left.Right)
		queue = append(queue, right.Left)
	}

	return true
}

// TreeMeasurements contains tree measurements using BFS
type TreeMeasurements struct{}

// MaxDepth finds maximum depth of tree
func (tm *TreeMeasurements) MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 0

	for len(queue) > 0 {
		depth++
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return depth
}

// MinDepth finds minimum depth to leaf node
func (tm *TreeMeasurements) MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 0

	for len(queue) > 0 {
		depth++
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// First leaf node found
			if node.Left == nil && node.Right == nil {
				return depth
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return depth
}

// NodeWithIndex represents a node with its index
type NodeWithIndex struct {
	Node  *TreeNode
	Index int
}

// MaxWidth finds maximum width of tree
func (tm *TreeMeasurements) MaxWidth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxWidth := 1
	queue := []NodeWithIndex{{root, 0}}

	for len(queue) > 0 {
		levelSize := len(queue)
		firstIndex := queue[0].Index
		lastIndex := queue[levelSize-1].Index

		if lastIndex-firstIndex+1 > maxWidth {
			maxWidth = lastIndex - firstIndex + 1
		}

		for i := 0; i < levelSize; i++ {
			item := queue[0]
			queue = queue[1:]

			if item.Node.Left != nil {
				queue = append(queue, NodeWithIndex{item.Node.Left, 2 * item.Index})
			}
			if item.Node.Right != nil {
				queue = append(queue, NodeWithIndex{item.Node.Right, 2*item.Index + 1})
			}
		}
	}

	return maxWidth
}

func main() {
	// Build a test tree:
	//       3
	//      / \
	//     9   20
	//        /  \
	//       15   7
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	// Test level-order traversal
	lot := &LevelOrderTraversal{}
	fmt.Println("Level order:", lot.LevelOrder(root))
	fmt.Println("Level order bottom up:", lot.LevelOrderBottom(root))
	fmt.Println("Zigzag level order:", lot.ZigzagLevelOrder(root))

	// Test tree level problems
	tlp := &TreeLevelProblems{}
	fmt.Println("Right side view:", tlp.RightSideView(root))
	fmt.Println("Average of levels:", tlp.AverageOfLevels(root))
	fmt.Println("Level with min nodes:", tlp.LevelWithMinimumNodes(root))

	// Test tree validation
	tv := &TreeValidation{}
	fmt.Println("Is complete tree:", tv.IsCompleteTree(root))
	fmt.Println("Is symmetric:", tv.IsSymmetric(root))

	// Test tree measurements
	tm := &TreeMeasurements{}
	fmt.Println("Max depth:", tm.MaxDepth(root))
	fmt.Println("Min depth:", tm.MinDepth(root))
	fmt.Println("Max width:", tm.MaxWidth(root))

	// Test with symmetric tree
	symmetricRoot := &TreeNode{Val: 1}
	symmetricRoot.Left = &TreeNode{Val: 2}
	symmetricRoot.Right = &TreeNode{Val: 2}
	symmetricRoot.Left.Left = &TreeNode{Val: 3}
	symmetricRoot.Left.Right = &TreeNode{Val: 4}
	symmetricRoot.Right.Left = &TreeNode{Val: 4}
	symmetricRoot.Right.Right = &TreeNode{Val: 3}

	fmt.Println("Is symmetric (symmetric tree):", tv.IsSymmetric(symmetricRoot))
}