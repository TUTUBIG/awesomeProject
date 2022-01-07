package main

import (
	"fmt"
	"math"
)

type BinaryTreeNode struct {
	data           int
	lChild, rChild *BinaryTreeNode
}

type BinaryTree struct {
	root *BinaryTreeNode
}

func (bt *BinaryTree) getDiameter() int {
	diameterSum := 0
	return diameter(bt.root, &diameterSum)
}

func diameter(root *BinaryTreeNode, diameterSum *int) int {
	if root == nil {
		return 0
	}
	lDiameter, rDiameter := diameter(root.lChild, diameterSum), diameter(root.rChild, diameterSum)
	*diameterSum = int(math.Max(float64(lDiameter)+float64(rDiameter), float64(*diameterSum)))
	return int(math.Max(float64(lDiameter), float64(rDiameter))) + 1
}

func (bt *BinaryTree) getMaxDepth() int {
	return maxDepth(bt.root)
}

func maxDepth(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(maxDepth(root.lChild)), float64(maxDepth(root.rChild)))) + 1
}

func (bt *BinaryTree) isBalanced() bool {
	return balanced(bt.root) != -1
}

func balanced(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}

	left, right := balanced(root.lChild), balanced(root.rChild)

	if left == -1 || right == -1 || math.Abs(float64(left)-float64(right)) > 1 {
		return -1
	}

	return 1 + int(math.Max(float64(left), float64(right)))
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
