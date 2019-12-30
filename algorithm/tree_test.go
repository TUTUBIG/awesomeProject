package algorithm

import "testing"

func TestTree(t *testing.T) {
	bt := initRoot(1)
	bt.insert(3)
	bt.insert(4)
	bt.insert(2)
	bt.print()
}
