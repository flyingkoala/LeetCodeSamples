/*
@Time : 2020/6/18 13:13
@Author : wkang
@File : sample20
@Description:
20. 有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true
*/
package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	if len(s) ==0{
		return true
	}
	if len(s)%2!=0{
		return false
	}
	if string(s[0])=="}"||string(s[0])==")"||string(s[0])=="]"{
		return false
	}

	hash := map[string]string{")":"(", "]":"[", "}":"{"}
	stack := make([]string, 0)

	for i:=0;i<len(s);i++ {

		hashnow:= hash[string(s[i])] //当前元素在hash表中的对应元素
		if hashnow == ""{			//当前元素是正括号 添加到栈中
			stack=append(stack,string(s[i]))
		}else {
			//当前元素是反括号
			pop:=stack[len(stack)-1]//栈中的最后一个元素
			if hash[pop] =="" && hashnow!=pop { //如果栈中的最后一个元素是正括号 并且和当前元素不成对，直接返回失败
				return false
			}
			if hashnow == pop{  //如果栈中的最后一个元素和当前元素成对
				stack=stack[:len(stack)-1]//修改栈 最后一个元素出栈
			} else if hashnow== stack[0]{ //如果栈中的第一个元素和当前元素成对
				stack=stack[1:]				//修改栈 第一个元素出栈
			}else {
				stack=append(stack,string(s[i])) //入栈
			}
		}
		//fmt.Println(stack)
	}
	if len(stack)==0{
		return true
	}
	return false

}

//方法2：依次去掉符合条件的括号 最终判断结果
func isValid1(s string) bool {
	for {
		old := s
		s = strings.Replace(s,"()","",-1)
		s = strings.Replace(s,"[]","",-1)
		s = strings.Replace(s,"{}","",-1)
		if s == "" {
			return true
		}
		if s == old {
			return false
		}
	}
	return false
}


func main()  {
	fmt.Println(isValid("{}[]([()])"))
	fmt.Println(isValid("([()])"))
	fmt.Println(isValid("([)]"))
}