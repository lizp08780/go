package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"lizp.com/algorithm/sort"
)

func main() {
	a := []int{1, 2, 0}
	sort.HeapSort(a)
	fmt.Println(a)
	beego.Run()
}
