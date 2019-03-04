package link

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

func Add(head *Node, data int) {
	point := head
	for point.Next != nil {
		point = point.Next
	}
	var node Node
	point.Next = &node
	node.Data = data
}

//删除index节点，并返回节点数据，c语言删除节点需要手动释放内存
func Delete(head *Node, index int) int {
	if index < 0 || index > GetLength(head) {
		fmt.Println("please check index")
		return 0
	} else {
		point := head
		for i:= 0; i < index-1; i++ {
			point = point.Next
		}
		point.Next = point.Next.Next
		data := point.Next.Data
		return data
	}
}

//给index位置插入data
func Insert(head *Node, index int, data int) {
	if index < 0 || index > GetLength(head) {
		fmt.Println("please check index")
	} else {
		point := head
		for i := 0; i < index - 1; i++ {
			point = point.Next
		}
		var node Node
		node.Data = data
		node.Next = point.Next
		point.Next = &node
	}
}

//获取长度
func GetLength(head *Node) int {
	point := head
	length := 0
	for point != nil {
		length++
		point = point.Next
	}
	return length
}

//获取元素位置
func Search(head *Node, data int) {
	point := head
	index := 0
	for point.Next != nil {
		if point.Data == data {
			fmt.Println(data, "exit at", index)
			break
		} else {
			index++
			point = point.Next
			if index > GetLength(head) - 1 {
				fmt.Println(data, "not exit")
				break
			}
			continue
		}
	}
}

//遍历头节点
func Travers(head *Node) {
	point := head
	for point != nil {
		fmt.Println(point.Data)
		point = point.Next
	}
	fmt.Println("traver finished")
}
