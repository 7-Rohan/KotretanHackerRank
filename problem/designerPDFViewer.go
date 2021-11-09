package problem

import (
	"fmt"
)

// link problem:
// https://www.hackerrank.com/challenges/designer-pdf-viewer/problem?isFullScreen=false&h_r=next-challenge&h_v=zen

func designerPdfViewer(h []int32, word string) int32 {
	// Write your code here
	var max int32
	for i := 0; i < len(word); i++ {
		if h[word[i]-97] > max {
			max = h[word[i]-97]
		}
	}
	return int32(len(word)) * max
}

func Main_designerPdfViewer() {
	h := []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7}
	word := "zaba"

	fmt.Println("==============designerPdfViewer=============")
	fmt.Println("Input\t: h = ", h)
	fmt.Println("\t  word = ", word)
	fmt.Println("Output\t:", designerPdfViewer(h, word))
}
