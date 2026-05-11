package main

func main() {

}

func separateDigits(nums []int) []int {
	var ans []int
	for _, num := range nums {
		tmp := []int{}
		n := num
		for n >= 10 {
			tmp = append([]int{n % 10}, tmp...)
			n /= 10
		}
		tmp = append([]int{n}, tmp...)
		ans = append(ans, tmp...)
	}
	return ans
}
