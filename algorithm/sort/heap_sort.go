package sort

// HeapSort sorts the slice using heap sort algorithm.
func HeapSort(arr SortInterface) {
	n := arr.Len()

	// Build max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Extract elements from heap one by one
	for i := n - 1; i > 0; i-- {
		arr.Swap(0, i)
		heapify(arr, i, 0)
	}
}

// heapify ensures the subtree rooted at i is a max heap
func heapify(arr SortInterface, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr.Less(largest, left) {
		largest = left
	}
	if right < n && arr.Less(largest, right) {
		largest = right
	}
	if largest != i {
		arr.Swap(i, largest)
		heapify(arr, n, largest)
	}
}
