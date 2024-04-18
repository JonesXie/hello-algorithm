package main

import (
	"fmt"
	. "hello-algo/pkg"
)

/* 广度优先遍历 */
// 使用邻接表来表示图，以便获取指定顶点的所有邻接顶点
func graphBFS(g *graphAdjList, startVet Vertex) []Vertex {
	// 顶点遍历序列
	res := make([]Vertex, 0)
	// 哈希表，用于记录已被访问过的顶点
	visited := make(map[Vertex]struct{})
	visited[startVet] = struct{}{}
	// 队列用于实现 BFS, 使用切片模拟队列
	queue := make([]Vertex, 0)
	queue = append(queue, startVet)
	// 以顶点 vet 为起点，循环直至访问完所有顶点
	for len(queue) > 0 {
		// 队首顶点出队
		vet := queue[0]
		queue = queue[1:]
		// 记录访问顶点
		res = append(res, vet)
		// 遍历该顶点的所有邻接顶点
		for _, adjVet := range g.adjList[vet] {
			_, isExist := visited[adjVet]
			// 只入队未访问的顶点
			if !isExist {
				queue = append(queue, adjVet)
				visited[adjVet] = struct{}{}
			}
		}
	}
	// 返回顶点遍历序列
	return res
}

/* 深度优先遍历辅助函数 */
func dfs(g *graphAdjList, visited map[Vertex]struct{}, res *[]Vertex, vet Vertex) {
	// append 操作会返回新的的引用，必须让原引用重新赋值为新slice的引用
	*res = append(*res, vet)
	visited[vet] = struct{}{}
	// 遍历该顶点的所有邻接顶点
	for _, adjVet := range g.adjList[vet] {
		_, isExist := visited[adjVet]
		// 递归访问邻接顶点
		if !isExist {
			dfs(g, visited, res, adjVet)
		}
	}
}

/* 深度优先遍历 */
// 使用邻接表来表示图，以便获取指定顶点的所有邻接顶点
func graphDFS(g *graphAdjList, startVet Vertex) []Vertex {
	// 顶点遍历序列
	res := make([]Vertex, 0)
	// 哈希表，用于记录已被访问过的顶点
	visited := make(map[Vertex]struct{})
	dfs(g, visited, &res, startVet)
	// 返回顶点遍历序列
	return res
}

func bfsTest() {
	/* 初始化无向图 */
	vets := ValsToVets([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	edges := [][]Vertex{
		{vets[0], vets[1]}, {vets[0], vets[3]}, {vets[1], vets[2]}, {vets[1], vets[4]},
		{vets[2], vets[5]}, {vets[3], vets[4]}, {vets[3], vets[6]}, {vets[4], vets[5]},
		{vets[4], vets[7]}, {vets[5], vets[8]}, {vets[6], vets[7]}, {vets[7], vets[8]}}
	graph := newGraphAdjList(edges)
	fmt.Println("初始化后，图为:")
	graph.print()

	/* 广度优先遍历 */
	res := graphBFS(graph, vets[0])
	fmt.Println("广度优先遍历（BFS）顶点序列为:")
	PrintSlice(VetsToVals(res))

}

func dfsTest() {
	/* 初始化无向图 */
	vets := ValsToVets([]int{0, 1, 2, 3, 4, 5, 6})
	edges := [][]Vertex{
		{vets[0], vets[1]}, {vets[0], vets[3]}, {vets[1], vets[2]},
		{vets[2], vets[5]}, {vets[4], vets[5]}, {vets[5], vets[6]}}
	graph := newGraphAdjList(edges)
	fmt.Println("初始化后，图为:")
	graph.print()

	/* 深度优先遍历 */
	res := graphDFS(graph, vets[0])
	fmt.Println("深度优先遍历（DFS）顶点序列为:")
	PrintSlice(VetsToVals(res))
}

func main() {
	bfsTest()
	dfsTest()
}
