package dataStructure

import (
	"fmt"
)

type Object interface {} // 空接口类型


type Node struct {
	Data Object
	Next *Node
}

type SingleLinkedList struct {
	headNode * Node // 头节点
}

// 判断链表是否为空, 结构体方法
func (linkedList * SingleLinkedList) IsEmpty() bool{
	if linkedList.headNode == nil{
		return true
	} else {
		return false
	}
}

// 获取链表长度
func (linkedList * SingleLinkedList) LengthOfList() int {
	if linkedList.IsEmpty() {
		return 0
	} else {
		count := 0
		cur := linkedList.headNode
		for cur != nil{
			count ++
			cur = cur.Next
		}
		return count
	}
}

// 从链表尾部添加元素
func (linkedList * SingleLinkedList) AppendNode(data Object) {
	node := &Node{Data: data} //
	if linkedList.IsEmpty() {
		// 将headNode指向第一个节点node
		linkedList.headNode = node
	} else {
		cur := linkedList.headNode
		for cur.Next != nil{
			cur = cur.Next
		}
		cur.Next = node
	}
}

// 遍历链表
func (linkedList * SingleLinkedList) ShowList(){
	if ! linkedList.IsEmpty() {
		cur := linkedList.headNode
		for cur != nil {
			fmt.Println(cur.Data)
			cur = cur.Next
		}
	}
}

// 删除链表指定值的元素
func (linkedList * SingleLinkedList) RemoveNode(data Object){
	if ! linkedList.IsEmpty() {
		cur := linkedList.headNode
		pre := linkedList.headNode
		for cur != nil {
			if data == cur.Data{ // 找到了
				if cur == linkedList.headNode{ // 是头节点
					linkedList.headNode = cur.Next
				} else { // 不是头节点
					pre.Next = cur.Next
				}
				break

			} else { // 未找到 继续往后找
				pre = cur
				cur = cur.Next
			}
		}
	}
}


// 反转链表
// 1->2->3->4->5
// 5->4->3->2->1
// 第一次  nil<-1
// 第二次  nil<-1<-2
// 第三次  nil<-1<-2<-3
// 第四次  nil<-1<-2<-3<-4
// 第五次  nil<-1<-2<-3<-4<-5
func (linkedList * SingleLinkedList) ReversedLinkedList() *SingleLinkedList{

	reverseHead := &SingleLinkedList{}
	curNode := linkedList.headNode
	var preNode *Node = nil
	for curNode != nil {

		nextNode := curNode.Next
		if nextNode == nil{
			reverseHead.headNode = curNode
		}

		curNode.Next = preNode
		preNode = curNode
		curNode = nextNode

	}

	return reverseHead
}