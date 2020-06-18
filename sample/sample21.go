/*
@Time : 2020/6/18 14:54
@Author : wkang
@File : sample21
@Description:21. 合并两个有序链表
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/
package main

type ListNode struct {
     Val int
     Next *ListNode
 }


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	pre:= &ListNode{}
	result:=pre
	for l1!=nil&&l2!=nil{
		if l1.Val<=l2.Val{
			pre.Next = l1
			l1=l1.Next
		}else {
			pre.Next = l2
			l2=l2.Next
		}
		pre=pre.Next
	}
	if l1!=nil{
		pre.Next=l1
	}
	if l2!=nil{
		pre.Next=l2
	}
	return result.Next

}