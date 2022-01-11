package algorithm

import "testing"

func exampleGraph() *graph {
	g := createGraph(8)
	g.add(0, 1)
	g.add(0, 3)
	g.add(1, 2)
	g.add(1, 4)
	g.add(3, 4)
	g.add(2, 5)
	g.add(4, 5)
	g.add(4, 6)
	g.add(5, 7)
	g.add(6, 7)
	return g
}

func TestGraph_BreadthFirstInSearch(t *testing.T) {
	g := exampleGraph()
	path := g.BreadthFirstInSearch(0, 7)
	printPath(path, 0, 7)
}

func TestGraph_DepthFirstInSearch(t *testing.T) {
	g := exampleGraph()
	path := g.DepthFirstInSearch(0, 7)
	printPath(path, 0, 7)
}
