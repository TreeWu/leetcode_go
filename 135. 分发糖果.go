package main

import "fmt"

/**
135. 分发糖果
老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。

你需要按照以下要求，帮助老师给这些孩子分发糖果：

每个孩子至少分配到 1 个糖果。
相邻的孩子中，评分高的孩子必须获得更多的糖果。
那么这样下来，老师至少需要准备多少颗糖果呢？

示例 1:

输入: [1,0,2]
dp[0] =1
dp[1] = 0
dp[2] = 1
 num =1
 1*3+ 1+0+1
5

输出: 5
解释: 你可以分别给这三个孩子分发 2、1、2 颗糖果。
示例 2:

输入: [1,4,3 ,2,1]
dp[0] =1     1
dp[1] = 2     5
dp[2] =1     4
dp[3] =0    3
dp[4]=-1   2
d[5] =-2  1


1+2+1 =4
输出: 4
解释: 你可以分别给这三个孩子分发 1、2、1 颗糖果。
     第三个孩子只得到 1 颗糖果，这已满足上述两个条件。
 */

func main() {
fmt.Println(candy([]int{1,3,2,2,1}))
}

func candy(ratings []int) int {
	dp:=make([]int,len(ratings))
	dp[0]=1
	min:=dp[0]
	for i:=1;i<len(ratings)-1;i++ {
		if ratings[i]> ratings[i-1] {
			dp[i] = dp[i-1]+1
		}else if ratings[i]== ratings[i-1]{
			dp[i] =1
		}else {
			dp[i] = dp[i-1] -1
			
		}
	}
	fmt.Println(dp,min)
	sum:=0
	for d :=range dp {
		sum+=dp[d]
	}


	if min<1 {
		sum+= (1-min)*len(dp)
	}



	return sum

}
