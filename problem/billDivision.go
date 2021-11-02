package problem

import "fmt"

// link problem:
// https://www.hackerrank.com/challenges/divisible-sum-pairs/problem?isFullScreen=true

func bonAppetit(bill []int32, k int32, b int32) {
	// Write your code here
	var annaBill []int32
	for i := 0; i < len(bill); i++ {
		if i != int(k) {
			c := bill[i]
			annaBill = append(annaBill, c)
		}
	}
	var sum int32
	for _, v := range annaBill {
		sum += v
	}
	bActual := sum / int32(2)
	if b-bActual == 0 {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(b - bActual)
	}
}

func Main_billDivision() {
	a := []int32{3, 10, 2, 9}
	k := 1
	b := 7
	fmt.Println("==============billDivision=============")
	fmt.Println("Input\t: a = ", a)
	fmt.Println("\t  k = ", k)
	fmt.Println("\t  b = ", b)
	fmt.Println("Output\t:")
	bonAppetit(a, int32(k), int32(b))
}
