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
	return g
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
	archs := make([][]ArcType, 0, this.vexNum)
	for i := 0; i < this.vexNum; i++ {
		cols := make([]ArcType, this.vexNum, this.vexNum)
		for j := 0; j < this.vexNum; j++ {
			cols[j] = MaxInt
		}
		archs = append(archs, cols)
	}
}
