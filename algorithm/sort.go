package algorithm

import (
	"fmt"
	"math/rand"
)

type Sort []int

func (s Sort) Init() {

	for i := range s {
		s[i] = rand.Intn(len(s))
	}

	fmt.Println("------Init------", s)

}

func (s Sort) Select() {
	if len(s) <= 1 {
		return
	}

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	fmt.Println("------Select------", s)
}

func (s Sort) Bubble() {
	if len(s) <= 1 {
		return
	}

	for j := len(s) - 1; j > 0; j-- {
		for i := 0; i < j; i++ {
			if s[i] > s[i+1] {
				s[i], s[i+1] = s[i+1], s[i]
			}
		}
	}
	fmt.Println("------Bubble------", s)
}

func partition(arr []int, low, high int) int {
	pivotV := arr[low]
	for low < high {
		for low < high && (arr[high] >= pivotV) {
			high--
		}
		arr[low], arr[high] = arr[high], arr[low]

		for low < high && arr[low] <= pivotV {
			low++
		}
		arr[low], arr[high] = arr[high], arr[low]
	}
	return low
}

func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	pivot := partition(arr, start, end)
	quickSort(arr, start, pivot)
	quickSort(arr, pivot+1, end)
}

func (s Sort) Quick() {
	if len(s) <= 1 {
		return
	}
	quickSort(s, 0, len(s)-1)

	fmt.Println("------Quick------", s)
}

func maxHeap(array []int, rootIndex int, size int) {
	if rootIndex < size {
		maxIndex := rootIndex
		leftIndex, rightIndex := 2*rootIndex+1, 2*rootIndex+2
		if leftIndex < size {
			if array[leftIndex] > array[maxIndex] {
				maxIndex = leftIndex
			}
		}

		if rightIndex < size {
			if array[rightIndex] > array[maxIndex] {
				maxIndex = rightIndex
			}
		}
		if maxIndex != rootIndex {
			array[rootIndex], array[maxIndex] = array[maxIndex], array[rootIndex]
		}
	}

}

func (s Sort) Heap() {
	if len(s) <= 1 {
		return
	}

	for j := len(s); j > 0; j-- {
		t := s[:j]
		for i := len(t) - 1; i >= 0; i-- {
			maxHeap(t, i, len(t))
		}
		t[0], t[len(t)-1] = t[len(t)-1], t[0]
	}

	fmt.Println("--------Heap-------", s)

}
