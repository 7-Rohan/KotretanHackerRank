package problem

import "fmt"

// link problem:
// https://www.hackerrank.com/challenges/kangaroo/problem?isFullScreen=true

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// Write your code here
	if v1 > v2 {
		var w1, w2 int32
		w1 = x1
		w2 = x2
		for w1 <= w2 {
			w1 += v1
			w2 += v2
			if w1 == w2 {
				return "YES"
			}
		}
	}
	return "NO"
}

func Main_kangaroo() {
	x1 := int32(0)
	v1 := int32(2)
	x2 := int32(5)
	v2 := int32(3)
	fmt.Println("==============kangaroo=============")
	fmt.Println("Input\t: x1 = ", x1)
	fmt.Println("\t  v1 = ", v1)
	fmt.Println("\t  x2 = ", x2)
	fmt.Println("\t  v2 = ", v2)
	fmt.Println("Output\t: ", kangaroo(x1, v1, x2, v2))

}
