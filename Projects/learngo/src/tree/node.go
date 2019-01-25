package tree

import "fmt"

type Node struct {
	Value int
	Left, Right *Node
}

func CreateNode(val int)(*Node){
	return &Node{Value:val}
}

func (node *Node)PrintNode() {
	fmt.Printf("node value: %d\n", node.Value)
}


func (node *Node) Traversal() {
	if node == nil{
		return
	}
	node.Left.Traversal()
	node.PrintNode()
	node.Right.Traversal()
}