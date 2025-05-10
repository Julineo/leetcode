# LeetCode Solutions in Go

This repository contains my solutions to LeetCode problems implemented in Go. Each problem is organized in its own directory with a clear structure for both the solution and test cases. Problem directories are prefixed with their LeetCode problem number for easy navigation.

## Project Structure

```
.
├── problems/           # All LeetCode problems
│   ├── 0001_two_sum/  # Problem #1: Two Sum
│   │   ├── solution.go    # Problem solution
│   │   └── solution_test.go   # Test cases
│   └── ...
└── README.md
```

## Running Tests

You can run tests in several ways:

1. Run all tests:
```bash
go test ./...
```

2. Run tests for a specific problem:
```bash
go test ./problems/0001_two_sum/...
```

3. Run a specific test:
```bash
go test ./problems/0001_two_sum/... -run TestTwoSum
```

## Adding a New Problem

1. Create a new directory under `problems/` with the format `XXXX_problem_name` where XXXX is the LeetCode problem number (padded with zeros)
2. Create `solution.go` with your implementation
3. Create `solution_test.go` with test cases
4. Follow the existing pattern for consistency

## Problem Template

Each problem should follow this structure:

```go
// solution.go
package problem_name

// Solution contains the main solution function
func Solution(input ...interface{}) interface{} {
    // Implementation
}

// solution_test.go
package problem_name

import "testing"

func TestSolution(t *testing.T) {
    // Test cases
}
``` 