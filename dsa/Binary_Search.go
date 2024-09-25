package dsa

// Expects a ordered array, and a target value, returns true if it exists and the index of it, if does not exists, return false and -1
func BinarySearch(arr []int, target int) (bool, int) {
	if len(arr) == 0 {
		return false, -1
	}
	l, h := 0, len(arr)-1

	for l <= h {
		mid := (l + h) / 2
		if arr[mid] > target {
			h = mid - 1
		} else if arr[mid] < target {
			l = mid + 1
		} else {
			return true, mid
		}
	}

	return false, -1
}
