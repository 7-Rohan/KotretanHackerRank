package problem

import "fmt"

func miniMaxSum(arr []int32) {
	// Write your code here
	var min int32
	var max int32
	var sum int
	min = arr[0]
	for _, v := range arr {
		sum += int(v)
		if v >= max {
			max = v
		}
		if v <= min {
			min = v
		}
	}
	fmt.Print(sum-int(max), " ", sum-int(min))
}

func Main_miniMaxSum() {
	n := []int32{254961783, 604179258, 462517083, 967304281, 860273491}
	fmt.Println("==============miniMaxSum=============")
	fmt.Println("Input\t: a = ", n)
	fmt.Println("Output\t: ")
	miniMaxSum(n)
}
