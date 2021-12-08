package CriticalPathAlgorithm

import (
	"errors"
	"fmt"
)

var ExistLoopErr = errors.New("图存在回路！")

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
			//sideVexIndex := this.getIndexByVex(side.V1)
			//this.pushSide(sideVexIndex, vexi, side.Weight)
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
func (this *ALGraph) AOVSort() ([]int, error) {
	//初始化事件最早发生时间数组
	this.ve = make([]ArcType, this.vexnum)

	queue := NewAOVQueue() //初始化辅助队列
	count := 0             //初始化累加器，用于判断是否有回路
	topResult := make([]int, 0, this.vexnum)
	//aovPath := ""
	for _, vex := range this.vertices {
		if vex.inNum == 0 {
			queue.pushQueue(vex)
		}
	}

	for !queue.emptyQueue() {
		vex := queue.popQueue()
		vexIndex := this.getIndexByVex(vex.data)
		topResult = append(topResult, vexIndex)
		//aovPath += string(vex.data) + " -> "
		//将vex的各个邻接点入度减1
		tmp := vex.arcNode
		for tmp != nil {
			//拓扑排序
			this.vertices[tmp.adjvex].inNum--
			if this.vertices[tmp.adjvex].inNum == 0 { //将新的入度为0的顶点加入队列
				queue.pushQueue(this.vertices[tmp.adjvex])
			}

			//事件最早发生时间计算
			if this.ve[vexIndex]+tmp.info > this.ve[tmp.adjvex] {
				this.ve[tmp.adjvex] = this.ve[vexIndex] + tmp.info
			}

			tmp = tmp.next
		}
		count++
	}

	if count < this.vexnum {
		return nil, ExistLoopErr
	}
	//fmt.Println("拓扑排序：", strings.TrimRight(aovPath, " -> "))
	return topResult, nil
}

func (this *ALGraph) CriticalPath() error {
	topResult, err := this.AOVSort()
	if err != nil {
		return err
	}

	//事件最迟发生时间
	//把topResult数组看成一个栈
	stackTop := len(topResult) - 1
	//弹出汇点的元素
	inVex := topResult[stackTop]
	stackTop--
	//初始化vl数组为ve结束点的值
	this.vl = make([]ArcType, this.vexnum)
	for i := 0; i < this.vexnum; i++ {
		this.vl[i] = this.ve[inVex]
	}

	//栈不为空时，按拓扑逆序求各个顶点的vl值
	for stackTop != -1 {
		inVex = topResult[stackTop]
		stackTop--

		tmp := this.vertices[inVex].arcNode
		for tmp != nil {
			if this.vl[inVex] > this.vl[tmp.adjvex]-tmp.info {
				this.vl[inVex] = this.vl[tmp.adjvex] - tmp.info
			}
			tmp = tmp.next
		}
	}

	//关键路径输出
	for in := 0; in < this.vexnum; in++ {
		vex := this.vertices[in]
		tmp := vex.arcNode
		for tmp != nil {
			ee := this.ve[in]
			el := this.vl[tmp.adjvex] - tmp.info
			if ee == el { //关键路径
				fmt.Printf("<%s,%s>weight:%d ", vex.data, this.vertices[tmp.adjvex].data, tmp.info)
			}
			tmp = tmp.next
		}
	}

	return nil
}
