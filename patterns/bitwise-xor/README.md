# Bitwise XOR Pattern

## Overview
The Bitwise XOR pattern leverages the unique properties of the XOR operation to solve problems involving finding unique elements, detecting duplicates, or manipulating bits efficiently. XOR has special mathematical properties that make it extremely useful for certain types of problems.

## XOR Properties
1. **Commutative**: a ⊕ b = b ⊕ a
2. **Associative**: (a ⊕ b) ⊕ c = a ⊕ (b ⊕ c)
3. **Identity**: a ⊕ 0 = a
4. **Self-Inverse**: a ⊕ a = 0
5. **Transitive**: if a ⊕ b = c, then a ⊕ c = b and b ⊕ c = a

## When to Use
- **Single Unique Element**: Find the one element that appears once while others appear twice
- **Missing Number**: Find missing number in a sequence
- **Bit Manipulation**: Toggle, set, or clear specific bits
- **Space-Efficient Solutions**: Solve problems without extra space
- **Encryption/Decryption**: Simple encryption using XOR operations

## Time/Space Complexity
- **Time**: O(n) - Single pass through data
- **Space**: O(1) - Constant space usage

## Pattern Variations

### 1. Find Single Unique Element
```python
def single_number(nums):
    result = 0
    for num in nums:
        result ^= num
    return result
```

### 2. XOR of Range
```python
def xor_range(start, end):
    def xor_from_1_to_n(n):
        if n % 4 == 1:
            return 1
        elif n % 4 == 2:
            return n + 1
        elif n % 4 == 3:
            return 0
        else:
            return n
    
    return xor_from_1_to_n(end) ^ xor_from_1_to_n(start - 1)
```

## Common Problem Patterns

### Pattern 1: Single Number
**Problem**: Find the number that appears only once while others appear twice.

```python
def single_number(nums):
    result = 0
    for num in nums:
        result ^= num
    return result

# Example usage
nums = [2, 2, 1, 3, 3]
print(single_number(nums))  # Output: 1
```

### Pattern 2: Two Single Numbers
**Problem**: Find two numbers that appear only once while others appear twice.

```python
def single_number_ii(nums):
    # Get XOR of both unique numbers
    xor_all = 0
    for num in nums:
        xor_all ^= num
    
    # Find rightmost set bit
    rightmost_bit = xor_all & (-xor_all)
    
    # Divide numbers into two groups and XOR separately
    num1 = num2 = 0
    for num in nums:
        if num & rightmost_bit:
            num1 ^= num
        else:
            num2 ^= num
    
    return [num1, num2]

# Example usage
nums = [1, 2, 1, 3, 2, 5]
print(single_number_ii(nums))  # Output: [3, 5]
```

### Pattern 3: Missing Number
**Problem**: Find the missing number in an array containing n distinct numbers from 0 to n.

```python
def missing_number(nums):
    n = len(nums)
    expected_xor = 0
    actual_xor = 0
    
    # XOR of all numbers from 0 to n
    for i in range(n + 1):
        expected_xor ^= i
    
    # XOR of all numbers in array
    for num in nums:
        actual_xor ^= num
    
    return expected_xor ^ actual_xor

# Alternative more concise version
def missing_number_v2(nums):
    result = len(nums)
    for i, num in enumerate(nums):
        result ^= i ^ num
    return result

# Example usage
nums = [3, 0, 1]
print(missing_number(nums))  # Output: 2
```

### Pattern 4: Single Number III (Appears Once, Others Thrice)
**Problem**: Find the number that appears once while others appear three times.

```python
def single_number_iii(nums):
    ones = twos = 0
    
    for num in nums:
        # ones stores bits that appear 1 time or 4 times or 7 times...
        # twos stores bits that appear 2 times or 5 times or 8 times...
        ones = (ones ^ num) & ~twos
        twos = (twos ^ num) & ~ones
    
    return ones

# Example usage
nums = [2, 2, 3, 2]
print(single_number_iii(nums))  # Output: 3
```

### Pattern 5: Find Duplicate Number
**Problem**: Find the duplicate number in an array using XOR (specific constraints).

