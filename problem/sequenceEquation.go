package problem

import (
	"fmt"
)

// link problem:
// https://www.hackerrank.com/challenges/permutation-equation/problem?isFullScreen=true&h_r=next-challenge&h_v=zen

func permutationEquation(p []int32) []int32 {
	// Write your code here
	var output []int32
	for i := 0; i < len(p); i++ {
		var asd int
	innerloop1:
		for t, v := range p {
			if v == int32(i)+1 {
				asd = t
				break innerloop1
			}
		}
	innerloop2:
		for u, w := range p {
			if w == int32(asd)+1 {
				output = append(output, int32(u)+1)
				break innerloop2
			}
		}
	}
	return output
}

func Main_sequenceEquation() {
	p := []int32{5, 2, 1, 3, 4}

	fmt.Println("==============sequenceEquation=============")
	fmt.Println("Input\t: p = ", p)
	fmt.Println("Output\t:", permutationEquation(p))
}
