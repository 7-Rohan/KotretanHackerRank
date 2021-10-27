package problem

import "fmt"

// link problem:
// https://www.hackerrank.com/challenges/apple-and-orange/problem?h_r=next-challenge&h_v=zen&isFullScreen=false

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	// Write your code here
	var countApples int
	var countOranges int
	for i := 0; i < len(apples); i++ {
		if apples[i]+a <= t && apples[i]+a >= s {
			countApples++
		}
	}
	for j := 0; j < len(oranges); j++ {
		if oranges[j]+b <= t && oranges[j]+b >= s {
			countOranges++
		}
	}
	fmt.Println(countApples)
	fmt.Println(countOranges)
}

func Main_countApplesAndOranges() {
	s := int32(7)
	t := int32(10)
	a := int32(4)
	b := int32(12)
	apples := []int32{2, 3, -4}
	oranges := []int32{3, -2, -4}
	fmt.Println("==============gradingStudents=============")
	fmt.Println("Input\t: s = ", s)
	fmt.Println("\t  t = ", t)
	fmt.Println("\t  a = ", a)
	fmt.Println("\t  b = ", b)
	fmt.Println("Output\t: ")
	countApplesAndOranges(s, t, a, b, apples, oranges)
}
