# 1550. Three Consecutive Odds

## Problem Description

Given an integer array `arr`, return `true` if there are three consecutive odd numbers in the array. Otherwise, return `false`.

## Examples

### Example 1:
```
Input: arr = [2,6,4,1]
Output: false
Explanation: There are no three consecutive odds.
```

### Example 2:
```
Input: arr = [1,2,34,3,4,5,7,23,12]
Output: true
Explanation: [5,7,23] are three consecutive odds.
```

## Constraints

- 1 <= arr.length <= 1000
- 1 <= arr[i] <= 1000

## Solution Approach

The solution uses a simple counter approach:
1. Initialize a counter for consecutive odd numbers
2. Iterate through the array
3. For each number:
   - If it's odd (using bitwise AND with 1), increment the counter
   - If we reach 3 consecutive odds, return true
   - If we find an even number, reset the counter to 0
4. Return false if we never find 3 consecutive odds

Time Complexity: O(n)
Space Complexity: O(1)

The solution uses a bitwise operation (`num&1 == 1`) to check if a number is odd, which is more efficient than using modulo (`num%2 != 0`). 