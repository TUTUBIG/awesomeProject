package algorithm

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
	}

	if lnl.tail != nil {
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
	if lnl.size < n {
		return 0
	}
	cur := lnl.head
	for i := 1; i < lnl.size; i++ {
		if i == n {
			return cur.data
		} else {
			cur = cur.next
		}
	}
	return 0
}

type graph struct {
	v   int                   // the number of vertex
	adj map[int]*listNodeList //adjacency list
}

func createGraph(v int) *graph {
	g := new(graph)
	g.v = v
	g.adj = make(map[int]*listNodeList, v)
	for i := 1; i <= v; i++ {
		g.adj[i] = new(listNodeList)
	}
	return g
}

func (g *graph) add(s, t int) {
	g.adj[s].add(t)
	g.adj[t].add(s)
}

func (g *graph) BreadthInSearch(s, t int) []int {
	if s == t {
		return nil
	}

	visited := make(map[int]bool, g.v)
	visited[s] = true
	queue := new(listNodeList)
	queue.add(s)
	prev := make([]int, g.v)
	for queue.length() > 0 {
		w := queue.poll()
		for i := 0; i < g.adj[w].length(); i++ {
			w1 := g.adj[w].get(i + 1)
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

func (g *graph) DepthInSearch(s, t int) []int {
	return nil
}
