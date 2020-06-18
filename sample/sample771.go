/*
@Time : 2020/4/22 11:19
@Author : wkang
@File : sample771
@Description: 给定字符串J 代表石头中宝石的类型，和字符串 S代表你拥有的石头。 S 中每个字符代表了一种你拥有的石头的类型，你想知道你拥有的石头中有多少是宝石。

J 中的字母不重复，J 和 S中的所有字符都是字母。字母区分大小写，因此"a"和"A"是不同类型的石头。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/jewels-and-stones
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

func numJewelsInStones(J string, S string) int {
	count := 0
	mark := map[rune]bool{}
	for _, c := range J {
		mark[c] = true
	}
	for _, s := range S {
		if mark[s] {
			count++
		}
	}
	return count
}

//执行
func Run771(){

	res:=numJewelsInStones("abcE","QWSQCQdwdqdacfwe")
	fmt.Println(res)

}