package problem

import "fmt"

// link problem:
// https://www.hackerrank.com/challenges/divisible-sum-pairs/problem?isFullScreen=true

var iteration int32

func minOnSlice(A []int32) ([]int32, int32) {
	min := A[0]
	for _, v := range A {
		if v <= min {
			min = v
		}
	}
	var As []int32
	var count int32
	for _, w := range A {
		if w != min {
			As = append(As, w)
		} else if count == 1 {
			As = append(As, w)
		} else {
			count++
		}
	}
	return As, min
}

func cookies(k int32, A []int32) int32 {
	// Write your code here
	if len(A) == 2 && A[0]+2*A[1] < k {
		return -1
	}
	As1, min1 := minOnSlice(A)
	As2, min2 := minOnSlice(As1)
	if min1 < k {
		ap := min1 + 2*min2
		As2 = append(As2, ap)
		iteration++
	} else {
		// fmt.Println(As2)
		return iteration
	}
	fmt.Println(As2)
	return cookies(k, As2)
}

func Main_jesseAndCookies() {
	a := []int32{1, 1, 1}
	k := 10
	fmt.Println("==============jesse&Cookies=============")
	fmt.Println("Input\t: A = ", a)
	fmt.Println("\t  k = ", k)
	fmt.Println("Output\t:", cookies(int32(k), a))
}

// 3 10
// 1 1 1

// 8 8467293
// 13 47 74 12 89 74 18 38
