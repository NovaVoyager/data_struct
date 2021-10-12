package main

import "fmt"

func main() {
	root := generateTreeSimple()

	fmt.Printf("DLR: ")
	preForeach(root)
	fmt.Printf("\n")

	fmt.Printf("LDR: ")
	midForeach(root)
	fmt.Printf("\n")

	fmt.Printf("LRD: ")
	backForeach(root)
	fmt.Printf("\n")

	fmt.Println("-------------------------------------------")

	fmt.Printf("LDR: ")
	midForeachWithStack(root)
	fmt.Printf("\n")
}

//generateTreeSimple 生成树，简单版
func generateTreeSimple() *Node {
	var a, b, c, d, e, f, g, h, i, j, head *Node

	a = &Node{data: "A"}
	b = &Node{data: "B"}
	c = &Node{data: "C"}
	d = &Node{data: "D"}
	e = &Node{data: "E"}
	f = &Node{data: "F"}
	g = &Node{data: "G"}
	h = &Node{data: "H"}
	i = &Node{data: "I"}
	j = &Node{data: "J"}

	head = a

	a.lNode = b
	a.rNode = g
	b.lNode = c
	b.rNode = e
	c.rNode = d
	e.lNode = f
	g.lNode = h
	g.rNode = j
	h.rNode = i

	return head
}

//preForeach 先序遍历
func preForeach(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("%s ", node.data)
	preForeach(node.lNode)
	preForeach(node.rNode)
}

//midForeach 中序遍历
func midForeach(node *Node) {
	if node == nil {
		return
	}
	midForeach(node.lNode)
	fmt.Printf("%s ", node.data)
	midForeach(node.rNode)
}

//backForeach 后序遍历
func backForeach(node *Node) {
	if node == nil {
		return
	}
	backForeach(node.lNode)
	backForeach(node.rNode)
	fmt.Printf("%s ", node.data)
}

//midForeachWithStack
func midForeachWithStack(tree *Node) {
	p := tree
	stack := generateStack()
	for p != nil || !emptyStack(stack) {
		if p != nil {
			pushStack(stack, p)
			p = p.lNode
		} else {
			q := popStack(stack)
			fmt.Printf("%s ", q.data)
			p = q.rNode
		}
	}
}
