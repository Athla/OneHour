package leetcode

// Consider the matrix as a flattend 1D array to allow O(log(n * m))
func searchMatrix(matrix [][]int, target int) bool {
	// find matrix in bounds

	m := len(matrix)    // rows
	n := len(matrix[0]) // columns

	l, r := 0, m*n-1 // 0, last matrix element

	for l <= r {
		mid := (l + r) / 2
		// mid/n -> row index
		// mid%n 0> col index
		mval := matrix[mid/n][mid%n]

		if mval == target {
			return true
		} else if mval < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}
