package main

import "fmt"

type Input struct {
	s1, s2 string
	except bool
}

func main() {
	tests := []*Input{
		{
			s1:     "abcdba",
			s2:     "cabdab",
			except: true,
		},
		{
			s1:     "abe",
			s2:     "bea",
			except: false,
		},
	}
	for _, test := range tests {
		fmt.Println("===========start")
		fmt.Println("s1:", test.s1)
		fmt.Println("s2:", test.s2)
		fmt.Println("result:", checkStrings(test.s1, test.s2) == test.except)
		fmt.Println("===========end")
	}
}

func checkStrings(s1 string, s2 string) bool {
	n := len(s1)
	var cnt1, cnt2 [2][26]int
	for i := range n {
		cnt1[i%2][int(s1[i]-'a')]++
		cnt2[i%2][int(s2[i]-'a')]++
	}
	return cnt1 == cnt2
}
