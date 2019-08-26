package main

func QuickSort1(arr []int) {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}
	quickSort(arr, 0, arrLen-1)
}

func quickSort1(arr []int, start, end int) {
	if start >= end {
		return
	}

	pivot := partition1(arr, start, end)
	quickSort1(arr, start, pivot)
	quickSort(arr, pivot+1, end)
}

func partition1(arr []int, low, high int) int {
	pivotV := arr[low]
	for low < high {
		for low < high && arr[high] > pivotV { //指针从右边开始向右找到一个比pivot小的数
			high--
		}
		arr[low] = arr[high] //将这个数放到low位，注意第一次这个位置放的是pivot值，所以不会丢

		for low < high && arr[low] < pivotV { //指针从左边开始向右找到第一个比pivot大的数
			low++
		}
		arr[high] = arr[low] //将这个数赋值给之前的high指针，因为之前high指针指向的数已经被一定，所以不会丢
	}

	//最后将pivot的值放入合适位置，此时low与high相等
	arr[low] = pivotV
	return low
}
