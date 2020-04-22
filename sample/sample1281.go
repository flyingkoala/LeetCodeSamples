/*
@Time : 2020/4/22 11:07
@Author : wkang
@File : sample1281
@Description:给你一个整数 n，请你帮忙计算并返回该整数「各位数字之积」与「各位数字之和」的差。
*/
package sample

import (
	"fmt"
	"strconv"
)

func subtractProductAndSum(n int) int {
	ji:=1
	he:=0
	str:=strconv.Itoa(n)
	for _,v:=range str{
		c := string(v)
		num,_:=strconv.Atoi(c)
		ji=ji*num
		he = he+num
	}
	return ji-he
}

//执行
func Run1281(){
	res:=subtractProductAndSum(54)
	fmt.Println(res)

}