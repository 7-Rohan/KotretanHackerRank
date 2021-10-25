package problem

import (
	"fmt"
	"strings"
)

func isBalanced(s string) string {
	// Write your code here
	if len(s)%2 != 0 {
		return "NO"
	}
	match := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}
	str := strings.Split(s, "")
	for i := 0; i < len(str)/2; i++ {
		switch str[i] {
		case "{":
			if str[len(str)-i-1] != match["{"] {
				return "NO"
			}
		case "[":
			if str[len(str)-i-1] != match["["] {
				return "NO"
			}
		case "(":
			if str[len(str)-i-1] != match["("] {
				return "NO"
			}
		}
	}
	return "YES"
}

func Main_balancedBracket() {
	s := "{[({})]}()"
	fmt.Println("==============balancedBracket=============")
	fmt.Println("Input\t: a = ", s)
	fmt.Println("Output\t: ", isBalanced(s))
}
