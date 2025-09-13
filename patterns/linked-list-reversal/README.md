# Linked List Reversal Pattern

## Overview
The Linked List Reversal pattern involves reversing links between nodes in a linked list. This fundamental technique is used to reverse entire lists, parts of lists, or reorganize list structures. The pattern is essential for many linked list manipulation problems and forms the foundation for more complex list operations.

## When to Use
- **Reverse Operations**: Reversing entire or partial linked lists
- **List Reorganization**: Rearranging list elements in specific patterns
- **Palindrome Checks**: Comparing list halves after reversal
- **K-Group Operations**: Reversing nodes in groups of k
- **Cycle Problems**: Some cycle-related problems use reversal concepts

## Time/Space Complexity
- **Time**: O(n) - Single pass through the list
- **Space**: O(1) - Only a few pointers needed (iterative), O(n) for recursive

## Basic List Node Structure
```python
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
```

## Pattern Variations

### 1. Iterative Reversal
```python
def reverse_list_iterative(head):
    prev = None
    current = head
    
    while current:
        next_temp = current.next  # Store next node
        current.next = prev       # Reverse the link
        prev = current           # Move prev forward
        current = next_temp      # Move current forward
    
    return prev  # prev is the new head
```

### 2. Recursive Reversal
```python
def reverse_list_recursive(head):
    # Base case
    if not head or not head.next:
        return head
    
    # Recursively reverse the rest
    new_head = reverse_list_recursive(head.next)
    
    # Reverse current connection
    head.next.next = head
    head.next = None
    
    return new_head
```

## Common Problem Patterns

### Pattern 1: Reverse Entire Linked List
**Problem**: Reverse a singly linked list.

```python
def reverse_list(head):
    prev = None
    current = head
    
    while current:
        next_node = current.next
        current.next = prev
        prev = current
        current = next_node
    
    return prev

# Example usage
# Input:  1 -> 2 -> 3 -> 4 -> 5 -> None
# Output: 5 -> 4 -> 3 -> 2 -> 1 -> None
```

### Pattern 2: Reverse List in Groups of K
**Problem**: Reverse nodes of linked list k at a time.

```python
def reverse_k_group(head, k):
    def reverse_group(start, end):
        prev = end.next
        current = start
        
        while current != end.next:
            next_temp = current.next
            current.next = prev
            prev = current
            current = next_temp
        
        return end, start  # New start and end after reversal
    
    def find_kth_node(node, k):
        while node and k > 1:
            node = node.next
            k -= 1
        return node
    
    dummy = ListNode(0)
    dummy.next = head
    prev_group_end = dummy
    
    while True:
        kth_node = find_kth_node(prev_group_end.next, k)
        if not kth_node:
            break
        
        next_group_start = kth_node.next
        
        # Reverse current group
        new_start, new_end = reverse_group(prev_group_end.next, kth_node)
        
        # Connect with previous group
        prev_group_end.next = new_start
        new_end.next = next_group_start
        prev_group_end = new_end
    
    return dummy.next

# Example usage: k = 2
# Input:  1 -> 2 -> 3 -> 4 -> 5 -> None
# Output: 2 -> 1 -> 4 -> 3 -> 5 -> None
```

### Pattern 3: Reverse Sublist
**Problem**: Reverse linked list from position m to n.

```python
def reverse_between(head, left, right):
    if not head or left == right:
        return head
    
    dummy = ListNode(0)
    dummy.next = head
    prev = dummy
    
    # Move to the node before the reversal starts
    for _ in range(left - 1):
        prev = prev.next
    
    # Start of reversal
    current = prev.next
    
    # Reverse the sublist
    for _ in range(right - left):
        next_temp = current.next
        current.next = next_temp.next
        next_temp.next = prev.next
        prev.next = next_temp
    
    return dummy.next

# Example usage: left = 2, right = 4
# Input:  1 -> 2 -> 3 -> 4 -> 5 -> None
# Output: 1 -> 4 -> 3 -> 2 -> 5 -> None
```

### Pattern 4: Rotate List
**Problem**: Rotate the list to the right by k places.

```python
def rotate_right(head, k):
    if not head or not head.next or k == 0:
        return head
    
    # Find length and make it circular
    length = 1
    tail = head
    while tail.next:
        tail = tail.next
        length += 1
    
    # Connect tail to head to make it circular
    tail.next = head
    
    # Find new tail (length - k % length - 1 steps from head)
    k = k % length
    steps_to_new_tail = length - k
    
    new_tail = head
    for _ in range(steps_to_new_tail - 1):
        new_tail = new_tail.next
    
    new_head = new_tail.next
    new_tail.next = None
    
    return new_head

# Example usage: k = 2
# Input:  1 -> 2 -> 3 -> 4 -> 5 -> None
# Output: 4 -> 5 -> 1 -> 2 -> 3 -> None
```

### Pattern 5: Palindrome Linked List
**Problem**: Check if a linked list is a palindrome.

