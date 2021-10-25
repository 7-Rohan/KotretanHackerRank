package problem

import "fmt"

func plusMinus(arr []int32) {
	// Write your code here
	var numeratorPlus int32
	var numeratorZero int32
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			numeratorZero++
		} else if arr[i] > 0 {
			numeratorPlus++
		}
	}
	fmt.Printf("%.6f\n", float64(numeratorPlus)/float64(len(arr)))
	fmt.Printf("%.6f\n", float64(int32(len(arr))-numeratorPlus-numeratorZero)/float64(len(arr)))
	fmt.Printf("%.6f\n", float64(numeratorZero)/float64(len(arr)))
}

func Main_plusMinus() {
	a := []int32{17, 28, 0}
	fmt.Println("==============plusMinus=============")
	fmt.Println("Input\t: a = ", a)
	fmt.Println("Output\t: ")
	plusMinus(a)
}
