package problem

import "fmt"

// link problem:
// https://www.hackerrank.com/challenges/the-birthday-bar/problem?isFullScreen=true&h_r=next-challenge&h_v=zen&h_r=next-challenge&h_v=zen&h_r=next-challenge&h_v=zen

func birthday(s []int32, d int32, m int32) int32 {
	// Write your code here
	var output int32
	for i := 0; i <= len(s)-int(m); i++ {
		sum := 0
		for _, v := range s[i : i+int(m)] {
			sum += int(v)
		}
		if sum == int(d) {
			output++
		}
	}
	return output
}

func Main_subarrayDivision() {
	a := []int32{1, 1, 1, 1, 1, 1}
	d := 3
	m := 2
	fmt.Println("==============subarrayDivision=============")
	fmt.Println("Input\t: a = ", a)
	fmt.Println("\t  d = ", d)
	fmt.Println("\t  m = ", m)
	fmt.Println("Output\t:", birthday(a, int32(d), int32(m)))
}
