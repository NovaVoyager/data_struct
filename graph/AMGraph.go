package main

import (
	"fmt"
	"sort"
	"strings"
)

const (
	//MaxInt 极大值，表示无穷大
	MaxInt = 32767

	ZeroFlag = 0
	OneFlag  = 1

	//directedFlag 有向标记
	directedFlag = true
	//unDirectedFlag 无向标记
	unDirectedFlag = false
)

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

//NewNoAMGraph 无向网
func NewNoAMGraph(vexs []VerTextType, elems []Elem) *AMGraph {
	if len(vexs) == 0 || len(elems) == 0 {
		return nil
	}
	g := &AMGraph{
		vexNum: len(vexs),
		arcNum: len(elems),
	}
	g.createVexs(vexs)
	g.initArcs(MaxInt)
	g.setAMGraphArcs(elems, unDirectedFlag)
	return g
}

//NewAMGraph 有向网
func NewAMGraph(vexs []VerTextType, elems []Elem) *AMGraph {
	if len(vexs) == 0 || len(elems) == 0 {
		return nil
	}
	g := &AMGraph{
		vexNum: len(vexs),
		arcNum: len(elems),
	}
	g.createVexs(vexs)
	g.initArcs(MaxInt)
	g.setAMGraphArcs(elems, directedFlag)
	return g
}

//NewNoAMFigure 无向图
func NewNoAMFigure(vexs []VerTextType, elems []Elem) *AMGraph {
	if len(vexs) == 0 || len(elems) == 0 {
		return nil
	}
	g := &AMGraph{
		vexNum: len(vexs),
		arcNum: len(elems),
	}
	g.createVexs(vexs)
	g.initArcs(0)
	g.setNoFigure(elems)
	return g
}

//setNoFigure 无向图
func (this *AMGraph) setNoFigure(elems []Elem) {
	for _, elem := range elems {
		v1i := this.getIndexByVex(elem.V1)
		v2i := this.getIndexByVex(elem.V2)
		this.arcs[v1i][v2i] = OneFlag
		this.arcs[v2i][v1i] = OneFlag
	}
}

//setAMGraphArcs 无向网
func (this *AMGraph) setAMGraphArcs(elems []Elem, direction bool) {
	for _, elem := range elems {
		v1i := this.getIndexByVex(elem.V1)
		v2i := this.getIndexByVex(elem.V2)
		this.arcs[v1i][v2i] = elem.Weight
		if !direction { //无向
			this.arcs[v2i][v1i] = elem.Weight
		}
	}
}

//createVexs 创建顶点表
func (this *AMGraph) createVexs(vexs []VerTextType) {
	this.vexs = vexs
}

