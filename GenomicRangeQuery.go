package solution

// you can also use imports, for example:
import "fmt"

// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(S string, P []int, Q []int) []int {
	// write your code in Go 1.4
	M := len(P)
	N := len(S)
	rangeA := make([]int, N)
	rangeC := make([]int, N)
	rangeG := make([]int, N)
	rangeT := make([]int, N)
	// Initialize the 4 arrays
	for i := 0; i < N; i++ {
		rangeA[i] = -1
		rangeC[i] = -1
		rangeG[i] = -1
		rangeT[i] = -1
	}

	result := make([]int, M)

	switch S[0] {
	case 'A':
		rangeA[0] = 0
	case 'C':
		rangeC[0] = 0
	case 'G':
		rangeG[0] = 0
	default:
		rangeT[0] = 0
	}
	for i := 1; i < N; i++ {
		switch S[i] {
		case 'A':
			rangeA[i] = i
			rangeC[i] = rangeC[i-1]
			rangeG[i] = rangeG[i-1]
			rangeT[i] = rangeT[i-1]
		case 'C':
			rangeC[i] = i
			rangeA[i] = rangeA[i-1]
			rangeG[i] = rangeG[i-1]
			rangeT[i] = rangeT[i-1]
		case 'G':
			rangeG[i] = i
			rangeA[i] = rangeA[i-1]
			rangeC[i] = rangeC[i-1]
			rangeT[i] = rangeT[i-1]
		default:
			rangeA[i] = rangeA[i-1]
			rangeG[i] = rangeG[i-1]
			rangeC[i] = rangeC[i-1]
			rangeT[i] = i
		}
	}

	for i := 0; i < M; i++ {
		fmt.Println(P[i], Q[i])
		if rangeA[Q[i]] >= P[i] {
			result[i] = 1
			continue
		}
		if rangeC[Q[i]] >= P[i] {
			result[i] = 2
			continue
		}
		if rangeG[Q[i]] >= P[i] {
			result[i] = 3
			continue
		}
		if rangeT[Q[i]] >= P[i] {
			result[i] = 4
		}
	}
	return result
}
