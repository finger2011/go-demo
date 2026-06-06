package main

import "fmt"

func main() {
	num1, num2 := 120, 130
	fmt.Println("num:", totalWaviness(num1, num2))
}

func totalWaviness(num1 int, num2 int) int {
	var ans int
	for i := num1; i <= num2; i++ {
		if i < 100 {
			continue
		}
		num := i
		f1 := num % 10
		num /= 10
		f2 := num % 10
		num /= 10
		for num > 0 {
			f3 := num % 10
			num /= 10
			if (f2 > f1 && f2 > f3) || (f2 < f1 && f2 < f3) {
				ans++
			}
			f1, f2 = f2, f3
		}
	}
	return ans
}

// num1, num2 : 1~10^15
func totalWaviness2(num1 int64, num2 int64) int64 {
	var ans int64
	for i := num1; i <= num2; i++ {
		if i < 100 {
			continue
		}
		num := i
		f1 := num % 10
		num /= 10
		f2 := num % 10
		num /= 10
		for num > 0 {
			f3 := num % 10
			num /= 10
			if (f2 > f1 && f2 > f3) || (f2 < f1 && f2 < f3) {
				ans++
			}
			f1, f2 = f2, f3
		}
	}
	return ans
}
