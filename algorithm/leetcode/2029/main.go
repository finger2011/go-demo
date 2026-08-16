package main

func main() {

}

func stoneGameIX(stones []int) bool {
	v0, v1, v2 := 0, 0, 0
	for _, val := range stones {
		switch val % 3 {
		case 0:
			v0++
		case 1:
			v1++
		case 2:
			v2++
		}
	}
	if v0%2 == 0 {
		return v1 >= 1 && v2 >= 1
	}
	return v1-v2 > 2 || v2-v1 > 2
}
