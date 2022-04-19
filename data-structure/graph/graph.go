package main

import "fmt"

type Graph struct {
	data [][]int
}

func newGraph(length int) *Graph {
	grpah := new(Graph)
	grpah.data = make([][]int, length)
	return grpah
}

func (g *Graph) addEdge(u, v int) {
	g.data[u] = append(g.data[u], v)
	g.data[v] = append(g.data[v], u)
}

func (g *Graph) print() {
	for i, v := range g.data {
		fmt.Println("graph representation with list of vertext ", v)
		fmt.Print("vertex : ", i)
		for _, vv := range v {
			fmt.Print(" -> ", vv)
		}
		fmt.Println("")
	}
}

// graph representation
// 0 --- 1
// |    | \
// |    |  2
// |    | /
// 4 ----3

func main() {
	graph := newGraph(5)
	graph.addEdge(0, 1)
	graph.addEdge(0, 4)
	graph.addEdge(1, 2)
	graph.addEdge(1, 3)
	graph.addEdge(1, 4)
	graph.addEdge(2, 3)
	graph.addEdge(3, 4)
	fmt.Println(graph.data)
	graph.print()
}
