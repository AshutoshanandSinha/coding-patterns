# Fast & Slow Pointers Pattern (Floyd's Cycle Detection)

## Overview
The Fast & Slow Pointers pattern, also known as the "Tortoise and Hare" algorithm or Floyd's Cycle Detection, uses two pointers moving at different speeds to detect cycles in data structures, find middle elements, or solve problems related to circular patterns.

## When to Use
- **Cycle Detection**: Detecting cycles in linked lists or sequences
- **Finding Middle**: Finding the middle of a linked list
- **Palindrome Check**: Checking if a linked list is a palindrome
- **Cycle Length**: Finding the length of a cycle
- **Start of Cycle**: Finding where a cycle begins
- **Nth from End**: Finding the nth node from the end

## Time/Space Complexity
- **Time**: O(n) - Each node visited at most twice
- **Space**: O(1) - Only two pointers used

## Core Concept
```python
def has_cycle_template(head):
    if not head or not head.next:
        return False
    
    slow = head      # Moves 1 step at a time
    fast = head      # Moves 2 steps at a time
    
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
        
        if slow == fast:  # Cycle detected
            return True
    
    return False
```

## Common Problem Patterns

### Pattern 1: Linked List Cycle Detection
**Problem**: Detect if a linked list has a cycle.

```python
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

def has_cycle(head):
    if not head or not head.next:
        return False
    
    slow = fast = head
    
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
        
        if slow == fast:
            return True
    
    return False

# Example usage
# Creating a cycle: 1 -> 2 -> 3 -> 4 -> 2 (cycle)
node1 = ListNode(1)
node2 = ListNode(2)
node3 = ListNode(3)
node4 = ListNode(4)

node1.next = node2
node2.next = node3
node3.next = node4
node4.next = node2  # Creates cycle

print(has_cycle(node1))  # Output: True
```

### Pattern 2: Find Start of Cycle
**Problem**: Find the node where the cycle begins.

```python
def detect_cycle_start(head):
    if not head or not head.next:
        return None
    
    # Phase 1: Detect if cycle exists
    slow = fast = head
    
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
        
        if slow == fast:
            break
    else:
        return None  # No cycle found
    
    # Phase 2: Find cycle start
    # Move one pointer to head, keep other at meeting point
    # Move both at same speed until they meet
    slow = head
    while slow != fast:
        slow = slow.next
        fast = fast.next
    
    return slow

# Example usage
# Same cycle as above
start_node = detect_cycle_start(node1)
print(start_node.val if start_node else None)  # Output: 2
```

### Pattern 3: Find Middle of Linked List
**Problem**: Find the middle node of a linked list.

```python
def find_middle(head):
    if not head:
        return None
    
    slow = fast = head
    
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
    
    return slow

# Example usage
# Creating list: 1 -> 2 -> 3 -> 4 -> 5
head = ListNode(1)
head.next = ListNode(2)
head.next.next = ListNode(3)
head.next.next.next = ListNode(4)
head.next.next.next.next = ListNode(5)

middle = find_middle(head)
print(middle.val)  # Output: 3
```

### Pattern 4: Palindrome Linked List
**Problem**: Check if a linked list is a palindrome.

```python
def is_palindrome_linked_list(head):
    if not head or not head.next:
        return True
    
    # Step 1: Find middle using fast & slow pointers
    slow = fast = head
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
    
    # Step 2: Reverse second half
    def reverse_list(node):
        prev = None
        while node:
            next_node = node.next
            node.next = prev
            prev = node
            node = next_node
        return prev
    
    second_half = reverse_list(slow)
    
    # Step 3: Compare first and second half
    first_half = head
    while second_half:
        if first_half.val != second_half.val:
            return False
        first_half = first_half.next
        second_half = second_half.next
    
    return True

# Example usage
# Creating palindrome: 1 -> 2 -> 3 -> 2 -> 1
head = ListNode(1)
head.next = ListNode(2)
head.next.next = ListNode(3)
head.next.next.next = ListNode(2)
head.next.next.next.next = ListNode(1)

print(is_palindrome_linked_list(head))  # Output: True
```

### Pattern 5: Happy Number
**Problem**: Determine if a number is happy (sum of squares of digits eventually leads to 1).

```python
def is_happy_number(n):
    def get_sum_of_squares(num):
        total = 0
        while num > 0:
            digit = num % 10
            total += digit * digit
            num //= 10
        return total
    
    slow = fast = n
    
    while True:
        slow = get_sum_of_squares(slow)
        fast = get_sum_of_squares(get_sum_of_squares(fast))
        
        if fast == 1:
            return True
        
        if slow == fast:  # Cycle detected, not happy
            return False

# Example usage
print(is_happy_number(19))  # Output: True
print(is_happy_number(2))   # Output: False
```

### Pattern 6: Remove Nth Node From End
**Problem**: Remove the nth node from the end of a linked list.

