/*
@Time : 2020/4/23 14:16
@Author : wkang
@File : sample9
@Description:判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
*/
package sample

import "strconv"

func isPalindrome(x int) bool {
	b:=true
	str:=strconv.Itoa(x)
	for i:=0;i<len(str);i++{
		if str[i]!=str[len(str)-1-i]{
			b=false
			break
		}
	}
	return b
}