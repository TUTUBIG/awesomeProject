package main

import "testing"

func TestTree(t *testing.T) {
	bt := initRoot(1)
	bt.insert(3)
	bt.insert(4)
	bt.insert(2)
	bt.print()
}

/*
*查找数据，也可以高效插入、删除数据
*每个节点的key值大于左子节点，小于右子节点
*不一定是完全二叉树
 */
func TestBinarySearchTree(t *testing.T) {

}