```python
def remove_nth_from_end(head, n):
    # Use two pointers with n gap between them
    dummy = ListNode(0)
    dummy.next = head
    
    fast = slow = dummy
    
    # Move fast pointer n+1 steps ahead
    for _ in range(n + 1):
        fast = fast.next
    
    # Move both pointers until fast reaches end
    while fast:
        slow = slow.next
        fast = fast.next
    
    # Remove nth node
    slow.next = slow.next.next
    
    return dummy.next

# Example usage
# Creating list: 1 -> 2 -> 3 -> 4 -> 5
head = ListNode(1)
head.next = ListNode(2)
head.next.next = ListNode(3)
head.next.next.next = ListNode(4)
head.next.next.next.next = ListNode(5)

# Remove 2nd from end (node with value 4)
new_head = remove_nth_from_end(head, 2)
# Result: 1 -> 2 -> 3 -> 5
```

### Pattern 7: Circular Array Loop
**Problem**: Check if there's a circular loop in an array with specific conditions.

```python
def circular_array_loop(nums):
    def next_index(i):
        return (i + nums[i]) % len(nums)
    
    for i in range(len(nums)):
        if nums[i] == 0:
            continue
            
        slow = fast = i
        
        # Check if all moves are in same direction
        forward = nums[i] > 0
        
        while True:
            # Move slow pointer
            slow = next_index(slow)
            if nums[slow] * nums[i] <= 0:  # Direction change
                break
                
            # Move fast pointer twice
            fast = next_index(fast)
            if nums[fast] * nums[i] <= 0:  # Direction change
                break
                
            fast = next_index(fast)
            if nums[fast] * nums[i] <= 0:  # Direction change
                break
            
            # Check if cycle found
            if slow == fast:
                # Check if cycle length > 1
                if slow == next_index(slow):
                    break  # Single element cycle
                return True
        
        # Mark visited elements
        slow = i
        while nums[slow] * nums[i] > 0:
            next_slow = next_index(slow)
            nums[slow] = 0
            slow = next_slow
    
    return False

# Example usage
nums = [2, -1, 1, 2, 2]
print(circular_array_loop(nums))  # Output: True
```

## Practice Problems

### Easy
1. **Linked List Cycle** - Detect cycle in linked list
2. **Middle of Linked List** - Find middle node
3. **Happy Number** - Check if number is happy

### Medium
1. **Linked List Cycle II** - Find start of cycle
2. **Remove Nth Node From End** - Remove specific node
3. **Palindrome Linked List** - Check if list is palindrome
4. **Reorder List** - Reorder list in specific pattern

### Hard
1. **Circular Array Loop** - Detect cycle in array with conditions
2. **Find Duplicate Number** - Find duplicate in array using cycle detection
3. **Intersection of Two Linked Lists** - Find intersection point

## Mathematical Foundation

### Why Does Floyd's Algorithm Work?

If there's a cycle:
- Let's say cycle length is `C`
- When slow pointer enters cycle, fast pointer is already `k` steps into the cycle
- Fast pointer gains 1 step per iteration on slow pointer
- They will meet after `C - k` steps (when fast catches up)

### Finding Cycle Start:
- When pointers meet, slow has traveled `d` steps
- Fast has traveled `2d` steps
- Fast has completed `n` full cycles more than slow: `2d = d + n*C`
- Therefore: `d = n*C`
- Distance from start to cycle entry equals distance from meeting point to cycle entry

## Tips and Tricks

1. **Null Checks**: Always check for null pointers before accessing `.next`
2. **Edge Cases**: Handle empty lists and single-node lists
3. **Direction Consistency**: For array problems, ensure consistent direction
4. **Cycle Length**: To find cycle length, keep one pointer fixed after meeting
5. **Multiple Cycles**: Some problems may have multiple potential cycles

## Common Mistakes

1. **Infinite Loops**: Not handling edge cases properly
2. **Wrong Speed**: Using incorrect pointer advancement (both moving 1 step)
3. **Null Pointer Access**: Not checking if pointers are null before accessing
4. **Off-by-One**: Incorrect positioning when finding nth elements
5. **Direction Changes**: Not handling direction changes in array problems

## Template Code

```python
def fast_slow_template(head):
    # Handle edge cases
    if not head or not head.next:
        return handle_edge_case()
    
    slow = fast = head
    
    # Phase 1: Detection
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
        
        if slow == fast:
            # Cycle detected or condition met
            break
    else:
        # No cycle or condition not met
        return no_cycle_result()
    
    # Phase 2: Additional processing if needed
    # (finding start, length, etc.)
    return process_result(slow, fast)
```

## Related Patterns

- **Two Pointers**: Different speeds vs. different directions
- **Sliding Window**: Window movement vs. cycle detection
- **Binary Search**: Different search strategies

## Implementation Notes

- **Python**: Use `and` operator for safe null checking
- **Java**: Check for `null` before accessing `.next`
- **JavaScript**: Use `&&` for safe navigation
- **C++**: Always validate pointers before dereferencing