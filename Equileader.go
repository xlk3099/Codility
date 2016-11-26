package solution

// you can also use imports, for example:
//import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	leader := 0
	max := 0
	result := 0
	set := make(map[int]int)
	// 1st find the leader of the whole array
	for _, value := range A {
		if _, okay := set[value]; okay {
			set[value]++
			if max < set[value] {
				max = set[value]
				leader = value
			}
		} else {
			set[value] = 1
		}
	}
	// in case no leader found
	if max < len(A)/2 {
		return -1
	}
	// recount
	left := 0
	for index, value := range A {
		if value == leader {
			left++
		}
		if left > (index+1)/2 && (max-left) > (len(A)-index-1)/2 {
			result++
		}
	}
	return result
}
