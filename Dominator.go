package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	set := make(map[int]int)

	for index, value := range A {
		if _, okay := set[value]; okay {
			set[value]++
			if set[value] > len(A)/2 {
				return index
			}
		} else {
			set[value] = 1
		}
	}
	return -1
}
