package solution

// Thanks to Falk Hüffner's idea: http://stackoverflow.com/questions/4801242/algorithm-to-calculate-number-of-intersecting-discs
// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	s := len(A)
	start := make([]int, s)
	end := make([]int, s)
	for i := 0; i < s; i++ {
		if (i - A[i]) < 0 {
			start[0]++
		} else {
			start[i-A[i]]++
		}
		if i+A[i] >= s {
			end[s-1]++
		} else {
			end[i+A[i]]++
		}
	}

	active := 0
	sum := 0
	for i := 0; i < s; i++ {
		sum += active*start[i] + start[i]*(start[i]-1)/2
		if sum > 10000000 {
			return -1
		}
		active += start[i] - end[i]
	}
	return sum
}
