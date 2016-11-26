package solution

// you can also use imports, for example:
// import "os"
// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	leftMax := make([]int, len(A))
	rightMax := make([]int, len(A))

	for i := 1; i < len(A)-1; i++ {
		if leftMax[i-1]+A[i] < 0 {
			leftMax[i] = 0
		} else {
			leftMax[i] = leftMax[i-1] + A[i]
		}
	}

	for i := len(A) - 2; i > 1; i-- {
		if rightMax[i+1]+A[i] < 0 {
			rightMax[i] = 0
		} else {
			rightMax[i] = rightMax[i+1] + A[i]
		}
	}
	result := 0
	for i := 1; i < len(A)-1; i++ {
		if current := leftMax[i-1] + rightMax[i+1]; result < current {
			result = current
		}
	}
	return result
}
