package main

import "fmt"

type Input struct {
	moves  string
	except bool
}

func main() {
	tests := []*Input{
		{
			moves:  "",
			except: true,
		},
	}
	for _, test := range tests {
		fmt.Println("======================start")
		fmt.Println("moves:", test.moves)
		fmt.Println("except:", test.except)
		ans := judgeCircle(test.moves)
		fmt.Println("return:", ans)
		fmt.Println("result:", ans == test.except)
		fmt.Println("======================end")
	}
}

func judgeCircle(moves string) bool {
	var h, v int
	for i := range moves {
		switch moves[i] {
		case 'U':
			v++
		case 'D':
			v--
		case 'L':
			h--
		case 'R':
			h++
		}
	}
	return h == 0 && v == 0
}
