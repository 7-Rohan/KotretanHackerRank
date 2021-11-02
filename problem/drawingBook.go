package problem

import (
	"fmt"
)

// link problem:
// https://www.hackerrank.com/challenges/drawing-book/problem?isFullScreen=true

func pageCount(n int32, p int32) int32 {
	// Write your code here
	if p/2 <= n/2-p/2 {
		return p / 2
	} else {
		return n/2 - p/2
	}
}

func Main_drawingBook() {
	n := 5
	p := 4
	fmt.Println("==============drawingBook=============")
	fmt.Println("Input\t: n = ", n)
	fmt.Println("\t  p = ", p)
	fmt.Println("Output\t: ", pageCount(int32(n), int32(p)))
}
