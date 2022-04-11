package main

import (
	"fmt"

	"github.com/RamazanZholdas/DataStructuresGolang/heaps"
)

func main() {
	// Binary heaps
	slice := []int{50, 16, 48, 14, 8, 34, 20, 9, 1, 5, 7}
	heap := heaps.BinaryHeap{Data: slice}
	fmt.Println("before insertion:", heap)
	heap.Insert(63)
	fmt.Println("after insertion:", heap)
	heap.ExtractMax()
	fmt.Println("after extraction:", heap)
}
