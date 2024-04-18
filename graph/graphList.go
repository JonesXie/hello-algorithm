package main

import (
	"fmt"
	. "hello-algo/pkg"
	"strconv"
	"strings"
)

/* 基于邻接表实现的无向图类 */
type graphAdjList struct {
	// 邻接表，key：顶点，value：该顶点的所有邻接顶点
	adjList map[Vertex][]Vertex
}

/* 构造函数 */
func newGraphAdjList(edges [][]Vertex) *graphAdjList {
	g := &graphAdjList{
		adjList: make(map[Vertex][]Vertex),
	}
	// 添加所有顶点和边
	for _, edge := range edges {
		g.addVertex(edge[0])
		g.addVertex(edge[1])
		g.addEdge(edge[0], edge[1])
	}
	return g
}

/* 获取顶点数量 */
func (g *graphAdjList) size() int {
	return len(g.adjList)
}

/* 添加边 */
func (g *graphAdjList) addEdge(vet1 Vertex, vet2 Vertex) {
	_, ok1 := g.adjList[vet1]
	_, ok2 := g.adjList[vet2]
	if !ok1 || !ok2 || vet1 == vet2 {
		panic("error")
	}
	// 添加边 vet1 - vet2, 添加匿名 struct{},
	g.adjList[vet1] = append(g.adjList[vet1], vet2)
	g.adjList[vet2] = append(g.adjList[vet2], vet1)
}

/* 删除边 */
func (g *graphAdjList) removeEdge(vet1 Vertex, vet2 Vertex) {
	_, ok1 := g.adjList[vet1]
	_, ok2 := g.adjList[vet2]
	if !ok1 || !ok2 || vet1 == vet2 {
		panic("error")
	}
	// 删除边 vet1 - vet2
	g.adjList[vet1] = DeleteSliceElms(g.adjList[vet1], vet2)
	g.adjList[vet2] = DeleteSliceElms(g.adjList[vet2], vet1)
}

/* 添加顶点 */
func (g *graphAdjList) addVertex(vet Vertex) {
	_, ok := g.adjList[vet]
	if ok {
		return
	}
	// 在邻接表中添加一个新链表
	g.adjList[vet] = make([]Vertex, 0)
}

/* 删除顶点 */
func (g *graphAdjList) removeVertex(vet Vertex) {
	_, ok := g.adjList[vet]
	if !ok {
		panic("error")
	}
	// 在邻接表中删除顶点 vet 对应的链表
	delete(g.adjList, vet)
	// 遍历其他顶点的链表，删除所有包含 vet 的边
	for v, list := range g.adjList {
		g.adjList[v] = DeleteSliceElms(list, vet)
	}
}

/* 打印邻接表 */
func (g *graphAdjList) print() {
	var builder strings.Builder
	fmt.Printf("邻接表 = \n")
	for k, v := range g.adjList {
		builder.WriteString("\t\t" + strconv.Itoa(k.Val) + ": ")
		for _, vet := range v {
			builder.WriteString(strconv.Itoa(vet.Val) + " ")
		}
		fmt.Println(builder.String())
		builder.Reset()
	}
}

func main() {
	/* 初始化无向图 */
	v := ValsToVets([]int{1, 3, 2, 5, 4})
	edges := [][]Vertex{{v[0], v[1]}, {v[0], v[3]}, {v[1], v[2]}, {v[2], v[3]}, {v[2], v[4]}, {v[3], v[4]}}
	graph := newGraphAdjList(edges)
	fmt.Println("初始化后，图为:")
	graph.print()

	/* 添加边 */
	// 顶点 1, 2 即 v[0], v[2]
	graph.addEdge(v[0], v[2])
	fmt.Println("\n添加边 1-2 后，图为")
	graph.print()

	/* 删除边 */
	// 顶点 1, 3 即 v[0], v[1]
	graph.removeEdge(v[0], v[1])
	fmt.Println("\n删除边 1-3 后，图为")
	graph.print()

	/* 添加顶点 */
	v5 := NewVertex(6)
	graph.addVertex(v5)
	fmt.Println("\n添加顶点 6 后，图为")
	graph.print()

	/* 删除顶点 */
	// 顶点 3 即 v[1]
	graph.removeVertex(v[1])
	fmt.Println("\n删除顶点 3 后，图为")
	graph.print()
}
