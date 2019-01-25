package interfaces

import "tree"

type NewNode tree.Node

func (nn *NewNode) PostTraversal() {
	if nn == nil {
		return
	}
	if (*nn).Left != nil {
		var left = NewNode(*nn.Left)
		left.PostTraversal()
	}
	if (*nn).Right != nil {
		var right = NewNode(*(nn.Right))
		right.PostTraversal()
	}
	var cur = tree.Node(*nn)
	cur.PrintNode()

}
