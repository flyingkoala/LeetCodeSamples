/*
@Time : 2020/4/22 16:24
@Author : wkang
@File : sample1365
@Description:给你一个数组 nums，对于其中每个元素 nums[i]，请你统计数组中比它小的所有数字的数目。

换而言之，对于每个 nums[i] 你必须计算出有效的 j 的数量，其中 j 满足 j != i 且 nums[j] < nums[i] 。

以数组形式返回答案

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/how-many-numbers-are-smaller-than-the-current-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package sample

import (
	"fmt"
	"sort"
)

func smallerNumbersThanCurrent(nums []int) []int {
	length:=len(nums)
	sortnums:=make([]int,length)
	copy(sortnums,nums)
	sort.Ints(sortnums)
	m := make(map[int]int)
	for k,v:= range sortnums{
		if _, ok := m[v]; !ok {
			m[v]=k
		}
	}
	var res []int
	for _,v:=range nums{
		res=append(res,m[v])
	}
	return res
}

//执行
func Run1365(){
	res:=smallerNumbersThanCurrent([]int{7,7,7,7})
	fmt.Println(res)

}