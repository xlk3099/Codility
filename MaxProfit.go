package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	if len(A) == 0 {
		return 0
	}
	// first create an arrary of daily diff
	diff := make([]int, len(A))
	// ignore first day
	diff[0] = 0
	for i := 1; i < len(A); i++ {
		diff[i] = A[i] - A[i-1]
	}
	currentProfit := 0
	max := 0
	for i := 0; i < len(diff); i++ {
		if currentProfit < 0 {
			currentProfit = diff[i]
		} else {
			currentProfit += diff[i]
		}
		if max < currentProfit {
			max = currentProfit
		}
	}
	return max
}
