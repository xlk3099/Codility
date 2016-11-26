package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	currentSum := A[0]
	max := A[0]

	for i := 1; i < len(A); i++ {
		if currentSum < 0 {
			currentSum = A[i]
		} else {
			currentSum += A[i]
		}
		if max < currentSum {
			max = currentSum
		}
	}
	return max

}
