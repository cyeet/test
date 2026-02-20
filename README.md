# test

## Merge Sort

A generic, stable merge sort implementation for any [ordered type](sort/mergesort.go) lives in the `sort` package.

### Usage

```go
import msort "example/hello/sort"

sorted := msort.Merge([]int{5, 2, 8, 1, 9, 3})
// sorted == [1 2 3 5 8 9]

words := msort.Merge([]string{"banana", "apple", "cherry"})
// words == ["apple" "banana" "cherry"]
```

### Complexity

| Property | Value |
|---|---|
| Time (all cases) | O(n log n) |
| Space (auxiliary) | O(n) |
| Stable | ✓ |

### Running tests

```sh
go test ./...
```