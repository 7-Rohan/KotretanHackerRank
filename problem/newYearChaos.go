package problem

import (
	"fmt"
)

// link problem:
// https://www.hackerrank.com/challenges/one-week-preparation-kit-new-year-chaos/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=preparation-kits&playlist_slugs%5B%5D=one-week-preparation-kit&playlist_slugs%5B%5D=one-week-day-four

// var output int

// func minimumBribes(q []int32) {
// 	// Write your code here
// 	boole := true
// 	var count int
// 	// var count2 int
// 	for i := len(q) - 1; i >= 0; i-- {
// 		// fmt.Println(q[i]-int32(i+1) > int32(2), q[i]-int32(i+1))
// 		if q[i]-int32(i+1) == int32(2) {
// 			f := q[i]
// 			b1 := q[i+1]
// 			b2 := q[i+2]
// 			q[i] = b1
// 			q[i+1] = b2
// 			q[i+2] = f
// 			output += 2
// 			break
// 		} else if q[i]-int32(i+1) == int32(1) {
// 			f := q[i]
// 			b1 := q[i+1]
// 			q[i] = b1
// 			q[i+1] = f
// 			output++
// 			break
// 		} else if q[i]-int32(i+1) > int32(2) {
// 			boole = false
// 			break
// 		} else if q[i]-int32(i+1) == int32(0) {
// 			count++
// 		}
// 	}
// 	// for i := 0; i < len(q); i++ {
// 	// 	if q[i]-int32(i+1) == int32(2) {
// 	// 		f := q[i]
// 	// 		b1 := q[i+1]
// 	// 		b2 := q[i+2]
// 	// 		q[i] = b1
// 	// 		q[i+1] = b2
// 	// 		q[i+2] = f
// 	// 		output += 2
// 	// 		break
// 	// 	} else if q[i]-int32(i+1) == int32(1) {
// 	// 		f := q[i]
// 	// 		b1 := q[i+1]
// 	// 		q[i] = b1
// 	// 		q[i+1] = f
// 	// 		output++
// 	// 		break
// 	// 	} else if q[i]-int32(i+1) > int32(2) {
// 	// 		boole = false
// 	// 		break
// 	// 	} else if q[i]-int32(i+1) == int32(0) {
// 	// 		count++
// 	// 	}
// 	// 	count2++
// 	// }
// 	if !boole {
// 		fmt.Println("Too chaotic")
// 		return
// 	}
// 	if count != len(q) {
// 		minimumBribes(q)
// 	} else {
// 		fmt.Println(output)
// 	}
// }

func minimumBribes(q []int32) {
	var testarray []int32
	for i := 0; i < len(q); i++ {
		testarray = append(testarray, int32(i+1))
	}
	var bribecount int32
	for j := 0; j < len(q); j++ {
		if j+2 < len(q) && q[j] == testarray[j+2] {
			a := testarray[j]
			b := testarray[j+1]
			c := testarray[j+2]
			testarray[j] = c
			testarray[j+1] = a
			testarray[j+2] = b
			bribecount += 2
		} else if j+1 < len(q) && q[j] == testarray[j+1] {
			a := testarray[j]
			b := testarray[j+1]
			testarray[j] = b
			testarray[j+1] = a
			bribecount++
		} else if q[j] != testarray[j] && (j+1 >= len(q) || q[j] != testarray[j+1]) && (j+2 >= len(q) || q[j] != testarray[j+2]) {
			fmt.Println("Too chaotic")
			return
		}
	}
	fmt.Println(bribecount)
}

func Main_newYearChaos() {
	q := []int32{1, 2, 5, 3, 7, 8, 6, 4}
	fmt.Println("==============newYearChaos=============")
	fmt.Println("Input\t: q = ", q)
	fmt.Println("Output\t:")
	minimumBribes(q)
}
