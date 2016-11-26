package solution

// you can also use imports, for example:
import "fmt"

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(H []int) int {
	// write your code in Go 1.4
	stack := []int{}
	stack = append(stack, H[0])
	n := len(H)
	result := 1
	for i := 1; i < n; i++ {
		fmt.Println(stack[len(stack)-1])
		for len(stack) > 0 && stack[len(stack)-1] > H[i] {
			stack = stack[:len(stack)-1]

		}
		if len(stack) > 0 && stack[len(stack)-1] == H[i] {
			continue
		}
		stack = append(stack, H[i])
		result++
	}
	return result
}
