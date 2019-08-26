package main

import (
	"math/rand"
	"fmt"
)

type Sort []int

func (s Sort) Init() {

	for i := range s {
		s[i] = rand.Intn(len(s))
	}
}

func (s Sort) UnKnow() {
	fmt.Println("------UnKnow------")
	if len(s) <= 1 {
		return
	}

	for i := 0;i < len(s);i++ {
		for j := i+1;j < len(s);j++ {
			if s[i] > s[j] {
				s[i],s[j] = s[j],s[i]
			}
		}
	}
}

func (s Sort) Bubble() {
	fmt.Println("------Bubble------")
	if len(s) <= 1 {
		return
	}

	for i := 0;i < len(s)-1;i++ {
		for j := i;j < len(s)-i-1;j++ {
			if s[j] > s[j+1] {
				s[j],s[j+1] = s[j+1],s[j]
			}
		}
	}
}

func partition(arr []int,low,high int) int {
	pivotV := arr[low]
	for low < high {
		for low < high && arr[high] > pivotV {
			high--
		}
		arr[low] = arr[high]

		for low < high && arr[low] < pivotV {
			low++
		}
		arr[high] = arr[low]
	}
	arr[low] = pivotV
	return low
}

func quickSort(arr []int,start,end int) {
	if start >= end {
		return
	}

	pivot := partition(arr,start,end)
	quickSort(arr,start,pivot)
	quickSort(arr,pivot+1,end)
}

func (s Sort) Quick() {
	fmt.Println("------Quick------")
	if len(s) <= 1 {
		return
	}
	quickSort(s,0,len(s)-1)

}

func main()  {

	/*s := make(Sort,20)
	s.Init()
	fmt.Println(s)
	s.UnKnow()
	fmt.Println(s)
	s.Init()
	fmt.Println(s)
	s.Bubble()
	fmt.Println(s)
	s.Init()
	fmt.Println(s)
	s.Quick()
	fmt.Println(s)*/

	fmt.Println("--------",1 << 50)

}
