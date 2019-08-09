package main

import (
	"fmt"
	"strings"
)

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true
*/

func main() {
	fmt.Println(isValid2("([)]"))

}

func isValid(s string) bool {
	stack := &Stack{}

	split := strings.Split(s, "")
	for i := 0; i < len(split); i++ {
		if stack.Size == 0 {
			stack.push(split[i])
		} else {
			peek, _ := stack.peek()
			if isDoubleCheck(split[i], peek.(string)) {
				stack.pop()
			} else {
				stack.push(split[i])
			}
		}

	}
	if stack.Size != 0 {
		return false
	}

	return true
}

func isValid2(s string) bool {
	m := make(map[string]int)
	split := strings.Split(s, "")
	for i := 0; i < len(split); i++ {
		s1 := split[i]
		if s1 == ")" {
			if i2 := m["("]; i2 >= 1 {
				m["("] = i2 - 1
			} else {
				return false
			}
		} else if s1 == "]" {
			if i2 := m["["]; i2 >= 1 {
				m["["] = i2 - 1
			} else {
				return false
			}
		} else if s1 == "}" {
			if i2 := m["{"]; i2 >= 1 {
				m["{"] = i2 - 1
			} else {
				return false
			}
		}else {
			m[s1]=m[s1]+1
		}
	}
	for _,value := range m {
		if value !=0 {
			return false
		}
	}
	return true
}

func isDoubleCheck(s, t string) bool {
	if s == ")" && t == "(" {
		return true
	}
	if s == "}" && t == "{" {
		return true
	}
	if s == "]" && t == "[" {
		return true
	}
	return false
}
