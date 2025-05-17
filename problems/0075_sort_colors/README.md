# 75. Sort Colors

## Problem Description

Given an array `nums` with `n` objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

We will use the integers `0`, `1`, and `2` to represent the color red, white, and blue, respectively.

You must solve this problem without using the library's sort function.

## Examples

### Example 1:
```
Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
```

### Example 2:
```
Input: nums = [2,0,1]
Output: [0,1,2]
```

## Constraints

- n == nums.length
- 1 <= n <= 300
- nums[i] is either 0, 1, or 2

## Solution Approach

The solution uses the Dutch National Flag algorithm (three-way partitioning):
1. Use three pointers:
   - p0: points to the rightmost boundary of 0s
   - curr: current element being processed
   - p2: points to the leftmost boundary of 2s
2. For each element:
   - If it's 2, swap with p2 and move p2 left
   - If it's 0, swap with p0 and move both p0 and curr right
   - If it's 1, just move curr right
3. Continue until curr > p2

Time Complexity: O(n)
Space Complexity: O(1)

## Follow-up Questions

1. **Counting Sort**: How would you solve this using counting sort? What would be the trade-offs?

2. **In-place Requirement**: Why is the in-place requirement important? What are the advantages of the current solution over a non-in-place solution?

3. **Stability**: Is this sorting algorithm stable? Why or why not? How would you modify it to make it stable?

4. **Multiple Colors**: How would you modify the solution if there were more than three colors? Would the time/space complexity change?

5. **Parallel Processing**: Could this algorithm be parallelized? What would be the challenges?

6. **Memory Access Patterns**: How does this solution perform in terms of cache efficiency? Could it be optimized further?

7. **Alternative Approaches**: What other approaches could be used to solve this problem? What are their trade-offs?

8. **Real-world Application**: In what real-world scenarios might you need to sort objects by color? How would the requirements differ? 