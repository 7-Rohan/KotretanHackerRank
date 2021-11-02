package problem

import (
	"fmt"
	"sort"
	"strings"
)

// link problem:
// https://www.hackerrank.com/challenges/one-week-preparation-kit-grid-challenge/problem?isFullScreen=true&h_l=interview&playlist_slugs%5B%5D=preparation-kits&playlist_slugs%5B%5D=one-week-preparation-kit&playlist_slugs%5B%5D=one-week-day-four

func gridChallenge(grid []string) string {
	// Write your code here
	for i, v := range grid {
		s := strings.Split(v, "")
		sort.Strings(s)
		grid[i] = strings.Join(s, "")
	}
	fmt.Println(grid)
	mapStrInt := map[string]int32{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
	}
	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid)-1; j++ {
			if mapStrInt[string(grid[j][i])] > mapStrInt[string(grid[j+1][i])] {
				return "NO"
			}
		}
	}
	return "YES"
}

func Main_gridChallenge() {
	grid := []string{"mpxz", "abcd", "wlmf"}
	fmt.Println("==============billDivision=============")
	fmt.Println("Input\t: grid = ", grid)
	fmt.Println("Output\t:", gridChallenge(grid))

}
