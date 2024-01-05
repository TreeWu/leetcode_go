package main

import (
	"container/list"
	"fmt"
)

/*
*
20. 有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true
*/
func main() {

	fmt.Println(isValidBrackets("()[]{}"))
	fmt.Println(有效的括号("()[]{}"))

	fmt.Println(isValidBrackets("([)]"))
	fmt.Println(有效的括号("([)]"))

}

func isValidBrackets(s string) bool {

	stock := list.New()

	seq := make(map[rune]rune)
	seq[')'] = '('
	seq['}'] = '{'
	seq[']'] = '['

	rs := []rune(s)
	for i := range rs {
		front := stock.Front()
		if front != nil {
			if seq[rs[i]] == front.Value {
				stock.Remove(front)
				continue
			}
		}
		stock.PushFront(rs[i])
	}
	return stock.Len() == 0
}

func 有效的括号(s string) bool {
	bytes := []rune(s)
	var stock []rune
	for i := range bytes {
		b := bytes[i]
		if b == '(' || b == '{' || b == '[' {
			stock = append(stock, b)
			continue
		}
		if len(stock) == 0 {
			return false
		}
		top := stock[len(stock)-1]

		if top == '(' && b == ')' {
			stock = stock[:len(stock)-1]
		}

		if top == '[' && b == ']' {
			stock = stock[:len(stock)-1]
		}
		if top == '{' && b == '}' {
			stock = stock[:len(stock)-1]
		}
	}
	return len(stock) == 0
}
