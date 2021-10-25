package problem

import (
	"fmt"
	"strconv"
	"strings"
)

// link problem:
// https://www.hackerrank.com/challenges/time-conversion/problem?isFullScreen=true&h_r=next-challenge&h_v=zen

func timeConversion(s string) string {
	// Write your code here
	var output []string
	strtime, _ := strconv.Atoi(string(s[:2]))
	// fmt.Println(s[len(s)-2:] == "AM", s[:2] == "12", s[len(s)-2:], string(s[:2]))
	if s[len(s)-2:] == "AM" && strtime == 12 {
		output = append(output, "00")
	} else if s[len(s)-2:] == "PM" && strtime < 12 {
		output = append(output, strconv.Itoa(strtime+12))
	} else {
		output = append(output, s[:2])
	}
	output = append(output, string(s[2:len(s)-2]))
	soutput := strings.Join(output, "")
	return soutput
}

func Main_timeConversion() {
	a := "12:05:45PM"
	fmt.Println("==============timeConversion=============")
	fmt.Println("Input\t: time = ", a)
	fmt.Println("Output\t: ", timeConversion(a))

}
