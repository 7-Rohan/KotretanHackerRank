package problem

import (
	"fmt"
)

// link:
// https://www.hackerrank.com/challenges/sock-merchant/problem?isFullScreen=true

func sockMerchant(n int32, ar []int32) int32 {
	// Write your code here
	mapWadah := make(map[int32]int32)
	for _, v := range ar {
		_, ok := mapWadah[v]
		if ok {
			mapWadah[v]++
		} else {
			mapWadah[v] = 1
		}
	}
	var output int32
	for _, v := range mapWadah {
		output += v / 2
	}
	return output
}

func Main_salesByMatch() {
	s := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	n := len(s)
	fmt.Println("==============sockMerchant=============")
	fmt.Println("Input\t: a = ", s)
	fmt.Println("Output\t: ", sockMerchant(int32(n), s))
}
