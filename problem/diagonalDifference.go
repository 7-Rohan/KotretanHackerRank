package problem

import (
	"fmt"
	"math"
)

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	var sum1 int32
	var sum2 int32
	for i := 0; i < len(arr); i++ {
		sum1 += arr[i][i]
		sum2 += arr[len(arr)-i-1][i]
	}
	return int32(math.Abs(float64(sum1) - float64(sum2)))
}

func Main_diagonalDifference() {
	arr := [][]int32{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}

	fmt.Println("==============diagonalDifference=============")
	fmt.Println("Input\t: a = ", arr)
	fmt.Println("Output\t: ", diagonalDifference(arr))
	fmt.Println("")
}
