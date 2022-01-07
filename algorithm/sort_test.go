package main

import "testing"

func TestSort(t *testing.T) {
	s := make(Sort, 10)
	/*s.Init()
	s.Bubble()
	s.Init()
	s.Select()*/
	s.Init()
	//s.Quick()
	s.Heap()
}
