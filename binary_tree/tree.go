package main

import (
	"fmt"
)

func sortTree() {
	//NewSortBinTree([]int{63, 90, 70, 55, 58})
	tree := NewSortBinTree([]int{38, 12, 34, 56, 13, 6, 98, 3, 17, 40, 78})
	tree.deleteBST(tree.tree, 40)
}

func main() {
	sortTree()
	/*fmt.Println("please input tree:")
	str := readIn()
	i := 0
	root := createTree(str, &i)

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

	fmt.Println("-------------------------------------------")
	fmt.Printf("level for:")
	levelForeach(root)
	fmt.Printf("\n")

	fmt.Printf("copy tree LDR:")
	copyT := copyTree(root)
	midForeachWithStack(copyT)
	fmt.Printf("\n")

	fmt.Printf("node num:")
	nodenum := nodeNum(root)
	fmt.Printf("%d\n", nodenum)

	fmt.Printf("leaf node num:")
	leafNodenum := leafNodeNum(root)
	fmt.Printf("%d\n", leafNodenum)

	fmt.Printf("tree depth:")
	depth := treeDepth(root)
	fmt.Printf("%d\n", depth)*/
}

//treeDepth 获取数深度
func treeDepth(node *Node) int {
	if node == nil {
		return 0
	}
	m := treeDepth(node.lNode)
	n := treeDepth(node.rNode)
	if m > n {
		return m + 1
	}
	return n + 1
}

//leafNodeNum 获取叶子节点数量
func leafNodeNum(node *Node) int {
	if node == nil {
		return 0
	}
	if node.lNode == nil && node.rNode == nil {
		return 1
	}
	return leafNodeNum(node.lNode) + leafNodeNum(node.rNode)
}

//nodeNum 获取节点数量
func nodeNum(node *Node) int {
	if node == nil {
		return 0
	}
	return nodeNum(node.lNode) + nodeNum(node.rNode) + 1
}

//copyTree 复制二叉树
func copyTree(node *Node) *Node {
	newNode := new(Node)
	if node == nil {
		return nil
	} else {
		newNode.data = node.data
		newNode.lNode = copyTree(node.lNode)
		newNode.rNode = copyTree(node.rNode)
	}
	return newNode
}

//createTree 创建二叉树
func createTree(str string, i *int) *Node {
	newNode := new(Node)
	s := string(str[*i])
	*i++
	if s == "#" {
		return nil
	} else {
		newNode.data = s
		newNode.lNode = createTree(str, i)
		newNode.rNode = createTree(str, i)
	}
	return newNode
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

//levelForeach 层级遍历
func levelForeach(tree *Node) {
	queue := newQueue()
	pushQueue(queue, tree)
	for !emptyQueue(queue) {
		node := popQueue(queue)
		fmt.Printf("%s ", node.data)
		pushQueue(queue, node.lNode)
		pushQueue(queue, node.rNode)
	}
}
