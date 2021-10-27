package problem

import "fmt"

// link problem:
// https://www.hackerrank.com/challenges/breaking-best-and-worst-records/problem?isFullScreen=true&h_r=next-challenge&h_v=zen&h_r=next-challenge&h_v=zen

func breakingRecords(scores []int32) []int32 {
	// Write your code here
	output := []int32{0, 0}
	max := scores[0]
	min := scores[0]
	for _, v := range scores {
		if v > max {
			output[0]++
			max = v
		}
		if v < min {
			output[1]++
			min = v
		}
	}
	return output
}

func Main_breakingRecords() {
	a := []int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42}
	fmt.Println("==============breakingTheRecords=============")
	fmt.Println("Input\t: a = ", a)
	fmt.Println("Output\t: ", breakingRecords(a))
}
