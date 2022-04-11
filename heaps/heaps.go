package heaps

import (
	"errors"
	"fmt"
)

type BinaryHeap struct {
	Data []int
}

func (h *BinaryHeap) Len() int {
	return len(h.Data)
}

func (h *BinaryHeap) Insert(val int) {
	h.Data = append(h.Data, val)
	h.hepify()
}

func (h *BinaryHeap) ExtractMax() (int, error) {
	if h.Len() == 0 {
		return 0, errors.New("slice is empty")
	}
	max := h.Data[0]
	h.Data = h.Data[1:]
	h.topDownHeapify()
	index, ok := h.checkHeap()
	if !ok {
		h.fixHeap(index)
	}
	return max, nil
}

func (h *BinaryHeap) topDownHeapify() {
	var temp int
	index := 0
	fmt.Println("initial data", h.Data)
	for {
		if h.left(index) > h.Len()-1 {
			break
		} else if h.Data[index] > h.Data[h.left(index)] && h.Data[index] > h.Data[h.right(index)] {
			break
		} else if h.Data[index] < h.Data[h.right(index)] {
			temp = h.Data[index]
			h.Data[index] = h.Data[h.right(index)]
			h.Data[h.right(index)] = temp
			index = h.right(index)
		} else if h.Data[index] < h.Data[h.left(index)] {
			temp = h.Data[index]
			h.Data[index] = h.Data[h.left(index)]
			h.Data[h.left(index)] = temp
			index = h.left(index)
		}
		fmt.Println(h.Data)
	}
}

func (h *BinaryHeap) checkHeap() (int, bool) {
	for i := 0; i < h.Len(); i++ {
		if h.left(i) < h.Len() {
			if h.Data[i] < h.Data[h.left(i)] {
				fmt.Println("left child is smaller than parent")
				return i, false
			}
		}
		if h.right(i) < h.Len() {
			if h.Data[i] < h.Data[h.right(i)] {
				fmt.Println("right child is smaller than parent")
				return i, false
			}
		}
	}
	return 0, true
}

func (h *BinaryHeap) fixHeap(index int) {
	var temp int
	for {
		if h.left(index) > h.Len()-1 {
			break
		} else if h.Data[index] > h.Data[h.left(index)] && h.Data[index] > h.Data[h.right(index)] {
			break
		} else if h.Data[index] < h.Data[h.right(index)] {
			temp = h.Data[index]
			h.Data[index] = h.Data[h.right(index)]
			h.Data[h.right(index)] = temp
			index = h.right(index)
		} else if h.Data[index] < h.Data[h.left(index)] {
			temp = h.Data[index]
			h.Data[index] = h.Data[h.left(index)]
			h.Data[h.left(index)] = temp
			index = h.left(index)
		}
	}
}

func (h *BinaryHeap) hepify() {
	var temp int
	index := h.Len() - 1
	for {
		if h.parent(index) < 0 {
			break
		} else if h.Data[index] < h.Data[h.parent(index)] {
			break
		} else if h.Data[index] > h.Data[h.parent(index)] {
			temp = h.Data[h.parent(index)]
			h.Data[h.parent(index)] = h.Data[index]
			h.Data[index] = temp
			index = h.parent(index)
		}
	}
}

func (h *BinaryHeap) parent(i int) (index int) {
	if i%2 == 0 {
		index = (i - 2) / 2
	} else {
		index = (i - 1) / 2
	}
	return index
}

func (h *BinaryHeap) left(i int) int {
	return 2*i + 1
}

func (h *BinaryHeap) right(i int) int {
	return 2*i + 2
}
