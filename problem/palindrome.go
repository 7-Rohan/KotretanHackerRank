package problem

import (
	"fmt"
)

func palindrome(s string) int32 {
	// Write your code here
	count := int32(1)
	for i := 1; i < len(s)-1; i++ {
		for j := 0; j <= i && j+i < len(s); j++ {
			if s[i-j] == s[i+j] {
				count++
			}
		}
	}
	return count
}

func Main_palindrome() {
	s := "aaababc"
	fmt.Println("==============palindrome=============")
	fmt.Println("Input\t: a = ", s)
	fmt.Println("Output\t: ", palindrome(s))
}
