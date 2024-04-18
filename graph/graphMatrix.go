package main

import "fmt"

// graph 有两种实现方式，邻接矩阵和邻接表

/* 基于邻接矩阵实现的无向图类 -- graph_adjacency_matrix*/

type graphAdjMat struct {
	// 顶点列表，元素代表“顶点值”，索引代表“顶点索引”
	vertices []int
	// 邻接矩阵，行列索引对应“顶点索引”
	adjMat [][]int
}

/* 获取顶点数量 */
func (g *graphAdjMat) size() int {
	return len(g.vertices)
}

/* 添加顶点 */
func (g *graphAdjMat) addVertex(val int) {
	n := g.size()
	// 向顶点列表中添加新顶点的值
	g.vertices = append(g.vertices, val)
	// 在邻接矩阵中添加一行
	newRow := make([]int, n)
	g.adjMat = append(g.adjMat, newRow)
	// 在邻接矩阵中添加一列
	for i := range g.adjMat {
		g.adjMat[i] = append(g.adjMat[i], 0)
	}
}

/* 删除顶点 */
func (g *graphAdjMat) removeVertex(index int) {
	if index >= g.size() {
		return
	}
	// 在顶点列表中移除索引 index 的顶点
	g.vertices = append(g.vertices[:index], g.vertices[index+1:]...)
	// 在邻接矩阵中删除索引 index 的行
	g.adjMat = append(g.adjMat[:index], g.adjMat[index+1:]...)
	// 在邻接矩阵中删除索引 index 的列
	for i := range g.adjMat {
		g.adjMat[i] = append(g.adjMat[i][:index], g.adjMat[i][index+1:]...)
	}
}

/* 添加边 */
// 参数 i, j 对应 vertices 元素索引
func (g *graphAdjMat) addEdge(i, j int) {
	// 索引越界与相等处理
	if i < 0 || j < 0 || i >= g.size() || j >= g.size() || i == j {
		fmt.Errorf("%s", "Index Out Of Bounds Exception")
	}
	// 在无向图中，邻接矩阵关于主对角线对称，即满足 (i, j) == (j, i)
	g.adjMat[i][j] = 1
	g.adjMat[j][i] = 1
}

/* 删除边 */
// 参数 i, j 对应 vertices 元素索引
func (g *graphAdjMat) removeEdge(i, j int) {
	// 索引越界与相等处理
	if i < 0 || j < 0 || i >= g.size() || j >= g.size() || i == j {
		fmt.Errorf("%s", "Index Out Of Bounds Exception")
	}
	g.adjMat[i][j] = 0
	g.adjMat[j][i] = 0
}

/* 打印邻接矩阵 */
func (g *graphAdjMat) print() {
	fmt.Printf("\t顶点列表 = %v\n", g.vertices)
	fmt.Printf("\t邻接矩阵 = \n")
	for i := range g.adjMat {
		fmt.Printf("\t\t\t%v\n", g.adjMat[i])
	}
}

// 初始化
func newGraphAdjMat(vertices []int, edges [][]int) *graphAdjMat {
	// 添加顶点
	n := len(vertices)
	// 矩阵
	adjMat := make([][]int, n)
	for i := range adjMat {
		adjMat[i] = make([]int, n)
	}
	// 初始化图
	g := &graphAdjMat{
		vertices: vertices,
		adjMat:   adjMat,
	}
	// 添加边
	// 请注意，edges 元素代表顶点索引，即对应 vertices 元素索引
	for i := range edges {
		g.addEdge(edges[i][0], edges[i][1])

	}
	return g
}

func main() {
	/* 初始化无向图 */
	// 请注意，edges 元素代表顶点索引，即对应 vertices 元素索引
	vertices := []int{1, 3, 2, 5, 4}
	edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 3}, {2, 4}, {3, 4}}
	graph := newGraphAdjMat(vertices, edges)
	fmt.Println("初始化后，图为:")
	graph.print()

	/* 添加边 */
	// 顶点 1, 2 的索引分别为 0, 2
	graph.addEdge(0, 2)
	fmt.Println("添加边 1-2 后，图为")
	graph.print()

	/* 删除边 */
	// 顶点 1, 3 的索引分别为 0, 1
	graph.removeEdge(0, 1)
	fmt.Println("删除边 1-3 后，图为")
	graph.print()

	/* 添加顶点 */
	graph.addVertex(6)
	fmt.Println("添加顶点 6 后，图为")
	graph.print()

	/* 删除顶点 */
	// 顶点 3 的索引为 1
	graph.removeVertex(1)
	fmt.Println("删除顶点 3 后，图为")
	graph.print()
}
