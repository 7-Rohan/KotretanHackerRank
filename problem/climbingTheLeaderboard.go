package problem

import (
	"fmt"
)

// link problem:
// https://www.hackerrank.com/challenges/climbing-the-leaderboard/problem?isFullScreen=false&h_r=next-challenge&h_v=zen

// func climbingLeaderboard(ranked []int32, player []int32) []int32 {
// 	// Write your code here
// 	var output []int32
// 	var previous int32
// 	var rank foo
// 	for _, v := range ranked {
// 		if previous == v {
// 			continue
// 		} else {
// 			rank = append(rank, v)
// 		}
// 		previous = v
// 	}

// 	fmt.Println(rank)

// 	for _, v := range player {
// 		insertedRank := append(rank, v)
// 		fmt.Println(insertedRank)
// 		sort.Sort(insertedRank)
// 		fmt.Println(insertedRank)
// 		count := 0
// 		for i, w := range insertedRank {
// 			if w == v && count == 0 {
// 				output = append(output, int32(i+1))
// 				count++
// 			}
// 		}
// 	}
// 	return output
// }

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	// Write your code here
	var output []int32
	ranked = append(ranked, int32(0))
	for _, v := range player {
		var previous int32
		rank := 1
	innerLoop:
		for i := 0; i < len(ranked); i++ {
			if previous == ranked[i] {
				continue
			} else if ranked[i] > v {
				rank++
			} else if ranked[i] <= v {
				output = append(output, int32(rank))
				break innerLoop
			}
			previous = ranked[i]
		}
	}
	return output
}

func Main_climbingTheLeaderboard() {
	ranked := []int32{997, 981, 957, 933, 930, 927, 926, 920, 916, 896, 887, 874, 863, 863, 858, 847, 815, 809, 803, 794, 789, 785, 783, 778, 764, 755, 751, 740, 737, 730, 691, 677, 652, 650, 587, 585, 583, 568, 546, 541, 540, 538, 531, 527, 506, 493, 457, 435, 430, 427, 422, 422, 414, 404, 400, 394, 387, 384, 374, 371, 369, 369, 368, 365, 363, 337, 336, 328, 325, 316, 314, 306, 282, 277, 230, 227, 212, 199, 179, 173, 171, 168, 136, 125, 124, 95, 92, 88, 85, 70, 68, 61, 60, 59, 44, 43, 28, 23, 13, 12}
	player := []int32{12, 20, 30, 32, 35, 37, 63, 72, 83, 85, 96, 98, 98, 118, 122, 125, 129, 132, 140, 144, 150, 164, 184, 191, 194, 198, 200, 220, 228, 229, 229, 236, 238, 246, 259, 271, 276, 281, 283, 287, 300, 302, 306, 307, 312, 318, 321, 325, 341, 341, 341, 344, 349, 351, 354, 356, 366, 369, 370, 379, 380, 380, 396, 405, 408, 417, 423, 429,
		433, 435, 438, 441, 442, 444, 445, 445, 452, 453, 465, 466, 467, 468, 469, 471, 475, 482, 489, 491, 492, 493, 498, 500, 501, 504, 506, 508, 523, 529, 530, 539, 543, 551, 552, 556, 568, 569, 571, 587, 591, 601, 602, 606, 607, 612, 614, 619, 620, 623, 625, 625, 627, 638, 645, 653, 661, 662, 669, 670, 676, 684, 689, 690, 709, 709, 710, 716, 724, 726, 730, 731, 733, 737, 744, 744, 747, 757, 764, 765, 765, 772, 773, 774, 777, 787, 794, 796, 797, 802, 805, 811, 814, 819, 819, 829, 830, 841, 842, 847, 857, 857, 859, 860, 866, 872, 879, 882, 895,
		900, 900, 903, 905, 915, 918, 918, 922, 925, 927, 928, 929, 931, 934, 937, 955, 960, 966, 974, 982, 988, 996, 996}
	fmt.Println("==============climbingTheLeaderboard=============")
	fmt.Println("Input\t: ranked = ", ranked)
	fmt.Println("\t  player = ", player)
	fmt.Println("Output\t:", climbingLeaderboard(ranked, player))
}

// kodingan sort []int32 dll

// type foo []int32

// func (f foo) Len() int {
// 	return len(f)
// }

// func (f foo) Less(i, j int) bool {
// 	return f[i] > f[j]
// }

// func (f foo) Swap(i, j int) {
// 	f[i], f[j] = f[j], f[i]
// }
