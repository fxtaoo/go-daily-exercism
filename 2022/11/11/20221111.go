// 20221111
// 字符串 判断两个给定的字符串排序后是否⼀致

// 问题描述
// 给定两个字符串，请编写程序，确定其中⼀个字符串的字符重新排列后，能否变成另⼀个字符串。
// 这⾥规定【⼤⼩写为不同字符】，且考虑字符串重点空格。
// 给定⼀个string s1和⼀个string s2，请返回⼀个bool，代表两串是否重新排列后可相同。
// 保证两串的⻓度都⼩于等于5000。

package day

import (
	"sort"
	"strings"
)

var lenMax int = 5000

func F20221111(s1, s2 string) bool {
	l1 := len(s1)
	l2 := len(s2)

	if l1 > lenMax || l2 > lenMax || l1 != l2 {
		return false
	}

	s1Slice := strings.Split(s1, "")
	sort.Strings(s1Slice)
	s1Sort := strings.Join(s1Slice, "")

	return s1Sort == s2

}

func F20221111_2(s1, s2 string) bool {
	l1 := len(s1)
	l2 := len(s2)

	if l1 > lenMax || l2 > lenMax || l1 != l2 {
		return false
	}

	for _, v := range s1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}

	return true
}
