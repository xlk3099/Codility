package solution

//Determine whether a triangle can be built from a given set of edges.
// you can also use imports, for example:
// import "fmt"
// import "os"
import "sort"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	sort.Ints(A)
	n := len(A)
	result := 0
	if n <= 2 || A[n-3] <= 0 {
		return result
	}
	for i := n - 1; i > 1; i-- {
		if A[i] < A[i-1]+A[i-2] {
			result = 1
		}
	}
	return result

}
