package problem

import "fmt"

//link problem:
//https://www.hackerrank.com/challenges/grading/problem?isFullScreen=true

func gradingStudents(grades []int32) []int32 {
	// Write your code here
	for i, v := range grades {
		if v >= 38 {
			if v%5 >= 3 {
				grades[i] += 5 - v%5
			}
		}
	}
	return grades
}

func Main_gradingStudents() {
	n := []int32{73, 67, 38, 33}
	fmt.Println("==============gradingStudents=============")
	fmt.Println("Input\t: candles = ", n)
	fmt.Println("Output\t: ", gradingStudents(n))

}
