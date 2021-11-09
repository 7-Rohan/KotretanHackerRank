package problem

import (
	"fmt"
)

// link problem:
// https://www.hackerrank.com/challenges/strange-advertising/problem?isFullScreen=false&h_r=next-challenge&h_v=zen&h_r=next-challenge&h_v=zen&h_r=next-challenge&h_v=zen&h_r=next-challenge&h_v=zen&h_r=next-challenge&h_v=zen

func viralAdvertising(n int32) int32 {
	// Write your code here
	var output int32
	a := int32(5)
	for i := 0; i < int(n); i++ {
		output += a / 2
		a = a / 2 * 3
	}
	return output
}

func Main_viralAdvertising() {
	n := 3

	fmt.Println("==============viralAdvertising=============")
	fmt.Println("Input\t: n = ", n)
	fmt.Println("Output\t:", viralAdvertising(int32(n)))
}
