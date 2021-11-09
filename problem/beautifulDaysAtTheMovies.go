package problem

import (
	"fmt"
	"strconv"
	"strings"
)

// link problem:
// https://www.hackerrank.com/challenges/beautiful-days-at-the-movies/problem?isFullScreen=false&h_r=next-challenge&h_v=zen&h_r=next-challenge&h_v=zen&h_r=next-challenge&h_v=zen&h_r=next-challenge&h_v=zen

func reverseStr(s string) string {
	var reverse []string

	for i := len(s) - 1; i >= 0; i-- {
		reverse = append(reverse, string(s[i]))
	}
	return strings.Join(reverse, "")
}

func beautifulDays(i int32, j int32, k int32) int32 {
	// Write your code here
	var output int32
	for p := int(i); p <= int(j); p++ {
		iStr := strconv.Itoa(int(p))
		reverseIStr := reverseStr(iStr)
		strI, _ := strconv.Atoi(reverseIStr)
		if (p-strI)%int(k) == 0 {
			output++
		}
	}
	return output
}

func Main_beautifulDaysAtTheMovies() {
	i := 20
	j := 23
	k := 6

	fmt.Println("==============beautifulDaysAtTheMovies=============")
	fmt.Println("Input\t: i = ", i)
	fmt.Println("\t  j = ", j)
	fmt.Println("\t  k = ", k)
	fmt.Println("Output\t:", beautifulDays(int32(i), int32(j), int32(k)))
}
