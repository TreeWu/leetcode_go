package main

import (
	"fmt"
	"math"
)

/**
1002. 查找常用字符
给定仅有小写字母组成的字符串数组 A，返回列表中的每个字符串中都显示的全部字符（包括重复字符）组成的列表。例如，如果一个字符在每个字符串中出现 3 次，但不是 4 次，则需要在最终答案中包含该字符 3 次。

你可以按任意顺序返回答案。



示例 1：

输入：["bella","label","roller"]
输出：["e","l","l"]
示例 2：

输入：["cool","lock","cook"]
输出：["c","o"]


提示：

1 <= A.length <= 100
1 <= A[i].length <= 100
A[i][j] 是小写字母
*/
func main() {
	fmt.Println(commonChars([]string{"bella", "label", "roller"}))
}

func commonChars(A []string) []string {

	var ans []string

	minSeq := [26]int{}
	for i := range minSeq {
		minSeq[i] = math.MaxInt64
	}
	for _, a := range A {
		curSeq := [26]int{}
		for i := range a {
			curSeq[a[i]-'a'] ++
		}
		for i := range curSeq {
			if curSeq[i] < minSeq[i] {
				minSeq[i] = curSeq[i]
			}
		}
	}
	for i := range minSeq {
		for j := 0; j < minSeq[i]; j++ {
			ans = append(ans, string('a'+i))
		}
	}
	return ans

}
