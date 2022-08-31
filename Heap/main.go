package main

import "fmt"

// Maxheap has a structure that holds the array
type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest key and removes it from the heap
func (h *MaxHeap) Extract() int {
	if len(h.array) == 0 {
		panic("Heap overflow !!!")
	}
	top := h.array[0]
	last := len(h.array) - 1
	h.array[0] = h.array[last]
	h.array = h.array[:last]
	h.maxHeapifyDown()
	return top
}


// get parent index
func getParent(index int) int {
	return index / 2
}

// get left child index 
func left(index int) int {
	return index * 2 + 1
}

// get right child index 
func right(index int) int {
	return index * 2 + 2
}

// swap two elements in the array 
func (h *MaxHeap) swap(x, y int) {
	h.array[x], h.array[y] = h.array[y], h.array[x]
}


// maintains the conditions of a heap after insertion
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[getParent((index))] < h.array[index] {
		h.swap(getParent(index), index)
		index = getParent(index)
	}
}

// maintains the conditions of a heap after deletion
func (h *MaxHeap) maxHeapifyDown() {
	index := 0
	last := len(h.array) - 1
	childToCompare := 0
	for left(index) <= last {
		if left(index) == last {
			childToCompare = last
		} else if h.array[left(index)] > h.array[right(index)] {
			childToCompare = left(index)
		} else {
			childToCompare = right(index)
		}
		if h.array[childToCompare] > h.array[index] {
			h.swap(childToCompare, index)
			index = childToCompare
		} else {
			return 
		}
	}
}

func main() {
	m := &MaxHeap{}
	buildHeap := []int{10, 20, 30, 15, 25, 35}	
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}
	for i := 0; i < 6;i++ {
		m.Extract()
		fmt.Println(m)
	}
}
