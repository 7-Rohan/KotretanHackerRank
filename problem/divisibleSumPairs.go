package problem

import "fmt"

// link problem:
// https://www.hackerrank.com/challenges/divisible-sum-pairs/problem?isFullScreen=true

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	// Write your code here
	var output int32
	for i := 0; i < len(ar); i++ {
		for j := i + 1; j < len(ar); j++ {
			if (ar[i]+ar[j])%k == 0 {
				output++
			}
		}
	}
	return output
}

func Main_divisibleSumPairs() {
	a := []int32{1, 3, 2, 6, 1, 2}
	n := len(a)
	k := 3
	fmt.Println("==============divisibleSumPairs=============")
	fmt.Println("Input\t: a = ", a)
	fmt.Println("\t  n = ", n)
	fmt.Println("\t  k = ", k)
	fmt.Println("Output\t:", divisibleSumPairs(int32(n), int32(k), a))
}
