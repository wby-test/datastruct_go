package link

//在O(1)时间删除链表节点，给定链表头指针和节点指针（用后一个节点替代前面的节点）
func DeleteRandomNode(cur *Node) {
	if cur == nil {
		panic("node is nil")
	}
	if cur.Next == nil {
		panic("node is tail")
	}
	point := cur.Next
	cur.Data = point.Data
	cur.Next = point.Next
}


//单链表转置
////1非递归算法
func ReverseByLoop(head *Node) *Node{
	if head == nil || head.Next == nil {
		return head
	}
	var pre *Node = nil
	var next *Node = nil
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

////2递归算法
func ReverseByRecusion(head *Node) *Node{
	if head == nil || head.Next == nil {
		return head
	}

	newHead  := ReverseByRecusion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}