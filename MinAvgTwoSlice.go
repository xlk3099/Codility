package solution

// you can also use imports, for example:
//import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	result := 0
	min := float64(1000001)
	for i := 0; i < len(A)-1; i++ {
		if float64(A[i]+A[i+1])/2.0 < min {
			min = float64((A[i] + A[i+1]) / 2.0)
			result = i
		}
		if i < len(A)-2 && float64(A[i]+A[i+1]+A[i+2])/3.0 < min {
			min = float64(A[i]+A[i+1]+A[i+2]) / 3.0
			result = i
		}
	}
	return result
}
