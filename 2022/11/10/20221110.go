// 20221110
// 字符串 翻转字符串

// 问题描述
// 请实现⼀个算法，在不使⽤【额外数据结构和储存空间】的情况下，翻转⼀个给定的字符串(可以使⽤单个过程变量)。
// 给定⼀个string，请返回⼀个string，为翻转后的字符串。
// 保证字符串的⻓度⼩于等于 5000。

// 反转字符串 -> 从末尾到首新建一个字符串

package day

import (
	"strings"
)

var maxLen int = 2500

func F20221110(str string) string {

	var b strings.Builder

	// 控制字符串长度
	n := 0
	if len(str) > maxLen {
		n = len(str) - maxLen
	}

	for i := len(str) - 1; i >= n; i-- {
		b.WriteByte(str[i])
	}
	return b.String()
}

func F20221110_2(str string) string {
	s := []rune(str)
	l := len(s)
	for i := 0; i < l/2; i++ {
		if i > maxLen {
			break
		}
		s[i], s[l-1-i] = s[l-1-i], s[i]
	}
	return string(s)
}
