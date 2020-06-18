/*
@Time : 2020/4/22 10:36
@Author : wkang
@File : sample1342
@Description:给你一个非负整数 num ，请你返回将它变成 0 所需要的步数。 如果当前数字是偶数，你需要把它除以 2 ；否则，减去 1 。
*/
package main

import (
	"fmt"

)

func numberOfSteps (num int) int {
	if num == 0 {
		return 0
	}
	res := 0
	for num > 0 {
		if num % 2 == 0{
			num = num / 2
		} else {
			num = num - 1
		}
		res++
	}
	return res
}

func Run1342(){
	res:=numberOfSteps(10)
	fmt.Println(res)

}