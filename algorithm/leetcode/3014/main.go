package main

func main() {

}

func minimumPushes(word string) int {
	var ans int
	for i := range word {
		ans += (i + 8) / 8
	}
	return ans

}