//initArcs 初始化邻接矩阵表
func (this *AMGraph) initArcs(weight ArcType) {
	this.arcs = make([][]ArcType, 0, this.vexNum)
	for i := 0; i < this.vexNum; i++ {
		cols := make([]ArcType, this.vexNum, this.vexNum)
		for j := 0; j < this.vexNum; j++ {
			if i == j {
				cols[j] = ZeroFlag
			} else {
				cols[j] = weight
			}
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

//forDepthAMGraph 深度优先遍历图
func (this *AMGraph) forDepthAMGraph(v int) {
	visited := make([]bool, len(this.vexs), len(this.vexs))
	this.DFS(v, visited)
}

//DFS 深度优先遍历
func (this *AMGraph) DFS(v int, visited []bool) {
	fmt.Printf("顶点：%d,value:%s\n", v, this.vexs[v])
	visited[v] = true
	for i := 0; i < this.vexNum; i++ {
		if this.arcs[v][i] != 0 && !visited[i] {
			this.DFS(i, visited)
		}
	}
}

//shortEdge 普里姆算法辅助数组
type shortEdge struct {
	//adjvex 顶点下标
	adjvex int

	//lowCost 两个顶点之间边的最小权值,当为0时，表示该顶点已被加入U集合
	lowCost ArcType
}

//Prim 普利姆算法 最小生成树
func (this *AMGraph) Prim(startV int) {
	if this.vexNum == 0 {
		fmt.Println("图为空!!!")
		return
	}

	//初始辅助数组
	shortEdges := make([]shortEdge, 0, this.vexNum)
	for i := 0; i < this.vexNum; i++ {
		data := shortEdge{
			adjvex:  startV,
			lowCost: this.arcs[startV][i],
		}
		shortEdges = append(shortEdges, data)
	}
	shortEdges[startV].lowCost = 0 //把初始顶点加入集合U

	for i := 0; i < this.vexNum-1; i++ {
		k := this.minEdge(shortEdges)    //寻找最短边的邻接点
		this.outputMST(k, shortEdges[k]) //输出最小生成树路径
		shortEdges[k].lowCost = 0        //将顶点k加入到集合U中
		for j := 0; j < this.vexNum; j++ {
			if this.arcs[k][j] < shortEdges[j].lowCost {
				shortEdges[j].lowCost = this.arcs[k][j]
				shortEdges[j].adjvex = k
			}
		}
	}
}

//minEdge 寻找最短边的邻接点
func (this *AMGraph) minEdge(se []shortEdge) int {
	min := 0
	var minVal ArcType
	for i, v := range se {
		if v.lowCost != 0 { //初始化不属于U集合的值
			min = i
			minVal = v.lowCost
		}
	}

	for i := 0; i < this.vexNum; i++ {
		if se[i].lowCost == 0 { //该顶点已属于U集合
			continue
		}
		if se[i].lowCost < minVal {
			min = i
			minVal = se[i].lowCost
		}
	}

	return min
}

//outputMST 输出MST
func (this *AMGraph) outputMST(k int, se shortEdge) {
	fmt.Println(fmt.Sprintf("(%d,%d)%d", se.adjvex, k, se.lowCost))
}

//Dijkstra 迪杰斯特拉算法
func (this *AMGraph) Dijkstra(startV int) {
	//初始化元素
	s := make([]int, this.vexNum, this.vexNum)    //保存是否被访问元素
	dist := make([]ArcType, 0, this.vexNum)       //保存顶点到到各顶点最小的权值
	path := make([]int, this.vexNum, this.vexNum) //保存各顶最小权值路径
	for i := 0; i < this.vexNum; i++ {
		dist = append(dist, this.arcs[startV][i])
		if this.arcs[startV][i] == MaxInt {
			path[i] = -1
		} else {
			path[i] = startV
		}
	}
	s[startV] = 1 //把初始顶点加入s集合

	num := 1
	for num < this.vexNum {
		num++
		//找到权值最小元素返回下标
		min := this.findMinDist(s, dist)
		s[min] = 1 //加入s集合

		//调整dist和path
		for i := 0; i < this.vexNum; i++ {
			if (s[i] == 0) && (dist[i] > dist[min]+this.arcs[min][i]) {
				dist[i] = dist[min] + this.arcs[min][i] //找到的最短路径去修改dist对应的顶点值
				path[i] = min                           //更新path为当前最小路径顶点值
			}
		}
	}

	//打印起始点到各个顶点最短路径
	this.displayPath(startV, dist, path)
}

//displayPath 打印起始点到各个顶点最短路径
func (this *AMGraph) displayPath(startV int, dist []ArcType, path []int) {
	for i := 0; i < this.vexNum; i++ {
		if i == startV { //访问的是起始顶点
			continue
		}
		fmt.Printf("顶点：%s,下标：%d -> 顶点：%s,下标：%d 的最短路径权值：%d  ", this.vexs[startV], startV, this.vexs[i],
			i, dist[i])
		func(pathIndex int) {
			paths := []int{pathIndex}
			tmp := path[pathIndex]
			paths = append(paths, tmp)
			for tmp != 0 {
				tmp = path[tmp]
				paths = append(paths, tmp)
			}

			//逆向输出
			var s VerTextType = ""
			for j := len(paths) - 1; j >= 0; j-- {
				s += this.vexs[paths[j]] + "->"
			}
			newStr := strings.TrimRight(string(s), "->")
			fmt.Printf("%s", newStr)
		}(i)

		fmt.Printf("\n")
	}

}

//findMinDist 获取最小权值下标
func (this *AMGraph) findMinDist(s []int, dist []ArcType) int {
	minI := 0
	var minV ArcType = MaxInt
	for i := 0; i < this.vexNum; i++ {
		if s[i] == 1 { //存在s集合中,跳过
			continue
		}
		if dist[i] < minV {
			minV = dist[i]
			minI = i
		}
	}
	return minI
}

//EdgeType 克鲁斯卡尔算法辅助数据结构
type EdgeType struct {
	//from 出发顶点
	from int

	//to 到达顶点
	to int

	//weight 边上权值
	weight ArcType
}

//EdgeGraph 克鲁斯卡尔算法图的存储结构
type EdgeGraph struct {

	//vertexs 顶点数组
	vertexs []VerTextType

	//edges 存放边的数组
	edges []EdgeType

	//vexNum 顶点数
	vexNum int

	//edgeNum 边数
	edgeNum int
}

//NewEdgeGraph 创建对象
func NewEdgeGraph(vertexs []VerTextType, edges []EdgeType) *EdgeGraph {
	if len(vertexs) == 0 || len(edges) == 0 {
		return nil
	}
	g := EdgeGraph{
		vertexs: vertexs,
		edges:   edges,
		vexNum:  len(vertexs),
		edgeNum: len(edges),
	}
	g.sort()
	return &g
}

//sort 对边按照权值 由小到大排序
func (this *EdgeGraph) sort() {
	sort.Slice(this.edges, func(i, j int) bool {
		if this.edges[i].weight < this.edges[j].weight {
			return true
		}
		return false
	})
}

//Kruskal 克鲁斯卡尔算法 最小生成树
func (this *EdgeGraph) Kruskal() {
	//初始化双亲
	parent := make([]int, 0, this.vexNum)
	for i := 0; i < this.vexNum; i++ {
		parent = append(parent, -1)
	}

	for i, num := 0, 0; i < this.vexNum; i++ {
		//找到所在生成树的根节点
		vex1 := this.findRoot(parent, this.edges[i].from)
		//找到所在生成树的根节点
		vex2 := this.findRoot(parent, this.edges[i].to)
		if vex1 != vex2 { //找到两个根节点不相同，说明在两个联通分量，不会构成环
			this.outputMST(this.edges[i])

			//合并生成树
			parent[vex2] = vex1
			num++
		}

		//判断边数为顶点数-1时，直接返回
		if num == this.vexNum-1 {
			break
		}
	}
}

//findRoot 找到所在生成树的根节点
func (this *EdgeGraph) findRoot(parent []int, v int) int {
	t := v
	for parent[t] > -1 {
		t = parent[t]
	}
	return t
}

//outputMST 输出MST
func (this *EdgeGraph) outputMST(edge EdgeType) {
	fmt.Println(fmt.Sprintf("(%d,%d)%d", edge.from, edge.to, edge.weight))
}
