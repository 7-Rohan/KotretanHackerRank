package problem

import "fmt"

// link problem:
// https://www.hackerrank.com/challenges/migratory-birds/problem?isFullScreen=true

func migratoryBirds(arr []int32) int32 {
	// Write your code here
	var output int32
	var mapCount [6]int32
	for _, v := range arr {
		mapCount[v]++
	}
	var max int32
	for _, w := range mapCount {
		if w >= max {
			max = w
		}
	}
	for i := 5; i >= 0; i-- {
		if mapCount[int32(i)] == max {
			output = int32(i)
		}
	}
	return output
}

func Main_migratoryBirds() {
	a := []int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}
	fmt.Println("==============migratoryBirds=============")
	fmt.Println("Input\t: a = ", a)
	fmt.Println("Output\t:", migratoryBirds(a))
}
