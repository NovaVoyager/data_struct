package main

//SortBinNode 节点
type SortBinNode struct {
	data         int
	lNode, rNode *SortBinNode
}

//SortBinTree 二叉排序树
type SortBinTree struct {
	tree *SortBinNode
}

func NewSortBinTree(nodeValues []int) *SortBinTree {
	tree := &SortBinTree{}
	tree.create(nodeValues)
	return tree
}

//create 创建二叉排序树
func (this *SortBinTree) create(nodeValues []int) {
	for _, v := range nodeValues {
		this.tree = this.insertData(this.tree, v)
	}
}

//insertData 插入节点
func (this *SortBinTree) insertData(node *SortBinNode, val int) *SortBinNode {
	if node == nil {
		node = &SortBinNode{
			data:  val,
			lNode: nil,
			rNode: nil,
		}
		return node
	} else {
		if val < node.data {
			node.lNode = this.insertData(node.lNode, val)
		} else {
			node.rNode = this.insertData(node.rNode, val)
		}
	}
	return node
}
