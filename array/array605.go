/*
@Time : 2020/6/19 11:34
@Author : wkang
@File : array605
@Description:
605. 种花问题
假设你有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花卉不能种植在相邻的地块上，它们会争夺水源，两者都会死去。
给定一个花坛（表示为一个数组包含0和1，其中0表示没种植花，1表示种植了花），和一个数 n 。能否在不打破种植规则的情况下种入 n 朵花？能则返回True，不能则返回False。
示例 1:
输入: flowerbed = [1,0,0,0,1], n = 1
输出: True
示例 2:
输入: flowerbed = [1,0,0,0,1], n = 2
输出: False
注意:
数组内已种好的花不会违反种植规则。
输入的数组长度范围为 [1, 20000]。
n 是非负整数，且不会超过输入数组的大小。
*/
package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	//可以插入的1的数量
	num:=0
	//只有一个空位而且为0 特殊处理
	if len(flowerbed)==1 &&flowerbed[0]==0{
		return true
	}
	//遍历数组 将值为1的数组下标记录
	indexs:=make([]int,0,10)
	for k,v:=range flowerbed{
		if v==1{
			indexs=append(indexs,k)
		}
	}
	//如果数组中没有1
	if len(indexs)==0{
		num=(len(flowerbed)+1)/2
	}else {
		//计算值为1的中间部分可以插入多少
		for i:=0;i<len(indexs)-1;i++{
			if len(indexs)<=1{
				//只有一个1或者没有1的话 直接退出循环
				break
			}
			diff:=indexs[i+1]-indexs[i]//求的1之间的坐标差
			f:=diff/2-1
			if f<0{
				f=0
			}
			num+=f
		}
		//分别单独处理两端
		if indexs[0]!=0{
			//原数组不是1开头
			num+=indexs[0]/2
		}
		if len(flowerbed)-1!=indexs[len(indexs)-1] {
			//原数组不是1结尾的
			num += (len(flowerbed) - 1 - indexs[len(indexs)-1]) / 2
		}
	}
	//fmt.Println(indexs)
	//fmt.Println(num)
	if num>=n{
		return true
	}
	return false

}

func main()  {
	canPlaceFlowers([]int{0,0,1,0,0,0,0,0,0,0,1,0,0,0,1,0,0,1},2)
}