package problem

import "fmt"

// link problem:
// https://www.hackerrank.com/challenges/electronics-shop/problem?h_r=next-challenge&h_v=zen&isFullScreen=false&h_r=next-challenge&h_v=zen&h_r=next-challenge&h_v=zen

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	/*
	 * Write your code here.
	 */
	var output int32
	for _, v := range keyboards {
		for _, w := range drives {
			if w+v > output && w+v <= b {
				output = w + v
			}
		}
	}
	if output == 0 {
		return -1
	} else {
		return output
	}
}

func Main_electronicsShop() {
	b := 5
	keyboards := []int32{4}
	drives := []int32{5}
	fmt.Println("==============electronicsShop=============")
	fmt.Println("Input\t: keyboards = ", keyboards)
	fmt.Println("\t  drives = ", drives)
	fmt.Println("\t  b = ", b)
	fmt.Println("Output\t:", getMoneySpent(keyboards, drives, int32(b)))
}

// kodingan sort []int32 dll

// type foo []int32

// func (f foo) Len() int {
// 	return len(f)
// }

// func (f foo) Less(i, j int) bool {
// 	return f[i] > f[j]
// }

// func (f foo) Swap(i, j int) {
// 	f[i], f[j] = f[j], f[i]
// }
