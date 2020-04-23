/*
@Time : 2020/4/23 13:24
@Author : wkang
@File : sample7
@Description:给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package sample

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var num []int
	i:=0
	for{
		temp:=x%10
		num=append(num,temp)
		x=x/10
		i++
		if x==0{
			break
		}
	}
	res:=0
	for _,v:=range num {
		res+=int(math.Pow10(i-1))*v
		i--
	}
	//if res<math.MinInt32||res>math.MaxInt32{
	//	return 0
	//}
	//判断int的范围 这样判断非常省内存
	if res > (1<<31-1) || res < (-1<<31) {
		return 0
	}


	return res
}

//执行
func Run7(){
	res:=reverse(-4537)
	fmt.Println(res)

}