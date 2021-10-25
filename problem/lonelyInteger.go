package problem

import "fmt"

func lonelyinteger(a []int32) int32 {
	// Write your code here
	var alone []bool
	for i := 0; i < len(a); i++ {
		alone = append(alone, false)
	}
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				alone[i] = true
				alone[j] = true
			}
		}
	}
	var index int
	for p, v := range alone {
		if !v {
			index = p
		}
	}
	return a[index]
}

func Main_LonelyInteger() {
	a := []int32{1, 2, 3, 4, 3, 2, 1}

	fmt.Println("==============lonelyIngeter=============")
	fmt.Println("Input\t: a = ", a)
	fmt.Println("Output\t: ", lonelyinteger(a))
	fmt.Println("")
}
