// 20221112
// 字符串 字符串替换问题

// 问题描述
// 请编写⼀个⽅法，将字符串中的空格全部替换为“%20”。
// 假定该字符串有⾜够的空间存放新增的字符，
// 并且知道字符串的真实⻓度(⼩于等于1000)，
// 同时保证字符串由【⼤⼩写的英⽂字⺟组成】。
// 给定⼀个string为原始的串，返回替换后的string。

// 替换空格
// 剔除非大小写英文字母

package day

import (
	"strings"
	"unicode"
)

func F20221112(str string) string {

	var newStr strings.Builder

	for _, e := range str {
		if unicode.IsLower(e) || unicode.IsUpper(e) || e == rune(' ') {
			newStr.WriteRune(e)
		}
	}
	return strings.ReplaceAll(newStr.String(), " ", "%20")
}
