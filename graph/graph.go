package main

const (
	//MaxInt 极大值，表示无穷大
	MaxInt = 32767
)

//VerTextType 顶点数据类型
type VerTextType string

//ArcType 边的权值数据类型
type ArcType int

//Elem 构建图的边元素
type Elem struct {
	//V1 顶点1
	V1 VerTextType

	//V2 顶点2
	V2 VerTextType

	//Weight 权值
	Weight ArcType
}

type AMGraph struct {
	//vexs 顶点表,记录顶点数据
	vexs []VerTextType

	//arcs 邻接矩阵
	arcs [][]ArcType

	//vexNum 顶点数
	vexNum int

	//arcNum 边数
	arcNum int
}

//NewAMGraph 无向网
func NewAMGraph(vexs []VerTextType, elems []Elem) *AMGraph {
	if len(vexs) == 0 || len(elems) == 0 {
		return nil
	}
	g := &AMGraph{
		vexNum: len(vexs),
		arcNum: len(elems),
	}
	g.createVexs(vexs)
	g.initArcs()
	g.setArcs(elems)
	return g
}

func (this *AMGraph) setArcs(elems []Elem) {
	for _, elem := range elems {
		v1i := this.getIndexByVex(elem.V1)
		v2i := this.getIndexByVex(elem.V2)
		this.arcs[v1i][v2i] = elem.Weight
		this.arcs[v2i][v1i] = elem.Weight
	}
}

//createVexs 创建顶点表
func (this *AMGraph) createVexs(vexs []VerTextType) {
	this.vexs = make([]VerTextType, 0, len(vexs))
	for _, vex := range vexs {
		this.vexs = append(this.vexs, vex)
	}
}

//initArcs 初始化邻接矩阵表
func (this *AMGraph) initArcs() {
	this.arcs = make([][]ArcType, 0, this.vexNum)
	for i := 0; i < this.vexNum; i++ {
		cols := make([]ArcType, this.vexNum, this.vexNum)
		for j := 0; j < this.vexNum; j++ {
			cols[j] = MaxInt
		}
		this.arcs = append(this.arcs, cols)
	}
}

//getIndexByVex 获取顶点表下标
func (this *AMGraph) getIndexByVex(vex VerTextType) int {
	for index, vexVal := range this.vexs {
		if vex == vexVal {
			return index
		}
	}
	return -1
}
