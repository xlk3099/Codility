package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import "sort"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	// Do sort first
	sort.Ints(A)
	n := len(A)
	result := A[n-1] * A[n-2] * A[n-3]
	if A[0]*A[1]*A[n-1] > A[n-1]*A[n-2]*A[n-3] {
		result = A[0] * A[1] * A[n-1]
	}
	return result
}
