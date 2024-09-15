
## QuickSort
- Given an array, defines it's high and low by it's len and starting index
- Check if Low < High:
    1. Case true -> Call the partition function:
    2.  1. Partition -> Given the array, and the low and high points, set a idx and pivot value, with those values, iterate over until the value of J is smaller than the high pointer. After that, perform checks regarding the value of J in the scope, if it's smaller than the pivot, realize the swap of i (the low ptr), with the high ptr in the array values:
```go
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
```
-   2. 2. By doing are correctly swapped, needing only one last swap to put the pivot in the correct poistion.
-   3. After that , the quick sort is called in the return values of the partition function (the new arr and the new pivot)
-   4. at the end of the process, the value is returned.
### Code:
```go
package algorithms

func partition(arr []int, low, high int) ([]int, int) {
	// set current pivot
	pivot := arr[high]
	i := low

	for j := low; j < high; j++ {
		// Do the necessary swaps in the current partition
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	// puts the pivot in the correct position
	arr[i], arr[high] = arr[high], arr[i]

	return arr, i

}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func QuickSort(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

```
## MergeSort
- Similarly to QuickSort, it's a divide and conquer algorithm that works by:
-   1. Given an array, check if it's len is <= 2. That's the stop point of the recursion, where the given array is returned.
-   2. If the len of the array is > 2, it's split in the middle via **[:len(arr)/2]** or **[len(arr)/2:]** and then given to the mergeSort, this process is responsible to create the smallest possible slice of the array.
-   3. After that, the _first_ and _second_ slices are passed to a merging function.
-   
```go
func merge(a, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		final = append(final, a[i])
	}

	for ; j < len(b); j++ {
		final = append(final, b[j])
	}

	return final
}
```
-   4. This function perform the following steps:
    - Check if values inside the partitions are greater than one or other, and then add those values to final array
    - After that, go through each partition and add the remainder values to them, in order
    - Return the final array.
-   5. The final function implementation looks like:

### Code
```go
func merge(a, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		final = append(final, a[i])
	}

	for ; j < len(b); j++ {
		final = append(final, b[j])
	}

	return final
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	first := mergeSort(arr[:len(arr)/2])
	second := mergeSort(arr[len(arr)/2:])
	return merge(first, second)
}

func MergeSort(arr []int) []int {
	return mergeSort(arr)
}
```

