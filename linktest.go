package main

import (
	"datastruct_go/link"
	"fmt"
)

func main() {
	head := link.Node{0, nil}
	for i := 1; i < 10; i++ {
		fmt.Println("i = ", i)
		link.Add(&head, i)
	}

	link.Travers(&head)

	cur := head.Next.Next.Next
	link.DeleteRandomNode(cur)
	link.Travers(&head)
}