```python
def find_duplicate_xor(nums):
    # This works when we know the range is 1 to n-1 and one number is duplicate
    n = len(nums) - 1
    array_xor = 0
    range_xor = 0
    
    # XOR all array elements
    for num in nums:
        array_xor ^= num
    
    # XOR all numbers from 1 to n
    for i in range(1, n + 1):
        range_xor ^= i
    
    return array_xor ^ range_xor

# Example usage (when array has duplicates of numbers 1 to n-1)
nums = [1, 3, 4, 2, 2]
print(find_duplicate_xor(nums))  # Output: 2
```

### Pattern 6: Complement of Base 10 Integer
**Problem**: Return complement of a given integer.

```python
def find_complement(num):
    # Find the bit length
    bit_length = num.bit_length()
    
    # Create mask with all 1s of the same length
    mask = (1 << bit_length) - 1
    
    # XOR with mask to flip all bits
    return num ^ mask

# Example usage
num = 5  # Binary: 101
print(find_complement(num))  # Output: 2 (Binary: 010)
```

### Pattern 7: Maximum XOR of Two Numbers
**Problem**: Find the maximum XOR of any two numbers in an array.

```python
def find_maximum_xor(nums):
    max_xor = 0
    mask = 0
    
    # Build the answer bit by bit from MSB to LSB
    for i in range(31, -1, -1):
        mask |= (1 << i)  # Add current bit to mask
        prefixes = {num & mask for num in nums}
        
        temp = max_xor | (1 << i)  # Try to make current bit 1
        
        # Check if we can achieve this temp value
        for prefix in prefixes:
            if temp ^ prefix in prefixes:
                max_xor = temp
                break
    
    return max_xor

# Example usage
nums = [3, 10, 5, 25, 2, 8]
print(find_maximum_xor(nums))  # Output: 28
```

### Pattern 8: Subarray XOR Queries
**Problem**: Answer XOR queries on subarrays efficiently.

```python
def xor_queries(arr, queries):
    # Build prefix XOR array
    prefix_xor = [0]
    for num in arr:
        prefix_xor.append(prefix_xor[-1] ^ num)
    
    result = []
    for left, right in queries:
        # XOR of subarray [left, right] = prefix_xor[right+1] ^ prefix_xor[left]
        result.append(prefix_xor[right + 1] ^ prefix_xor[left])
    
    return result

# Example usage
arr = [1, 3, 4, 8]
queries = [[0, 1], [1, 2], [0, 3], [3, 3]]
print(xor_queries(arr, queries))  # Output: [2, 7, 14, 8]
```

## Practice Problems

### Easy
1. **Single Number** - Find unique element among duplicates
2. **Missing Number** - Find missing number in sequence
3. **Number Complement** - Find complement of integer
4. **Hamming Distance** - Count differing bits

### Medium
1. **Single Number II** - Unique element among triplets
2. **Single Number III** - Two unique elements
3. **Maximum XOR of Two Numbers** - Find maximum XOR pair
4. **Subarray XOR Queries** - Range XOR queries

### Hard
1. **Maximum XOR With an Element From Array** - Complex XOR queries
2. **Count Triplets That Can Form Two Arrays of Equal XOR** - Advanced XOR
3. **Minimum XOR Sum of Two Arrays** - Optimization with XOR
4. **Maximum Genetic Difference Query** - Tree + XOR

## Tips and Tricks

1. **XOR Properties**: Master the fundamental properties of XOR
2. **Bit Manipulation**: Understand how to isolate and work with individual bits
3. **Prefix XOR**: Use cumulative XOR for range queries
4. **Grouping Strategy**: Group elements by their XOR properties
5. **Trie for XOR**: Use trie data structure for complex XOR problems

## Common Mistakes

1. **Misunderstanding XOR**: Not fully grasping XOR properties
2. **Bit Length Issues**: Not handling different bit lengths correctly
3. **Overflow**: Not considering integer overflow in some languages
4. **Sign Extension**: Issues with signed integers and bit operations

## Related Patterns

- **Bit Manipulation**: XOR is a fundamental bit operation
- **Hash Tables**: Sometimes XOR can replace hash-based solutions
- **Mathematical Properties**: XOR has unique algebraic properties
- **Trie**: Advanced XOR problems often use trie data structure

## Implementation Languages

XOR works consistently across languages:
- **Python**: Use `^` operator, built-in `bit_length()` method
- **Java**: Use `^` operator, `Integer.bitCount()` for bit operations
- **JavaScript**: Use `^` operator, be careful with 32-bit integer limits
- **C++**: Use `^` operator, `__builtin_clz()` for bit operations
- **Go**: Use `^` operator, `bits` package for advanced operations