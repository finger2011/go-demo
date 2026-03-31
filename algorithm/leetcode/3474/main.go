package main

import (
	"bytes"
	"fmt"
)

type Input struct {
	s1, s2, except string
}

func main() {
	tests := []*Input{
		{
			s1:     "TFTF",
			s2:     "ab",
			except: "ababa",
		},
		{
			s1:     "TFTF",
			s2:     "abc",
			except: "",
		},
		{
			s1:     "FFTFFF",
			s2:     "a",
			except: "bbabbb",
		},
		{
			s1:     "FFFTFTTFTTFTTTTFFTTFFFT",
			s2:     "XX",
			except: "",
		},
	}
	for _, test := range tests {
		fmt.Println("start:")
		fmt.Println("s1:", test.s1)
		fmt.Println("s2:", test.s2)
		fmt.Println("result ==========> ", generateString(test.s1, test.s2) == test.except)
		fmt.Println("end:")
	}
}

func generateString(str1 string, str2 string) string {
	n, m := len(str1), len(str2)
	t := []byte(str2)
	ans := make([]byte, n+m-1)
	confirm := make([]byte, n+m-1)
	for i := range n {
		if str1[i] == 'T' {
			for j := range m {
				if confirm[i+j] != 0 && ans[i+j] != str2[j] {
					return ""
				}
				ans[i+j] = str2[j]
				confirm[i+j] = '1'
			}
		}
	}
	for i := range ans {
		if ans[i] == 0 {
			ans[i] = 'a'
		}
	}
	for i := range n {
		if str1[i] == 'F' && bytes.Equal(ans[i:i+m], t) {
			idx := -1
			for j := m - 1; j >= 0; j-- {
				if confirm[i+j] == 0 {
					idx = j
					break
				}
			}
			if idx == -1 {
				return ""
			}
			ans[i+idx] = 'b'
		}
	}

	return string(ans)
}
