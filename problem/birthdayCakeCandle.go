package problem

import "fmt"

// link problem:
// https://www.hackerrank.com/challenges/birthday-cake-candles/problem?isFullScreen=true

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	var max int
	var count int32
	for _, v := range candles {
		if v >= int32(max) {
			max = int(v)
		}
	}
	for _, v := range candles {
		if v == int32(max) {
			count++
		}
	}
	return count
}

func Main_birthdayCakeCandles() {
	n := []int32{3, 2, 1, 3}
	fmt.Println("==============birthdayCakeCandles=============")
	fmt.Println("Input\t: candles = ", n)
	fmt.Println("Output\t: ", birthdayCakeCandles(n))

}
