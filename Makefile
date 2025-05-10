.PHONY: test test-all test-problem clean

# Run all tests
test-all:
	go test ./...

# Run tests for a specific problem
# Usage: make test-problem PROBLEM=0001_two_sum
test-problem:
	go test ./problems/$(PROBLEM)/...

# Run a specific test
# Usage: make test TEST=TestTwoSum
test:
	go test ./... -run $(TEST)

# Clean up any temporary files
clean:
	go clean
	find . -name "*.test" -type f -delete 