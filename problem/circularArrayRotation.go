package problem

import (
	"fmt"
)

// link problem:
// https://www.hackerrank.com/challenges/circular-array-rotation/problem?isFullScreen=true

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	// Write your code here
	var output []int32
	var rotated []int32
	for i := 0; i < len(a); i++ {
		p := a[(len(a)+i-int(k))%len(a)]
		rotated = append(rotated, p)
	}
	for _, v := range queries {
		x := rotated[v]
		output = append(output, x)
	}
	return output
}

func Main_circularArrayRotation() {
	a := []int32{1, 2, 3}
	k := 5
	queries := []int32{0, 1, 2}

	fmt.Println("==============circularArrayRotation=============")
	fmt.Println("Input\t: a = ", a)
	fmt.Println("\t  k = ", k)
	fmt.Println("\t  queries = ", queries)
	fmt.Println("Output\t:", circularArrayRotation(a, int32(k), queries))
}
