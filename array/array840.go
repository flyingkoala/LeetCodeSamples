/*
@Time : 2020/6/23 16:09
@Author : wkang
@File : array840
@Description:
840. 矩阵中的幻方
3 x 3 的幻方是一个填充有从 1 到 9 的不同数字的 3 x 3 矩阵，其中每行，每列以及两条对角线上的各数之和都相等。
给定一个由整数组成的 grid，其中有多少个 3 × 3 的 “幻方” 子矩阵？（每个子矩阵都是连续的）。

示例：

输入: [[4,3,8,4],
      [9,5,1,9],
      [2,7,6,2]]
输出: 1
解释:
下面的子矩阵是一个 3 x 3 的幻方：
438
951
276

而这一个不是：
384
519
762
总的来说，在本示例所给定的矩阵中只有一个 3 x 3 的幻方子矩阵。
提示:

1 <= grid.length <= 10
1 <= grid[0].length <= 10
0 <= grid[i][j] <= 15
*/
package main

import "fmt"

func numMagicSquaresInside(grid [][]int) int {
	zong:=len(grid[0])
	heng:=len(grid)
	if zong<3||heng<3{
		return 0
	}
	count:=0
	for i:=0;i<=heng-3;i++{
		for j:=0;j<=zong-3;j++{
			a:=grid[i][j]
			b:=grid[i+1][j]
			c:=grid[i+2][j]
			d:=grid[i][j+1]
			e:=grid[i+1][j+1]
			f:=grid[i+2][j+1]
			g:=grid[i][j+2]
			h:=grid[i+1][j+2]
			i:=grid[i+2][j+2]

			if e!=5{
				continue
			}
			if e==5&&a==5&&i==5{
				continue
			}
			if a>9||b>9||c>9||d>9||e>9||f>9||g>9||h>9||i>9{
				continue
			}
			sum:=a+b+c
			if sum==d+e+f&&sum==g+h+i&&sum==a+d+g&&sum==b+e+h&&sum==e+f+i&&sum==a+e+i&&sum==c+e+g{
				count++
			}
		}
	}

	return count
}

func main(){
	fmt.Println(numMagicSquaresInside([][]int{ {5,5,5},{5,5,5},{5,5,5}}))
}