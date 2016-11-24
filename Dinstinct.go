package solution

// you can also use imports, for example:
// import "os"
import "sort"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	sort.Ints(A)
	if len(A) == 0 {
		return 0
	}
	result := 1
	for i := 1; i < len(A); i++ {
		if A[i] != A[i-1] {
			result++
		}
	}
	return result
}
