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
	var wadah []string
	for i := 0; i < len(str); i++ {
		_, ok := match[str[i]]
		if ok {
			a := str[i]
			wadah = append(wadah, a)
		}
		if str[i] == match[wadah[len(wadah)-1]] {
			wadah = wadah[:len(wadah)-1]
		}
	}
	if len(wadah) == 0 {
		return "YES"
	} else {
		return "NO"
	}
}

func Main_balancedBracket() {
	s := "{[({}())]}()"
	fmt.Println("==============balancedBracket=============")
	fmt.Println("Input\t: a = ", s)
	fmt.Println("Output\t: ", isBalanced(s))
}
