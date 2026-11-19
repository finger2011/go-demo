package main

func main() {

}

// 第i位和第i+1位必须相同，不同则修改册数加一
func minChanges(s string) int {
	var ant int
	for i := 0; i < len(s); i += 2 {
		if s[i] != s[i+1] {
			ant++
		}
	}
	return ant
}
