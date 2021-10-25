package problem

import "fmt"

func compareTriplets(a []int32, b []int32) []int32 {
	// Write your code here
	var output []int32
	output = append(output, 0)
	output = append(output, 0)
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] < b[i] {
			output[1]++
		} else if a[i] > b[i] {
			output[0]++
		}
	}
	return output
}

func Main_compareTriplets() {
	a := []int32{17, 28, 30}
	b := []int32{99, 16, 8}
	fmt.Println("==============simpleArraySum=============")
	fmt.Println("Input\t: a = ", a, "\n\t  b = ", b)
	fmt.Println("Output\t: ", compareTriplets(a, b))
}
