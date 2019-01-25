package main

import (
	"fmt"
	"interfaces"
	"tree"
)

func main() {
	var root tree.Node
	root.Left = tree.CreateNode(1)
	root.Right = tree.CreateNode(2)
	root.Traversal()
	fmt.Println(1)

	var root1 = interfaces.NewNode{3, nil, nil}
	root1.Left = tree.CreateNode(1)
	root1.Right = tree.CreateNode(2)

	root1.PostTraversal()
}
