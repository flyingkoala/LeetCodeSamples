/*
@Time : 2020/4/22 10:28
@Author : wkang
@File : sampleLcp01
@Description:小A 和 小B 在玩猜数字。小B 每次从 1, 2, 3 中随机选择一个，小A 每次也从 1, 2, 3 中选择一个猜。他们一共进行三次这个游戏，请返回 小A 猜对了几次？
输入的guess数组为 小A 每次的猜测，answer数组为 小B 每次的选择。guess和answer的长度都等于3。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/guess-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package sample

import "fmt"

func game(guess []int, answer []int) int {
	num:=0
	for k,v:=range guess{
		if v== answer[k]{
			num++
		}
	}
	return num
}
func RunLcp01(){
	res:=game([]int{1,3,2},[]int{1,2,2})
	fmt.Println(res)

}