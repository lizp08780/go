package main

import (
	"fmt"
	"lizp.com/algorithm/sort"
)

func main() {
	a := []int{1, 2, 0, 4, 3, 3}
	//sort.HeapSort(a)
	//sort.BubbleSort2
	//sort.InsertSort(a)
	sort.SelectSort(a)
	fmt.Println(a)
	b := sort.MergeSort(a)
	//b := sort.QuickSort(a)
	fmt.Println(b)
	//beego.Run()
}
