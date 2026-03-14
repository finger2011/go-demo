package main

import "fmt"

func main() {
	// n, k := 3, 9 // cab
	// n, k := 3, 8 // bcb
	// n, k := 10, 100 // abacbabacb
	testArr := [][2]int{
		{3, 9},
		{3, 8},
		{4, 20}, // cacb
		{10, 100},
	}
	for _, arr := range testArr {
		fmt.Println("getHappyString (", arr[0], "; ", arr[1], ") ===> ", getHappyString(arr[0], arr[1]))
		fmt.Println("getHappyString2(", arr[0], "; ", arr[1], ") ===> ", getHappyString2(arr[0], arr[1]))
	}

}

// 总数量为3 * 2^(n - 1)
// 先判断第一位
// 随后的每一位都只有2种选法
func getHappyString(n int, k int) string {
	var ans string
	num := 1 << (n - 1)
	if k > 3*num {
		return ans
	}
	dicMap := map[byte][2]string{
		'a': {"b", "c"},
		'b': {"a", "c"},
		'c': {"a", "b"},
	}
	if k <= num {
		ans = "a"
	} else if k <= num*2 {
		ans = "b"
		k -= num
	} else {
		ans = "c"
		k -= num * 2
	}
	ops := 1
	for ops < n {
		num >>= 1
		if num == 1 {
			ans += dicMap[ans[ops-1]][k-1]
			break
		}
		if k <= num {
			ans += dicMap[ans[ops-1]][0]
		} else {
			ans += dicMap[ans[ops-1]][1]
			k -= num
		}
		ops++
	}
	return ans
}

// 先把k减一，变成从0开始(k0 = k - 1)，那么对应的就是k0的二进制上的每位代表需要填充的
func getHappyString2(n int, k int) string {
	if k > 3<<(n-1) {
		return ""
	}
	ans := make([]byte, n)
	k--
	ans[0] = 'a' + byte(k>>(n-1))
	for i := 1; i < n; i++ {
		ans[i] = 'a' + byte(k>>(n-1-i)&1)
		if ans[i] >= ans[i-1] {
			ans[i]++
		}
	}
	return string(ans)
}
