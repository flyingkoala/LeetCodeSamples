/*
@Time : 2020/6/18 15:14
@Author : wkang
@File : array665
@Description:
给你一个长度为 n 的整数数组，请你判断在 最多 改变 1 个元素的情况下，该数组能否变成一个非递减数列。
我们是这样定义一个非递减数列的： 对于数组中所有的 i (0 <= i <= n-2)，总满足 nums[i] <= nums[i + 1]。

示例 1:
输入: nums = [4,2,3]
输出: true
解释: 你可以通过把第一个4变成1来使得它成为一个非递减数列。
示例 2:
输入: nums = [4,2,1]
输出: false
解释: 你不能在只改变一个元素的情况下将其变为非递减数列。
说明：
1 <= n <= 10 ^ 4
- 10 ^ 5 <= nums[i] <= 10 ^ 5
*/
package main

import (
	"fmt"
)

func checkPossibility(nums []int) bool {
	if len(nums)<=2{
		return true
	}
	if len(nums)==3{
		if nums[0]>nums[1]&&nums[1]>nums[2]{
			return false
		}else{
			return true
		}
	}
	flag:=0
	for i:=0;i<len(nums)-1;i++{
		tmp:=nums[i]
		//如果当前元素比下一个元素大
		if nums[i]>nums[i+1]{
			//如果当前元素是第一个元素
			if i==0 {
				//当前元素值=下一个元素
				nums[i]=nums[i+1]
			}else{
				//当前元素值=上一个元素
				nums[i]=nums[i-1]
			}
			//修改了之后的当前元素继续比较下一个元素
			if nums[i]>nums[i+1]{
				//如果还比下一个元素大 把当前元素改回来 就修改下一个元素吧
				nums[i] = tmp
				nums[i+1]=tmp

			}
			flag++
			if flag>=2{
				return false
			}
		}

	}
	return true
}

func main(){
	fmt.Println(checkPossibility([]int{3,4,2,3}))
}