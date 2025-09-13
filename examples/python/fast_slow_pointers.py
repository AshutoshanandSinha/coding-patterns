#!/usr/bin/env python3
"""
Fast & Slow Pointers Pattern Examples

The Fast & Slow Pointers technique uses two pointers moving at different speeds
to solve problems involving cyclic data structures or finding middle elements.

Time Complexity: O(n) for most problems
Space Complexity: O(1) - constant space usage
"""

from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def has_cycle(head: Optional[ListNode]) -> bool:
    """
    Detect if a linked list has a cycle using Floyd's algorithm.
    
    Args:
        head: Head of the linked list
        
    Returns:
        True if cycle exists, False otherwise
        
    Time: O(n), Space: O(1)
    """
    if not head or not head.next:
        return False
    
    slow = fast = head
    
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
        
        if slow == fast:
            return True
    
    return False


def find_cycle_start(head: Optional[ListNode]) -> Optional[ListNode]:
    """
    Find the start of the cycle in a linked list.
    
    Args:
        head: Head of the linked list
        
    Returns:
        Node where cycle starts, or None if no cycle
        
    Time: O(n), Space: O(1)
    """
    if not head or not head.next:
        return None
    
    slow = fast = head
    
    # Find if cycle exists
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
        
        if slow == fast:
            break
    else:
        return None  # No cycle found
    
    # Find cycle start
    current = head
    while current != slow:
        current = current.next
        slow = slow.next
    
    return current


def find_middle_node(head: Optional[ListNode]) -> Optional[ListNode]:
    """
    Find the middle node of a linked list.
    
    Args:
        head: Head of the linked list
        
    Returns:
        Middle node (for even length, returns second middle)
        
    Time: O(n), Space: O(1)
    """
    if not head:
        return None
    
    slow = fast = head
    
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
    
    return slow


def is_palindrome_linked_list(head: Optional[ListNode]) -> bool:
    """
    Check if a linked list is a palindrome.
    
    Args:
        head: Head of the linked list
        
    Returns:
        True if palindrome, False otherwise
        
    Time: O(n), Space: O(1)
    """
    if not head or not head.next:
        return True
    
    # Find middle using fast/slow pointers
    slow = fast = head
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
    
    # Reverse second half
    def reverse_list(node):
        prev = None
        while node:
            next_node = node.next
            node.next = prev
            prev = node
            node = next_node
        return prev
    
    second_half = reverse_list(slow)
    
    # Compare first and second half
    first_half = head
    while second_half:
        if first_half.val != second_half.val:
            return False
        first_half = first_half.next
        second_half = second_half.next
    
    return True


def find_happy_number(n: int) -> bool:
    """
    Determine if a number is happy using fast/slow pointer approach.
    A happy number is defined by the process: starting with any positive integer,
    replace the number by the sum of the square of its digits, repeat until
    the number equals 1 (where it will stay), or it loops endlessly in a cycle.
    
    Args:
        n: The number to check
        
    Returns:
        True if happy number, False otherwise
        
    Time: O(log n), Space: O(1)
    """
    def get_next(number):
        total_sum = 0
        while number > 0:
            digit = number % 10
            total_sum += digit * digit
            number //= 10
        return total_sum
    
    slow = fast = n
    
    while True:
        slow = get_next(slow)
        fast = get_next(get_next(fast))
        
        if fast == 1:
            return True
        
        if slow == fast:
            return False


def circular_array_loop(nums: list[int]) -> bool:
    """
    Check if there exists a cycle in a circular array.
    
    Args:
        nums: Array of integers representing jumps
        
    Returns:
        True if valid cycle exists, False otherwise
        
    Time: O(n), Space: O(1)
    """
    def get_next_index(i):
        return (i + nums[i]) % len(nums)
    
    for i in range(len(nums)):
        is_forward = nums[i] >= 0
        slow = fast = i
        
        while True:
            slow = get_next_index(slow)
            fast = get_next_index(get_next_index(fast))
            
            # Check if direction changed or single element cycle
            if (nums[slow] >= 0) != is_forward or \
               (nums[fast] >= 0) != is_forward or \
               (nums[get_next_index(fast)] >= 0) != is_forward:
                break
            
            if slow == fast:
                # Check if it's not a single element cycle
                if slow == get_next_index(slow):
                    break
                return True
    
    return False


# Test functions
def test_has_cycle():
    """Test cycle detection"""
    # Create a cycle: 1 -> 2 -> 3 -> 2
    node1 = ListNode(1)
    node2 = ListNode(2)
    node3 = ListNode(3)
    node1.next = node2
    node2.next = node3
    node3.next = node2  # Creates cycle
    
    assert has_cycle(node1) == True
    
    # No cycle: 1 -> 2 -> 3 -> None
    node1.next = node2
    node2.next = node3
    node3.next = None
    
    assert has_cycle(node1) == False
    print("✅ Cycle detection tests passed")


def test_find_middle():
    """Test finding middle node"""
    # Odd length: 1 -> 2 -> 3 -> 4 -> 5
    nodes = [ListNode(i) for i in range(1, 6)]
    for i in range(4):
        nodes[i].next = nodes[i + 1]
    
    middle = find_middle_node(nodes[0])
    assert middle.val == 3
    
    # Even length: 1 -> 2 -> 3 -> 4
    nodes = [ListNode(i) for i in range(1, 5)]
    for i in range(3):
        nodes[i].next = nodes[i + 1]
    
    middle = find_middle_node(nodes[0])
    assert middle.val == 3
    print("✅ Middle node tests passed")


def test_happy_number():
    """Test happy number detection"""
    assert find_happy_number(19) == True  # 1^2 + 9^2 = 82, 8^2 + 2^2 = 68, 6^2 + 8^2 = 100, 1^2 + 0^2 + 0^2 = 1
    assert find_happy_number(2) == False   # Will cycle
    assert find_happy_number(7) == True    # Happy number
    print("✅ Happy number tests passed")


def test_circular_array():
    """Test circular array loop detection"""
    assert circular_array_loop([2, -1, 1, 2, 2]) == True
    assert circular_array_loop([-1, 2]) == False
    assert circular_array_loop([-2, 1, -1, -2, -2]) == False
    print("✅ Circular array tests passed")


if __name__ == "__main__":
    print("Testing Fast & Slow Pointers Pattern...")
    test_has_cycle()
    test_find_middle()
    test_happy_number()
    test_circular_array()
    print("\n🎉 All Fast & Slow Pointers tests passed!")