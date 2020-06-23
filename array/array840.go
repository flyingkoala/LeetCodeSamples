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
			x1:=grid[i][j]+grid[i+1][j]+grid[i+2][j]
			x2:=grid[i][j+1]+grid[i+1][j+1]+grid[i+2][j+1]
			x3:=grid[i][j+2]+grid[i+1][j+2]+grid[i+2][j+2]
			y1:=grid[i][j]+grid[i][j+1]+grid[i][j+2]
			y2:=grid[i+1][j]+grid[i+1][j+1]+grid[i+1][j+2]
			y3:=grid[i+2][j]+grid[i+2][j+1]+grid[i+2][j+2]
			z1:=grid[i][j]+grid[i+1][j+1]+grid[i+2][j+2]
			z2:=grid[i+2][j]+grid[i+1][j+1]+grid[i][j+2]
			if x1==x2&&x1==x3&&x1==y1&&x1==y2&&x1==y3&&x1==z1&&x1==z2{
				count++
			}
		}
	}

	return count
}

func main(){
	fmt.Println(numMagicSquaresInside([][]int{ {5,5,5},{5,5,5},{5,5,5}}))
}