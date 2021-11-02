package problem

import (
	"fmt"
	"strconv"
	"strings"
)

// link problem:
// https://www.hackerrank.com/challenges/one-week-preparation-kit-recursive-digit-sum/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=preparation-kits&playlist_slugs%5B%5D=one-week-preparation-kit&playlist_slugs%5B%5D=one-week-day-four

func superDigit(n string, k int32) int32 {
	// Write your code here
	var wadahP []string
	for k%10 == 0 {
		k = k / 10
	}
	for i := 0; i < int(k); i++ {
		wadahP = append(wadahP, n)
	}
	p := strings.Join(wadahP, "")
	arrStr := strings.Split(p, "")
	var sum int32
	for _, v := range arrStr {
		intArr, _ := strconv.Atoi(v)
		sum += int32(intArr)
	}
	sumstr := strconv.Itoa(int(sum))
	if len(sumstr) > 1 {
		return superDigit(sumstr, 1)
	}
	return sum
}

func Main_recursiveDigitSum() {
	n := "9875"
	k := 50
	fmt.Println("==============superDigit=============")
	fmt.Println("Input\t: n = ", n)
	fmt.Println("\t  k = ", k)
	fmt.Println("Output\t:", superDigit(n, int32(k)))
}
