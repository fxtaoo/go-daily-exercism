// 20221109
// 字符串 字符串中字符是否全都不同

// 问题描述
// 请实现⼀个算法，确定⼀个字符串的所有字符【是否全都不同】。
// 这⾥我们要求【不允许使⽤额外的存储结构】。
// 给定⼀个string，请返回⼀个bool值,true代表所有字符全都不同，false代表存在相同的字符。
// 保证字符串中的字符为【ASCII字符】。
// 字符串的⻓度⼩于等于【3000】。

// 思路
// 全都不同 -> 只出现一次
// strings.Count 一个字符串包含另一个字符串数量
// strings.Index strings.LastIndex 一个字符串在另一个字符串，第一次与最后一次索引

package day

import "strings"

func F20221109(str string) bool {
	if len(str) > 3000 {
		return false
	}

	for _, c := range str {
		if c > 127 {
			return false
		}

		if strings.Count(str, string(c)) > 1 {
			return false
		}
	}
	return true
}

func F20221109_2(str string) bool {
	if len(str) > 3000 {
		return false
	}

	for k, c := range str {
		if c > 127 {
			return false
		}

		if strings.Index(str, string(c)) != k {
			return false
		}
	}
	return true
}

func F20221109_3(str string) bool {
	if len(str) > 3000 {
		return false
	}

	for k, c := range str {
		if c > 127 {
			return false
		}

		if strings.LastIndex(str, string(c)) != k {
			return false
		}
	}
	return true
}
