package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(S string) int {
	// write your code in Go 1.4
	result := 0
	if len(S) == 0 {
		return 1
	}
	left := []byte{}
	for _, value := range S {
		if value == '(' || value == '[' || value == '{' {
			left = append(left, byte(value))
		} else {
			if len(left) <= 0 {
				return 0
			}
			top := left[len(left)-1]
			switch value {
			case ')':
				if top != '(' {
					return 0
				}
				left = left[:len(left)-1]
			case ']':
				if top != '[' {
					return 0
				}
				left = left[:len(left)-1]

			case '}':
				if top != '{' {
					return 0
				}
				left = left[:len(left)-1]
			default:
				return result
			}
		}
	}
	if len(left) == 0 {
		result = 1
	}
	return result
}
