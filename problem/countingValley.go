package problem

import "fmt"

// link problem:
// https://www.hackerrank.com/challenges/counting-valleys/problem?h_r=next-challenge&h_v=zen&isFullScreen=false&h_r=next-challenge&h_v=zen

func countingValleys(steps int32, path string) int32 {
	// Write your code here
	var output int32
	var condition int32
	for i := 0; i < len(path); i++ {
		a := condition
		if string(path[i]) == "D" {
			condition--
		} else if string(path[i]) == "U" {
			condition++
		}
		if a == 0 && condition == -1 {
			output++
		}
	}
	return output
}

func Main_countingValley() {
	path := "UDDDUDUU"
	step := len(path)
	fmt.Println("==============countingValley=============")
	fmt.Println("Input\t: path = ", path)
	fmt.Println("\t  step = ", step)
	fmt.Println("Output\t:", countingValleys(int32(step), path))
}
