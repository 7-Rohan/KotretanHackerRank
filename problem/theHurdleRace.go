package problem

import (
	"fmt"
)

// link problem:
// https://www.hackerrank.com/challenges/the-hurdle-race/problem?isFullScreen=false

func hurdleRace(k int32, height []int32) int32 {
	// Write your code here
	var output int32
	for _, v := range height {
		if v-k > output {
			output = v - k
		}
	}
	return output
}

func Main_theHurdleRace() {
	k := 1
	height := []int32{1, 2, 3, 3, 2}

	fmt.Println("==============theHurdleRace=============")
	fmt.Println("Input\t: height = ", height)
	fmt.Println("\t  k = ", k)
	fmt.Println("Output\t:", hurdleRace(int32(k), height))
}
