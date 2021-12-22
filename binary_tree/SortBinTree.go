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

//searchNode 节点查找
func (this *SortBinTree) searchNode(node *SortBinNode, key int) *SortBinNode {
	if node == nil {
		return nil
	}
	if node.data == key {
		return node
	}
	if key < node.data {
		return this.searchNode(node.lNode, key)
	}
	return this.searchNode(node.rNode, key)
}

//deleteBST 删除节点
func (this *SortBinTree) deleteBST(node *SortBinNode, key int) bool {
	//判断未找到
	//判断已找到
	//	删除节点
	//		叶子节点
	//		只有左子树或只有右子树
	//		既有左子树也有右子树

	if node == nil {
		return false
	}
	if node.data == key {
		node = this.deleteNode(node)
	} else if key < node.data {
		this.deleteBST(node.lNode, key)
	} else {
		this.deleteBST(node.rNode, key)
	}
	return true
}

//deleteNode 删除节点
func (this *SortBinTree) deleteNode(node *SortBinNode) *SortBinNode {
	if node.lNode == nil && node.rNode == nil { //叶子节点
		return nil
	}
	if node.rNode == nil {
		return node.lNode
	} else if node.lNode == nil {
		return node.rNode
	} else {
		parent := node
		pre := node.lNode
		for pre.rNode != nil {
			parent = pre
			pre = pre.rNode
		}
		node.data = pre.data
		if parent == node {
			return pre.lNode
		} else {
			return pre.rNode
		}
	}
}
