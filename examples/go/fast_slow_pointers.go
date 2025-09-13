package main

import "fmt"

/*
Fast & Slow Pointers Pattern Examples

The Fast & Slow Pointers technique uses two pointers moving at different speeds
to solve problems involving cyclic data structures or finding middle elements.

Time Complexity: O(n) for most problems
Space Complexity: O(1) - constant space usage
*/

// ListNode represents a node in a linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// HasCycle detects if a linked list has a cycle using Floyd's algorithm
// Time: O(n), Space: O(1)
func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

// FindCycleStart finds the start of the cycle in a linked list
// Time: O(n), Space: O(1)
func FindCycleStart(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head

	// Find if cycle exists
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

	// No cycle found
	if fast == nil || fast.Next == nil {
		return nil
	}

	// Find cycle start
	current := head
	for current != slow {
		current = current.Next
		slow = slow.Next
	}

	return current
}

// FindMiddleNode finds the middle node of a linked list
// Time: O(n), Space: O(1)
func FindMiddleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// IsPalindromeLinkedList checks if a linked list is a palindrome
// Time: O(n), Space: O(1)
func IsPalindromeLinkedList(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// Find middle using fast/slow pointers
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Reverse second half
	reverseList := func(node *ListNode) *ListNode {
		var prev *ListNode
		for node != nil {
			nextNode := node.Next
			node.Next = prev
			prev = node
			node = nextNode
		}
		return prev
	}

	secondHalf := reverseList(slow)

	// Compare first and second half
	firstHalf := head
	for secondHalf != nil {
		if firstHalf.Val != secondHalf.Val {
			return false
		}
		firstHalf = firstHalf.Next
		secondHalf = secondHalf.Next
	}

	return true
}

// FindHappyNumber determines if a number is happy using fast/slow pointer approach
// Time: O(log n), Space: O(1)
func FindHappyNumber(n int) bool {
	getNext := func(number int) int {
		totalSum := 0
		for number > 0 {
			digit := number % 10
			totalSum += digit * digit
			number /= 10
		}
		return totalSum
	}

	slow, fast := n, n

	for {
		slow = getNext(slow)
		fast = getNext(getNext(fast))

		if fast == 1 {
			return true
		}

		if slow == fast {
			return false
		}
	}
}

// CircularArrayLoop checks if there exists a cycle in a circular array
// Time: O(n), Space: O(1)
func CircularArrayLoop(nums []int) bool {
	getNextIndex := func(i int) int {
		return (i + nums[i] + len(nums)) % len(nums)
	}

	for i := 0; i < len(nums); i++ {
		isForward := nums[i] >= 0
		slow, fast := i, i

		for {
			slow = getNextIndex(slow)
			fast = getNextIndex(getNextIndex(fast))

			// Check if direction changed or single element cycle
			if (nums[slow] >= 0) != isForward ||
				(nums[fast] >= 0) != isForward ||
				(nums[getNextIndex(fast)] >= 0) != isForward {
				break
			}

			if slow == fast {
				// Check if it's not a single element cycle
				if slow == getNextIndex(slow) {
					break
				}
				return true
			}
		}
	}

	return false
}

// Helper function to create a linked list from slice
func createLinkedList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	current := head

	for i := 1; i < len(vals); i++ {
		current.Next = &ListNode{Val: vals[i]}
		current = current.Next
	}

	return head
}

// Test functions
func testHasCycle() {
	fmt.Println("Testing cycle detection...")

	// Create a cycle: 1 -> 2 -> 3 -> 2
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node2 // Creates cycle

	if !HasCycle(node1) {
		panic("Expected cycle to be detected")
	}

	// No cycle: 1 -> 2 -> 3 -> nil
	node3.Next = nil
	if HasCycle(node1) {
		panic("Expected no cycle")
	}

	fmt.Println("✅ Cycle detection tests passed")
}

func testFindMiddle() {
	fmt.Println("Testing find middle node...")

	// Odd length: 1 -> 2 -> 3 -> 4 -> 5
	head := createLinkedList([]int{1, 2, 3, 4, 5})
	middle := FindMiddleNode(head)
	if middle.Val != 3 {
		panic(fmt.Sprintf("Expected middle value 3, got %d", middle.Val))
	}

	// Even length: 1 -> 2 -> 3 -> 4
	head = createLinkedList([]int{1, 2, 3, 4})
	middle = FindMiddleNode(head)
	if middle.Val != 3 {
		panic(fmt.Sprintf("Expected middle value 3, got %d", middle.Val))
	}

	fmt.Println("✅ Middle node tests passed")
}

func testHappyNumber() {
	fmt.Println("Testing happy number detection...")

	testCases := []struct {
		input    int
		expected bool
	}{
		{19, true}, // 1^2 + 9^2 = 82, 8^2 + 2^2 = 68, 6^2 + 8^2 = 100, 1^2 + 0^2 + 0^2 = 1
		{2, false}, // Will cycle
		{7, true},  // Happy number
	}

	for _, tc := range testCases {
		result := FindHappyNumber(tc.input)
		if result != tc.expected {
			panic(fmt.Sprintf("Happy number test failed for %d: expected %v, got %v", 
				tc.input, tc.expected, result))
		}
	}

	fmt.Println("✅ Happy number tests passed")
}

func testCircularArray() {
	fmt.Println("Testing circular array loop detection...")

	testCases := []struct {
		input    []int
		expected bool
	}{
		{[]int{2, -1, 1, 2, 2}, true},
		{[]int{-1, 2}, false},
		{[]int{-2, 1, -1, -2, -2}, false},
	}

	for _, tc := range testCases {
		result := CircularArrayLoop(tc.input)
		if result != tc.expected {
			panic(fmt.Sprintf("Circular array test failed for %v: expected %v, got %v", 
				tc.input, tc.expected, result))
		}
	}

	fmt.Println("✅ Circular array tests passed")
}

func main() {
	fmt.Println("Testing Fast & Slow Pointers Pattern...")
	testHasCycle()
	testFindMiddle()
	testHappyNumber()
	testCircularArray()
	fmt.Println("\n🎉 All Fast & Slow Pointers tests passed!")
}