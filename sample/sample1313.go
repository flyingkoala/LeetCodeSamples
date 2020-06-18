/*
@Time : 2020/4/22 15:45
@Author : wkang
@File : sample1313
@Description:给你一个以行程长度编码压缩的整数列表 nums 。

考虑每对相邻的两个元素 freq, val] = [nums[2*i], nums[2*i+1]] （其中 i >= 0 ），每一对都表示解压后子列表中有 freq 个值为 val 的元素，你需要从左到右连接所有子列表以生成解压后的列表。

请你返回解压后的列表。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/decompress-run-length-encoded-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

func decompressRLElist(nums []int) []int {
	var res []int
	for k,v:=range nums{
		if k%2==0{
			for i:=0;i<v;i++{
				res=append(res, nums[k+1])
			}
		}
	}
	return res
}