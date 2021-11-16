package main

import "fmt"

//ArcNode 邻接表，边的存储结构
type ArcNode struct {
	adjvex int
	next   *ArcNode
	info   ArcType
}

//VexNode 顶点表
type VexNode struct {
	data    VerTextType
	arcNode *ArcNode
}

//ALGraph 邻接表
type ALGraph struct {
	vertices       []VexNode
	vexnum, arcnum int
}

//ALElem 构建图的边元素
type ALElem struct {
	//V1 顶点1
	V1 VerTextType

	//Weight 权值
	Weight ArcType
}

//ALGraphElem 邻接表创建元素
type ALGraphElem struct {
	vex  VerTextType
	side []ALElem
}

//NewALGraph 实例化邻接表
func NewALGraph(vexs []VerTextType, elems []ALGraphElem) *ALGraph {
	if len(vexs) == 0 || len(elems) == 0 {
		return nil
	}
	al := ALGraph{
		vexnum: len(vexs),
	}
	al.createALGraph(vexs, elems)
	return &al
}

//createALGraph 创建邻接表
func (this *ALGraph) createALGraph(vexs []VerTextType, elems []ALGraphElem) {
	//创建顶点表
	this.createVex(vexs)
	//创建边表
	this.createArc(elems)
}

//createVex 创建顶点表
func (this *ALGraph) createVex(vexs []VerTextType) {
	this.vertices = make([]VexNode, 0, this.vexnum)
	for _, vex := range vexs {
		data := VexNode{
			data:    vex,
			arcNode: nil,
		}
		this.vertices = append(this.vertices, data)
	}
}

//createArc 创建边表
func (this *ALGraph) createArc(elems []ALGraphElem) {
	for _, elem := range elems {
		//找到顶点
		//遍历边，生产边的链表
		//顶点指向链表第一个节点
		vexi := this.getIndexByVex(elem.vex)

		for _, side := range elem.side {
			//顶点到边
			v1i := this.getIndexByVex(side.V1)
			this.pushSide(vexi, v1i, side.Weight)
			//边到顶点
			sideVexIndex := this.getIndexByVex(side.V1)
			this.pushSide(sideVexIndex, vexi, side.Weight)
			this.arcnum++
		}
	}
}

//pushSide 添加邻接边表
func (this *ALGraph) pushSide(vexi, sidei int, weight ArcType) {
	data := ArcNode{
		adjvex: sidei,
		next:   nil,
		info:   weight,
	}

	vexNode := this.vertices[vexi].arcNode
	if vexNode != nil {
		data.next = vexNode
	}
	vexNode = &data
	this.vertices[vexi].arcNode = vexNode
}

//getIndexByVex 获取顶点表下标
func (this *ALGraph) getIndexByVex(vex VerTextType) int {
	for index, vexVal := range this.vertices {
		if vex == vexVal.data {
			return index
		}
	}
	return -1
}

//forBreadthALGraph 广度优先搜索遍历
func (this *ALGraph) forBreadthALGraph() {
	visited := make([]bool, this.vexnum, this.vexnum)
	this.BFS(0, visited)
}

//BFS 广度优先搜索
func (this *ALGraph) BFS(v int, visited []bool) {
	queue := NewQueue()
	queue.pushQueue(v)
	for !queue.emptyQueue() {
		u := queue.popQueue()
		arcNode := this.vertices[u].arcNode
		for arcNode != nil {
			if !visited[arcNode.adjvex] {
				fmt.Printf("顶点：%d,value:%s\n", arcNode.adjvex, this.vertices[arcNode.adjvex].data)
				visited[arcNode.adjvex] = true
				queue.pushQueue(arcNode.adjvex)
			}
			arcNode = arcNode.next
		}
	}
}
