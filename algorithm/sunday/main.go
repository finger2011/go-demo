package main

import "fmt"

func main() {
	fmt.Println(strStrSunday("Here is a little Hao little", "little"))
}

func strStrSunday(haystack, needle string) int {
	//先判断两个字符串的合法性
	if len(haystack) < len(needle) {
		return -1
	}
	if haystack == needle {
		return 0
	}
	//定义最终位置的索引
	index := -1
	i := 0
	//定义目标匹配索引
	needleIndex := 0
	for i < len(haystack) {
		//逐字节判断是否相等
		if haystack[i] == needle[needleIndex] {
			//只有当index为-1时，说明是首次匹配到字符
			if index == -1 {
				index = i
			}
			//主串索引和模式串索引都自增
			i++
			needleIndex++
			//判断是否完成匹配
			if needleIndex >= len(needle) {
				break
			}
			continue
		}
		//走到这里说明没有匹配成功，将匹配目标索引置为默认
		index = -1
		//计算主串需要移动的位置
		i = i + len(needle) - needleIndex
		//如果主串索引大于了主串实际长度则返回
		if i >= len(haystack) {
			return index
		}
		//计算下一个字符在模式串最右的位置
		offset := 1
		for j := len(needle) - 1; j > 0; j-- {
			if haystack[i] == needle[j] {
				offset = j
				break
			}
		}
		//将主串的索引左移指定长度，使当前的字符和模式串中最右匹配到的字符串对齐
		i = i - offset
		//将模式串的索引重置
		needleIndex = 0
	}
	return index
}
