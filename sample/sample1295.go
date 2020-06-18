/*
@Time : 2020/4/22 13:44
@Author : wkang
@File : sample1295
@Description: 给你一个整数数组 nums，请你返回其中位数为 偶数 的数字的个数。
*/
package main

import (
	"strconv"
)

func findNumbers(nums []int) int {
	num:=0
	for _,v:=range nums{
		str:=strconv.Itoa(v)
		if len(str)%2==0{
			num++
		}
	}
	return  num
}