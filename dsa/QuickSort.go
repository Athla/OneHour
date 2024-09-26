package dsa

func QuickSort(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		arr, p := partition(arr, low, high)
		arr = quickSort(arr, 0, p-1)
		arr = quickSort(arr, p+1, high)
	}

	return arr
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]

	return arr, i
}

// func partition(arr []int, low, high int) ([]int, int) {
// 	// set current pivot
// 	pivot := arr[high]
// 	i := low
//
// 	for j := low; j < high; j++ {
// 		// Do the necessary swaps in the current partition
// 		if arr[j] < pivot {
// 			arr[i], arr[j] = arr[j], arr[i]
// 			i++
// 		}
// 	}
//
// 	// puts the pivot in the correct position
// 	arr[i], arr[high] = arr[high], arr[i]
//
// 	return arr, i
//
// }
//
// func quickSort(arr []int, low, high int) []int {
// 	if low < high {
// 		var p int
// 		arr, p = partition(arr, low, high)
// 		arr = quickSort(arr, low, p-1)
// 		arr = quickSort(arr, p+1, high)
// 	}
// 	return arr
// }
//
// func QuickSort(arr []int) []int {
// 	return quickSort(arr, 0, len(arr)-1)
// }
