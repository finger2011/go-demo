package main

import (
	"fmt"
)

type Input struct {
	s1, s2 string
	except bool
}

func main() {
	tests := []*Input{
		{
			s1:     "abcd",
			s2:     "cbad",
			except: true,
		},
		{
			s1:     "abcd",
			s2:     "dacb",
			except: false,
		},
	}
	for _, test := range tests {
		fmt.Println("===========start")
		fmt.Println("s1:", test.s1)
		fmt.Println("s2:", test.s2)
		fmt.Println("result:", canBeEqual(test.s1, test.s2) == test.except)
		fmt.Println("===========end")
	}
}

func canBeEqual(s1 string, s2 string) bool {
	n := len(s1)
	b1, b2 := []byte(s1), []byte(s2)
	for i := range n {
		if b1[i] != b2[i] {
			if i+2 >= n {
				return false
			}
			if b1[i] != b2[i+2] {
				return false
			}
			b2[i], b2[i+2] = b2[i+2], b2[i]
		}

	}
	return true
}
