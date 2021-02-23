package main

import (
	"fmt"

	"dataStructure"
)

func main(){
	// 先创建头节点
	list := dataStructure.SingleLinkedList{}
	r1 := list.IsEmpty()
	fmt.Println("IsEmpty:", r1)
	list.AppendNode(1)
	list.AppendNode(2)
	list.AppendNode(3)
	list.AppendNode(4)
	list.AppendNode(5)
	list.AppendNode(6)

	r2 :=list.IsEmpty()
	fmt.Println("IsEmpty:", r2)
	list.ShowList()
	fmt.Println("=================")
	list.RemoveNode(1)
	list.ShowList()
	l := list.LengthOfList()
	fmt.Println("LengthOfList:", l)

}