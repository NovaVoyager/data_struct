package CriticalPathAlgorithm

import (
	"fmt"
	"strings"
)

//VerTextType 顶点数据类型
type VerTextType string

//ArcType 边的权值数据类型
type ArcType int

//ArcNode 邻接表，边的存储结构
type ArcNode struct {
	adjvex int
	next   *ArcNode
	info   ArcType
}

//VexNode 顶点表
type VexNode struct {
	//inNum 入度数量
	inNum int

	data    VerTextType
	arcNode *ArcNode
}

//ALGraph 邻接表
type ALGraph struct {
	vertices       []VexNode
	vexnum, arcnum int

	//ve 事件最早发生时间
	ve []ArcType
	//vl 事件最迟发生时间
	vl []ArcType
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
	Vex  VerTextType
	Side []ALElem
}

type Vexs struct {
	V     VerTextType
	InNum int
}

//NewALGraph 实例化邻接表
func NewALGraph(vexs []Vexs, elems []ALGraphElem) *ALGraph {
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
func (this *ALGraph) createALGraph(vexs []Vexs, elems []ALGraphElem) {
	//创建顶点表
	this.createVex(vexs)
	//创建边表
	this.createArc(elems)
}

//createVex 创建顶点表
func (this *ALGraph) createVex(vexs []Vexs) {
	this.vertices = make([]VexNode, 0, this.vexnum)
	for _, vex := range vexs {
		data := VexNode{
			data:    vex.V,
			arcNode: nil,
			inNum:   vex.InNum,
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
		vexi := this.getIndexByVex(elem.Vex)

		for _, side := range elem.Side {
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

//AOVSort 拓扑排序
func (this *ALGraph) AOVSort() {
	queue := NewAOVQueue() //初始化辅助队列
	count := 0             //初始化累加器，用于判断是否有回路
	aovPath := ""
	for _, vex := range this.vertices {
		if vex.inNum == 0 {
			queue.pushQueue(vex)
		}
	}

	for !queue.emptyQueue() {
		vex := queue.popQueue()
		aovPath += string(vex.data) + " -> "
		//将vex的各个邻接点入度减1
		tmp := vex.arcNode
		for tmp != nil {
			this.vertices[tmp.adjvex].inNum--
			if this.vertices[tmp.adjvex].inNum == 0 { //将新的入度为0的顶点加入队列
				queue.pushQueue(this.vertices[tmp.adjvex])
			}
			tmp = tmp.next
		}
		count++
	}

	if count < this.vexnum {
		fmt.Println("存在回路！")
	}
	fmt.Println("拓扑排序：", strings.TrimRight(aovPath, " -> "))
}
