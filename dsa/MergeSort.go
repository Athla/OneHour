package dsa

func MergeSort(arr []int) []int {
	return mergeSort(arr)
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left := mergeSort(arr[len(arr)/2:])
	right := mergeSort(arr[:len(arr)/2])
	return merge(left, right)
}

func merge(a, b []int) []int {
	var result []int
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		result = append(result, a[i])
	}
	for ; j < len(b); j++ {
		result = append(result, b[j])
	}

	return result

}
