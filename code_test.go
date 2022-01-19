package code

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// assume that input is sorted
func bruteforce(input []int) (int, [][3]int) {
	n := 0
	mmap := make(map[[3]int]bool)
	var out [][3]int
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			for k := j + 1; k < len(input); k++ {
				n++
				if input[i]+input[j]+input[k] == 0 && i != j && i != k && j != k {
					rs := [3]int{input[i], input[j], input[k]}
					if !mmap[rs] {
						mmap[rs] = true
						out = append(out, rs)
					}
				}
			}
		}
	}
	return n, out
}

// assume that input is sorted
// 1) Find first positive index
// 2) check negative part with two negative numbers and one positive
// 3) check positive part with two positive and one negative number

func withSplit(input []int) (int, [][3]int) {
	positiveIndex := 0
	for n := range input {
		if input[n] >= 0 {
			positiveIndex = n
			break
		}
	}

	n := 0
	mmap := make(map[[3]int]bool)
	var out [][3]int

	// negative part
	for i := 0; i < positiveIndex; i++ {
		for j := i + 1; j < positiveIndex; j++ {
			for k := positiveIndex; k < len(input); k++ {
				n++
				if input[i]+input[j]+input[k] == 0 && i != j && i != k && j != k {
					rs := [3]int{input[i], input[j], input[k]}
					if !mmap[rs] {
						mmap[rs] = true
						out = append(out, rs)
					}
				}
				// no need to go further
				if input[i]+input[j]+input[k] > 0 {
					break
				}
			}
		}
	}

	// positive part
	for i := positiveIndex; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			for k := positiveIndex - 1; k >= 0; k-- {
				n++
				if input[i]+input[j]+input[k] == 0 && i != j && i != k && j != k {
					rs := [3]int{input[k], input[i], input[j]}
					if !mmap[rs] {
						mmap[rs] = true
						out = append(out, rs)
					}
				}
				// no need to go further
				if input[i]+input[j]+input[k] < 0 {
					break
				}
			}
		}
	}
	return n, out
}

// sorted array
var arr = []int{-4, -1, -1, 0, 1, 2}

//var arr = []int{-6, -5, -4, -1, -1, 0, 1, 2, 3, 4} // on this array it will be 120 iteration on bruteforce and 89 with split

func TestBruteforce(t *testing.T) {
	n, rs := bruteforce(arr)
	require.Equal(t, 20, n) // num of iterations
	require.Equal(t, [][3]int{{-1, -1, 2}, {-1, 0, 1}}, rs)

}

func TestWithSplit(t *testing.T) {
	n, rs := withSplit(arr)
	require.Equal(t, 18, n) // num of iterations
	require.Equal(t, [][3]int{{-1, -1, 2}, {-1, 0, 1}}, rs)

}
