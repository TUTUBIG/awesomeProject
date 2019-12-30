package algorithm

import "fmt"

type BinaryTreeNode struct {
	data           int
	lChild, rChild *BinaryTreeNode
}

type BinaryTree struct {
	*BinaryTreeNode
}

func initRoot(data int) BinaryTree {
	bt := new(BinaryTree)
	bt.BinaryTreeNode = new(BinaryTreeNode)
	bt.data = data
	return *bt
}

func (bt *BinaryTree) insert(data int) {
	if bt == nil {
		return
	}

	cur := bt.BinaryTreeNode

	for cur != nil {
		if data < cur.data {
			if cur.lChild == nil {
				cur.lChild = &BinaryTreeNode{
					data: data,
				}
				break
			}
			cur = cur.lChild
			continue
		}
		if data > cur.data {
			if cur.rChild == nil {
				cur.rChild = &BinaryTreeNode{
					data: data,
				}
				break
			}
			cur = cur.rChild
			continue
		}
	}
}

func (bt *BinaryTree) print() {
	if bt == nil {
		return
	}

	stack := new(Stack)
	cur := bt.BinaryTreeNode

	for cur != nil {
		if cur.lChild != nil {
			stack.Push(cur)
			cur = cur.lChild
			continue
		}
		fmt.Print(cur.data, "->")
		father, ok := stack.Pop().(*BinaryTreeNode)
		if ok {
			fmt.Print(father.data, "->")
			cur = father
			cur.lChild = nil
		} else if cur.rChild == nil {
			break
		}

		if cur.rChild != nil {
			cur = cur.rChild
		}
	}
}
