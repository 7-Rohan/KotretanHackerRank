package problem

import "fmt"

func simpleArraySum(ar []int32) int32 {
	// Write your code here
	var output int32
	for _, v := range ar {
		output += v
	}
	return output
}

func Main_simpleArraySum() {
	ar := []int32{1, 2, 3, 4, 10, 11}
	fmt.Println("==============simpleArraySum=============")
	fmt.Println("Input\t: ar = ", ar)
	fmt.Println("Output\t: ", simpleArraySum(ar))
}
