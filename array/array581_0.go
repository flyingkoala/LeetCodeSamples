/*
@Time : 2020/6/29 13:17
@Author : wkang
@File : array581
@Description:
581. 最短无序连续子数组
给定一个整数数组，你需要寻找一个连续的子数组，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
你找到的子数组应是最短的，请输出它的长度。
示例 1:

输入: [2, 6, 4, 8, 10, 9, 15]
输出: 5
解释: 你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
说明 :
输入的数组长度范围在 [1, 10,000]。
输入的数组可能包含重复元素 ，所以升序的意思是<=。
*/
package main

import "fmt"

func findUnsortedSubarray(nums []int) int {
	if len(nums)<=1{
		return 0
	}
	if len(nums)==2 &&nums[0]>nums[1]{
		return 2
	}
	var flags []int
	var indexs []int
	for i:=0;i<len(nums)-1;i++{
		if nums[i]>nums[i+1] && len(flags)==0{
			flags=append(flags,nums[i])
			indexs=append(indexs,i)
		}
		if len(flags)>0{
			if nums[i+1]>=flags[len(flags)-1] {
				flags=append(flags,nums[i+1])
				indexs=append(indexs,i+1)
			}
		}
	}
	if len(flags)<=0{
		return 0
	}
	if nums[len(nums)-1]<flags[len(flags)-1]{
		flags=append(flags,nums[len(nums)-1])
		indexs=append(indexs,len(nums))
	}

	fmt.Println(flags)
	fmt.Println(indexs)
	if len(indexs)<=0{
		return 0
	}
	if len(indexs)==1 && indexs[0]==0{
		return len(nums)
	}
	for i:=len(indexs)-1;i>=1;i--{
		if indexs[i]-1!=indexs[i-1]{
			return indexs[i]-indexs[0]
		}

	}

	return 0

}

func main(){
	//fmt.Println(findUnsortedSubarray([]int{1,3,2,4,5,6}))
	//fmt.Println(findUnsortedSubarray([]int{2,6,4,8,10,9,15}))
	fmt.Println(findUnsortedSubarray([]int{1,3,2,2,2}))
}