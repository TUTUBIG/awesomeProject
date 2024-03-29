package algorithm

import "fmt"

type listNode struct {
	data int
	next *listNode
}

type listNodeList struct {
	head *listNode
	tail *listNode
	size int
}

func (lnl *listNodeList) length() int {
	return lnl.size
}

func (lnl *listNodeList) add(d int) {
	if lnl.head == nil {
		lnl.head = &listNode{data: d}
		lnl.tail = lnl.head
	} else if lnl.tail != nil {
		lnl.tail.next = &listNode{data: d}
		lnl.tail = lnl.tail.next
	}

	lnl.size++
}

func (lnl *listNodeList) poll() int {
	if lnl.size == 0 {
		return 0
	}
	d := lnl.head.data
	lnl.head = lnl.head.next
	lnl.size--
	return d
}

func (lnl *listNodeList) get(n int) int {
	if lnl.size < n+1 {
		return 0
	}
	cur := lnl.head
	for i := 0; i < lnl.size; i++ {
		if i == n {
			return cur.data
		} else {
			cur = cur.next
		}
	}
	return 0
}

type graph struct {
	v     int                   // the number of vertex
	adj   map[int]*listNodeList //adjacency list
	found bool
}

func createGraph(v int) *graph {
	g := new(graph)
	g.v = v
	g.adj = make(map[int]*listNodeList, v)
	for i := 0; i < v; i++ {
		g.adj[i] = new(listNodeList)
	}
	return g
}

func (g *graph) add(s, t int) {
	g.adj[s].add(t)
	g.adj[t].add(s)
}

func (g *graph) BreadthFirstInSearch(s, t int) []int {
	if s == t {
		return nil
	}

	visited := make(map[int]bool, g.v)
	visited[s] = true
	queue := new(listNodeList)
	queue.add(s)
	prev := make([]int, g.v)
	for i := 0; i < g.v; i++ {
		prev[i] = -1
	}
	for queue.length() > 0 {
		w := queue.poll()
		for i := 0; i < g.adj[w].length(); i++ {
			w1 := g.adj[w].get(i)
			if !visited[w1] {
				prev[w1] = w
				if w1 == t {
					return prev
				}
				visited[w1] = true
				queue.add(w1)
			}
		}
	}
	return prev
}

func (g *graph) DepthFirstInSearch(s, t int) []int {
	prev := make([]int, g.v)
	visited := make([]bool, g.v)
	for i := 0; i < g.v; i++ {
		prev[i] = -1
	}
	g.recurDfs(s, t, visited, prev)
	return prev
}

func (g *graph) recurDfs(w, t int, visited []bool, prev []int) {
	if g.found {
		return
	}
	visited[w] = true
	if w == t {
		g.found = true
		return
	}
	for i := 0; i < g.adj[w].length(); i++ {
		w1 := g.adj[w].get(i)
		if !visited[w1] {
			prev[w1] = w
			g.recurDfs(w1, t, visited, prev)
		}
	}
}

func printPath(path []int, s, t int) {
	if len(path) < t {
		return
	}
	if path[t] != -1 && s != t {
		printPath(path, s, path[t])
	}
	fmt.Print(t, ">")
}
