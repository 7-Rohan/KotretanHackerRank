package problem

import (
	"fmt"
	"math"
)

// link:
// https://www.hackerrank.com/challenges/cats-and-a-mouse/problem?isFullScreen=true

// Complete the catAndMouse function below.
func catAndMouse(x int32, y int32, z int32) string {
	if math.Abs(float64(x-z)) < math.Abs(float64(y-z)) {
		return "Cat A"
	} else if math.Abs(float64(x-z)) == math.Abs(float64(y-z)) {
		return "Mouse C"
	} else {
		return "Cat B"
	}
}

func Main_catsAndAMouse() {
	x := 1
	y := 3
	z := 2
	fmt.Println("==============catsAndAMouse=============")
	fmt.Println("Input\t: x = ", x)
	fmt.Println("\t  y = ", y)
	fmt.Println("\t  z = ", z)
	fmt.Println("Output\t:", catAndMouse(int32(x), int32(y), int32(z)))
}
