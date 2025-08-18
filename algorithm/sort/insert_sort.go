package sort

// InsertSort sorts the slice using insertion sort algorithm.
func InsertSort(arr SortInterface) {
	for i := 1; i < arr.Len(); i++ {
		for j := i; j > 0 && arr.Less(j, j-1); j-- {
			if arr.Less(j, j-1) {
				// Assuming a method to swap elements exists
				arr.Swap(j, j-1)
			}
		}
	}
}
