package main

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)-1]
	left := make([]int, 0)
	right := make([]int, 0)

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)

	return append(append(left, pivot), right...)
}
