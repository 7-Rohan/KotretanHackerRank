package problem

import (
	"fmt"
)

// link problem:
// https://www.hackerrank.com/challenges/jumping-on-the-clouds-revisited/problem?h_r=next-challenge&h_r=next-challenge&h_v=zen&h_v=zen&isFullScreen=false

func jumpingOnClouds(c []int32, k int32) int32 {
	e := int32(100)
	var count int
	for i := int(k); count < 1; i += int(k) {
		if i%len(c) == 0 {
			count++
		}
		if c[i%len(c)] == 1 {
			e -= 2
		}
		e--
		fmt.Println(count)
		fmt.Println(i)
		fmt.Println(i % len(c))
		fmt.Println(e)
		fmt.Println(c[i%len(c)] == 1, c[i%len(c)])
	}
	return e
}

func Main_jumpingOnTheCloud() {
	c := []int32{0, 0, 1, 0, 0, 1, 1, 0}
	k := 2

	fmt.Println("==============jumpingOnTheCloud=============")
	fmt.Println("Input\t: c = ", c)
	fmt.Println("\t  k = ", k)
	fmt.Println("Output\t:", jumpingOnClouds(c, int32(k)))
}
