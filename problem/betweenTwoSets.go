package problem

import "fmt"

// link problem:
// https://www.hackerrank.com/challenges/between-two-sets/problem?isFullScreen=true&h_r=next-challenge&h_v=zen

func getTotalX(a []int32, b []int32) int32 {
	// Write your code here
	var maxa int32
	var minb int32
	minb = b[0]
	for _, v := range a {
		if v >= maxa {
			maxa = v
		}
	}
	for _, w := range b {
		if w <= minb {
			minb = w
		}
	}
	var arrOutput []int32
	for i := 0; i <= int(minb-maxa); i++ {
		p := maxa
		p += int32(i)
		if p <= minb {
			passa := 0
			passb := 0
			for _, x := range b {
				if x%p == 0 {
					passb++
				}
			}
			for _, y := range a {
				if p%y == 0 {
					passa++
				}
			}
			if passa == len(a) && passb == len(b) {
				arrOutput = append(arrOutput, p)
			}
		}
	}
	fmt.Println(arrOutput)
	return int32(len(arrOutput))
}

func Main_betweenTwoSets() {
	a := []int32{3, 4}
	b := []int32{24, 48}
	fmt.Println("==============BetweenTwoSets=============")
	fmt.Println("Input\t: a = ", a)
	fmt.Println("\t  b = ", b)
	fmt.Println("Output\t: ", getTotalX(a, b))
}
