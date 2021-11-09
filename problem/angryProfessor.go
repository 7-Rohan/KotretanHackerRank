package problem

import (
	"fmt"
)

// link problem:
// https://www.hackerrank.com/challenges/angry-professor/problem?isFullScreen=false&h_r=next-challenge&h_v=zen&h_r=next-challenge&h_v=zen&h_r=next-challenge&h_v=zen

func angryProfessor(k int32, a []int32) string {
	// Write your code here
	var count int32
	for _, v := range a {
		if v <= 0 {
			count++
		}
	}
	if count < k {
		return "YES"
	} else {
		return "NO"
	}
}

func Main_angryProfessor() {
	k := 2
	a := []int32{0, -1, 2, 1}

	fmt.Println("==============angryProfessor=============")
	fmt.Println("Input\t: k = ", k)
	fmt.Println("\t  a", a)
	fmt.Println("Output\t:", angryProfessor(int32(k), a))
}
