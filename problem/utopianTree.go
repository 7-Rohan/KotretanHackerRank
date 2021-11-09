package problem

import (
	"fmt"
)

// link problem:
// https://www.hackerrank.com/challenges/utopian-tree/problem?isFullScreen=false&h_r=next-challenge&h_v=zen&h_r=next-challenge&h_v=zen

func utopianTree(n int32) int32 {
	// Write your code here
	var output int32
	for i := 0; i <= int(n); i++ {
		if i%2 == 0 {
			output++
		} else {
			output *= 2
		}
	}
	return output
}

func Main_utopianTree() {
	n := 4

	fmt.Println("==============utopianTree=============")
	fmt.Println("Input\t: n = ", n)
	fmt.Println("Output\t:", utopianTree(int32(n)))
}
