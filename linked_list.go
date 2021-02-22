package dataStructure

import (
	"fmt"
)

type Node struct {
	Val int
	Next *Node
}


// 在链表的最后插入节点
func InsertNode(head * Node, newNode * Node){
	// 思路，先找当最后一个节点，将新的节点加入到链表最后
	curNode := head
	for {
		if curNode.Next == nil{
			break
		}
		curNode = curNode.Next
	}
	// 将新的节点加到链表最后
	curNode.Next = newNode

}

// 遍历链表
func ListNode(head *Node) {
	if head.Next != nil {
		curNode := head.Next
		for {
			fmt.Println(curNode.Val)
			if curNode.Next == nil {
				break
			}
			curNode = curNode.Next
		}
	}
}

// 统计链表长度
func LengthOfList(head *Node) int{
	if head.Next == nil{
		return 0
	} else {
		length := 0
		curNode := head.Next // 表示第一个节点
		for {
			if curNode == nil {
				break
			}
			curNode = curNode.Next
			length ++
		}
		return length
	}
}

// 删除节点
func DeleteNode(head *Node, value int) {
	if head.Next != nil{
		curNode := head.Next // 当前节点
		preCurNode := head.Next // 上一个节点
		for {
			if curNode.Val == value{ // 找到了
				if curNode == head.Next{ // 是否是头节点
					head.Next = curNode.Next
					break
				} else {
					preCurNode.Next = curNode.Next
					break
				}
			} else { // 未找到，继续往下个节点找
				preCurNode = curNode
				curNode = curNode.Next
				if curNode == nil{
					break
				}
			}
		}
	}
}