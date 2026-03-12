package main

import (
	"fmt"
	"sort"
)

// Edge 表示图中的一条边
type Edge struct {
	U, V, Weight int
}

// Graph 表示图
type Graph struct {
	Vertices int
	Edges    []Edge
}

// UnionFind 并查集结构
type UnionFind struct {
	parent []int
	rank   []int
}

// NewUnionFind 创建并查集
func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}
	return &UnionFind{parent, rank}
}

// Find 查找元素的根节点（带路径压缩）
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // 路径压缩
	}
	return uf.parent[x]
}

// Union 合并两个集合（按秩合并）
func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX != rootY {
		// 按秩合并
		if uf.rank[rootX] < uf.rank[rootY] {
			uf.parent[rootX] = rootY
		} else if uf.rank[rootX] > uf.rank[rootY] {
			uf.parent[rootY] = rootX
		} else {
			uf.parent[rootY] = rootX
			uf.rank[rootX]++
		}
	}
}

// Kruskal 实现Kruskal算法求最小生成树
func Kruskal(g *Graph) []Edge {
	// 结果切片，存储最小生成树的边
	result := make([]Edge, 0)

	// 1. 按照边的权重排序
	sortedEdges := make([]Edge, len(g.Edges))
	copy(sortedEdges, g.Edges)
	sort.Slice(sortedEdges, func(i, j int) bool {
		return sortedEdges[i].Weight < sortedEdges[j].Weight
	})

	// 2. 初始化并查集
	uf := NewUnionFind(g.Vertices)

	// 3. 遍历排序后的边
	for _, edge := range sortedEdges {
		// 检查这条边连接的两个顶点是否属于同一集合
		rootU := uf.Find(edge.U)
		rootV := uf.Find(edge.V)

		// 如果不在同一集合，不会形成环
		if rootU != rootV {
			result = append(result, edge)
			uf.Union(rootU, rootV)

			// 最小生成树有 V-1 条边时即可停止
			if len(result) == g.Vertices-1 {
				break
			}
		}
	}

	return result
}

// 计算最小生成树的总权重
func totalWeight(edges []Edge) int {
	sum := 0
	for _, e := range edges {
		sum += e.Weight
	}
	return sum
}

func main() {
	// 示例：创建一个图
	// 顶点编号从0开始
	g := &Graph{
		Vertices: 6,
		Edges: []Edge{
			{0, 1, 4}, // 边 0-1 权重4
			{0, 2, 3}, // 边 0-2 权重3
			{1, 2, 1}, // 边 1-2 权重1
			{1, 3, 2}, // 边 1-3 权重2
			{2, 3, 4}, // 边 2-3 权重4
			{2, 4, 3}, // 边 2-4 权重3
			{3, 4, 5}, // 边 3-4 权重5
			{3, 5, 6}, // 边 3-5 权重6
			{4, 5, 7}, // 边 4-5 权重7
		},
	}

	fmt.Println("Kruskal算法求最小生成树")
	fmt.Println("图的边数:", len(g.Edges))
	fmt.Println("顶点数:", g.Vertices)
	fmt.Println()

	mst := Kruskal(g)

	fmt.Println("最小生成树的边:")
	for _, edge := range mst {
		fmt.Printf("%d -- %d 权重: %d\n", edge.U, edge.V, edge.Weight)
	}

	fmt.Printf("\n最小生成树总权重: %d\n", totalWeight(mst))
}
