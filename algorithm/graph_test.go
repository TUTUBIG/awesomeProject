package algorithm

import "testing"

func TestGraph_BreadthInSearch(t *testing.T) {
	g := createGraph(5)
	g.add(1, 2)
	g.add(1, 4)
	g.add(2, 3)
	g.add(2, 4)
	g.add(2, 5)
	g.add(3, 5)
	g.add(4, 5)
	t.Log(g.BreadthInSearch(1, 3))
}
