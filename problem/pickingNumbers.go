package problem

import (
	"fmt"
)

// link:
// https://www.hackerrank.com/challenges/picking-numbers/problem?isFullScreen=true

// Complete the catAndMouse function below.
func pickingNumbers(a []int32) int32 {
	// Write your code here
	var output int32
	for _, v := range a {
		var count int32
		for _, w := range a {
			if w-v == 1 || w-v == 0 {
				count++
			}
		}
		if count >= output {
			output = count
		}
	}
	return output
}

func Main_pickingNumbers() {
	a := []int32{1, 2, 2, 3, 1, 2}
	fmt.Println("==============pickingNumbers=============")
	fmt.Println("Input\t: a = ", a)
	fmt.Println("Output\t:", pickingNumbers(a))
}
