package solution

// you can also use imports, for example:
//import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, B []int) int {
	// write your code in Go 1.4
	stack := []int{}
	result := 0
	for i := 0; i < len(A); i++ {
		if B[i] == 1 {
			stack = append(stack, A[i])
			continue
		}
		for len(stack) != 0 {
			if stack[len(stack)-1] > A[i] {
				break
			} else {
				stack = stack[:len(stack)-1]
			}
		}
		if len(stack) == 0 && B[i] == 0 {
			result++
		}
		//fmt.Println(result)
	}
	return result + len(stack)
}
