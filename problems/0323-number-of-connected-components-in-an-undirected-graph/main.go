package main

import "leetcode/helper"

// 无向图中连通分量的数目
func countComponents(n int, edges [][]int) int {
	graph := buildGraph(n, edges)
	visited := make([]bool, n)

	var ans int
	for i := range graph {
		if !visited[i] {
			dfs(i, graph, visited)
			ans += 1
		}
	}
	return ans
}

func buildGraph(n int, dislikes [][]int) [][]int {
	graph := make([][]int, n)
	for _, pair := range dislikes {
		i, j := pair[0], pair[1]
		graph[i] = append(graph[i], j)
		graph[j] = append(graph[j], i)
	}
	return graph
}

func dfs(vertex int, graph [][]int, visited []bool) {
	visited[vertex] = true
	for _, neighbor := range graph[vertex] {
		if !visited[neighbor] {
			dfs(neighbor, graph, visited)
		}
	}
}

func countComponents2(n int, edges [][]int) int {
	graph := buildGraph(n, edges)
	visited := make([]bool, n)

	var ans int
	for i := range graph {
		if !visited[i] {
			bfs(i, graph, visited)
			ans += 1
		}
	}
	return ans
}

func bfs(vertex int, graph [][]int, visited []bool) {
	visited[vertex] = true
	queue := []int{vertex}
	for i := 0; i < len(queue); i++ {
		node := queue[i]
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}

func countComponents3(n int, edges [][]int) int {
	uf := newUnionFind(n)
	for _, pair := range edges {
		i, j := pair[0], pair[1]
		uf.connect(i, j)
	}
	return uf.count()
}

type unionFind struct {
	ancestor []int
	n        int
}

func newUnionFind(n int) *unionFind {
	ancestor := make([]int, n)
	for i := range n {
		ancestor[i] = i
	}
	return &unionFind{n: n, ancestor: ancestor}
}

func (o *unionFind) find(x int) int {
	for o.ancestor[x] != x {
		o.ancestor[x] = o.find(o.ancestor[x])
	}
	return o.ancestor[x]
}

func (o *unionFind) count() int {
	return o.n
}

func (o *unionFind) connect(p, q int) {
	p = o.find(p)
	q = o.find(q)

	if p != q {
		o.ancestor[p] = q
		o.n--
	}
}

func testOne(n int, in string, ans int) {
	{
		matrix := helper.ParseIntMatrix(in)
		helper.Assert(countComponents(n, matrix) == ans)
	}

	{
		matrix := helper.ParseIntMatrix(in)
		helper.Assert(countComponents2(n, matrix) == ans)
	}

	{
		matrix := helper.ParseIntMatrix(in)
		helper.Assert(countComponents3(n, matrix) == ans)
	}
}

func main() {
	testOne(5, "[[0, 1], [1, 2], [3, 4]]", 2)
	testOne(5, "[[0,1], [1,2], [2,3], [3,4]]", 1)
}