```python
def is_palindrome(head):
    if not head or not head.next:
        return True
    
    # Find the middle of the linked list
    slow = fast = head
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
    
    # Reverse the second half
    def reverse_list(node):
        prev = None
        current = node
        while current:
            next_temp = current.next
            current.next = prev
            prev = current
            current = next_temp
        return prev
    
    second_half = reverse_list(slow)
    
    # Compare both halves
    first_half = head
    while second_half:
        if first_half.val != second_half.val:
            return False
        first_half = first_half.next
        second_half = second_half.next
    
    return True

# Example usage
# Input:  1 -> 2 -> 2 -> 1 -> None
# Output: True
```

### Pattern 6: Add Two Numbers (Reverse Order)
**Problem**: Add two numbers represented as linked lists (digits in reverse order).

```python
def add_two_numbers(l1, l2):
    dummy = ListNode(0)
    current = dummy
    carry = 0
    
    while l1 or l2 or carry:
        val1 = l1.val if l1 else 0
        val2 = l2.val if l2 else 0
        
        total = val1 + val2 + carry
        carry = total // 10
        current.next = ListNode(total % 10)
        
        current = current.next
        l1 = l1.next if l1 else None
        l2 = l2.next if l2 else None
    
    return dummy.next

# Example usage
# l1: 2 -> 4 -> 3 (represents 342)
# l2: 5 -> 6 -> 4 (represents 465)
# Output: 7 -> 0 -> 8 (represents 807)
```

### Pattern 7: Reverse Alternate K Nodes
**Problem**: Reverse alternate groups of k nodes.

```python
def reverse_alternate_k_group(head, k):
    def reverse_group(node, count):
        prev = None
        current = node
        
        while current and count > 0:
            next_temp = current.next
            current.next = prev
            prev = current
            current = next_temp
            count -= 1
        
        return prev, current
    
    def skip_k_nodes(node, k):
        while node and k > 0:
            node = node.next
            k -= 1
        return node
    
    if not head or k <= 1:
        return head
    
    # Reverse first k nodes
    new_head, next_node = reverse_group(head, k)
    head.next = next_node
    
    current = next_node
    while current:
        # Skip k nodes
        current = skip_k_nodes(current, k)
        
        if not current:
            break
        
        # Reverse next k nodes
        prev_tail = current
        for _ in range(k - 1):
            if not prev_tail:
                break
            prev_tail = prev_tail.next
        
        if not prev_tail:
            break
        
        next_start = prev_tail.next
        reversed_head, _ = reverse_group(current, k)
        
        # Find the end of current group after reversal
        tail = reversed_head
        for _ in range(k - 1):
            tail = tail.next
        
        tail.next = next_start
        current = next_start
    
    return new_head

# Example usage: k = 2
# Input:  1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> None
# Output: 2 -> 1 -> 3 -> 4 -> 6 -> 5 -> 7 -> 8 -> None
```

### Pattern 8: Swap Nodes in Pairs
**Problem**: Swap every two adjacent nodes.

```python
def swap_pairs(head):
    dummy = ListNode(0)
    dummy.next = head
    prev = dummy
    
    while prev.next and prev.next.next:
        # Nodes to be swapped
        first = prev.next
        second = prev.next.next
        
        # Swapping
        prev.next = second
        first.next = second.next
        second.next = first
        
        # Move prev pointer
        prev = first
    
    return dummy.next

# Example usage
# Input:  1 -> 2 -> 3 -> 4 -> None
# Output: 2 -> 1 -> 4 -> 3 -> None
```

## Practice Problems

### Easy
1. **Reverse Linked List** - Basic reversal
2. **Palindrome Linked List** - Check palindrome using reversal
3. **Middle of Linked List** - Find middle node

### Medium
1. **Reverse Linked List II** - Reverse sublist
2. **Swap Nodes in Pairs** - Swap adjacent nodes
3. **Rotate List** - Rotate by k positions
4. **Add Two Numbers** - Addition with reversal
5. **Reverse Nodes in k-Group** - Group reversal

### Hard
1. **Reverse Alternate K Nodes** - Complex reversal pattern
2. **Clone List with Random Pointer** - Deep copy with reversal
3. **Merge k Sorted Lists** - Uses reversal concepts
4. **LRU Cache** - Uses doubly linked list operations

## Tips and Tricks

1. **Three Pointers**: Use prev, current, and next for iterative reversal
2. **Dummy Nodes**: Use dummy nodes to simplify edge cases
3. **Draw It Out**: Visualize the links being changed
4. **Boundary Cases**: Handle null lists and single node lists
5. **Two-Phase Approach**: Sometimes reverse, then process, then reverse back

## Common Mistakes

1. **Lost References**: Losing reference to next node before changing links
2. **Null Pointer Errors**: Not checking for null before accessing next
3. **Infinite Loops**: Creating cycles accidentally
4. **Wrong Return Value**: Returning old head instead of new head
5. **Index Errors**: Off-by-one errors in position-based reversals

## Related Patterns

- **Fast & Slow Pointers**: Used to find middle for palindrome checks
- **Two Pointers**: Used in various list manipulation problems
- **Stack**: Alternative approach for some reversal problems
- **Recursion**: Recursive approach to list reversal

## Implementation Languages

The pattern works across languages with pointer/reference support:
- **Python**: Use object references
- **Java**: Object references with careful null checking
- **JavaScript**: Object references similar to Python
- **C++**: Raw pointers or smart pointers
- **Go**: Pointers with automatic memory management