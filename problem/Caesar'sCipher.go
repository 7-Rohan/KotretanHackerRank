package problem

import (
	"fmt"
	"strings"
)

func caesarCipher(s string, k int32) string {
	// Write your code here
	st := "abcdefghijklmnopqrstuvwxyz"
	capSt := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	basestr := strings.Split(st, "")
	capBasestr := strings.Split(capSt, "")
	str := strings.Split(s, "")
	for i := 0; i < len(str); i++ {
	innerloop:
		for j := 0; j < len(basestr); j++ {
			if str[i] == basestr[j] {
				str[i] = basestr[(j+int(k))%len(basestr)]
				break innerloop
			}
			if str[i] == capBasestr[j] {
				str[i] = capBasestr[(j+int(k))%len(capBasestr)]
				break innerloop
			}
		}
	}
	return strings.Join(str, "")
}

func Main_caesarCipher() {
	s := "www.abc.xy"
	k := 87
	fmt.Println("==============caesarCipher=============")
	fmt.Println("Input\t: a = ", s, "k = ", k)
	fmt.Println("Output\t: ", caesarCipher(s, int32(k)))

}
